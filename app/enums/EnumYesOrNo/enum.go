package EnumYesOrNo

import "github.com/dengpju/higo-enum/enum"

var e YesOrNo

func Inspect(value int) error {
	return e.Inspect(value)
}

//是否
type YesOrNo int

func (this YesOrNo) Name() string {
	return "YesOrNo"
}

func (this YesOrNo) Inspect(value interface{}) error {
	return enum.Inspect(this, value)
}

func (this YesOrNo) Message() string {
	return enum.String(this)
}

const (
	Yes YesOrNo = 1 //是
	No  YesOrNo = 2 //否
)

func (this YesOrNo) Register() enum.Message {
	return make(enum.Message).
		Put(Yes, "是").
		Put(No, "否")
}

func Bool(value int) bool {
	return value == int(Yes)
}
