package middlewares

import (
	"fmt"
	"github.com/dengpju/higo-gin/higo"
	"github.com/gin-gonic/gin"
	"higo-framework/app/const/StaticConst"
	"regexp"
)

// 运行日志
type Admin struct{}

// 构造函数
func NewAdmin() *Admin {
	return &Admin{}
}

func (this *Admin) Middle(hg *higo.Higo) gin.HandlerFunc {
	return func(cxt *gin.Context) {
		ok, err := regexp.Match("^"+StaticConst.IndexRelativePath+"/", []byte(cxt.Request.URL.Path))
		if nil != err {
			panic(err)
		}
		if !ok {
			fmt.Printf("Admin中间件:%s\n", higo.RouterContainer.Get(cxt.Request.Method, cxt.Request.URL.Path).Desc())
		}
		cxt.Next()
	}
}
