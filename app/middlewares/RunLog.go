package middlewares

import (
	"github.com/dengpju/higo-gin/higo"
	"github.com/gin-gonic/gin"
)

// 运行日志
type RunLog struct{}

// 构造函数
func NewRunLog() *RunLog {
	return &RunLog{}
}

func (this *RunLog) Middle(hg *higo.Higo) gin.HandlerFunc {
	return func(cxt *gin.Context) {
		// TODO::处理日志
		cxt.Next()
	}
}
