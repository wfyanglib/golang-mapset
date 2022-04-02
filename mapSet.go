package mapSet

// 存在的问题：
// 1: 存储结构为map[interface{}]struct{},虽然元素加入到map中就是遍历时就是无序的,但这并不是真正的随机。

// 此接口分为线程安全,线程不安全两种实现
type MapSet interface {
	// Add给MapSet中添加一个元素
	Add(i interface{}) bool

	// 返回MapSet中的元素个数
	RetElementCount() int

	// 清空MapSet中的所有元素
	Clear()

	// 复制所有的键值克隆一个相同的MapSet
	Clone() MapSet

	// 给定一系列元素，判断这些元素是否都在MapSet中
	Contains(i ...interface{}) bool

	// 从MapSet中删除一个元素
	Remove(i interface{})

	// 随机返回MapSet中的一个元素,注意并不是弹出
	RandomReturn() interface{}

	// 把MapSet中的成员作为切片返回
	ToSlice() []interface{}

	// 判断两个MapSet是否相等,如果元素数量相等且两个MapSet中的元素都是一一对应则两个MapSet相等
	Equal(other MapSet) bool

	// 遍历MapSet中的每个元素,并对每个元素传递一个方法，如果传递的 func 返回 true，则停止迭代
	Each(func(interface{}) bool)

	// 返回MapSet的所有Key组成的字符串，可以指定sep为分隔字符
	String(sep string) string

	// MapSet中随机返回一个元素,并再MapSet中删除这个元素
	Pop() interface{}
}

// NewMapSet创建并返回一个空的MapSet
func NewMapSet(s ...interface{}) MapSet {
	set := newThreadSafeSet()
	for _, value := range s {
		set.Add(value)
	}
	return &set
}

// NewMapSetWith创建一个给定元素的新MapSet，相当于给MapSet初始赋值
func NewMapSetWith(elts ...interface{}) MapSet {
	return NewSetFromSlice(elts)
}

// NewSetFromSlice从给定的s切片中添加元素
func NewSetFromSlice(s []interface{}) MapSet {
	a := NewMapSet(s...)
	return a
}

// 返回一个空的MapSet
func NewThreadUnsafeSet() MapSet {
	set := newThreadUnsafeSet()
	return &set
}

// 根据参数s初始化一个MapSet
func NewThreadUnsafeSetFromSlice(s []interface{}) MapSet {
	a := NewThreadUnsafeSet()
	for _, item := range s {
		a.Add(item)
	}
	return a
}