package orderedSlice

import (
	"fmt"
)

type OrderedSlice struct {
	count   int // 总元素数量
	data    []*SliceNode
	sortWay Sorter
}

type SliceNode struct {
	K int
	V interface{}
}

func (s *SliceNode) String() string {
	return fmt.Sprintf("key:%d, value:%v", s.K, s.V)
}

func NewOrderedSlice(sortWay OrderedListSortedWay) *OrderedSlice {
	slice := make([]*SliceNode, 0)

	return &OrderedSlice{
		count:   0,
		data:    slice,
		sortWay: newSorter(sortWay),
	}
}

func (o *OrderedSlice) String() string {

	fmtString := "orderedSlice : "

	if o.IsEmpty() {
		return fmtString + "is empty!"
	}

	for i := 0; i < o.count; i++ {
		fmtString += fmt.Sprintf("[%d:{%s}] ", i+1, o.data[i])
	}

	return fmtString
}

func (o *OrderedSlice) IsEmpty() bool {
	return o.count == 0
}

func (o *OrderedSlice) Insert(k int, v interface{}) {
	if o.IsEmpty() {
		o.data = append(o.data, &SliceNode{
			K: k,
			V: v,
		})
		o.count++
		return
	}

	node, insertIndex, has := o.findKey(0, o.count, k)
	if has {
		node.V = v
		return
	}

	newSliceNode := &SliceNode{K: k, V: v}

	o.data = append(o.data, newSliceNode)

	copy(o.data[insertIndex+1:], o.data[insertIndex:])
	o.data[insertIndex] = newSliceNode
	o.count++

	return
}

func (o *OrderedSlice) SearchKey(k int) (*SliceNode, bool) {
	if o.IsEmpty() {
		return nil, false
	}

	node, _, has := o.findKey(0, o.count, k)

	return node, has

}

// findKey 找到 index
// insert 调用就是找写入的 index
// delete/search 调用就是找具体 key 对应的 index
func (o *OrderedSlice) findKey(start, end int, key int) (*SliceNode, int, bool) {

	middleIndex := o.findMiddleIndex(start, end)

	if middleIndex == start {
		if !o.sortWay.Less(o.data[middleIndex].K, key) {
			middleIndex += 1
		}

		return o.data[middleIndex], middleIndex, o.data[middleIndex].K == key
	}

	if o.sortWay.Less(o.data[middleIndex].K, key) {
		return o.findKey(start, middleIndex, key)
	}

	if o.data[middleIndex].K == key {
		return o.data[middleIndex], middleIndex, true
	}

	return o.findKey(middleIndex, end, key)
}

// findMiddleIndex 找 slice 中间的下标
func (o *OrderedSlice) findMiddleIndex(start, end int) int {

	if start >= end {
		return start
	}

	middle := (end - start) / 2

	if middle < 1 {
		return start
	}

	return middle
}

// Delete 删除 key
func (o *OrderedSlice) Delete(k int) (*SliceNode, int) {
	if o.IsEmpty() {
		return nil, -1
	}

	node, index, has := o.findKey(0, o.count, k)
	if !has {
		return nil, -1
	}

	fmt.Println("index delete : ", index)

	if index == o.count-1 {
		o.data = o.data[:o.count-1]
		o.count--
		return node, index
	}

	copy(o.data[index:], o.data[index+1:])
	o.count--
	return node, index
}
