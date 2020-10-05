package wechat

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"server_mall/configs"
)

var (
	APPID  = configs.Dft.Get().Weacht.Appid
	SECRET = configs.Dft.Get().Weacht.Secret
)

type AuthAccessToken struct {
	Errmsg       string `json:"errmsg"`
	Errcode      int    `json:"errcode"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

func GetAuthAccessToken(code string) (ret AuthAccessToken, err error) {
	rurl := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code")
	_, _, errs := gorequest.New().Get(rurl).EndStruct(&ret)
	if ret.Errcode != 0 || len(errs) > 0 {
		err = fmt.Errorf("wechat get auth access token err:%v code:%d msg:%s", errs, ret.Errcode, ret.Errmsg)
		return
	}
	return
}

type UserInfo struct {
	Errmsg     string `json:"errmsg"`
	Errcode    int    `json:"errcode"`
	Openid     string `json:"openid"`
	Nickname   string `json:"nickname"`
	Headimgurl string `json:"headimgurl"`
}

func GetUserInfo(openid string) (ret UserInfo, err error) {
	tokenRes, e := GetAccessToken()
	if e != nil {
		err = e
		return
	}
	rurl := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", tokenRes.AccessToken, openid)
	_, _, errs := gorequest.New().Get(rurl).EndStruct(&ret)
	if ret.Errcode != 0 || len(errs) > 0 {
		err = fmt.Errorf("wechat get userinfo err:%v code:%d msg:%s", errs, ret.Errcode, ret.Errmsg)
		return
	}
	return
}
