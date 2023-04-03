package beans

import (
	"github.com/dengpju/higo-gin/higo"
)

type Bean struct{ higo.Bean }

func NewBean() *Bean {
	return &Bean{}
}
