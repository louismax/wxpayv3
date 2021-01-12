package wxpayv3

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"strings"
)

type ReqUploadImage struct {
	Filename string `json:"filename"`
	Sha256   string `json:"sha256"`
}

// APIUrl ReqUploadImage APIURL
func (this ReqUploadImage) APIUrl() string {
	return "/v3/merchant/media/upload"
}

// Method ReqUploadImage Method
func (this ReqUploadImage) Method() string {
	return "POST"
}

// Params ReqUploadImage Params
func (this ReqUploadImage) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

// RawJsonStr ReqUploadImage RawJsonStr
func (this ReqUploadImage) RawJsonStr() string {
	jsons, errs := json.Marshal(this)
	if errs != nil {
		fmt.Println(errs.Error())
		return "Error"
	}
	return string(jsons)
}

type RespUploadImage struct {
	MediaId string `json:"media_id"`
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func CreateFormFile(fieldname, filename, contentType string, w *multipart.Writer) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))
	h.Set("Content-Type", contentType)
	return w.CreatePart(h)
}

func CreateForm(key, contentType string, w *multipart.Writer) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s";`,
			escapeQuotes(key)))
	h.Set("Content-Type", contentType)
	return w.CreatePart(h)
}
