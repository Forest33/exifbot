package usecase

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/dsoprea/go-exif/v3"
	"github.com/dsoprea/go-exif/v3/common"

	"github.com/forest33/exifbot/business/entity"
	"github.com/forest33/exifbot/pkg/logger"
	"github.com/forest33/exifbot/pkg/structs"
)

const (
	userAgent = "exifbot (https://github.com/forest33/exifbot)"
)

type ParserUseCase struct {
	ctx         context.Context
	cfg         *entity.ParserConfig
	log         *logger.Zerolog
	workerChan  chan *entity.ParserJob
	exifTagsMap map[string]entity.ExifTag
}

func NewParserUseCase(ctx context.Context, cfg *entity.ParserConfig, log *logger.Zerolog) (*ParserUseCase, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	uc := &ParserUseCase{
		ctx:         ctx,
		cfg:         cfg,
		log:         log,
		workerChan:  make(chan *entity.ParserJob, cfg.WorkersPoolSize),
		exifTagsMap: structs.SliceToMap(entity.ExifTags, func(e entity.ExifTag) string { return e.Name }),
	}

	uc.init()

	return uc, nil
}

func (uc *ParserUseCase) init() {
	for _ = range uc.cfg.WorkersPoolSize {
		go uc.worker()
	}
}

func (uc *ParserUseCase) worker() {
	for {
		select {
		case <-uc.ctx.Done():
			return
		case job, ok := <-uc.workerChan:
			if !ok {
				return
			}

			res, err := uc.Parse(job.URL)
			job.CallbackChan <- &entity.ParserJobResult{
				URL:          job.URL,
				Exif:         res,
				Err:          err,
				CallbackData: job.CallbackData,
			}
		}
	}
}

func (uc *ParserUseCase) Parse(url string) (map[string]*entity.ExifResultItem, error) {
	data, err := uc.loadURL(url)
	if err != nil {
		uc.log.Error().Err(err).Msg("failed to receive image")
		return nil, entity.ErrLoadURL
	}

	rawExif, err := exif.SearchAndExtractExif(data)
	if err != nil {
		return nil, entity.ErrNoExif
	}

	tags, _, err := exif.GetFlatExifDataUniversalSearch(rawExif, nil, true)
	if err != nil || tags == nil {
		return nil, entity.ErrNoExif
	}

	tagsMap := structs.SliceToMap(tags, func(t exif.ExifTag) string { return t.TagName })
	res := make(map[string]*entity.ExifResultItem, len(tags))

	for tagName := range tagsMap {
		tag, ok := uc.exifTagsMap[tagName]
		if !ok {
			continue
		}

		row := &entity.ExifResultItem{}
		if tag.Handler != nil {
			row = tag.Handler(tagsMap[tagName])
		} else {
			row.Value = tagsMap[tagName].FormattedFirst
		}
		if row.IsEmpty() {
			continue
		}

		res[tagName] = row
	}

	if gps := uc.parseGPS(tagsMap); gps != nil {
		res[entity.ExifGPSTag] = gps
	}

	return res, nil
}

func (uc *ParserUseCase) parseGPS(tags map[string]exif.ExifTag) *entity.ExifResultItem {
	lat, ok := tags["GPSLatitude"]
	if !ok {
		return nil
	}
	lng, ok := tags["GPSLongitude"]
	if !ok {
		return nil
	}
	if lat.TagTypeId != exifcommon.TypeRational || lng.TagTypeId != exifcommon.TypeRational ||
		len(lat.Value.([]exifcommon.Rational)) != 3 || len(lng.Value.([]exifcommon.Rational)) != 3 {
		return nil
	}

	latVal := lat.Value.([]exifcommon.Rational)
	lngVal := lng.Value.([]exifcommon.Rational)
	latFloat := (float64(latVal[0].Numerator) / float64(latVal[0].Denominator)) + (float64(latVal[1].Numerator)/float64(latVal[1].Denominator))/60 + (float64(latVal[2].Numerator)/float64(latVal[2].Denominator))/3600
	lngFloat := (float64(lngVal[0].Numerator) / float64(lngVal[0].Denominator)) + (float64(lngVal[1].Numerator)/float64(lngVal[1].Denominator))/60 + (float64(lngVal[2].Numerator)/float64(lngVal[2].Denominator))/3600

	return &entity.ExifResultItem{
		Value: fmt.Sprintf("%f,%f", latFloat, lngFloat),
	}
}

func (uc *ParserUseCase) loadURL(url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(uc.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			uc.log.Error().Err(err).Msg("failed to close response body")
		}
	}()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, entity.ErrLoadURL
	}

	return io.ReadAll(resp.Body)
}

func (uc *ParserUseCase) GetChannel() chan *entity.ParserJob {
	return uc.workerChan
}
