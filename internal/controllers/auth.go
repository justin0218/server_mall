package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"server_mall/api"
	"server_mall/internal/services"
	"server_mall/pkg/resp"
	"server_mall/pkg/wechat"
	"time"
)

type AuthController struct {
	authService *services.AuthService
}

func (s *AuthController) Login(c *gin.Context) {
	code := c.Query("code")
	ret, err := s.authService.Login(code)
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}

func (s *AuthController) SaveCache(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	err := s.authService.SaveCache(key, value)
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, nil)
	return
}

func (s *AuthController) GetCache(c *gin.Context) {
	key := c.Query("key")
	value, err := s.authService.GetCache(key)
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, value)
	return
}

func (s *AuthController) AccessToken(c *gin.Context) {
	ret := wechat.AccessToken{}
	rk := fmt.Sprintf("%s_access_token", wechat.APPID)
	cacheRes, err := api.Rds.Get().Get(rk).Bytes()
	if err == nil {
		err = json.Unmarshal(cacheRes, &ret)
		if err == nil {
			resp.RespOk(c, ret)
			return
		}
	}
	rurl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", wechat.APPID, wechat.SECRET)
	_, bytesRes, errs := gorequest.New().Get(rurl).EndStruct(&ret)
	if ret.Errcode != 0 || len(errs) > 0 {
		resp.RespInternalErr(c, fmt.Errorf("wechat get access_token err:%v code:%d msg:%s", errs, ret.Errcode, ret.Errmsg).Error())
		return
	}
	api.Rds.Get().Set(rk, bytesRes, time.Second*7000)
	resp.RespOk(c, ret)
	return
}
