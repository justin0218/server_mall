package services

import (
	"server_mall/pkg/wechat"
)

type PayService struct {
}

func (s *PayService) Pay(openid string, outTradeNo string, body string, spbillCreateIp string, totalFee int, notifyUrl string, tradeType string) (ret wechat.JsApiPayRet, err error) {
	payRet, _, e := wechat.DoPay(openid, outTradeNo, body, spbillCreateIp, totalFee, notifyUrl, tradeType)
	if e != nil {
		err = e
		return
	}
	jsapi := wechat.GetJsapiSign("prepay_id="+payRet.PrepayId, payRet.NonceStr)
	ret.Timestamp = jsapi.TimeStamp
	ret.NonceStr = jsapi.NonceStr
	ret.Package = jsapi.Package
	ret.SignType = jsapi.SignType
	ret.PaySign = jsapi.PaySign
	return
}
