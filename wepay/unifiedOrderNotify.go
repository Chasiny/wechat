package WePay

type UnifiedOrderNotify struct {
	AppID              CDATA `xml:"appid"`                //公众账号ID
	Attach             CDATA `xml:"attach"`               //商家数据包
	BankType           CDATA `xml:"bank_type"`            //付款银行
	CashFee            CDATA `xml:"cash_fee"`             //现金支付金额
	DeviceInfo         CDATA `xml:"device_info"`          //设备号
	FeeType            CDATA `xml:"fee_type"`             //货币种类
	IsSubscribe        CDATA `xml:"is_subscribe"`         //是否关注公众账号
	MchID              CDATA `xml:"mch_id"`               //商户号
	NonceStr           CDATA `xml:"nonce_str"`            //随机字符串
	OpenID             CDATA `xml:"openid"`               //用户标识
	OutTradeNO         CDATA `xml:"out_trade_no"`         //商户订单号
	ResultCode         CDATA `xml:"result_code"`          //业务结果
	ReturnCode         CDATA `xml:"return_code"`          //返回错误码
	Sign               CDATA `xml:"sign"`                 //签名
	TimeEnd            CDATA `xml:"time_end"`             //支付完成时间
	TotalFee           int   `xml:"total_fee"`            //订单金额
	TradeType          CDATA `xml:"trade_type"`           //交易类型
	TransactionID      CDATA `xml:"transaction_id"`       //微信支付订单号
	ReturnMsg          CDATA `xml:"return_msg"`           //返回信息
	SignType           CDATA `xml:"sign_type"`            //签名类型
	ErrCode            CDATA `xml:"err_code"`             //错误代码
	ErrCodeDes         CDATA `xml:"err_code_des"`         //错误代码描述
	SettlementTotalFee CDATA `xml:"settlement_total_fee"` //应结订单金额
	CashFeeType        CDATA `xml:"cash_fee_type"`        //现金支付货币类型
}

type NotifyRespon struct {
	ReturnCode CDATA `xml:"return_code"` //返回错误码
	ReturnMsg  CDATA `xml:"return_msg"`  //返回信息
}
