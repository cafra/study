package list

type IList interface {
	PushHead(interface{}) IList
	Iterate(h func(interface{}))
	Reverse() IList
	Print()
}
