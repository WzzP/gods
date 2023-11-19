package containers

// EnumerableWithIndex 为容器提供函数操作, 可以通过索引来获取值
type EnumerableWithIndex interface {

	//Each 遍历该容器,传递的参数分别为索引和值
	Each(func(index int, value any))
	//Any 将容器每个元素传递给函数,如过改函数返回true则返回true
	Any(func(index int, value any) bool) bool
	//All 将容器每个元素传递给函数,如果所有函数结果都返回true则返回true
	All(func(index int, value any) bool) bool
	//Find 查找跟value相匹配的值,遇到返回true的则返回该元素的索引和值
	Find(func(index int, value any) bool) (int, any)
}

type EnumerableWithKey interface {
	Each(func(key any, value any))

	Any(func(key any, value any) bool) bool

	All(func(key any, value any) bool) bool

	Find(func(key any, value any) bool) (any, any)
}
