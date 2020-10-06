package services

import (
	"github.com/jinzhu/gorm"
	"server_mall/api"
	"server_mall/internal/models/user"
	"server_mall/pkg/jwt"
	"server_mall/pkg/wechat"
	"time"
)

type AuthService struct {
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
	db := api.Mysql.Get()
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

func (s *AuthService) SaveCache(key, value string) (err error) {
	client := api.Rds.Get()
	err = client.Set(key, value, time.Second*60).Err()
	return
}

func (s *AuthService) GetCache(key string) (value string, err error) {
	client := api.Rds.Get()
	value, err = client.Get(key).Result()
	return
}
