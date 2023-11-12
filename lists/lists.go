package lists

import "gods/containers"

// List list的接口，所有list实现都基于此接口
type List interface {
	Get(index int) (any, bool)
	Remove(index int)
	Add(values ...any)
	Contains(values ...any) bool
	Swap(index1, index2 int)
	Insert(index int, values ...any)
	Set(index int, value any)

	// Container 接口的组合
	containers.Container
}
