package orderedSlice

import (
	"fmt"
)

type OrderedSlice struct {
	count int // 总元素数量
	data  []*SliceNode
}

type SliceNode struct {
	K int
	V interface{}
}

func (s *SliceNode) String() string {
	return fmt.Sprintf("key:%d, value:%v", s.K, s.V)
}

func NewOrderedSlice() *OrderedSlice {
	slice := make([]*SliceNode, 0)

	return &OrderedSlice{
		count: 0,
		data:  slice,
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

	insertIndex, action := o.findKeyInsertIndex(0, o.count-1, k)

	fmt.Println(insertIndex)

	newSliceNode := &SliceNode{K: k, V: v}

	switch action {
	case Nothing:
		o.data = append(o.data, newSliceNode)
		o.count++
	case Update:
		o.data[insertIndex] = newSliceNode
	case MoveInsertToNext:
		o.data = append(o.data, newSliceNode)
		copy(o.data[insertIndex+1:], o.data[insertIndex:])
		o.data[insertIndex] = newSliceNode
		o.count++
	}

	return
}

func (o *OrderedSlice) SearchKey(k int) (*SliceNode, bool) {
	if o.IsEmpty() {
		return nil, false
	}

	node, _, has := o.findKey(0, o.count-1, k)

	return node, has
}

func (o *OrderedSlice) SearchKeyRange(keyStart, keyEnd int) []*SliceNode {

	var result = []*SliceNode{}

	if keyStart > keyEnd {
		return result
	}

	if o.IsEmpty() {
		return result
	}

	_, startIndex, has := o.findKey(0, o.count-1, keyStart)

	if has {
		result = append(result, o.data[startIndex])
	}

	if startIndex == o.count-1 {
		return result
	}

	for i := startIndex + 1; i < o.count; i++ {
		if o.data[i].K > keyEnd {
			break
		}

		result = append(result, o.data[i])
	}

	return result
}

// findKey 找到 index
// insert 调用就是找写入的 index
// delete/search 调用就是找具体 key 对应的 index
func (o *OrderedSlice) findKey(start, end int, key int) (*SliceNode, int, bool) {

	middleIndex := o.findMiddleIndex(start, end)

	if middleIndex == start || middleIndex == end {
		return o.data[middleIndex], middleIndex, o.data[middleIndex].K == key
	}

	if o.data[middleIndex].K > key {
		return o.findKey(start, middleIndex-1, key)
	}

	if o.data[middleIndex].K == key {
		return o.data[middleIndex], middleIndex, true
	}

	return o.findKey(middleIndex+1, end, key)
}

type insertAction int

const (
	Nothing insertAction = iota
	Update
	MoveInsertToNext
)

// findKeyInsertIndex 找到插入的index
func (o *OrderedSlice) findKeyInsertIndex(start, end int, key int) (int, insertAction) {

	middleIndex := o.findMiddleIndex(start, end)

	if o.data[middleIndex].K == key {
		return middleIndex, Update
	}

	if middleIndex == start {
		if o.data[end].K > key && o.data[middleIndex].K < key {
			return middleIndex + 1, MoveInsertToNext
		}

		return middleIndex + 1, Nothing
	}

	if middleIndex == end {
		if o.data[start].K < key && o.data[middleIndex].K < key {
			return middleIndex + 1, MoveInsertToNext
		}

		return middleIndex + 1, Nothing
	}

	if o.data[middleIndex].K > key {
		return o.findKeyInsertIndex(start, middleIndex-1, key)
	}

	return o.findKeyInsertIndex(middleIndex+1, end, key)
}

// findMiddleIndex 找 slice 中间的下标
func (o *OrderedSlice) findMiddleIndex(start, end int) int {

	if start >= end {
		return end
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

	node, index, has := o.findKey(0, o.count-1, k)
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
