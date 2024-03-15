package usecase

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"

	"github.com/forest33/exifbot/business/entity/locales"
)

const defaultLanguage = "en"

func (uc *BotUseCase) initLocales() error {
	bundle := i18n.NewBundle(language.English)
	if err := bundle.AddMessages(language.English, locales.EN...); err != nil {
		return err
	}
	if err := bundle.AddMessages(language.Russian, locales.RU...); err != nil {
		return err
	}

	uc.localizer = map[string]*i18n.Localizer{
		"en": i18n.NewLocalizer(bundle, language.English.String()),
		"ru": i18n.NewLocalizer(bundle, language.Russian.String()),
	}

	return nil
}

func (uc *BotUseCase) localize(lang, messageID string) string {
	if messageID == "" {
		return ""
	}

	loc, ok := uc.localizer[lang]
	if !ok {
		loc = uc.localizer[defaultLanguage]
	}

	res, err := loc.Localize(&i18n.LocalizeConfig{MessageID: messageID})
	if err != nil {
		uc.log.Warn().Err(err).
			Str("language", lang).
			Str("message_id", messageID).
			Msg("failed to get message")

		if lang != defaultLanguage {
			res, err = uc.localizer[defaultLanguage].Localize(&i18n.LocalizeConfig{MessageID: messageID})
			if err != nil {
				uc.log.Error().Err(err).
					Str("language", defaultLanguage).
					Str("message_id", messageID).
					Msg("failed to get message")
			}
		}
	}

	return res
}
