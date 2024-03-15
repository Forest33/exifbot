package locales

import "github.com/nicksnyder/go-i18n/v2/i18n"

var EN = []*i18n.Message{
	{
		ID: LIdMessageStart,
		Other: `Hello, I'm a EXIF bot! I can extract EXIF data from images. Just send me the image or URL to it.
To get EXIF data in the Google Chrome browser, you can use the extension <a href="https://chromewebstore.google.com/detail/exifer/laolhpdfddlcnbooiiknokijildkcdgo">Exifer</a>`,
	},
	{
		ID: LIdMessageHelp,
		Other: `Hello, I'm a EXIF bot! I can extract EXIF data from images. Just send me the image or URL to it.
To get EXIF data in the Google Chrome browser, you can use the extension <a href="https://chromewebstore.google.com/detail/exifer/laolhpdfddlcnbooiiknokijildkcdgo">Exifer</a>`,
	},
	{
		ID:    LIdMessageUnknownCommand,
		Other: "Invalid command, enter /help to get a list of commands.",
	},
	{
		ID:    LIdMessageInternalError,
		Other: "Oops! something went wrong, please try again later.",
	},
	{
		ID:    LIdMessageWrongURL,
		Other: "Invalid image URL",
	},
	{
		ID:    LIdMessageNoExif,
		Other: "Image does not contain EXIF data",
	},
	{
		ID:    LIdMessageErrLoadURL,
		Other: "Failed to load image",
	},

	{
		ID:    LIdUnknownValue,
		Other: "Unknown",
	},
	{
		ID:    LIdLengthMM,
		Other: "mm",
	},
	{
		ID:    LIdShowGPS,
		Other: "Show shooting location",
	},

	{
		ID:    LIdExifTagDateTimeOriginal,
		Other: "Date",
	},
	{
		ID:    LIdExifTagExposureTime,
		Other: "Exposure",
	},
	{
		ID:    LIdExifTagExposureBiasValue,
		Other: "Exposure Bias Value",
	},
	{
		ID:    LIdExifTagExposureMode,
		Other: "Exposure Mode",
	},
	{
		ID:    LIdExifTagApertureValue,
		Other: "Aperture Value",
	},
	{
		ID:    LIdExifTagExposureProgram,
		Other: "Mode",
	},
	{
		ID:    LIdExifTagISOSpeedRatings,
		Other: "ISO Value",
	},
	{
		ID:    LIdExifTagFlash,
		Other: "Flash",
	},
	{
		ID:    LIdExifTagFocalLength,
		Other: "Focal Length",
	},
	{
		ID:    LIdExifTagDigitalZoomRatio,
		Other: "Digital Zoom",
	},
	{
		ID:    LIdExifTagMake,
		Other: "Make",
	},
	{
		ID:    LIdExifTagModel,
		Other: "Model",
	},
	{
		ID:    LIdExifTagLensModel,
		Other: "Lens Model",
	},
	{
		ID:    LIdExifTagBodySerialNumber,
		Other: "Body Serial Number",
	},
	{
		ID:    LIdExifTagLensSerialNumber,
		Other: "Lens Serial Number",
	},
	{
		ID:    LIdExifTagMeteringMode,
		Other: "Metering Mode",
	},
	{
		ID:    LIdExifTagWhiteBalance,
		Other: "White Balance",
	},
	{
		ID:    LIdExifTagSoftware,
		Other: "Software",
	},
	{
		ID:    LIdExifTagColorSpace,
		Other: "Color Space",
	},
	{
		ID:    LIdExifTagArtist,
		Other: "Artist",
	},
	{
		ID:    LIdExifTagCopyright,
		Other: "Copyright",
	},
	{
		ID:    LIdExifTagDimensions,
		Other: "Dimensions",
	},

	{
		ID:    LIdExifExposureModeAuto,
		Other: "Auto exposure",
	},
	{
		ID:    LIdExifExposureModeManual,
		Other: "Manual exposure",
	},
	{
		ID:    LIdExifExposureModeAutoBracket,
		Other: "Auto bracket",
	},

	{
		ID:    LIdExifExposureProgramUndefined,
		Other: "Undefined",
	},
	{
		ID:    LIdExifExposureProgramManual,
		Other: "Manual",
	},
	{
		ID:    LIdExifExposureProgramNormalProgram,
		Other: "Normal program",
	},
	{
		ID:    LIdExifExposureProgramAperturePriority,
		Other: "Aperture priority",
	},
	{
		ID:    LIdExifExposureProgramShutterPriority,
		Other: "Shutter priority",
	},
	{
		ID:    LIdExifExposureProgramCreativeProgram,
		Other: "Creative program",
	},
	{
		ID:    LIdExifExposureProgramActionProgram,
		Other: "Action program",
	},
	{
		ID:    LIdExifExposureProgramPortraitMode,
		Other: "Portrait mode",
	},
	{
		ID:    LIdExifExposureProgramLandscapeMode,
		Other: "Landscape mode",
	},
	{
		ID:    LIdExifExposureProgramBulb,
		Other: "Bulb",
	},

	{
		ID:    LIdExifMeteringModeAverage,
		Other: "Average",
	},
	{
		ID:    LIdExifMeteringModeCenterWeightedAverage,
		Other: "Center weighted average",
	},
	{
		ID:    LIdExifMeteringModeSpot,
		Other: "Spot",
	},
	{
		ID:    LIdExifMeteringModeMultiSpot,
		Other: "Multi spot",
	},
	{
		ID:    LIdExifMeteringModePattern,
		Other: "Pattern",
	},
	{
		ID:    LIdExifMeteringModePartial,
		Other: "Partial",
	},
	{
		ID:    LIdExifMeteringModeOther,
		Other: "Other",
	},

	{
		ID:    LIdExifFlash0x00,
		Other: "Flash did not fire",
	},
	{
		ID:    LIdExifFlash0x01,
		Other: "Flash fired",
	},
	{
		ID:    LIdExifFlash0x05,
		Other: "Strobe return light not detected",
	},
	{
		ID:    LIdExifFlash0x07,
		Other: "Strobe return light detected",
	},
	{
		ID:    LIdExifFlash0x09,
		Other: "Flash fired, compulsory flash mode",
	},
	{
		ID:    LIdExifFlash0x0d,
		Other: "Flash fired, compulsory flash mode, return light not detected",
	},
	{
		ID:    LIdExifFlash0x0f,
		Other: "Flash fired, compulsory flash mode, return light detected",
	},
	{
		ID:    LIdExifFlash0x10,
		Other: "Flash did not fire, compulsory flash mode",
	},
	{
		ID:    LIdExifFlash0x18,
		Other: "Flash did not fire, auto mode",
	},
	{
		ID:    LIdExifFlash0x19,
		Other: "Flash fired, auto mode",
	},
	{
		ID:    LIdExifFlash0x1d,
		Other: "Flash fired, auto mode, return light not detected",
	},
	{
		ID:    LIdExifFlash0x1f,
		Other: "Flash fired, auto mode, return light detected",
	},
	{
		ID:    LIdExifFlash0x20,
		Other: "No flash function",
	},
	{
		ID:    LIdExifFlash0x41,
		Other: "Flash fired, red-eye reduction mode",
	},
	{
		ID:    LIdExifFlash0x45,
		Other: "Flash fired, red-eye reduction mode, return light not detected",
	},
	{
		ID:    LIdExifFlash0x47,
		Other: "Flash fired, red-eye reduction mode, return light detected",
	},
	{
		ID:    LIdExifFlash0x49,
		Other: "Flash fired, compulsory flash mode, red-eye reduction mode",
	},
	{
		ID:    LIdExifFlash0x4d,
		Other: "Flash fired, compulsory flash mode, red-eye reduction mode, return light not detected",
	},
	{
		ID:    LIdExifFlash0x4f,
		Other: "Flash fired, compulsory flash mode, red-eye reduction mode, return light detected",
	},
	{
		ID:    LIdExifFlash0x59,
		Other: "Flash fired, auto mode, red-eye reduction mode",
	},
	{
		ID:    LIdExifFlash0x5d,
		Other: "Flash fired, auto mode, return light not detected, red-eye reduction mode",
	},
	{
		ID:    LIdExifFlash0x5f,
		Other: "Flash fired, auto mode, return light detected, red-eye reduction mode",
	},

	{
		ID:    LIdExifColorSpaceSRGB,
		Other: "sRGB",
	},
	{
		ID:    LIdExifColorSpaceUncalibrated,
		Other: "Uncalibrated",
	},

	{
		ID:    LIdExifWhiteBalanceAuto,
		Other: "Auto white balance",
	},
	{
		ID:    LIdExifWhiteBalanceManual,
		Other: "Manual white balance",
	},
}
