package entity

import (
	"fmt"
	"math"
	"strings"

	"github.com/dsoprea/go-exif/v3"
	"github.com/dsoprea/go-exif/v3/common"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/forest33/exifbot/business/entity/locales"
)

const (
	ExifGPSTag = "gps"
)

type ExifTag struct {
	Name    string
	LabelID string
	Handler func(v exif.ExifTag) *ExifResultItem
}

type ExifResultItem struct {
	ValueID  string
	SuffixID string
	Value    string
}

func (r *ExifResultItem) IsEmpty() bool {
	return r == nil || (len(r.Value) == 0 && len(r.ValueID) == 0)
}

var ExifTags = []ExifTag{
	{
		Name:    "DateTimeOriginal",
		LabelID: locales.LIdExifTagDateTimeOriginal,
		Handler: nil,
	},
	{
		Name:    "ExposureTime",
		LabelID: locales.LIdExifTagExposureTime,
		Handler: nil,
	},
	{
		Name:    "ExposureBiasValue",
		LabelID: locales.LIdExifTagExposureBiasValue,
		Handler: prepareExposureBiasValue,
	},
	{
		Name:    "ExposureMode",
		LabelID: locales.LIdExifTagExposureMode,
		Handler: prepareExposureMode,
	},
	{
		Name:    "ApertureValue",
		LabelID: locales.LIdExifTagApertureValue,
		Handler: prepareAperture,
	},
	{
		Name:    "ExposureProgram",
		LabelID: locales.LIdExifTagExposureProgram,
		Handler: prepareExposureProgram,
	},
	{
		Name:    "ISOSpeedRatings",
		LabelID: locales.LIdExifTagISOSpeedRatings,
		Handler: nil,
	},
	{
		Name:    "Flash",
		LabelID: locales.LIdExifTagFlash,
		Handler: prepareFlash,
	},
	{
		Name:    "FocalLength",
		LabelID: locales.LIdExifTagFocalLength,
		Handler: prepareFocalLength,
	},
	{
		Name:    "DigitalZoomRatio",
		LabelID: locales.LIdExifTagDigitalZoomRatio,
		Handler: prepareDigitalZoomRatio,
	},
	{
		Name:    "Make",
		LabelID: locales.LIdExifTagMake,
		Handler: prepareMake,
	},
	{
		Name:    "Model",
		LabelID: locales.LIdExifTagModel,
	},
	{
		Name:    "LensModel",
		LabelID: locales.LIdExifTagLensModel,
	},
	{
		Name:    "BodySerialNumber",
		LabelID: locales.LIdExifTagBodySerialNumber,
	},
	{
		Name:    "LensSerialNumber",
		LabelID: locales.LIdExifTagLensSerialNumber,
	},
	{
		Name:    "MeteringMode",
		LabelID: locales.LIdExifTagMeteringMode,
		Handler: prepareMeteringMode,
	},
	{
		Name:    "WhiteBalance",
		LabelID: locales.LIdExifTagWhiteBalance,
		Handler: prepareWhiteBalance,
	},
	{
		Name:    "Software",
		LabelID: locales.LIdExifTagSoftware,
	},
	{
		Name:    "ColorSpace",
		LabelID: locales.LIdExifTagColorSpace,
		Handler: prepareColorSpace,
	},
	{
		Name:    "Artist",
		LabelID: locales.LIdExifTagArtist,
	},
	{
		Name:    "Copyright",
		LabelID: locales.LIdExifTagCopyright,
	},
}

func getShort(v exif.ExifTag) *uint16 {
	if v.TagTypeId != exifcommon.TypeShort {
		return nil
	}
	value, ok := v.Value.([]uint16)
	if !ok || len(value) != 1 {
		return nil
	}
	return &value[0]
}

func prepareExposureMode(v exif.ExifTag) *ExifResultItem {
	tv := getShort(v)
	if tv == nil {
		return nil
	}
	switch *tv {
	case 0:
		return &ExifResultItem{ValueID: locales.LIdExifExposureModeAuto}
	case 1:
		return &ExifResultItem{ValueID: locales.LIdExifExposureModeManual}
	case 2:
		return &ExifResultItem{ValueID: locales.LIdExifExposureModeAutoBracket}
	default:
		return &ExifResultItem{ValueID: locales.LIdUnknownValue}
	}
}

func prepareExposureBiasValue(v exif.ExifTag) *ExifResultItem {
	if v.TagTypeId != exifcommon.TypeSignedRational {
		return nil
	}
	value, ok := v.Value.([]exifcommon.SignedRational)
	if !ok || len(value) != 1 {
		return nil
	}
	if value[0].Numerator == 0 {
		return nil
	}
	return &ExifResultItem{Value: fmt.Sprintf("%.2f EV", float32(value[0].Numerator)/float32(value[0].Denominator))}
}

func prepareExposureProgram(v exif.ExifTag) *ExifResultItem {
	tv := getShort(v)
	if tv == nil {
		return nil
	}
	switch *tv {
	case 0:
		return &ExifResultItem{ValueID: locales.LIdExifExposureProgramUndefined}
	case 1:
		return &ExifResultItem{ValueID: locales.LIdExifExposureProgramManual}
	case 2:
		return &ExifResultItem{ValueID: locales.LIdExifExposureProgramNormalProgram}
	case 3:
		return &ExifResultItem{ValueID: locales.LIdExifExposureProgramAperturePriority}
	case 4:
		return &ExifResultItem{ValueID: locales.LIdExifExposureProgramShutterPriority}
	case 5:
		return &ExifResultItem{ValueID: locales.LIdExifExposureProgramCreativeProgram}
	case 6:
		return &ExifResultItem{ValueID: locales.LIdExifExposureProgramActionProgram}
	case 7:
		return &ExifResultItem{ValueID: locales.LIdExifExposureProgramPortraitMode}
	case 8:
		return &ExifResultItem{ValueID: locales.LIdExifExposureProgramLandscapeMode}
	case 9:
		return &ExifResultItem{ValueID: locales.LIdExifExposureProgramBulb}
	default:
		return &ExifResultItem{ValueID: locales.LIdUnknownValue}
	}
}

