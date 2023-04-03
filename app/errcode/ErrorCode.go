package errcode

import (
	"gitee.com/dengpju/higo-code/code"
)

// 错误码
type ErrorCode int64

func (this ErrorCode) Message(variables ...interface{}) string {
	return code.Get(this, variables...)
}

func (this ErrorCode) Register() code.Message {
	autoload()
	return code.Container()
}
