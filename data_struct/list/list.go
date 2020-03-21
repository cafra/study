package list

import (
	"fmt"
)

type IList interface {
	PushHead(interface{}) IList
	Iterate(h func(interface{}))
	Print()
}

type Element struct {
	next *Element
	list *List

	Value interface{}
}

type List struct {
	root *Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int      // current list length excluding (this) sentinel element
}

func (l *List) Init() *List {
	l.root = nil
	l.len = 0
	return l
}

func (l *List) PushHead(data interface{}) *List {
	node := &Element{
		Value: data,
		list:  l,
		next:  l.root,
	}
	l.root = node
	l.len += 1
	return l
}

func (l *List) Iterate(h func(interface{})) {
	node := l.root
	for node != nil {
		h(node.Value)
		node = node.next
	}
}
func (l *List) Print() {
	l.Iterate(func(i interface{}) {
		fmt.Println(i)
	})
}
