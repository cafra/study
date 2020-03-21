package list

import "fmt"



type SortElement struct {
	next *SortElement
	list *SortList

	Value ISort
}

type SortList struct {
	root *SortElement // sentinel list element, only &root, root.prev, and root.next are used
	len  int          // current list length excluding (this) sentinel element
}

func NewSortableList() ISortList {
	l := new(SortList)
	l.root = nil
	l.len = 0
	return l
}

func (l *SortList) PushHead(data interface{}) IList {
	node := &SortElement{
		Value: data.(ISort),
		list:  l,
		next:  l.root,
	}
	l.root = node
	l.len += 1
	return l
}

func (l *SortList) Iterate(h func(interface{})) {
	node := l.root
	for node != nil {
		h(node.Value)
		node = node.next
	}
}
func (l *SortList) Print() {
	l.Iterate(func(i interface{}) {
		fmt.Println(i)
	})
}
func (l *SortList) Reverse() IList {
	if l.root == nil || l.root.next == nil {
		return l
	}
	var pre *SortElement
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

func (l1 *SortList) Merge(l2 ISortList) ISortList {
	return &SortList{
		root: Merge(l1.root, l2.(*SortList).root),
		len:  l1.len + l2.(*SortList).len,
	}
}

func Merge(l1, l2 *SortElement) *SortElement {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Value.Less(l2.Value) {
		l1.next = Merge(l1.next, l2)
		return l1
	} else {
		l2.next = Merge(l1, l2.next)
		return l2
	}
}
