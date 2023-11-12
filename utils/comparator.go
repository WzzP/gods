package utils

// 比较工具函数
// if a > b 1
// if a = b 0
// if a < b -1
type Comparator func(a, b interface{}) int
