package controllers

import (
	"github.com/gin-gonic/gin"
	"server_mall/internal/services"
	"server_mall/pkg/resp"
)

type GoodsController struct {
	goodsService services.GoodsService
}

func (s *GoodsController) Detail(c *gin.Context) {
	goodsId := c.Query("goods_id")
	ret, err := s.goodsService.Detail(goodsId)
	if err != nil {
		resp.RespParamErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}
