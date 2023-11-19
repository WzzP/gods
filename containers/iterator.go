package containers

// IteratorWithIndex
type IteratorWithIndex interface {
	//Next 将迭代器移动到下一个元素,如果元素存在则返回true
	Next() bool
	//Value 返回当前的元素值
	Value() any
	//Index 返回当前的索引值
	Index() int
	//Begin 将迭代器重置为初始状态
	Begin()
	//First 将迭代器移动到第一个元素上,如果存在第一个元素则返回true
	First() bool
	//NextTo 将迭代器移动到满足函数条件的元素上,如果NextTo返回true 则可以使用Value()Index()来获取元素值和索引值
	NextTo(func(index int, value any) bool) bool
}

type ReverseIteratorWithIndex interface {

	//Prev 将迭代器移动前一个元素,如果存在就返回true
	Prev() bool
	//End 将迭代器移动到最后一个元素
	End()

	//Last 将迭代器移动到后一个元素,如果存在就返回true
	Last() bool

	//PrevTo 将迭代器移动到满足函数条件的元素上,如果PrevTo返回true 则可以使用Value()Index()来获取元素值和索引值
	PrevTo(func(index int, value any) bool) bool

	IteratorWithIndex
}
