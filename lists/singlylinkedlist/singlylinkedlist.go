package singlylinkedlist

import (
	"gods/lists"
	"gods/utils"
)

var _ lists.List = (*List)(nil)

// 单链表
type List struct {
	first *element
	last  *element
	size  int
}

type element struct {
	value any
	next  *element
}

func New(values ...any) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List) Get(index int) (any, bool) {
	if !list.withinRang(index) {
		return nil, false
	}

	element := list.first
	for e := 0; e != index; e, element = e+1, element.next {

	}
	return element.value, true
}

func (list *List) Remove(index int) {
	if !list.withinRang(index) {
		return
	}
	if list.size == 1 {
		list.Clear()
		return
	}

	var prevElement *element
	var element *element
	//判断index位置,靠前正序遍历,考后倒序遍历
	for e := 0; e != index; e, element = e+1, element.next {
		prevElement = element
	}
	//开始元素
	if element == list.first {
		list.first = element.next
	}
	//结束元素
	if element == list.last {
		list.last = prevElement
	}

	//中间元素,把前面元素的next 执行 被删除元素的next
	if prevElement != nil {
		prevElement.next = element.next
	}
	element = nil
	list.size--

}

func (list *List) Add(values ...any) {
	for _, value := range values {
		newElement := &element{value: value}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			//first       last
			// A ->  B -> C
			// add D
			// C -> D
			// last = D
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

func (list *List) Contains(values ...any) bool {
	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}
	for _, value := range values {
		found := false
		for element := list.first; element != nil; element = element.next {
			if element.value == value {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (list *List) Swap(i, j int) {
	if list.withinRang(i) && list.withinRang(j) && i != j {
		var element1, element2 *element
		for e, currentElement := 0, list.first; element1 == nil || element2 == nil; e, currentElement = e+1, currentElement.next {
			switch e {
			case i:
				element1 = currentElement
			case j:
				element2 = currentElement
			}
		}
		element1.value, element2.value = element2.value, element1.value
	}
}

func (list *List) Insert(index int, values ...any) {
	if len(values) == 0 {
		return
	}
	if !list.withinRang(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	list.size += len(values)
	var beforeElement *element
	foundElement := list.first
	for e := 0; e != index; e, foundElement = e+1, foundElement.next {

		beforeElement = foundElement
	}
	if foundElement == list.first {
		oldNextElement := list.first
		for i, value := range values {
			newElement := &element{value: value}
			if i == 0 {
				list.first = newElement
			} else {
				beforeElement.next = newElement
			}
			beforeElement = newElement
		}
		beforeElement.next = oldNextElement
	} else {
		oldNextElement := beforeElement.next
		for _, value := range values {
			newElement := &element{value: value}
			beforeElement.next = newElement
			beforeElement = newElement
		}
		beforeElement.next = oldNextElement
	}

}

func (list *List) Set(index int, value any) {
	if !list.withinRang(index) {
		if index == list.size {
			list.Add(value)
		}
		return
	}
	element := list.first
	for e := 0; e != index; e, element = e+1, element.next {

	}
	element.value = value
}

func (list *List) Sort(comparator utils.Comparator) {

}

func (list *List) Empty() bool {
	//TODO implement me
	panic("implement me")
}

func (list *List) Size() int {
	//TODO implement me
	panic("implement me")
}

func (list *List) Clear() {
	//TODO implement me
	panic("implement me")
}

func (list *List) Values() []any {
	//TODO implement me
	panic("implement me")
}

func (list *List) String() string {
	//TODO implement me
	panic("implement me")
}

func (list *List) withinRang(index int) bool {
	return index >= 0 && index < list.size
}
