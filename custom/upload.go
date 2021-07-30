package custom

type ReqUploadImage struct {
	Filename string `json:"filename"`
	Sha256   string `json:"sha256"`
}

type RespUploadImage struct {
	MediaId string `json:"media_id"`
}
