package UserDao

import (
	"github.com/dengpju/higo-gin/higo"
	"github.com/dengpju/higo-gin/higo/sql"
	"higo-framework/app/entity/UserEntity"
	"higo-framework/app/errcode"
	"higo-framework/app/exception/DaoException"
	"higo-framework/app/models/UserModel"
)

type Dao struct {
	model *UserModel.Impl
}

func New() *Dao {
	return &Dao{model: UserModel.New()}
}

func (this *Dao) Orm() *higo.Orm {
	return this.model.Orm
}

func (this *Dao) Model() *UserModel.Impl {
	return UserModel.New()
}

func (this *Dao) SetData(userEntity *UserEntity.Impl) {
	if userEntity.IsEdit() { //编辑
		if userEntity.PriEmpty() {
			DaoException.Throw("UserId"+errcode.PrimaryIdError.Message(), int(errcode.PrimaryIdError))
		}
		if !this.GetByUserId(userEntity.UserId).Exist() {
			DaoException.Throw(errcode.NotExistError.Message(), int(errcode.NotExistError))
		}
		builder := this.model.Update(this.model.TableName()).Where("user_id", userEntity.UserId)
		if UserEntity.FlagDelete == userEntity.Flag() {

		} else {
			//TODO::设置字段
		}
		builder.Set("update_time", userEntity.UpdateTime)
	} else { //新增
		this.model.Insert(this.model.TableName()).
			Set("create_time", userEntity.CreateTime).
			Set("update_time", userEntity.UpdateTime)
	}
	this.model.Build()
}

//添加
func (this *Dao) Add() int64 {
	higo.Result(this.model.Mapper(this.model.GetBuilder()).InsertGetId().Error).Unwrap()
	return this.model.LastInsertId()
}

//更新
func (this *Dao) Update() bool {
	higo.Result(this.model.Mapper(this.model.GetBuilder()).Exec().Error).Unwrap()
	return true
}

//id查询
func (this *Dao) GetByUserId(userId int64, fields ...string) *UserModel.Impl {
	if len(fields) == 0 {
		fields = append(fields, "*")
	}
	model := this.Model()
	this.model.Mapper(sql.Select(fields...).
		From(this.model.TableName()).
		Where("`user_id` = ?", userId).
		Where("isnull(`delete_time`)").
		ToSql()).Query().Scan(&model)
	return model
}

//id集查询
func (this *Dao) GetByUserIds(userIds []interface{}, fields ...string) []*UserModel.Impl {
	if len(fields) == 0 {
		fields = append(fields, "*")
	}
	models := make([]*UserModel.Impl, 0)
	this.model.Mapper(sql.Select(fields...).
		From(this.model.TableName()).
		Where("`user_id` IN (?)", userIds).
		Where("isnull(`delete_time`)").
		ToSql()).Query().Scan(&models)
	return models
}

//硬删除
func (this *Dao) DeleteByUserId(userId int64) {
	higo.Result(this.model.Mapper(sql.Delete(this.model.TableName()).
		Where("`user_id` = ?", userId).
		ToSql()).Exec().Error).Unwrap()
}

func (this *Dao) GetByMobile(mobile string, fields ...string) *UserModel.Impl {
	if len(fields) == 0 {
		fields = append(fields, "*")
	}
	model := this.Model()
	this.model.Mapper(sql.Select(fields...).
		From(this.model.TableName()).
		Where("`mobile` = ?", mobile).
		Where("isnull(`delete_time`)").
		ToSql()).Query().Scan(&model)
	return model
}
