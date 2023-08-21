package list

import "fmt"

type ArrayList struct {
	total int
	size  int
	data  []*ArrayListKV
}

type ArrayListKV struct {
	k int
	v interface{}
}

func (a *ArrayListKV) String() string {
	return fmt.Sprintf("[%d:%v]", a.k, a.v)
}

func NewArrayListRoot() *ArrayList {
	root := new(ArrayList)
	return root
}

func NewArrayListRootWithSize(size int) *ArrayList {
	l := new(ArrayList)
	l.size = size
	l.data = make([]*ArrayListKV, 0, size)
	return l
}

func newArrayListKV(key int, v interface{}) *ArrayListKV {
	return &ArrayListKV{
		k: key,
		v: v,
	}
}

func (a *ArrayList) Get(key int) (interface{}, bool) {
	kv, _, ok := a.search(key)
	if !ok {
		return nil, false
	}
	return kv.v, ok
}

func (a *ArrayList) Set(key int, v interface{}) {

	kv, _, has := a.search(key)
	if !has {
		kv = newArrayListKV(key, v)
		a.total++
		a.data = append(a.data, kv)
		return
	}

	kv.v = v
	return
}

func (a *ArrayList) Delete(key int) {

	_, index, has := a.search(key)
	if !has {
		return
	}

	if index == 0 {
		_, _ = a.DeleteHead()
		return
	}

	if index == a.total-1 {
		_, _ = a.DeleteTail()
		return
	}

	a.data = append(a.data[:index], a.data[index+1:]...)
	a.total--

	return
}

// IsEmpty 是否为空
func (a *ArrayList) IsEmpty() bool {
	return a.total == 0
}

// DeleteTail 删除链表尾
func (a *ArrayList) DeleteTail() (interface{}, bool) {

	if a.IsEmpty() {
		return nil, false
	}

	var tailKV = a.data[a.total-1]

	a.total--

	if a.IsEmpty() {
		a.data = make([]*ArrayListKV, 0, a.size)
		return tailKV.v, true
	}

	a.data = a.data[:a.total]
	return tailKV.v, true
}

// DeleteHead 删除链表头
func (a *ArrayList) DeleteHead() (interface{}, bool) {

	if a.IsEmpty() {
		return nil, false
	}

	var headKV = a.data[0]

	a.total--

	if a.IsEmpty() {
		a.data = make([]*ArrayListKV, 0, a.size)
		return headKV.v, true
	}

	a.data = a.data[1:]

	return headKV.v, true
}

func (a *ArrayList) String() string {
	if a.IsEmpty() {
		return "empty ArrayList"
	}

	var listString = ""

	for i := 0; i < a.total; i++ {
		listString += fmt.Sprintf("-->%s", a.data[i])
	}

	return listString
}

func (a *ArrayList) search(key int) (*ArrayListKV, int, bool) {
	if a.IsEmpty() {
		return nil, 0, false
	}

	for i := 0; i < a.total; i++ {
		if a.data[i].k == key {
			return a.data[i], i, true
		}
	}

	return nil, 0, false
}

func (a *ArrayList) Search(key int) (interface{}, bool) {
	node, _, has := a.search(key)
	if !has {
		return nil, false
	}

	return node.v, true
}
