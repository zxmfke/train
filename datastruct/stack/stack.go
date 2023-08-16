package stack

import "github.com/zxmfke/train/datastruct/list"

type ImplementStack interface {
	Push(k int, v interface{})

	Pop() (interface{}, bool)

	Print() string
}

type ListStack struct {
	l     *list.List
	total int
}

func NewListStack() *ListStack {
	return &ListStack{l: list.NewListRoot()}
}

// Push O(1)
func (l *ListStack) Push(k int, v interface{}) {
	l.l.Set(k, v)
	l.total++
}

// Pop O(1)
func (l *ListStack) Pop() (interface{}, bool) {
	v, has := l.l.DeleteTail()
	if !has {
		return nil, false
	}

	l.total--

	return v, true
}

func (l *ListStack) Print() string {
	return l.l.String()
}
