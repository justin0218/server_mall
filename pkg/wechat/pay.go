package wechat

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"reflect"
	"server_mall/pkg/tool"
	"sort"
	"strings"
)

type JsApiPayRet struct {
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Package   string `json:"package"`
	SignType  string `json:"sign_type"`
	PaySign   string `json:"pay_sign"`
}

type PayRes struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	PrepayId   string `xml:"prepay_id"`
	Sign       string `xml:"sign"`
	NonceStr   string `xml:"nonce_str"`
}

type Pay struct {
	Appid          string `xml:"appid"`
	MchId          string `xml:"mch_id"`
	Openid         string `xml:"openid"`
	NonceStr       string `xml:"nonce_str"`
	OutTradeNo     string `xml:"out_trade_no"`
	Sign           string `xml:"sign"`
	Body           string `xml:"body"`
	TotalFee       int    `xml:"total_fee"`
	SpbillCreateIp string `xml:"spbill_create_ip"`
	NotifyUrl      string `xml:"notify_url"`
	TradeType      string `xml:"trade_type"`
}

func getWxPaySign(req interface{}) string {
	objT := reflect.TypeOf(req)
	objV := reflect.ValueOf(req)
	keyArr := make([]string, 0)
	valMap := make(map[string]interface{})
	for i := 0; i < objT.NumField(); i++ {
		filed := objT.Field(i).Tag.Get("xml")
		val := objV.Field(i).Interface()
		if filed != "" {
			keyArr = append(keyArr, filed)
			valMap[filed] = val
		}
	}
	sort.Strings(keyArr)
	urlValues := []string{}
	for _, val := range keyArr {
		if fmt.Sprintf("%v", valMap[val]) != "" && fmt.Sprintf("%v", valMap[val]) != "0" {
			urlValues = append(urlValues, fmt.Sprintf("%s=%v", val, valMap[val]))
		}
	}
	body := strings.Join(urlValues, "&") + "&key=" + MchApiKey
	has := md5.Sum([]byte(body))
	md5str := fmt.Sprintf("%x", has)
	return strings.ToUpper(md5str)
}

func DoPay(openid string, outTradeNo string, body string, spbillCreateIp string, totalFee int, notifyUrl string, tradeType string) (ret PayRes, req Pay, err error) {
	req.Appid = APPID
	req.MchId = MchID
	req.Openid = openid
	req.NonceStr = tool.RandomStr(16)
	req.Body = body
	req.OutTradeNo = outTradeNo
	req.TotalFee = totalFee
	req.SpbillCreateIp = spbillCreateIp
	req.NotifyUrl = notifyUrl
	req.TradeType = tradeType
	req.Sign = getWxPaySign(req)
	postData, e := xml.Marshal(req)
	if e != nil {
		err = e
		return
	}
	_, resp, errs := gorequest.New().Post("https://api.mch.weixin.qq.com/pay/unifiedorder").Set("Content-Type", "application/xml").SendString(string(postData)).EndBytes()
	if len(errs) > 0 {
		err = errs[0]
		return
	}
	err = xml.Unmarshal(resp, &ret)
	return
}
