package wechat

import (
	"crypto/sha1"
	"fmt"
	"io"
	"server_mall/pkg/tool"
	"sort"
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
	ret.Noncestr = tool.RandomStr(16)
	ret.Timestamp = time.Now().Unix()
	signStr := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticket.Data.Ticket, ret.Noncestr, ret.Timestamp, url)
	ret.Signature = signature(signStr)
	return
}
func signature(params ...string) string {
	sort.Strings(params)
	h := sha1.New()
	for _, s := range params {
		_, _ = io.WriteString(h, s)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
