package containers

// Container 容器接口
type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []any
	String() string
}
