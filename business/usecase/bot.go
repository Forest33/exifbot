// Package usecase provides business logic.
package usecase

import (
	"context"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/forest33/exifbot/business/entity"
	"github.com/forest33/exifbot/pkg/database"
	"github.com/forest33/exifbot/pkg/logger"
)

// BotUseCase object capable of interacting with BotUseCase
type BotUseCase struct {
	ctx              context.Context
	cfg              *entity.BotConfig
	log              *logger.Zerolog
	db               *database.Database
	usersRepo        UsersRepo
	bot              *tgbotapi.BotAPI
	updatesChan      tgbotapi.UpdatesChannel
	commandChan      chan *commandJob
	senderChan       chan *senderJob
	parserResultChan chan *entity.ParserJobResult
	localizer        map[string]*i18n.Localizer
}

// UsersRepo is the common interface implemented UsersRepository methods
type UsersRepo interface {
	Create(ctx context.Context, in *entity.User) (*entity.User, error)
	Get(ctx context.Context, filter *entity.UsersFilter) ([]*entity.User, error)
}

type commandJob struct {
	msg *tgbotapi.Message
}

type senderJob struct {
	msg  *tgbotapi.Message
	text string
}

// NewBotUseCase creates a new BotUseCase
func NewBotUseCase(ctx context.Context, cfg *entity.BotConfig, log *logger.Zerolog, db *database.Database, usersRepo UsersRepo) (*BotUseCase, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	uc := &BotUseCase{
		ctx:              ctx,
		cfg:              cfg,
		log:              log,
		db:               db,
		usersRepo:        usersRepo,
		commandChan:      make(chan *commandJob, cfg.CommandWorkersPoolSize),
		senderChan:       make(chan *senderJob, cfg.SenderWorkersPoolSize),
		parserResultChan: make(chan *entity.ParserJobResult, cfg.SenderWorkersPoolSize),
	}

	if err := uc.initLocales(); err != nil {
		return nil, err
	}

	return uc, uc.init()
}

func (uc *BotUseCase) Start() {
	go uc.loop()
}

func (uc *BotUseCase) init() error {
	var err error
	uc.bot, err = tgbotapi.NewBotAPI(uc.cfg.Token)
	if err != nil {
		return err
	}
	uc.bot.Debug = uc.cfg.Debug

	u := tgbotapi.NewUpdate(0)
	u.Timeout = uc.cfg.UpdateTimeout
	uc.updatesChan = uc.bot.GetUpdatesChan(u)

	for i := 0; i < uc.cfg.CommandWorkersPoolSize; i++ {
		go uc.commandWorker()
	}

	for i := 0; i < uc.cfg.SenderWorkersPoolSize; i++ {
		go uc.senderWorker()
		go uc.parserWorker()
	}

	return nil
}

func (uc *BotUseCase) loop() {
	defer func() {
		uc.bot.StopReceivingUpdates()
	}()

	for {
		select {
		case <-uc.ctx.Done():
			return
		case update, ok := <-uc.updatesChan:
			if !ok {
				return
			}
			if update.Message == nil {
				continue
			}

			uc.commandChan <- &commandJob{
				msg: update.Message,
			}
		}
	}
}

func (uc *BotUseCase) commandWorker() {
	errFunc := func(msg *tgbotapi.Message, err error) {
		uc.log.Error().
			Err(err).
			Int64("user_id", msg.From.ID).
			Str("text", msg.Text).
			Msg("processing command error")
	}

	for {
		select {
		case <-uc.ctx.Done():
			return
		case job, ok := <-uc.commandChan:
			if !ok {
				return
			}

			err := uc.commandHandler(job.msg)
			if err != nil {
				errFunc(job.msg, err)
				continue
			}

			uc.log.Debug().
				Int64("user_id", job.msg.From.ID).
				Str("text", job.msg.Text).
				Msg("processing command successful")
		}
	}
}

func (uc *BotUseCase) senderWorker() {
	for {
		select {
		case <-uc.ctx.Done():
			return
		case job, ok := <-uc.senderChan:
			if !ok {
				return
			}
			if err := uc.reply(job.msg, job.text); err != nil {
				uc.log.Error().Err(err).
					Int64("user_id", job.msg.From.ID).
					Str("user_name", job.msg.From.FirstName).
					Msg("failed to send result message")
			}
		}
	}
}

func (uc *BotUseCase) commandHandler(msg *tgbotapi.Message) error {
	sendError := func(err error) {
		if err := uc.reply(msg, uc.localize(msg.From.LanguageCode, entity.GetErrorMessageID(err))); err != nil {
			uc.log.Error().Err(err).
				Int64("user_id", msg.From.ID).
				Str("user_name", msg.From.FirstName).
				Msg("failed to send error message")
		}
	}

	cmd, args, err := uc.parseCommand(msg)
	if err != nil {
		sendError(err)
		return err
	}

	switch cmd {
	case entity.CommandStart:
		err = uc.createUser(msg, args)
	case entity.CommandHelp:
		err = uc.reply(msg, uc.localize(msg.From.LanguageCode, "messageHelp"))
	case entity.CommandGetFromURL:
		uc.getExif(msg, args)
	default:
		err = entity.ErrUnknownBotCommand
	}

	if err != nil {
		sendError(err)
	}

	return err
}

func (uc *BotUseCase) parseCommand(m *tgbotapi.Message) (string, []string, error) {
	text := strings.TrimSpace(m.Text)
	args := strings.Split(text, " ")

	if m.Photo != nil && len(m.Photo) != 0 {
		if url, err := uc.bot.GetFileDirectURL(m.Photo[len(m.Photo)-1].FileID); err == nil {
			return entity.CommandGetFromURL, []string{url}, nil
		}
		return "", nil, entity.ErrInternal
	} else if m.Document != nil {
		if url, err := uc.bot.GetFileDirectURL(m.Document.FileID); err == nil {
			return entity.CommandGetFromURL, []string{url}, nil
		}
		return "", nil, entity.ErrInternal
	} else if m.Sticker != nil {
		if url, err := uc.bot.GetFileDirectURL(m.Sticker.FileID); err == nil {
			return entity.CommandGetFromURL, []string{url}, nil
		}
		return "", nil, entity.ErrInternal
	} else if !strings.HasPrefix(args[0], "/") {
		if entity.ValidateURL(text) != nil {
			return "", nil, entity.ErrWrongURL
		}
		return entity.CommandGetFromURL, []string{text}, nil
	}

	args[0] = strings.ToLower(args[0])

	if len(args) > 1 {
		return args[0], args[1:], nil
	}

	return args[0], nil, nil
}

func (uc *BotUseCase) reply(msg *tgbotapi.Message, text string) error {
	if len(text) == 0 {
		return nil
	}

	reply := tgbotapi.NewMessage(msg.Chat.ID, text)
	reply.ReplyToMessageID = msg.MessageID
	reply.DisableWebPagePreview = true
	reply.ParseMode = tgbotapi.ModeHTML

	_, err := uc.bot.Send(reply)

	return err
}

func (uc *BotUseCase) send(msg *tgbotapi.Message, text string) {
	uc.senderChan <- &senderJob{
		msg:  msg,
		text: text,
	}
}
