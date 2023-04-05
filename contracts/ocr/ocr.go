package ocr

type Ocr interface {
	OcrImg(base string) (data ReturnData, err error)
	GetIOcrImg(templateId string, base string) (data ReturnIOcrData, err error)
}

type ReturnData struct {
	Log_id           int64
	Words_result     []map[string]string
	Words_result_num int
}

type ReturnIOcrData struct {
	Logid      int64
	Data       interface{}
	Error_code int
	Error_msg  string
}
