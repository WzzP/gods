package arraylist

import "gods/containers"

var _ containers.Container = (*List)(nil)

type runable func(index int, value any)

func (list *List) Each(f runable) {
}
