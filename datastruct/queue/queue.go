package queue

import "github.com/zxmfke/train/datastruct/list"

type ImplementQueue interface {
	Enqueue(k int, v interface{})

	Dequeue() interface{}
}

type ListQueue struct {
	l     *list.List
	total int
}

func NewListQueue() *ListQueue {
	return &ListQueue{l: list.NewListRoot()}
}

// Enqueue O(1)
func (l *ListQueue) Enqueue(k int, v interface{}) {
	l.l.Set(k, v)
	l.total++
}

// Dequeue O(1)
func (l *ListQueue) Dequeue() (interface{}, bool) {
	v, has := l.l.DeleteHead()
	if !has {
		return nil, false
	}

	l.total--

	return v, true
}

func (l *ListQueue) Print() string {
	return l.l.String()
}
