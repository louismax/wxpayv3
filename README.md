# wxpayv3 (wechatpay-api-v3)
[简体中文](README.md)
## 微信支付API V3 SDK for GO

[![Go Report Card](https://goreportcard.com/badge/github.com/louismax/wxpayv3)](https://goreportcard.com/report/github.com/louismax/wxpayv3)
[![GoDoc](https://godoc.org/github.com/louismax/wxpayv3?status.svg)](https://godoc.org/github.com/louismax/wxpayv3)
[![GitHub release](https://img.shields.io/github/tag/louismax/wxpayv3.svg)](https://github.com/louismax/wxpayv3/releases)
[![GitHub license](https://img.shields.io/github/license/louismax/wxpayv3.svg)](https://github.com/louismax/wxpayv3/blob/master/LICENSE)
[![GitHub Repo Size](https://img.shields.io/github/repo-size/louismax/wxpayv3.svg)](https://img.shields.io/github/repo-size/louismax/wxpayv3.svg)
[![GitHub Last Commit](https://img.shields.io/github/last-commit/louismax/wxpayv3.svg)](https://img.shields.io/github/last-commit/louismax/wxpayv3.svg)

## 安装
`go get -v github.com/louismax/wxpayv3`

## 实现能力（持续更新）
### 基础
- 获取平台证书  ✔️
- 上传图片获取MediaId  ✔️
- 上传视频获取MediaId  ❌

### 商户进件(仅支持服务商,支持小微商户进件、特约商户进件)
- 提交申请单  ✔️
- 查询申请单状态  ✔️
- 修改结算账号  ✔️
- 查询结算账号  ✔️

### 基础支付(JSAPI支付、小程序支付，支持服务商版、普通商户版)
- JSAPI下单  ❌
- 查询订单  ❌
- 关闭订单  ❌
- 申请退款  ✔️
- 查询退款  ❌
- 申请交易账单(特约商户、普通商户)  ✔️
- 申请资金账单  ✔️
- 下载账单  ✔️

### 分账(支持服务商版)
- 请求分账  ✔️
- 查询分账结果  ✔️
- 请求分账回退  ✔️
- 查询分账回退结果  ✔️
- 解冻剩余资金  ✔️
- 查询剩余待分金额  ✔️
- 查询最大分账比例  ✔️
- 添加分账接收方  ✔️
- 删除分账接收方  ✔️
- 申请分账账单  ✔️
- 下载账单  ✔️

### 教培续费通(支持服务商版、普通商户版)
- 预签约  ✔️
- 协议号查询签约  ✔️
- 用户标识查询签约  ✔️
- 解约  ✔️
- 发送预扣款通知  ✔️
- 扣款受理  ✔️
- 微信订单号查单  ✔️
- 商户订单号查单  ✔️

### 校园刷脸代扣(仅支持服务商)
- 获取机构信息(根据机构ID)  ✔️
- 获取机构信息(根据机构名称)  ✔️
- 获取授权凭证  ✔️
- 查询刷脸用户信息  ✔️
- 修改刷脸用户信息  ✔️
- 解除刷脸用户签约关系  ✔️
- 预签约  ✔️
- 申请扣款  ✔️
- 签约查询  ✔️
- 人脸报文(签约解约)消息解密  ✔️
- 查询重采用户列表  ✔️
- 查询重采  ✔️
- 离线人脸团餐专属查单  ✔️
- 获取AuthInfo  ✔️
- 获取还款链接  ✔️







## 参考资料
* [微信支付商户平台API文档](https://pay.weixin.qq.com/wiki/doc/apiv3/index.shtml)
* [微信支付服务商平台API文档](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/index.shtml)
* [微信支付API V3接口规范](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/wechatpay/wechatpay-1.shtml)
* [微信支付教培续费通](https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/edu-papay/chapter1_1.shtml)

## 协议
MIT 许可证（MIT）。有关更多信息，请参见[协议文件](LICENSE)。

