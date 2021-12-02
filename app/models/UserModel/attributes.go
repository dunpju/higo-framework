package UserModel

import (
	"github.com/dengpju/higo-gin/higo"
    "time"
)

const (
    UserId = "user_id"  //主键
    UserSn = "user_sn"  //编号
    Name = "name"  //名称
    Alias = "alias"  //别名
    Nickname = "nickname"  //昵称
    Mobile = "mobile"  //手机号码
    Gender = "gender"  //性别:0-未知,1-男,2-女
    Email = "email"  //邮箱
    Telephone = "telephone"  //座机
    Avatar = "avatar"  //微信头像
    State = "state"  //状态:1-启用,2-禁用
    Address = "address"  //地址
    Openid = "openid"  //微信小程序openid
    Unionid = "unionid"  //微信小程序unionid
    Signature = "signature"  //签名
    BackgroundImage = "background_image"  //背景图
    LastLoginMode = "last_login_mode"  //最后登录方式:1-小程序,2-pc端后台,3-pc端企业微信扫码,4-admin端
    LastLoginTime = "last_login_time"  //最后登录时间
    CreateTime = "create_time"  //创建时间
    UpdateTime = "update_time"  //更新时间
    DeleteTime = "delete_time"  //删除时间
)
func WithUserId(v int64) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).UserId = v
	}
}

func WithUserSn(v string) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).UserSn = v
	}
}

func WithName(v string) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).Name = v
	}
}

func WithAlias(v string) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).Alias = v
	}
}

func WithNickname(v string) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).Nickname = v
	}
}

func WithMobile(v string) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).Mobile = v
	}
}

func WithGender(v int) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).Gender = v
	}
}

func WithEmail(v string) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).Email = v
	}
}

func WithTelephone(v string) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).Telephone = v
	}
}

func WithAvatar(v string) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).Avatar = v
	}
}

func WithState(v int) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).State = v
	}
}

func WithAddress(v string) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).Address = v
	}
}

func WithOpenid(v string) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).Openid = v
	}
}

func WithUnionid(v string) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).Unionid = v
	}
}

func WithSignature(v string) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).Signature = v
	}
}

func WithBackgroundImage(v string) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).BackgroundImage = v
	}
}

func WithLastLoginMode(v int) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).LastLoginMode = v
	}
}

func WithLastLoginTime(v interface{}) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).LastLoginTime = v
	}
}

func WithCreateTime(v time.Time) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).CreateTime = v
	}
}

func WithUpdateTime(v time.Time) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).UpdateTime = v
	}
}

func WithDeleteTime(v interface{}) higo.Property {
	return func(class higo.IClass) {
		class.(*Impl).DeleteTime = v
	}
}

