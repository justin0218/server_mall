package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server_mall/internal/services"
	"server_mall/pkg/resp"
	"time"
)

type PayController struct {
	payService services.PayService
}

func (s *PayController) Pay(c *gin.Context) {
	ret, err := s.payService.Pay("oBYAkw3URP9pAQekMZ1GYmuNfFfQ", fmt.Sprintf("%d", time.Now().Unix()), "测试", "127.0.0.1", 100, "https://baidu.com", "JSAPI")
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}
