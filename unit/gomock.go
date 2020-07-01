package unit

import (
	"github.com/golang/mock/gomock"
)

type MyInterface interface {
	HelloWorld(x int64, y string)
}
