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

// ApiDomain ApiDomain
const ApiDomain = "https://api.mch.weixin.qq.com/"

// 基础类接口
const (
	ApiCertification = "/v3/certificates"          // 平台证书下载
	ApiUploadImage   = "/v3/merchant/media/upload" //图片上传
)

// 服务商特约商户进件类接口
const (
	APIIncomingSubmitApplication          = "/v3/applyment4sub/applyment/"                              //提交进件申请单
	APIModifySettlement                   = "/v3/apply4sub/sub_merchants/{sub_mchid}/modify-settlement" //修改结算账号
	APIQuerySettlementAccount             = "/v3/apply4sub/sub_merchants/{sub_mchid}/settlement"        //查询结算账户
	APIGetStatusRepairOrderByBusinessCode = "/v3/applyment4sub/applyment/business_code/{business_code}" //通过业务申请编号查询申请状态
	APIGetStatusRepairOrderByApplymentId  = "/v3/applyment4sub/applyment/applyment_id/{applyment_id}"   //通过申请单号查询申请状态
)

const (
	APIPaymentPartnerQueryOrderByTransactionId = "/v3/pay/partner/transactions/id/{transaction_id}"               //服务商通过微信订单号查询订单
	APIPaymentQueryOrderByTransactionId        = "/v3/pay/transactions/id/{transaction_id}?mchid={mchid}"         //直连商户通过微信订单号查询订单
	APIPaymentPartnerQueryOrderByOutTradeNo    = "/v3/pay/partner/transactions/out-trade-no/{out-trade-no}"       //服务商通过微信订单号查询订单
	APIPaymentQueryOrderByOutTradeNo           = "/v3/pay/transactions/out-trade-no/{out-trade-no}?mchid={mchid}" //直连商户通过微信订单号查询订单
	APIPaymentPartnerCloseOrder                = "/v3/pay/partner/transactions/out-trade-no/{out_trade_no}/close" //服务商关闭订单
	APIPaymentCloseOrder                       = "/v3/pay/transactions/out-trade-no/{out_trade_no}/close"         //直连商户关闭订单
	APIPaymentRefund                           = "/v3/refund/domestic/refunds"                                    //基础支付退款
	APIPaymentQueryRefund                      = "/v3/refund/domestic/refunds/{out_refund_no}"                    //退款查询
	APIApplyTransactionBill                    = "/v3/bill/tradebill"                                             //申请交易账单
	APIApplyFundBill                           = "/v3/bill/fundflowbill"                                          //申请资金账单
	APIApplyProfitSharingBill                  = "/v3/profitsharing/bills"                                        //申请分账账单
	APIJSAPIOrdersForPartner                   = "/v3/pay/partner/transactions/jsapi"                             //服务商JSAPI下单
	APIJSAPIOrders                             = "/v3/pay/transactions/jsapi"                                     //直连商户JSAPI下单
)

// 教培续费通相关接口
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

// K12离线团餐类接口
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

const (
	APIInitiateProfitSharing             = "/v3/profitsharing/orders"                                                                          //请求分账
	APIQueryProfitSharingResult          = "/v3/profitsharing/orders/{out_order_no}?sub_mchid={sub_mchid}&transaction_id={transaction_id}"     //查询分账结果
	APIInitiateProfitSharingReturnOrders = "/v3/profitsharing/return-orders"                                                                   //请求分账回退
	APIQueryProfitSharingReturnOrders    = "/v3/profitsharing/return-orders/{out_return_no}?sub_mchid={sub_mchid}&out_order_no={out_order_no}" //查询分账回退结果
	APIUnfreezeRemainingFunds            = "/v3/profitsharing/orders/unfreeze"                                                                 //解冻剩余资金
	APIQueryRemainingFrozenAmount        = "/v3/profitsharing/transactions/{transaction_id}/amounts"                                           //查询剩余待分金额
	APIQueryMaximumSplitRatio            = "/v3/profitsharing/merchant-configs/{sub_mchid}"                                                    //查询查询最大分账比例
	APIAddProfitSharingReceiver          = "/v3/profitsharing/receivers/add"                                                                   //添加分账接收方
	APIDeleteProfitSharingReceiver       = "/v3/profitsharing/receivers/delete"                                                                //删除分账接收方
)

const (
	APISmartGuideRegister = "/v3/smartguide/guides"                   //服务人员注册
	APISmartGuideAssign   = "/v3/smartguide/guides/{guide_id}/assign" //服务人员分配
	APISmartGuideQuery    = "/v3/smartguide/guides"                   //服务人员查询
	APISmartGuideUpdate   = "/v3/smartguide/guides/{guide_id}"        //服务人员信息更新
)

const (
	APIEduSchoolPayPreSign                   = "/v3/eduschoolpay/contracts/presign"                        //校园轻松付预签约
	APIEduSchoolPayContractQueryById         = "/v3/eduschoolpay/contracts/{contract_id}"                  //通过协议号查询签约
	APIDissolveEduSchoolPayContract          = "/v3/eduschoolpay/contracts/{contract_id}/terminate"        //解约
	APIEduSchoolPayContractQueryByOpenId     = "/v3/eduschoolpay/users/{openid}/contracts"                 //查询用户签约列表
	APIEduSchoolPayTransactions              = "/v3/eduschoolpay/transactions"                             //扣款
	APIEduSchoolPayQueryOrderByTransactionId = "/v3/eduschoolpay/transactions/id/{transaction_id}"         //使用微信单号查单
	APIEduSchoolPayQueryOrderByOutTradeNo    = "/v3/eduschoolpay/transactions/out-trade-no/{out_trade_no}" //使用商户订单号查单
)

const APIViolationNotifications = "/v3/merchant-risk-manage/violation-notifications" //商户违规通知

const (
	APIComplaintsList = "/v3/merchant-service/complaints-v2" //申请交易账单
)
