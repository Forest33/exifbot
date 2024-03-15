package locales

import "github.com/nicksnyder/go-i18n/v2/i18n"

var RU = []*i18n.Message{
	{
		ID: LIdMessageStart,
		Other: `Привет, я EXIF-бот! Я могу извлечь данные EXIF из изображений. Просто пришлите мне изображение или его URL.
Для получения данных EXIF в браузере Google Chrome, можно воспользоваться расширением <a href="https://chromewebstore.google.com/detail/exifer/laolhpdfddlcnbooiiknokijildkcdgo">Exifer</a>.`,
	},
	{
		ID: LIdMessageHelp,
		Other: `Привет, я EXIF-бот! Я могу извлечь данные EXIF из изображений. Просто пришлите мне изображение или его URL.
Для получения данных EXIF в браузере Google Chrome, можно воспользоваться расширением <a href="https://chromewebstore.google.com/detail/exifer/laolhpdfddlcnbooiiknokijildkcdgo">Exifer</a>.`,
	},
	{
		ID:    LIdMessageUnknownCommand,
		Other: "Неверная команда. Введите /help, чтобы получить список команд.",
	},
	{
		ID:    LIdMessageInternalError,
		Other: "Упс! Что-то пошло не так. Пожалуйста, повторите попытку позже.",
	},
	{
		ID:    LIdMessageWrongURL,
		Other: "Некорректный URL изображения",
	},
	{
		ID:    LIdMessageNoExif,
		Other: "Изображение не содержит данных EXIF",
	},
	{
		ID:    LIdMessageErrLoadURL,
		Other: "Не удалось загрузить изображение",
	},

	{
		ID:    LIdUnknownValue,
		Other: "н/д",
	},
	{
		ID:    LIdLengthMM,
		Other: "мм",
	},
	{
		ID:    LIdShowGPS,
		Other: "Показать место съемки",
	},

	{
		ID:    LIdExifTagDateTimeOriginal,
		Other: "Дата съемки",
	},
	{
		ID:    LIdExifTagExposureTime,
		Other: "Экспозиция",
	},
	{
		ID:    LIdExifTagExposureBiasValue,
		Other: "Компенсация экспозиции",
	},
	{
		ID:    LIdExifTagExposureMode,
		Other: "Режим экспозиции",
	},
	{
		ID:    LIdExifTagApertureValue,
		Other: "Диафрагма",
	},
	{
		ID:    LIdExifTagExposureProgram,
		Other: "Режим",
	},
	{
		ID:    LIdExifTagISOSpeedRatings,
		Other: "Величина ISO",
	},
	{
		ID:    LIdExifTagFlash,
		Other: "Вспышка",
	},
	{
		ID:    LIdExifTagFocalLength,
		Other: "Фокусное расстояние",
	},
	{
		ID:    LIdExifTagDigitalZoomRatio,
		Other: "Цифровое масштабирование",
	},
	{
		ID:    LIdExifTagMake,
		Other: "Изготовитель камеры",
	},
	{
		ID:    LIdExifTagModel,
		Other: "Модель камеры",
	},
	{
		ID:    LIdExifTagLensModel,
		Other: "Модель объектива",
	},
	{
		ID:    LIdExifTagBodySerialNumber,
		Other: "Серийный номер камеры",
	},
	{
		ID:    LIdExifTagLensSerialNumber,
		Other: "Серийный номер объектива",
	},
	{
		ID:    LIdExifTagMeteringMode,
		Other: "Режим замера экспозиции",
	},
	{
		ID:    LIdExifTagWhiteBalance,
		Other: "Баланс белого",
	},
	{
		ID:    LIdExifTagSoftware,
		Other: "Программное обеспечение",
	},
	{
		ID:    LIdExifTagColorSpace,
		Other: "Цветовое пространство",
	},
	{
		ID:    LIdExifTagArtist,
		Other: "Автор",
	},
	{
		ID:    LIdExifTagCopyright,
		Other: "Авторские права",
	},
	{
		ID:    LIdExifTagDimensions,
		Other: "Размер изображения",
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
