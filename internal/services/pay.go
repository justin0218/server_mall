package services

import (
	"server_mall/pkg/wechat"
	"time"
)

type PayService struct {
}

func (s *PayService) Pay(openid string, outTradeNo string, body string, spbillCreateIp string, totalFee int, notifyUrl string, tradeType string) (ret wechat.JsApiPayRet, err error) {
	payRet, _, e := wechat.DoPay(openid, outTradeNo, body, spbillCreateIp, totalFee, notifyUrl, tradeType)
	if e != nil {
		err = e
		return
	}
	ret.Timestamp = time.Now().Unix()
	ret.NonceStr = payRet.NonceStr
	ret.Package = "prepay_id=" + payRet.PrepayId
	ret.SignType = "MD5"
	ret.PaySign = payRet.Sign
	return
}
