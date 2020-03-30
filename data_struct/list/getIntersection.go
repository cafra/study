package list

func GetIntersectionNode(list1, list2 List) *Element {
	l1 := list1.root
	l2 := list2.root
	for l1 != l2 {
		if l1 == nil {
			l1 = list2.root
		} else {
			l1 = l1.next
		}
		if l2 == nil {
			l2 = list1.root
		} else {
			l2 = l2.next
		}
	}
	return l1
}
