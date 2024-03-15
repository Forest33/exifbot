package usecase

import (
	"fmt"
	"strings"
	"unicode/utf8"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/forest33/exifbot/business/entity"
	"github.com/forest33/exifbot/business/entity/locales"
)

const resultSeparator = "\t\t"

func (uc *BotUseCase) getExif(msg *tgbotapi.Message, url []string) {
	parserUseCase.GetChannel() <- &entity.ParserJob{
		URL:          url[0],
		CallbackData: msg,
		CallbackChan: uc.parserResultChan,
	}
}

func (uc *BotUseCase) parserWorker() {
	for {
		select {
		case <-uc.ctx.Done():
			return
		case job, ok := <-uc.parserResultChan:
			if !ok {
				return
			}
			if job.Err != nil {
				uc.log.Error().Err(job.Err).Str("url", job.URL).Msg("failed to parse image")
				uc.send(job.CallbackData.(*tgbotapi.Message), uc.localize(job.CallbackData.(*tgbotapi.Message).From.LanguageCode, entity.GetErrorMessageID(job.Err)))
				continue
			}

			uc.log.Info().Str("url", job.URL).Msg("parsing image successful")

			msg := uc.CreateExifMessage(job.Exif, job.CallbackData.(*tgbotapi.Message).From.LanguageCode)

			uc.send(job.CallbackData.(*tgbotapi.Message), msg)
		}
	}
}

func (uc *BotUseCase) CreateExifMessage(exif map[string]*entity.ExifResultItem, lang string) string {
	var (
		st             = strings.Builder{}
		maxLabelLength int
	)

	for _, tag := range entity.ExifTags {
		if _, ok := exif[tag.Name]; !ok {
			continue
		}
		label := uc.localize(lang, tag.LabelID)
		if l := utf8.RuneCountInString(label); l > maxLabelLength {
			maxLabelLength = l
		}
	}

	for _, tag := range entity.ExifTags {
		data, ok := exif[tag.Name]
		if !ok {
			continue
		}

		label := uc.localize(lang, tag.LabelID)
		st.WriteString(label)
		st.WriteString(":")
		st.WriteString(strings.Repeat(" ", maxLabelLength-utf8.RuneCountInString(label)))
		st.WriteString(resultSeparator)

		value := data.Value
		if len(data.ValueID) != 0 {
			value = uc.localize(lang, data.ValueID)
		}
		st.WriteString(value)

		if len(data.SuffixID) != 0 {
			st.WriteString(" ")
			st.WriteString(uc.localize(lang, data.SuffixID))
		}

		st.WriteString("\n")
	}

	text := "<pre>" + st.String() + "</pre>"

	if gps, ok := exif[entity.ExifGPSTag]; ok {
		label := uc.localize(lang, locales.LIdShowGPS)
		text += fmt.Sprintf("<b>%s: </b><u><a href=\"https://www.google.com/maps/search/?api=1&query=%s\">%s</a></u> 📌", label, gps.Value, gps.Value)
	}

	return text
}
