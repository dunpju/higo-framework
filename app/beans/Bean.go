package beans

import (
	"github.com/dengpju/higo-gin/higo"
)

type MyBean struct{ higo.Bean }

func NewMyBean() *MyBean {
	return &MyBean{}
}
