package list

import "fmt"

type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next *Element

	// The list to which this element belongs.
	list *List

	// The value stored with this element.
	Value interface{}
}

type List struct {
	root *Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int      // current list length excluding (this) sentinel element
}

func (l *List) Init() *List {
	//l.root = new(Element)
	//l.root.next = nil
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
