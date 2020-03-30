package list

import (
	"fmt"
	"testing"
)

func TestList_PushHead(t *testing.T) {
	list := New()

	for i := 0; i < 10; i++ {
		list.PushHead(i)
	}
	list.Print()
	fmt.Println("===========")
	list.Reverse().Print()

	list = New()
	for i := 0; i < 3; i++ {
		list.PushHead(i)
	}
	list.Reverse().Print()
}
