package errcode

import "gitee.com/dengpju/higo-code/code"

const (
	TokenEmpty        CodeErrorCode = iota + 400001 //token不存在
	TokenError                                      //token错误
	TokenSetError                                   //设置token失败
	TokenExpiredError                               //无效token
	PrimaryIdError                                  //主键id错误
	NotExistError                                   //不存在
	NotFound                                        //资源不存在
)

func code400001() {
	code.Container().
		Put(TokenEmpty, "token不存在").
		Put(TokenError, "token错误").
		Put(TokenSetError, "设置token失败").
		Put(TokenExpiredError, "无效token").
		Put(PrimaryIdError, "主键id错误").
		Put(NotExistError, "不存在").
		Put(NotFound, "资源不存在")
}