func prepareMeteringMode(v exif.ExifTag) *ExifResultItem {
	tv := getShort(v)
	if tv == nil {
		return nil
	}
	switch *tv {
	case 1:
		return &ExifResultItem{ValueID: locales.LIdExifMeteringModeAverage}
	case 2:
		return &ExifResultItem{ValueID: locales.LIdExifMeteringModeCenterWeightedAverage}
	case 3:
		return &ExifResultItem{ValueID: locales.LIdExifMeteringModeSpot}
	case 4:
		return &ExifResultItem{ValueID: locales.LIdExifMeteringModeMultiSpot}
	case 5:
		return &ExifResultItem{ValueID: locales.LIdExifMeteringModePattern}
	case 6:
		return &ExifResultItem{ValueID: locales.LIdExifMeteringModePartial}
	case 255:
		return &ExifResultItem{ValueID: locales.LIdExifMeteringModeOther}
	default:
		return &ExifResultItem{ValueID: locales.LIdUnknownValue}
	}
}

func prepareAperture(v exif.ExifTag) *ExifResultItem {
	if v.TagTypeId != exifcommon.TypeRational {
		return nil
	}
	value, ok := v.Value.([]exifcommon.Rational)
	if !ok || len(value) != 1 {
		return nil
	}

	a := math.Pow(math.Sqrt(2), float64(value[0].Numerator)/float64(value[0].Denominator))
	strVal := fmt.Sprintf("f/%.2f", a)
	if strings.HasSuffix(strVal, "00") {
		strVal = fmt.Sprintf("f/%d", int(math.Round(a)))
	} else if strings.HasSuffix(strVal, "0") {
		strVal = fmt.Sprintf("f/%.1f", a)
	}

	return &ExifResultItem{Value: strVal}
}

func prepareFocalLength(v exif.ExifTag) *ExifResultItem {
	if v.TagTypeId != exifcommon.TypeRational {
		return nil
	}
	value, ok := v.Value.([]exifcommon.Rational)
	if !ok || len(value) != 1 {
		return nil
	}
	return &ExifResultItem{Value: fmt.Sprintf("%d", value[0].Numerator/value[0].Denominator), SuffixID: locales.LIdLengthMM}
}

func prepareFlash(v exif.ExifTag) *ExifResultItem {
	tv := getShort(v)
	if tv == nil {
		return nil
	}
	switch *tv {
	case 0x00:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x00}
	case 0x01:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x01}
	case 0x05:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x05}
	case 0x07:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x07}
	case 0x09:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x09}
	case 0x0d:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x0d}
	case 0x0f:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x0f}
	case 0x10:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x10}
	case 0x18:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x18}
	case 0x19:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x19}
	case 0x1d:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x1d}
	case 0x1f:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x1f}
	case 0x20:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x20}
	case 0x41:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x41}
	case 0x45:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x45}
	case 0x47:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x47}
	case 0x49:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x49}
	case 0x4d:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x4d}
	case 0x4f:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x4f}
	case 0x59:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x59}
	case 0x5d:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x5d}
	case 0x5f:
		return &ExifResultItem{ValueID: locales.LIdExifFlash0x5f}
	default:
		return &ExifResultItem{ValueID: locales.LIdUnknownValue}
	}
}

func prepareColorSpace(v exif.ExifTag) *ExifResultItem {
	tv := getShort(v)
	if tv == nil {
		return nil
	}
	switch *tv {
	case 0x01:
		return &ExifResultItem{ValueID: locales.LIdExifColorSpaceSRGB}
	case 0xffff:
		return &ExifResultItem{ValueID: locales.LIdExifColorSpaceUncalibrated}
	default:
		return &ExifResultItem{ValueID: locales.LIdUnknownValue}
	}
}

func prepareWhiteBalance(v exif.ExifTag) *ExifResultItem {
	tv := getShort(v)
	if tv == nil {
		return nil
	}
	switch *tv {
	case 0x00:
		return &ExifResultItem{ValueID: locales.LIdExifWhiteBalanceAuto}
	case 0x01:
		return &ExifResultItem{ValueID: locales.LIdExifWhiteBalanceManual}
	default:
		return &ExifResultItem{ValueID: locales.LIdUnknownValue}
	}
}

func prepareMake(v exif.ExifTag) *ExifResultItem {
	return &ExifResultItem{Value: cases.Title(language.English, cases.Compact).String(v.FormattedFirst)}
}

func prepareDigitalZoomRatio(v exif.ExifTag) *ExifResultItem {
	if v.TagTypeId != exifcommon.TypeRational {
		return nil
	}
	value, ok := v.Value.([]exifcommon.Rational)
	if !ok || len(value) != 1 {
		return nil
	}
	return &ExifResultItem{Value: fmt.Sprintf("×%.1f", float64(value[0].Numerator)/float64(value[0].Denominator))}
}
