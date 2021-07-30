package utils

import (
	"github.com/louismax/wxpayv3/custom"
	"testing"
)

func TestGenerateNonce(t *testing.T) {
	t.Log(GenerateNonce())
}

func TestFaceMessageDecryption(t *testing.T) {
	r, e := FaceMessageDecryption(custom.FaceMessageCiphertext{
		ID:           "",
		CreateTime:   "",
		EventType:    "",
		ResourceType: "",
		Resource: struct {
			Algorithm      string `json:"algorithm"`
			Ciphertext     string `json:"ciphertext"`
			OriginalType   string `json:"original_type"`
			AssociatedData string `json:"associated_data"`
			Nonce          string `json:"nonce"`
		}{},
		Summary: "",
	}, "")
	if e != nil {
		t.Log(e)
		return
	}
	t.Log(r)
}
