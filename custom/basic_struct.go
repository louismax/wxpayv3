package custom

import "crypto/x509"

// CertificateResp CertificateResp
type CertificateResp struct {
	Data []*CertificateData `json:"data"`
}

// EncryptCertificate encryptCertificate
type EncryptCertificate struct {
	Algorithm      string `json:"algorithm"`
	Nonce          string `json:"nonce"`
	AssociatedData string `json:"associated_data"`
	Ciphertext     string `json:"ciphertext"`
}

// CertificateData CertificateData
type CertificateData struct {
	EncryptCertificate *EncryptCertificate `json:"encrypt_certificate"`
	DecryptCertificate string              `json:"decrypt_certificate"`
	SerialNo           string              `json:"serial_no"`
	EffectiveTime      string              `json:"effective_time"`
	ExpireTime         string              `json:"expire_time"`
}

type CertificateDataList struct {
	SerialNo    string
	Certificate *x509.Certificate
}
