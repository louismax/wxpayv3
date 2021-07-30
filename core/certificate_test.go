package core

import (
	"crypto/x509"
	"testing"
	"time"
)

func TestLoadCertificateWithPath(t *testing.T) {
	pv, err := LoadCertificateWithPath("")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(GetCertificateSerialNumber(*pv))
}

func TestLoadPublicKeyWithPath(t *testing.T) {
	pv, err := LoadPublicKeyWithPath("")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(pv)
}

func TestIsCertValid(t *testing.T) {
	t.Log(IsCertValid(x509.Certificate{}, time.Now()))
}
