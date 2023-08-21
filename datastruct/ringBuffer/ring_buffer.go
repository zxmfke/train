package ringBuffer

import (
	"errors"
	"fmt"
	"github.com/zxmfke/train/datastruct/list"
	"strings"
)

var (
	errFull  = errors.New("ring buffer is full")
	errEmpty = errors.New("ring buffer is empty")
)

type RingBuffer struct {
	head      int
	tail      int
	arrayList *list.ArrayList
	size      int
	total     int
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		head:      0,
		tail:      0,
		arrayList: list.NewArrayListRootWithSize(size),
		size:      size,
	}
}

func (r *RingBuffer) IsFull() bool {
	if r.total == 0 {
		return false
	}

	return r.head == r.tail
}

func (r *RingBuffer) Get() (interface{}, error) {

	if r.arrayList.IsEmpty() {
		return nil, errEmpty
	}

	r.total--
	v, _ := r.arrayList.DeleteHead()
	r.TailStepIn()
	return v, nil
}

func (r *RingBuffer) Set(key int, v interface{}) error {
	if r.IsFull() {
		return errFull
	}

	r.total++
	r.arrayList.Set(key, v)
	r.HeadStepIn()
	return nil
}

func (r *RingBuffer) HeadStepIn() {
	r.head++
	if r.head == r.size {
		r.head = 0
	}
}

func (r *RingBuffer) TailStepIn() {
	r.tail++
	if r.tail == r.size {
		r.tail = 0
	}
}

func (r *RingBuffer) String() string {
	if r.arrayList.IsEmpty() {
		return "empty ring buffer"
	}

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("\nring buffer size : %d\n", r.size))
	sb.WriteString(fmt.Sprintf("ring buffer used : %d\n", r.total))
	sb.WriteString(fmt.Sprintf("ring buffer last : %d\n", r.size-r.total))
	sb.WriteString(fmt.Sprintf("ring buffer data : %s\n", r.arrayList))

	return sb.String()
}
