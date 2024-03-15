package entity

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const (
	CommandStart      = "/start"
	CommandGetFromURL = "/url"
	CommandHelp       = "/help"
)

func ValidateURL(url string) error {
	return validation.Validate(url, validation.Required, is.URL)
}
