package arraylist

import "gods/containers"

var _ containers.EnumerableWithIndex = (*List)(nil)

func (list *List) Each(f func(index int, value any)) {
	iterator := list.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

func (list *List) Any(f func(index int, value any) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return true
		}
	}
	return false
}

func (list *List) All(f func(index int, value any) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {

		if !f(iterator.Index(), iterator.Value()) {
			return false
		}
	}
	return true
}

func (list *List) Find(f func(index int, value any) bool) (int, any) {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return iterator.Index(), iterator.Value()
		}
	}
	return -1, nil
}

func (list *List) Select(f func(index int, value any) bool) *List {
	newList := &List{}
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			newList.Add(iterator.Value())
		}
	}
	return newList
}

func (list *List) Map(f func(index int, value any) any) *List {
	newList := &List{}
	iterator := list.Iterator()
	for iterator.Next() {
		newList.Add(f(iterator.Index(), iterator.Value()))
	}
	return newList
}
