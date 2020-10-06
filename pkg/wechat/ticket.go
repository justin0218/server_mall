package wechat

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type Ticket struct {
	Errmsg  string `json:"errmsg"`
	Errcode int    `json:"errcode"`
	Ticket  string `json:"ticket"`
}

func GetTicket() (ret Ticket, err error) {
	rurl := fmt.Sprintf("http://momoman.cn/mall/v1/server/ticket")
	_, _, errs := gorequest.New().Get(rurl).EndStruct(&ret)
	if ret.Errcode != 0 || len(errs) > 0 {
		err = fmt.Errorf("wechat get ticket err:%v code:%d msg:%s", errs, ret.Errcode, ret.Errmsg)
		return
	}
	return
}
