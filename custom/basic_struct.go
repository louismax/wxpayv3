package custom

import "time"

//CertificateResp CertificateResp
type CertificateResp struct {
	Data []*CertificateData `json:"data"`
}

//EncryptCertificate EncryptCertificate
type EncryptCertificate struct {
	Algorithm      string `json:"algorithm"`
	Nonce          string `json:"nonce"`
	AssociatedData string `json:"associated_data"`
	Ciphertext     string `json:"ciphertext"`
}

//CertificateData CertificateData
type CertificateData struct {
	EncryptCertificate *EncryptCertificate `json:"encrypt_certificate"`
	DecryptCertificate string              `json:"decrypt_certificate"`
	SerialNo           string              `json:"serial_no"`
	EffectiveTime      time.Time           `json:"effective_time "`
	ExpireTime         time.Time           `json:"expire_time "`
}
