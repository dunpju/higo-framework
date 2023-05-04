package errcode

import "github.com/dunpju/higo-gin/higo/errcode"

func init() {
	errcode.Autoload = autoload
}
