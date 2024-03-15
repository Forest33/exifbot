// Package main exifbot main package
package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/forest33/exifbot/business/entity"
	"github.com/forest33/exifbot/business/usecase"
	"github.com/forest33/exifbot/pkg/database"
	"github.com/forest33/exifbot/pkg/logger"

	db "github.com/forest33/exifbot/adapter/database"
)

var (
	log *logger.Zerolog
)

var (
	cfg = &entity.Config{}
	dbi *database.Database

	usersRepo  *db.UsersRepository
	botUseCase *usecase.BotUseCase

	ctx    context.Context
	cancel context.CancelFunc
)

func init() {
	if cfg.Runtime.GoMaxProcs == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	} else {
		runtime.GOMAXPROCS(cfg.Runtime.GoMaxProcs)
	}

	ctx, cancel = context.WithCancel(context.Background())
}

func main() {
	defer shutdown()

	log = logger.NewZerolog(logger.ZeroConfig{
		Level:             cfg.Logger.Level,
		TimeFieldFormat:   cfg.Logger.TimeFieldFormat,
		PrettyPrint:       cfg.Logger.PrettyPrint,
		DisableSampling:   cfg.Logger.DisableSampling,
		RedirectStdLogger: cfg.Logger.RedirectStdLogger,
		ErrorStack:        cfg.Logger.ErrorStack,
		ShowCaller:        cfg.Logger.ShowCaller,
	})

	initDatabase()
	initAdapters()
	initUseCases()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func initAdapters() {
	usersRepo = db.NewUsersRepository(dbi, log)
}

func initUseCases() {
	botUseCase, err := usecase.NewBotUseCase(ctx, cfg.Bot, log, dbi, usersRepo)
	if err != nil {
		log.Fatal(err)
	}

	exifUseCase, err := usecase.NewParserUseCase(ctx, cfg.Parser, log)
	if err != nil {
		log.Fatal(err)
	}

	botUseCase.Start()

	usecase.SetParserUseCase(exifUseCase)
}

func shutdown() {
	cancel()
	dbi.Close()
}
