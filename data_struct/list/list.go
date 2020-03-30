package list

import "fmt"

type Element struct {
	next *Element
	list *List

	Value interface{}
}

type List struct {
	root *Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int      // current list length excluding (this) sentinel element
}

func New() IList {
	l := new(List)
	l.root = nil
	l.len = 0
	return l
}

func (l *List) PushHead(data interface{}) IList {
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

func (l *List) Reverse() IList {
	if l.root == nil || l.root.next == nil {
		return l
	}
	var pre *Element
	cur := l.root

	for cur != nil {
		next := cur.next
		cur.next = pre
		pre = cur
		cur = next
	}
	l.root = pre
	return l
}
