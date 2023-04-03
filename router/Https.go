package router

import (
	"github.com/dengpju/higo-gin/higo"
	"github.com/gin-gonic/gin"
)

// https api 接口
type Https struct {
	*higo.Serve `inject:"Bean.NewServe('env.serve.HTTPS_HOST')"`
}

func NewHttps() *Https {
	return &Https{}
}

// 路由装载器
func (this *Https) Loader(hg *higo.Higo) {
	this.Api(hg)
}

// api 路由
func (this *Https) Api(hg *higo.Higo) {
	hg.GET("/", func(context *gin.Context) {
		_, _ = context.Writer.Write([]byte("Hello higo"))
	})
}
