package UserEntity

import (
	"github.com/dengpju/higo-gin/higo"
    "time"
)

type Impl struct {
	isEdit      bool
	currentFlag higo.Flag
	UserId    int64    `gorm:"column:user_id" json:"user_id" comment:"主键"`
	UserSn    string    `gorm:"column:user_sn" json:"user_sn" comment:"编号"`
	Name    string    `gorm:"column:name" json:"name" comment:"名称"`
	Alias    string    `gorm:"column:alias" json:"alias" comment:"别名"`
	Nickname    string    `gorm:"column:nickname" json:"nickname" comment:"昵称"`
	Mobile    string    `gorm:"column:mobile" json:"mobile" comment:"手机号码"`
	Gender    int    `gorm:"column:gender" json:"gender" comment:"性别:0-未知,1-男,2-女"`
	Email    string    `gorm:"column:email" json:"email" comment:"邮箱"`
	Telephone    string    `gorm:"column:telephone" json:"telephone" comment:"座机"`
	Avatar    string    `gorm:"column:avatar" json:"avatar" comment:"微信头像"`
	State    int    `gorm:"column:state" json:"state" comment:"状态:1-启用,2-禁用"`
	Address    string    `gorm:"column:address" json:"address" comment:"地址"`
	Openid    string    `gorm:"column:openid" json:"openid" comment:"微信小程序openid"`
	Unionid    string    `gorm:"column:unionid" json:"unionid" comment:"微信小程序unionid"`
	Signature    string    `gorm:"column:signature" json:"signature" comment:"签名"`
	BackgroundImage    string    `gorm:"column:background_image" json:"background_image" comment:"背景图"`
	LastLoginMode    int    `gorm:"column:last_login_mode" json:"last_login_mode" comment:"最后登录方式:1-小程序,2-pc端后台,3-pc端企业微信扫码,4-admin端"`
	LastLoginTime    interface{}    `gorm:"column:last_login_time" json:"last_login_time" comment:"最后登录时间"`
	CreateTime    time.Time    `gorm:"column:create_time" json:"create_time" comment:"创建时间"`
	UpdateTime    time.Time    `gorm:"column:update_time" json:"update_time" comment:"更新时间"`
	DeleteTime    interface{}    `gorm:"column:delete_time" json:"delete_time" comment:"删除时间"`
}

func New() *Impl {
	tn := time.Now()
    return &Impl{CreateTime: tn, UpdateTime: tn}

}

func (this *Impl) IsEdit() bool {
	return this.isEdit
}

func (this *Impl) SetIsEdit(isEdit bool) {
	this.isEdit = isEdit
}

func (this *Impl) SetFlag(flag higo.Flag) {
	this.currentFlag = flag
	this.isEdit = true
}

func (this *Impl) Flag() higo.Flag {
	return this.currentFlag
}

func (this *Impl) PriEmpty() bool {
	return this.UserId == 0
}
