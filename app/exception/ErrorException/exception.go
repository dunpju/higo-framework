package ErrorException

import (
	"github.com/dunpju/higo-gin/higo"
	"github.com/dunpju/higo-gin/higo/errcode"
	"github.com/dunpju/higo-throw/exception"
)

func Throw(err errcode.ErrorCode) {
	higo.Throw(exception.Code(int(err)), exception.Message(err.Message()), exception.Data(""))
}
