module server_mall

go 1.15

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200423211502-4bdfaf469ed5
	golang.org/x/image => github.com/golang/image v0.0.0-20200430140353-33d19683fad8
	golang.org/x/mod => github.com/golang/mod v0.2.0
	golang.org/x/net => github.com/golang/net v0.0.0-20200421231249-e086a090c8fd
	golang.org/x/sync => github.com/golang/sync v0.0.0-20200317015054-43a5402ce75a
	golang.org/x/sys => github.com/golang/sys v0.0.0-20200420163511-1957bb5e6d1f
	golang.org/x/text => github.com/golang/text v0.3.3
	golang.org/x/tools => github.com/golang/tools v0.0.0-20200425043458-8463f397d07c
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20191204190536-9bdfabe68543
	google.golang.org/appengine => github.com/golang/appengine v1.6.6
)

require (
	github.com/astaxie/beego v1.12.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.2
	github.com/gorilla/websocket v1.4.2
	github.com/itsjamie/gin-cors v0.0.0-20160420130702-97b4a9da7933
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/protobuf v0.0.0-20180321161605-ebd3be6d4fdb
	github.com/mojocn/base64Captcha v1.3.1
	github.com/parnurzeal/gorequest v0.2.16
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/tencentyun/cos-go-sdk-v5 v0.7.10
	gopkg.in/yaml.v2 v2.2.8
	moul.io/http2curl v1.0.0 // indirect
)
