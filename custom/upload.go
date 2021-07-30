package custom

//ReqUploadImage ReqUploadImage
type ReqUploadImage struct {
	Filename string `json:"filename"`
	Sha256   string `json:"sha256"`
}

//RespUploadImage RespUploadImage
type RespUploadImage struct {
	MediaId string `json:"media_id"`
}
