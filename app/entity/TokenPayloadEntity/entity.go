package TokenPayloadEntity

import (
	"higo-children/app/entity/EducationTeacherClassJoinEntity"
	"higo-children/app/entity/StaffPayloadEntity"
)

type Impl struct {
	LastLoginMode      int                                                         `json:"last_login_mode"`
	LastLoginTime      interface{}                                                 `json:"last_login_time"`
	UserId             int64                                                       `json:"user_id"`
	UserSn             string                                                      `json:"user_sn"`
	Name               string                                                      `json:"name"`
	Token              string                                                      `json:"token"`
}

func New() *Impl {
	impl := &Impl{}
	return impl
}

func (this *Impl) Exist() bool {
	return this.UserId > 0
}
