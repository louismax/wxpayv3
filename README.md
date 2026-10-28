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
- JSAPI下单(兼容服务商模式、直连商户模式)  ✔️
- Native下单(服务商模式，`NativeOrdersForPartner`)  ✔️
- Native下单(直连商户模式，`NativeOrders`)  ✔️
- 查询订单-通过微信订单号(兼容服务商模式、直连商户模式)  ✔️
- 查询订单-通过商户订单号(兼容服务商模式、直连商户模式)  ✔️
- 关闭订单(支持JSAPI、Native，`PaymentCloseOrder`)  ✔️
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







### 消费者投诉2.0（服务商及直连商户）
- 投诉列表查询（服务商不传商户号时查询名下全部子商户） `QueryComplaintsList`
- 投诉详情 `QueryComplaintDetail`
- 协商历史及留言 `QueryComplaintNegotiationHistory`
- 文字、图片回复 `ResponseComplaint`
- 即时服务投诉回复（返回 `log_id`） `ResponseComplaintImmediateService`
- 反馈处理完成 `CompleteComplaint`
- 投诉专用图片上传 `UploadComplaintImage`
- 按渠道返回的原始图片 URL 下载 `DownloadComplaintImage`
- 通知原始报文验签及解密 `ParseComplaintNotification`

```go
// client 以服务商商户号、证书私钥及平台公钥等现有选项初始化。
complaints, err := client.QueryComplaintsList("2026-10-01", "2026-10-27", 50, 0)
// 指定子商户时，最后一个参数传被诉商户号。
complaints, err = client.QueryComplaintsList("2026-10-01", "2026-10-27", 50, 0, "子商户号")

err = client.ResponseComplaint("投诉单号", custom.ReqResponseComplaint{
    ComplaintedMchid: "子商户号",
    ResponseContent: "已联系用户核实处理",
})
```

使用说明：
- 所有处理请求继续使用服务商身份签名，`complainted_mchid` 填被诉子商户号。
- 图片先通过投诉专用上传接口取得 `media_id` 再用于回复；普通商户资料上传接口不能替代。
- `DownloadComplaintImage` 使用详情或协商历史返回的微信支付图片 URL，保留原始转义字符；返回图片二进制内容。
- `ParseComplaintNotification` 接收 `http.Header` 和未经重新序列化的原始请求体；必须配置匹配的微信支付公钥或平台证书及服务商 APIv3 密钥。返回通知信封与解密后的投诉单号、动作类型。通知去重、保存及 HTTP 应答由调用方负责。
- 普通回复最多 500 字符、4 张图片；图片支持 JPG、PNG、BMP，最大 2M。
- 回复、即时服务回复、反馈处理完成遇到网络错误不自动重发，避免重复操作；调用方应查询确认处理结果。
- 本次未增加退款审批、退款操作或回调地址配置接口。

官方合作伙伴文档：[投诉业务示例](https://pay.wechatpay.cn/doc/v3/partner/4015933338)、[即时服务回复](https://pay.wechatpay.cn/doc/v3/partner/4017151726)、[投诉通知](https://pay.wechatpay.cn/doc/v3/partner/4012076174)。

## 参考资料
* [微信支付商户平台API文档](https://pay.weixin.qq.com/wiki/doc/apiv3/index.shtml)
* [微信支付服务商平台API文档](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/index.shtml)
* [微信支付API V3接口规范](https://pay.weixin.qq.com/wiki/doc/apiv3_partner/wechatpay/wechatpay-1.shtml)
* [微信支付教培续费通](https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/edu-papay/chapter1_1.shtml)

## 协议
MIT 许可证（MIT）。有关更多信息，请参见[协议文件](LICENSE)。

