package hashTable

import (
	"errors"
	"github.com/zxmfke/train/datastruct/list"
)

var (
	ErrHSKeyOutOfRange       = errors.New("hash func key out of range")
	ErrHSMapValueSizeInvalid = errors.New("hsMapValue size cant less than 10")
)

type HSTable struct {
	data     []*list.LinkedList
	hashFunc HashFun
	max      int
}

type HashFun func(int) int

func NewHSTable(hf HashFun, size int) (*HSTable, error) {
	if size < 10 {
		return nil, ErrHSMapValueSizeInvalid
	}

	hsm := new(HSTable)
	hsm.data = make([]*list.LinkedList, size)
	for i := 0; i < size; i++ {
		hsm.data[i] = list.NewLinkedListRoot()
	}

	hsm.max = size

	hsm.hashFunc = defaultHashFunc
	if hf != nil {
		hsm.hashFunc = hf
	}

	return hsm, nil
}

func defaultHashFunc(key int) int {
	return key % 10
}

func (h *HSTable) Set(key int, value interface{}) error {
	hsKey := h.hashFunc(key)

	if hsKey >= h.max {
		return ErrHSKeyOutOfRange
	}

	h.data[hsKey].Set(key, value)

	return nil
}

func (h *HSTable) Get(key int) (interface{}, bool) {
	hsKey := h.hashFunc(key)

	if hsKey >= h.max {
		return nil, false
	}

	return h.data[hsKey].Get(key)
}

func (h *HSTable) Delete(key int) error {
	hsKey := h.hashFunc(key)

	if hsKey >= h.max {
		return ErrHSKeyOutOfRange
	}

	h.data[hsKey].Delete(h.hashFunc(key))

	return nil
}
