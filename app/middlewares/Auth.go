package middlewares

import (
	"fmt"
	"github.com/dengpju/higo-gin/higo"
	"github.com/dengpju/higo-throw/exception"
	"github.com/gin-gonic/gin"
	"higo-framework/app/const/StaticConst"
	"higo-framework/app/errcode"
	"regexp"
)

// 鉴权
type Auth struct{}

// 构造函数
func NewAuth() *Auth {
	return &Auth{}
}

func (this *Auth) Middle(hg *higo.Higo) gin.HandlerFunc {
	return func(cxt *gin.Context) {
		ok, err := regexp.Match("^"+StaticConst.IndexRelativePath+"/", []byte(cxt.Request.URL.Path))
		if nil != err {
			panic(err)
		}
		if !ok {
			if route, ok := hg.GetRoute(cxt.Request.Method, cxt.Request.URL.Path); ok {
				fmt.Println(route.IsAuth())
				if ! higo.IsNotAuth(route.Flag()) && !route.IsStatic() {
					if "" == cxt.GetHeader("Token") {
						exception.Throw(exception.Message(errcode.TokenEmpty.Message()), exception.Code(int(errcode.TokenEmpty)))
					}
				}
			} else {
				exception.Throw(exception.Message(errcode.NotFound.Message()), exception.Code(int(errcode.NotFound)))
			}
		}
	}
}
