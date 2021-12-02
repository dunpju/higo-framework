package UserService

import (
	"github.com/dengpju/higo-gin/higo"
	"higo-children/app/dao/PrivilegeRoleDao"
	"higo-children/app/dao/PrivilegeRoleFlagDao"
	"higo-children/app/dao/PrivilegeUserRoleDao"
	"higo-children/app/dao/UserFlagDao"
	"higo-children/app/entity/PrivilegeUserRoleEntity"
	"higo-children/app/entity/UserFlagEntity"
	"higo-children/app/errcode"
	"higo-children/app/exception/ErrorException"
	"higo-children/app/params/ParamUser"
)

type Impl struct {
}

func New() *Impl {
	return &Impl{}
}

func (this *Impl) SetRole(params *ParamUser.SetRole) {
	//准备数据
	privilegeUserRoleEntity := PrivilegeUserRoleEntity.New()
	privilegeUserRoleEntity.UserId = params.UserId
	privilegeRoleIds := params.PrivilegeRoleId

	privilegeRoleDao := PrivilegeRoleDao.New()
	role := privilegeRoleDao.GetPrivilegeRoleByPrivilegeRoleIds(privilegeRoleIds, "privilege_role_id")
	if len(role) != len(privilegeRoleIds) {
		ErrorException.Throw(errcode.PrivilegeRoleError)
	}
	userFlagEntity := UserFlagEntity.New()
	userFlagEntity.UserId = privilegeUserRoleEntity.UserId
	privilegeRoleFlagDao := PrivilegeRoleFlagDao.New()
	privilegeRoleFlag := privilegeRoleFlagDao.GetJoinPrivilegeRoleFlagByPrivilegeRoleIds(privilegeRoleIds)
	for _, v := range privilegeRoleFlag {
		userFlagEntity.PrivilegeFlagIds = append(userFlagEntity.PrivilegeFlagIds, v.PrivilegeFlagId)
	}
	privilegeUserRoleDao := PrivilegeUserRoleDao.New()
	userFlagDao := UserFlagDao.New()
	//开启事务
	higo.Begin(privilegeUserRoleDao.Orm(), userFlagDao.Orm()).Transaction(func() error {
		//先删后增
		privilegeUserRoleDao.DeleteByUserId(privilegeUserRoleEntity.UserId)
		//插入tl_privilege_user_role表
		for _, privilegeRoleId := range privilegeRoleIds {
			privilegeUserRoleEntity.PrivilegeRoleId = privilegeRoleId
			privilegeUserRoleDao.SetData(privilegeUserRoleEntity)
			privilegeUserRoleDao.Add()
		}
		//先删后增
		userFlagDao := UserFlagDao.New()
		userFlagDao.DeleteByUserId(privilegeUserRoleEntity.UserId)
		//插入tl_user_flag表
		for _, privilegeRoleId := range userFlagEntity.PrivilegeFlagIds {
			userFlagEntity.PrivilegeFlagId = privilegeRoleId
			userFlagDao.SetData(userFlagEntity)
			userFlagDao.Add()
		}
		return nil
	})
}
