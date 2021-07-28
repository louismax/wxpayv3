package core

import "testing"

func TestLoadCertificateWithPath(t *testing.T) {
	pv, err := LoadCertificateWithPath("../apiclient_cert.pem")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(GetCertificateSerialNumber(*pv))
}
