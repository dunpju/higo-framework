package router

import (
	"fmt"
	"github.com/dengpju/higo-gin/higo"
	"github.com/dengpju/higo-utils/utils"
	"higo-framework/app/const/StaticConst"
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
	// 静态文件
	hg.Static(StaticConst.IndexRelativePath, fmt.Sprintf("%sdist", hg.GetRoot().Join(utils.PathSeparator())))
	this.Api(hg)
}

// api 路由
func (this *Https) Api(hg *higo.Higo) {
	//router.AddGroup()
}
