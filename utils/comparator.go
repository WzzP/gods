package utils

import "time"

// 比较工具函数
// if a > b 1
// if a = b 0
// if a < b -1
type Comparator func(a, b interface{}) int

func StringComparator(a, b any) int {
	s1 := a.(string)
	s2 := b.(string)

	//以长度的小的为遍历长度,保证不溢出
	min := len(s2)
	if len(s1) < len(s2) {
		min = len(s1)
	}
	diff := 0

	for i := 0; i < min && diff == 0; i++ {
		//将字符转为ASCII对应的整数,两者相减. 相等为0 大于1 s1 大 小于1 s2 大
		diff = int(s1[i]) - int(s2[i])
	}

	//如果前面的都相等,则比较长度
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	//如果s1长度大于s2,则返回1
	if diff < 0 {
		return -1
	}
	//如果s1长度小于s2,则返回-1
	if diff > 0 {
		return 1
	}
	//如果s1长度等于s2,则返回0
	return 0

}

// IntComparator provides a basic comparison on int
func IntComparator[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](a, b T) int {

	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

func TimeComparator(a, b any) int {
	aAsserted := a.(time.Time)
	bAsserted := b.(time.Time)
	switch {
	case aAsserted.After(bAsserted):
		return 1
	case aAsserted.Before(bAsserted):
		return -1
	default:
		return 0

	}
}
