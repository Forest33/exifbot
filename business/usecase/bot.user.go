package usecase

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/forest33/exifbot/business/entity"
)

func (uc *BotUseCase) createUser(msg *tgbotapi.Message, _ []string) error {
	ctx, err := uc.db.BeginTransaction(uc.ctx)
	if err != nil {
		uc.log.Error().Err(err).Msg("failed to begin transaction")
		return entity.ErrInternal
	}
	defer uc.db.CommitTransaction(ctx, err)

	_, err = uc.usersRepo.Create(ctx, &entity.User{
		ID:        msg.From.ID,
		FirstName: msg.From.FirstName,
		LastName:  msg.From.LastName,
		UserName:  msg.From.UserName,
		Language:  msg.From.LanguageCode,
	})
	if err != nil {
		return err
	}

	err = uc.reply(msg, uc.localize(msg.From.LanguageCode, "messageStart"))
	if err != nil {
		uc.log.Error().Err(err).
			Int64("user_id", msg.From.ID).
			Str("user_name", msg.From.FirstName).
			Msg("failed to send start message")
	}

	return err
}
