package list

func (l *List) Reverse() *List {
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
