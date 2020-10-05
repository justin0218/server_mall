package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"server_mall/internal/models/admin"
	"server_mall/internal/services"
	"server_mall/pkg/resp"
	"server_mall/pkg/upload"
	"strings"
)

type AdminController struct {
	adminService *services.AdminService
}

func (s *AdminController) Login(c *gin.Context) {
	req := admin.LoginReq{}
	err := c.BindJSON(&req)
	if err != nil {
		resp.RespParamErr(c)
		return
	}
	if req.Username == "" || req.Password == "" {
		resp.RespParamErr(c)
		return
	}
	ret, err := s.adminService.Login(req.Username, req.Password)
	if err != nil {
		resp.RespInternalErr(c, "账号或密码错误")
		return
	}
	resp.RespOk(c, ret)
	return
}

func (s *AdminController) UploadFile(c *gin.Context) {
	f, err := c.FormFile("editormd-image-file")
	if err != nil {
		resp.RespParamErr(c)
		return
	}
	path := c.DefaultQuery("path", "images")
	flen := strings.Split(f.Filename, ".")
	if len(flen) != 2 {
		resp.RespParamErr(c, "格式不正确")
		return
	}
	suffix := flen[1]
	fname := uuid.New().String()
	obj := fmt.Sprintf("%s/%s.%s", path, fname, suffix)
	file, err := f.Open()
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	url, err := upload.UploadFile(obj, fileBytes)
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	mapRes := make(map[string]string)
	mapRes["url"] = url
	mapRes["success"] = "1"
	c.JSON(http.StatusOK, mapRes)
	return
}

func (s *AdminController) FileRead(c *gin.Context) {
	key := c.Query("key")
	ret, err := s.adminService.FileRead(key)
	if err != nil {
		resp.RespInternalErr(c, err.Error())
		return
	}
	resp.RespOk(c, ret)
	return
}
