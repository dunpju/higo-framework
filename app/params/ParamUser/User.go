package ParamUser

import (
	"github.com/dengpju/higo-gin/higo"
	"higo-children/app/errcode"
)

type SetRole struct {
	UserId          int64   `json:"user_id" binding:"user_id"`
	PrivilegeRoleId []int64 `json:"privilege_role_id" binding:"privilege_role_id"`
}

func NewSetRole() *SetRole {
	return &SetRole{PrivilegeRoleId: make([]int64, 0)}
}

func (this *SetRole) RegisterValidator() *higo.Verify {
	return higo.Verifier().
		Tag("user_id",
			higo.Rule("required", errcode.MobileError)).
		Tag("privilege_role_id",
			higo.Rule("required", errcode.PrivilegeRoleIdError))
}
