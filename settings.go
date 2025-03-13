package wxpayv3

import (
	"fmt"
	"github.com/louismax/wxpayv3/core"
	"github.com/louismax/wxpayv3/custom"
)

// InjectWxPayMchParam 注入微信支付商户参数(商户号, 商户APIv3密钥, 商户API证书序列号, 商户私钥文件路径)
func InjectWxPayMchParam(mchID, apiV3Key, apiSerialNo, pvtKeyFilePath string) core.ClientOption {
	pvt, err := core.LoadPrivateKeyWithPath(pvtKeyFilePath)
	if err != nil {
		fmt.Printf("\033[31m%s\n", "[Error]--WxPayV3:通过私钥的文件路径加载商户私钥失败！"+err.Error())
		return core.ErrorOption{Error: err}
	}
	return core.BasicInformation{
		MchID:       mchID,
		MchAPIv3Key: apiV3Key,
		ApiCert: core.ApiCert{
			ApiSerialNo:   apiSerialNo,
			ApiPrivateKey: pvt,
		},
	}
}

// InjectWxPayMchParamFull 注入微信支付商户参数(商户号, 商户证书(需自己解析), 商户APIv3密钥)
func InjectWxPayMchParamFull(mchID string, apiCert core.ApiCert, apiV3Key string) core.ClientOption {
	return core.BasicInformation{
		MchID:       mchID,
		MchAPIv3Key: apiV3Key,
		ApiCert:     apiCert,
	}
}

// InjectWxPayMchParamExtra 注入微信支付商户参数(商户号, 商户APIv3密钥, 商户私钥文件路径, 商户证书文件路径)
func InjectWxPayMchParamExtra(mchID string, apiV3Key string, privateKeyPath string, certificatePath string) core.ClientOption {
	pvt, err := core.LoadPrivateKeyWithPath(privateKeyPath)
	if err != nil {
		fmt.Printf("\033[31m%s\n", "[Error]--WxPayV3:通过私钥的文件路径加载商户私钥失败！"+err.Error())
		return core.ErrorOption{Error: err}
	}
	cert, err := core.LoadCertificateWithPath(certificatePath)
	if err != nil {
		fmt.Printf("\033[31m%s\n", "[Error]--WxPayV3:通过证书的文件路径加载商户商户失败！"+err.Error())
		return core.ErrorOption{Error: err}
	}
	return core.BasicInformation{
		MchID:       mchID,
		MchAPIv3Key: apiV3Key,
		ApiCert: core.ApiCert{
			ApiSerialNo:    core.GetCertificateSerialNumber(*cert),
			ApiPrivateKey:  pvt,
			ApiCertificate: cert,
		},
	}
}

// InjectWxPayPlatformCert 注入微信支付平台证书
func InjectWxPayPlatformCert(certificateStr []string) core.ClientOption {
	defNo := ""
	list := make([]custom.CertificateDataList, 0)
	for _, v := range certificateStr {
		ct, err := core.LoadCertificate(v)
		if err != nil {
			fmt.Printf("\033[31m%s\n", "[Error]--WxPayV3:通过平台证书的文本内容加载微信支付平台证书失败！"+err.Error())
			return core.ErrorOption{Error: err}
		}
		if defNo == "" {
			defNo = core.GetCertificateSerialNumber(*ct)
		}
		list = append(list, custom.CertificateDataList{
			SerialNo:    core.GetCertificateSerialNumber(*ct),
			Certificate: ct,
		})
	}

	return core.PlatformCert{
		DefaultSerialNo: defNo,
		CertList:        list,
	}
}

// InjectWxPayPlatformCertUseCertPath 注入微信支付平台证书(本地平台证书文件路径)
func InjectWxPayPlatformCertUseCertPath(platformCertificatePath []string) core.ClientOption {
	list := make([]custom.CertificateDataList, 0)
	defNo := ""
	for _, v := range platformCertificatePath {
		ct, err := core.LoadCertificateWithPath(v)
		if err != nil {
			fmt.Printf("\033[31m%s\n", "[Error]--WxPayV3:通过证书的文件路径加载微信支付平台证书失败！"+err.Error())
			return core.ErrorOption{Error: err}
		}
		if defNo == "" {
			defNo = core.GetCertificateSerialNumber(*ct)
		}
		list = append(list, custom.CertificateDataList{
			SerialNo:    core.GetCertificateSerialNumber(*ct),
			Certificate: ct,
		})
	}
	return core.PlatformCert{
		DefaultSerialNo: defNo,
		CertList:        list,
	}
}

// InjectWxPayPlatformPubKey 注入微信支付平台公钥
func InjectWxPayPlatformPubKey(pubKeyId, pubKey string) core.ClientOption {
	pub, err := core.LoadPublicKey(pubKey)
	if err != nil {
		fmt.Printf("\033[31m%s\n", "[Error]--WxPayV3:通过公钥文本内容加载微信支付平台公钥失败！"+err.Error())
		return core.ErrorOption{Error: err}
	}
	return core.PlatformPubKey{
		PubKeyId: pubKeyId,
		PubKey:   pub,
	}
}

func InjectWxPayPlatformPubKeyUsePath(pubKeyId, pubKeyPath string) core.ClientOption {
	pub, err := core.LoadPublicKeyWithPath(pubKeyPath)
	if err != nil {
		fmt.Printf("\033[31m%s\n", "[Error]--WxPayV3:通过公钥文件路径加载微信支付平台公钥失败！"+err.Error())
		return core.ErrorOption{Error: err}
	}
	return core.PlatformPubKey{
		PubKeyId: pubKeyId,
		PubKey:   pub,
	}
}
