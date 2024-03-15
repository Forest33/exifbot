package entity

import (
	"errors"

	"github.com/forest33/exifbot/business/entity/locales"
)

var (
	// ErrUserNotExists error user not exists
	ErrUserNotExists = errors.New("user not exists")
	// ErrUnknownBotCommand error user not exists
	ErrUnknownBotCommand = errors.New("unknown bot command")
	// ErrNotCommand error not a command
	ErrNotCommand = errors.New("not a command")
	// ErrWrongNumberOfArguments error wrong number of arguments
	ErrWrongNumberOfArguments = errors.New("wrong number of arguments")
	// ErrWrongURL error wrong URL
	ErrWrongURL = errors.New("wrong URL")
	// ErrLoadURL failed to load URL
	ErrLoadURL = errors.New("failed to load URL")
	// ErrInternal internal error
	ErrInternal = errors.New("internal error")
	// ErrNoExif no EXIF data found
	ErrNoExif = errors.New("no EXIF data found")
)

func GetErrorMessageID(err error) string {
	switch {
	case errors.Is(err, ErrUnknownBotCommand), errors.Is(err, ErrNotCommand), errors.Is(err, ErrWrongNumberOfArguments):
		return locales.LIdMessageUnknownCommand
	case errors.Is(err, ErrWrongURL):
		return locales.LIdMessageWrongURL
	case errors.Is(err, ErrInternal):
		return locales.LIdMessageInternalError
	case errors.Is(err, ErrNoExif):
		return locales.LIdMessageNoExif
	case errors.Is(err, ErrLoadURL):
		return locales.LIdMessageErrLoadURL
	default:
		return ""
	}
}
