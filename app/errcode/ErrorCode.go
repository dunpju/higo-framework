package errcode

import (
	"gitee.com/dengpju/higo-code/code"
)

//错误码
type CodeErrorCode int64

func (this CodeErrorCode) Message() string {
	return code.Get(this)
}

func (this CodeErrorCode) Register() code.Message {
	autoload()
	return code.Container()
}
