package arraylist

import "gods/containers"

var _ containers.ReverseIteratorWithIndex = (*Iterator)(nil)

type Iterator struct {
	list  *List
	index inx
}

func (i Iterator) Prev() bool {
	//TODO implement me
	panic("implement me")
}

func (i Iterator) End() {
	//TODO implement me
	panic("implement me")
}

func (i Iterator) Last() bool {
	//TODO implement me
	panic("implement me")
}

func (i Iterator) PrevTo(f func(index int, value any) bool) bool {
	//TODO implement me
	panic("implement me")
}

func (i Iterator) Next() {
	//TODO implement me
	panic("implement me")
}

func (i Iterator) Value() any {
	//TODO implement me
	panic("implement me")
}

func (i Iterator) Index() int {
	//TODO implement me
	panic("implement me")
}

func (i Iterator) Begin() int {
	//TODO implement me
	panic("implement me")
}

func (i Iterator) First() bool {
	//TODO implement me
	panic("implement me")
}

func (i Iterator) NextTo(f func(index int, value any) bool) bool {
	//TODO implement me
	panic("implement me")
}
