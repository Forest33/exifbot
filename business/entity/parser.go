package entity

type ParserJob struct {
	URL          string
	CallbackData interface{}
	CallbackChan chan *ParserJobResult
}

type ParserJobResult struct {
	URL          string
	Exif         map[string]*ExifResultItem
	Err          error
	CallbackData interface{}
}
