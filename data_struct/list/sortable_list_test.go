package list

import (
	"testing"
)

type MyData int

func (i MyData) Less(j ISort) bool {
	return i < j.(MyData)
}

func TestSortList_Merge(t *testing.T) {
	l1 := new(SortList).Init()
	for _, v := range []MyData{1, 3, 5} {
		l1.PushHead(v)
	}
	l1.Reverse().Print()

	l2 := new(SortList).Init()
	for _, v := range []MyData{2, 4, 6} {
		l2.PushHead(v)
	}
	l2.Reverse().Print()

	l1.Merge(l2).Print()
}
