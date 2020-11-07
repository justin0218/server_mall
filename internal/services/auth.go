package services

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/parnurzeal/gorequest"
	"server_mall/internal/models/user"
	"server_mall/pkg/jwt"
	"server_mall/pkg/wechat"
	"time"
)

type AuthService struct {
	baseService
}

func (s *AuthService) Login(code string) (ret user.LoginRes, err error) {
	authToken, e := wechat.GetAuthAccessToken(code)
	if e != nil {
		err = e
		return
	}
	wuserInfo, e := wechat.GetUserInfo(authToken.Openid, authToken.AccessToken)
	if e != nil {
		err = e
		return
	}
	db := s.Mysql.Get()
	olduser, e := user.NewModel(db).GetByOpenid(authToken.Openid)
	if e == gorm.ErrRecordNotFound { //未注册
		newuser, e := user.NewModel(db).Create(user.User{
			Openid:   wuserInfo.Openid,
			Avatar:   wuserInfo.Headimgurl,
			Nickname: wuserInfo.Nickname,
		})
		if e != nil {
			err = e
			return
		}
		ret.User = newuser
		ret.Token, err = jwt.CreateToken(int64(newuser.Id))
		return
	}
	if e != nil {
		err = e
		return
	}
	err = user.NewModel(db).UpdateById(user.User{
		Id:       olduser.Id,
		Avatar:   wuserInfo.Headimgurl,
		Nickname: wuserInfo.Nickname,
	})
	if err != nil {
		return
	}
	ret.User = olduser
	ret.Token, err = jwt.CreateToken(int64(olduser.Id))
	return
}

func (s *AuthService) AccessToken() (ret wechat.AccessToken, err error) {
	rk := fmt.Sprintf("%s_access_token", wechat.APPID)
	cacheRes, _ := s.Redis.Get().Get(rk).Result()
	if cacheRes != "" {
		err = json.Unmarshal([]byte(cacheRes), &ret)
		if err == nil {
			return
		}
	}
	rurl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", wechat.APPID, wechat.SECRET)
	_, bytesRes, errs := gorequest.New().Get(rurl).EndBytes()
	if len(errs) > 0 {
		err = fmt.Errorf("wechat get access_token err:%v", errs)
		return
	}
	err = json.Unmarshal(bytesRes, &ret)
	if err != nil {
		err = fmt.Errorf("wechat get access_token err:%v", err)
		return
	}
	if ret.Errcode != 0 {
		err = fmt.Errorf("wechat get access_token err msg:%s", ret.Errmsg)
		return
	}
	s.Redis.Get().Set(rk, bytesRes, time.Second*7000)
	return
}

func (s *AuthService) Ticket() (ret wechat.Ticket, err error) {
	rk := fmt.Sprintf("%s_ticket", wechat.APPID)
	cacheRes, _ := s.Redis.Get().Get(rk).Result()
	if cacheRes != "" {
		err = json.Unmarshal([]byte(cacheRes), &ret)
		if err == nil {
			return
		}
	}
	accessToken, e := wechat.GetAccessToken()
	if e != nil {
		err = e
		return
	}
	rurl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi", accessToken.Data.AccessToken)
	_, bytesRes, errs := gorequest.New().Get(rurl).EndBytes()
	if len(errs) > 0 {
		err = fmt.Errorf("wechat get ticket err:%v", errs)
		return
	}
	err = json.Unmarshal(bytesRes, &ret)
	if err != nil {
		return
	}
	if ret.Errcode != 0 {
		err = fmt.Errorf(ret.Errmsg)
		return
	}
	s.Redis.Get().Set(rk, bytesRes, time.Second*7000)
	return
}

func (s *AuthService) Jssdk(url string) (ret wechat.Jssdk, err error) {
	return wechat.GetJssdk(url)
}

func (s *AuthService) GetUserInfo(uid int) (ret user.User, err error) {
	db := s.Mysql.Get()
	return user.NewModel(db).GetByUid(uid)
}

func (s *AuthService) GetShortUrl(lurl string) (ret wechat.ShorUrl, err error) {
	return wechat.GetShortUrl(lurl)
}
