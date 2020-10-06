package wechat

import (
	"crypto/sha1"
	"fmt"
	"io"
	"time"
)

type Jssdk struct {
	Appid     string `json:"appid"`
	Noncestr  string `json:"noncestr"`
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
}

//jsapi_ticket=sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg&
//=Wm3WZYTPz0wzccnW&timestamp=1414587457&url=http://mp.weixin.qq.com?params=value

func GetJssdk(url string) (ret Jssdk, err error) {
	ticket, e := GetTicket()
	if e != nil {
		err = e
		return
	}
	ret.Appid = APPID
	ret.Noncestr = fmt.Sprintf("%d", time.Now().Unix())
	ret.Timestamp = time.Now().Unix()
	signStr := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticket.Ticket, ret.Noncestr, ret.Timestamp, url)
	h := sha1.New()
	_, err = io.WriteString(h, signStr)
	if err != nil {
		return
	}
	ret.Signature = fmt.Sprintf("%x", h.Sum(nil))
	return
}
