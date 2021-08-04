package constant

import "time"

const (
	//AlgorithmAEADAES256GCM AlgorithmAEADAES256GCM
	AlgorithmAEADAES256GCM = "AEAD_AES_256_GCM"
)
const (
	//DefaultTimeout HTTP 请求默认超时时间
	DefaultTimeout = 30 * time.Second
)

//ApiDomain ApiDomain
const ApiDomain = "https://api.mch.weixin.qq.com/"

//基础类接口
const (
	ApiCertification = "/v3/certificates"          // 平台证书下载
	ApiUploadImage   = "/v3/merchant/media/upload" //图片上传
)

//服务商特约商户进件类接口
const (
	ApiQuerySettlementAccount             = "/v3/apply4sub/sub_merchants/{sub_mchid}/settlement"        //查询结算账户
	APIGetStatusRepairOrderByBusinessCode = "/v3/applyment4sub/applyment/business_code/{business_code}" //通过业务申请编号查询申请状态
	APIGetStatusRepairOrderByApplymentId  = "/v3/applyment4sub/applyment/applyment_id/{applyment_id}"   //通过申请单号查询申请状态
)

//教培续费通相关接口
const (
	APIEduPaPayPresign                   = "/v3/edu-papay/contracts/presign"                        //预签约
	APIEduPaPayContractQueryById         = "/v3/edu-papay/contracts/id/{contract_id}"               //通过签约ID查询签约
	APIEduPaPayContractQueryByOpenId     = "/v3/edu-papay/user/{openid}/contracts"                  //通过用户标识查询签约
	APIDissolveEduPaPayContract          = "/v3/edu-papay/contracts/{contract_id}"                  //解约
	APISendEduPaPayNotifications         = "/v3/edu-papay/user-notifications/{contract_id}/send"    //发送扣款预通知
	APIEduPaPayTransactions              = "/v3/edu-papay/transactions"                             //教培通扣款受理
	APIEduPaPayQueryOrderByTransactionId = "/v3/edu-papay/transactions/id/{transaction_id}"         //教培通使用微信单号查单
	APIEduPaPayQueryOrderByOutTradeNo    = "/v3/edu-papay/transactions/out-trade-no/{out_trade_no}" //教培通使用商户订单号查单
)

//K12离线团餐类接口
const (
	APIQueryOrganizationInfoById   = "/v3/offlinefacemch/organizations?organization_id={organization_id}"                                                                           //根据机构ID查询机构信息
	APIQueryOrganizationInfoByName = "/v3/offlinefacemch/organizations?organization_name={organization_name}"                                                                       //根据机构名称查询机构信息
	APIObtainAuthToken             = "/v3/offlinefacemch/tokens"                                                                                                                    //获取授权凭证
	APIPayCredential               = "/v3/offlinefacemch/paycredential"                                                                                                             //旧扣款接口(已废弃)
	APIQueryFaceUserInfo           = "/v3/offlinefacemch/organizations/{organization_id}/users/out-user-id/{out_user_id}"                                                           //刷脸用户信息查询
	APIUpdateFaceUserInfo          = "/v3/offlinefacemch/organizations/{organization_id}/users/out-user-id/{out_user_id}"                                                           //刷脸用户信息修改
	APIDissolveFaceUserContract    = "/v3/offlinefacemch/organizations/{organization_id}/users/user-id/{user_id}/terminate-contract"                                                //解除刷脸用户签约关系
	APIPreSignature                = "/v3/offlineface/contracts/presign"                                                                                                            //预签约
	APIOfflinefaceTransactions     = "/v3/offlineface/transactions"                                                                                                                 //申请扣款
	APIContractQuery               = "/v3/offlineface/contracts/{contract_id}?appid={appid}"                                                                                        //签约查询
	APIQueryRepurchaseUsersList    = "/v3/offlineface/face-collections?organization_id={organization_id}&offset={offset}&limit={limit}"                                             //查询重采用户列表
	APIQueryRetake                 = "/v3/offlineface/face-collections/{collection_id}"                                                                                             //查询重采
	APIQueryOfflineFaceOrders      = "/v3/offlineface/transactions/out-trade-no/{out_trade_no}?sp_mchid={sp_mchid}&sub_mchid={sub_mchid}&business_product_id={business_product_id}" //离线人脸团餐专属查单
	APIGetAuthInfo                 = "/v3/offlineface/authinfo"                                                                                                                     //获取AuthInfo
	APIGetRepaymentUrl             = "/v3/offlineface/repayment-url"                                                                                                                //获取还款链接
)
