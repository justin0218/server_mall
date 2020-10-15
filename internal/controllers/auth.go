package controllers

import (
	"github.com/gin-gonic/gin"
	"server_mall/internal/services"
	"server_mall/pkg/resp"
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

func (s *AuthController) Jssdk(c *gin.Context) {
	url := c.Query("url")
	ret, err := s.authService.Jssdk(url)
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}

func (s *AuthController) AccessToken(c *gin.Context) {
	ret, err := s.authService.AccessToken()
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}

func (s *AuthController) Ticket(c *gin.Context) {
	ret, err := s.authService.Ticket()
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}

func (s *AuthController) Shorturl(c *gin.Context) {
	lurl := c.Query("lurl")
	if lurl == "" {
		resp.RespParamErr(c)
		return
	}
	ret, err := s.authService.GetShorUrl(lurl)
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}
