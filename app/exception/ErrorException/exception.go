package ErrorException

import (
	"github.com/dengpju/higo-gin/higo"
	"github.com/dengpju/higo-throw/exception"
	"higo-framework/app/errcode"
)

func Throw(err errcode.CodeErrorCode) {
	higo.Throw(exception.Code(int(err)), exception.Message(err.Message()), exception.Data(""))
}
