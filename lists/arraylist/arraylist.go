package arraylist

import (
	"fmt"
	"gods/lists"
	"gods/utils"
	"strings"
)

// 将接口和结构体绑定，指定结构体要实现接口方法
var _ lists.List = (*List)(nil)

const (
	growthFactor = float32(2.0)  //增长超过100%
	shrinkFactor = float32(0.25) //当容量为25%时进行收缩，0%不收缩
)

// List 使用slice保存元素
type List struct {
	elements []interface{}
	size     int
}

func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Get 获取元素
func (list *List) Get(index int) (any, bool) {

	if !list.withinRang(index) {
		return nil, false
	}
	return list.elements[index], true
}

// Remove 删除元素
func (list *List) Remove(index int) {
	if !list.withinRang(index) {
		return
	}

	list.elements[index] = nil //清除引用
	//向左移动
	copy(list.elements[index:], list.elements[index+1:list.size]) //向左移动
	list.size--

	//动态收缩
	list.shrink()
}

// Add 添加元素
func (list *List) Add(values ...any) {
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

// Contains 判断元素是否存在
func (list *List) Contains(values ...any) bool {
	for _, searchValue := range values {
		found := false
		for index := 0; index < list.size; index++ {
			if list.elements[index] == searchValue {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Swap 交换index1 和 index2 元素的数据
func (list *List) Swap(index1, index2 int) {
	if list.withinRang(index1) && list.withinRang(index2) {
		list.elements[index1], list.elements[index2] = list.elements[index2], list.elements[index1]
	}
}

// Insert 插入元素
func (list *List) Insert(index int, values ...any) {
	if !list.withinRang(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}
	l := len(values)
	list.growBy(l)
	list.size += l
	//算是slice的一个隐藏特性，slice切片出来的数组还是指向原切片的数组，所以当修改新切片的数据，原切片也同样也会发生变化
	// s := make(int[], 100)
	// s2 := s[:]
	// s[1] = 100
	// s2[2] = 100
	// s[2] == 100
	// 但是随着不断的append，由于扩容后会产生新的数组，这时就不在指向同一个数组了
	copy(list.elements[index+l:], list.elements[index:list.size-l])
	copy(list.elements[index:], values)
}

// Set 指定index值为value
func (list *List) Set(index int, value any) {
	if !list.withinRang(index) {
		if index == list.size {
			list.Add(value)
		}
		//超过当前长度就不在设置
		return
	}
	list.elements[index] = value
}

// Empty 判断是否为空
func (list *List) Empty() bool {
	return list.size == 0
}

// Size 获取数组长度
func (list *List) Size() int {
	return list.size
}

// Clear 清空
func (list *List) Clear() {
	list.size = 0
	list.elements = []any{}
}

// Values 返回数组
func (list *List) Values() []any {
	newElements := make([]any, list.size)
	copy(newElements, list.elements)
	return newElements
}

func (list *List) String() string {
	str := "ArrayList\n"
	var values []string
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// 判断是否越界
func (list *List) withinRang(index int) bool {
	return index >= 0 && index < list.size
}

// 收缩slice，当数组长度小于slice 容量的25%时，进行收缩调整
func (list *List) shrink() {
	if shrinkFactor == 0.0 {
		return
	}

	currentCapacity := cap(list.elements)
	//siz 20  0.25 * 40
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}

// resize 调整slice容量
func (list *List) resize(cap int) {
	newElements := make([]interface{}, cap)

	copy(newElements, list.elements)
	list.elements = newElements
}

// 判断slice是否需要扩容，扩容的大小为当前容量的两倍
func (list *List) growBy(n int) {

	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

// IndexOf 返回value所处的索引位置
func (list *List) IndexOf(value any) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.elements {
		if element == value {
			return index
		}
	}
	return -1
}

func (list *List) Sort(comparator utils.Comparator) {
	if list.size < 2 {
		return
	}
	utils.Sort(list.elements[:list.size], comparator)

}
