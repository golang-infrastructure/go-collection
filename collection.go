package collection

import (
	"encoding/json"
	compare_anything "github.com/golang-infrastructure/go-compare-anything"
	"github.com/golang-infrastructure/go-iterator"
)

// Collection 集合的定义
type Collection[T any] interface {

	// Equalsable 能够比较两个集合是否相等
	compare_anything.Equalsable[T]

	// Iterable 集合都是可迭代的，可以返回一个迭代器来遍历集合
	iterator.Iterable[T]

	// Marshaler 集合是可以被JSON序列化和反序列化的
	json.Marshaler
	json.Unmarshaler

	// Add 往集合中增加一个元素
	Add(value T)

	// AddAll 往集合中一次增加多个元素
	AddAll(values []T)

	// Clear 清空集合
	Clear()

	// Contains 判断集合是否包含给定的元素
	Contains(value T) bool

	// ContainsAll 判断集合是否包含全部给定的元素
	ContainsAll(values []T) bool

	// ContainsAny 判断集合是否包含给定的一批元素中的任意一个
	ContainsAny(value []T) bool

	// String 集合转为字符串，注意这个方法不一定是可逆的，通常只是为了把集合打印日志之类的
	String() string

	// IsEmpty 判断集合是否为空
	IsEmpty() bool

	// Remove 从集合中移除给定元素
	Remove(value T)

	// RemoveAll 从集合中移除给定的一批元素
	RemoveAll(value []T)

	// RetainAll 当前集合减去其它集合，求差集
	RetainAll(collection Collection[T])

	// Union 返回两个集合的并集
	Union(collection Collection[T]) Collection[T]

	// Intersection 返回两个集合的交集
	Intersection(collection Collection[T]) Collection[T]

	// Size 返回集合的大小，如果集合为空则返回0
	Size() int

	// ToSlice 将集合转为切片
	ToSlice() []T

	// Copy 复制集合本身为一个新的集合
	Copy() Collection[T]
}
