package UserModel

import (
	"github.com/Masterminds/squirrel"
	"github.com/dengpju/higo-gin/higo"
	"github.com/dengpju/higo-ioc/injector"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)//"gitee.com/dengpju/higo-code/code"


type Impl struct {
	*higo.Orm       `inject:"Bean.NewOrm()"`
	UserId          int64       `gorm:"column:user_id" json:"user_id" comment:"主键"`
	UserSn          string      `gorm:"column:user_sn" json:"user_sn" comment:"编号"`
	Name            string      `gorm:"column:name" json:"name" comment:"名称"`
	Alias           string      `gorm:"column:alias" json:"alias" comment:"别名"`
	Nickname        string      `gorm:"column:nickname" json:"nickname" comment:"昵称"`
	Mobile          string      `gorm:"column:mobile" json:"mobile" comment:"手机号码"`
	Gender          int         `gorm:"column:gender" json:"gender" comment:"性别:0-未知,1-男,2-女"`
	Email           string      `gorm:"column:email" json:"email" comment:"邮箱"`
	Telephone       string      `gorm:"column:telephone" json:"telephone" comment:"座机"`
	Avatar          string      `gorm:"column:avatar" json:"avatar" comment:"微信头像"`
	State           int         `gorm:"column:state" json:"state" comment:"状态:1-启用,2-禁用"`
	Address         string      `gorm:"column:address" json:"address" comment:"地址"`
	Openid          string      `gorm:"column:openid" json:"openid" comment:"微信小程序openid"`
	Unionid         string      `gorm:"column:unionid" json:"unionid" comment:"微信小程序unionid"`
	Signature       string      `gorm:"column:signature" json:"signature" comment:"签名"`
	BackgroundImage string      `gorm:"column:background_image" json:"background_image" comment:"背景图"`
	LastLoginMode   int         `gorm:"column:last_login_mode" json:"last_login_mode" comment:"最后登录方式:1-小程序,2-pc端后台,3-pc端企业微信扫码,4-admin端"`
	LastLoginTime   interface{} `gorm:"column:last_login_time" json:"last_login_time" comment:"最后登录时间"`
	CreateTime      time.Time   `gorm:"column:create_time" json:"create_time" comment:"创建时间"`
	UpdateTime      time.Time   `gorm:"column:update_time" json:"update_time" comment:"更新时间"`
	DeleteTime      interface{} `gorm:"column:delete_time" json:"delete_time" comment:"删除时间"`
}

//init Validator
func init() {
	New().RegisterValidator()
}

func New(attrs ...higo.Property) *Impl {
	impl := &Impl{}
	higo.Propertys(attrs).Apply(impl)
	injector.BeanFactory.Apply(impl)
	return impl
}

func (this *Impl) New() higo.IClass {
	return New()
}

func (this *Impl) TableName() string {
	return "tl_user"
}

func (this *Impl) TableAlias(alias string) string {
	return this.TableName() + " AS " + alias
}

func (this *Impl) Mutate(attrs ...higo.Property) higo.Model {
	higo.Propertys(attrs).Apply(this)
	return this
}

//The custom tag, binding the tag eg: binding:"custom_tag_name"
//require import "gitee.com/dengpju/higo-code/code"
//
//example code:
//func (this *Impl) RegisterValidator() higo.Valid {
//	return higo.RegisterValid(this).
//		Tag("custom_tag_name",
//			higo.Rule("required", Codes.Success),
//			higo.Rule("min=5", Codes.Success))
//  Or
//  return higo.Verifier() // Manual call Register Validate: higo.Validate(verifier)
//}
func (this *Impl) RegisterValidator() *higo.Verify {
	return higo.RegisterValidator(this)
}

func (this *Impl) Exist() bool {
	return this.UserId > 0
}

func (this *Impl) GetByUserId(UserId int64, columns ...string) *gorm.DB {
	return this.Mapper(squirrel.Select(columns...).From(this.TableName()).Where("`user_id` = ?", UserId).ToSql()).Query()
}

func (this *Impl) GetByUserIds(UserIds []string, columns ...string) *gorm.DB {
	return this.Mapper(squirrel.Select(columns...).From(this.TableName()).Where("`user_id` IN(?)", strings.Join(UserIds, ",")).ToSql()).Query()
}

func (this *Impl) Paginate(perPage, page uint64) *higo.Pager {
	models := make([]*Impl, 0)
	pager := higo.NewPager(perPage, page)
	this.Table(this.TableName()).Paginate(pager).Find(&models)
	pager.Items = models
	return pager
}
