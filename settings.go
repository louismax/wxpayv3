package wxpayv3

import (
	"crypto/x509"
	"github.com/louismax/wxpayv3/core"
)

func InjectWechatPayParameter(mchID string, apiCert core.ApiCert, mchAPIv3Key string) core.ClientOption {
	return core.BasicInformation{
		MchID:mchID,
		MchAPIv3Key: mchAPIv3Key,
		ApiCert: apiCert,
	}
}

func InjectWechatPayParameterUseCertPath(mchID string, mchAPIv3Key string,privateKeyPath string,certificatePath string) core.ClientOption {
	pv,err := core.LoadPrivateKeyWithPath(privateKeyPath)
	if err != nil{
		return core.ErrorOption{Error:err}
	}
	ct,err := core.LoadCertificateWithPath(certificatePath)
	if err != nil{
		return core.ErrorOption{Error:err}
	}
	return core.BasicInformation{
		MchID:mchID,
		MchAPIv3Key: mchAPIv3Key,
		ApiCert:core.ApiCert{
			ApiSerialNo:   core.GetCertificateSerialNumber(*ct),
			ApiPrivateKey: pv,
			ApiCertificate: ct,
		},
	}
}

func InjectWechatPayPlatformCert(platformSerialNo  string,platformCertificate *x509.Certificate) core.ClientOption {
	return core.PlatformCert{
		PlatformSerialNo:platformSerialNo,
		PlatformCertificate:platformCertificate,
	}
}

func InjectWechatPayPlatformCertUseCertPath(platformCertificatePath string) core.ClientOption {
	ct,err := core.LoadCertificateWithPath(platformCertificatePath)
	if err != nil{
		return core.ErrorOption{Error:err}
	}

	return core.PlatformCert{
		PlatformSerialNo:core.GetCertificateSerialNumber(*ct),
		PlatformCertificate:ct,
	}
}