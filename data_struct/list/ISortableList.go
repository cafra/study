package list

type ISortList interface {
	IList
	Merge(ISortList) ISortList // 合并有序链表
}

type ISort interface {
	Less(ISort) bool
}
