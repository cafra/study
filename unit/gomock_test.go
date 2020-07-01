package unit

import (
	"github.com/cafra/study/unit/mock_unit"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mock_unit.NewMockMyInterface(ctrl)

	// Asserts that the first and only call to Bar() is passed 99.
	// Anything else will fail.
	m.EXPECT().HelloWorld("t1", "t2")

	SUT(m)
}
