package arraylist

import "gods/containers"

var _ containers.ReverseIteratorWithIndex = (*Iterator)(nil)

type Iterator struct {
	list  *List
	index int
}

func (list *List) Iterator() Iterator {
	return Iterator{list: list, index: -1}
}

func (i *Iterator) Prev() bool {
	if i.index >= 0 {
		i.index--
	}
	return i.list.withinRang(i.index)
}

func (i *Iterator) End() {
	i.index = i.list.size
}

func (i *Iterator) Last() bool {
	i.End()
	return i.Prev()
}

func (i *Iterator) PrevTo(f func(index int, value any) bool) bool {

	for i.Prev() {
		index, value := i.Index(), i.Value()
		if f(index, value) {
			return true
		}
	}
	return false
}

func (i *Iterator) Next() bool {
	if i.index < i.list.size {
		i.index++
	}
	return i.list.withinRang(i.index)
}

func (i *Iterator) Value() any {
	return i.list.elements[i.index]
}

func (i *Iterator) Index() int {
	return i.index
}

func (i *Iterator) Begin() {
	i.index = -1
}

func (i *Iterator) First() bool {
	i.Begin()
	return i.Next()
}

func (i *Iterator) NextTo(f func(index int, value any) bool) bool {
	for i.Next() {
		index, value := i.Index(), i.Value()
		if f(index, value) {
			return true
		}
	}
	return false
}
