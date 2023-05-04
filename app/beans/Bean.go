package beans

import (
	"github.com/dunpju/higo-gin/higo"
)

type Bean struct{ higo.Bean }

func NewBean() *Bean {
	return &Bean{}
}
