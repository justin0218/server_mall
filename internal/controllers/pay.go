package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server_mall/internal/services"
	"server_mall/pkg/jwt"
	"server_mall/pkg/resp"
	"time"
)

type PayController struct {
	payService   services.PayService
	authService  services.AuthService
	goodsService services.GoodsService
}

func (s *PayController) Pay(c *gin.Context) {
	uid := jwt.GetUid(c)
	uinfo, err := s.authService.GetUserInfo(uid)
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	goodsId := c.Query("goods_id")
	goodsInfo, err := s.goodsService.Detail(goodsId)
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	if goodsInfo.Status == 0 {
		resp.RespInternalErr(c, "商品已下架")
		return
	}
	ret, err := s.payService.Pay(uinfo.Openid, fmt.Sprintf("%d", time.Now().Unix()), goodsInfo.Name, c.ClientIP(), goodsInfo.Price, "https://baidu.com", "JSAPI")
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}
