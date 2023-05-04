package middlewares

import (
	"gitee.com/dengpju/higo-code/code"
	"github.com/dunpju/higo-gin/higo"
	"github.com/dunpju/higo-logger/logger"
	"github.com/dunpju/higo-throw/exception"
	"github.com/dunpju/higo-utils/utils/maputil"
	"github.com/dunpju/higo-utils/utils/runtimeutil"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RecoverHandle(cxt *gin.Context, r interface{}) {
	//记录debug调用栈
	goid, _ := runtimeutil.GoroutineID()
	logger.LoggerStack(r, goid)

	//封装通用json返回
	if h, ok := r.(gin.H); ok {
		cxt.JSON(http.StatusOK, h)
	} else if cd, ok := r.(*code.CodeMessage); ok {
		cxt.JSON(http.StatusOK, gin.H{
			"code":    cd.Code,
			"message": cd.Message,
			"data":    nil,
			"type":    "error",
		})
	} else if arrayMap, ok := r.(*maputil.ArrayMap); ok {
		cxt.JSON(http.StatusOK, arrayMap.Value())
	} else if validate, ok := r.(*higo.ValidateError); ok {
		cxt.JSON(http.StatusOK, gin.H{
			"code":    validate.Get().Code,
			"message": validate.Get().Message,
			"data":    nil,
			"type":    "error",
		})
	} else if err, ok := r.(error); ok {
		cxt.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest, // TODO::错误码自行修改
			"message": exception.ErrorToString(err),
			"data":    nil,
			"type":    "error",
		})
	} else {
		cxt.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest, // TODO::错误码自行修改
			"message": r,
			"data":    nil,
			"type":    "error",
		})
	}
}
