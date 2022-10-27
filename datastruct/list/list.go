package list

import "fmt"

type List struct {
	head *listNode
	tail *listNode
	root *listNode
}

type listNode struct {
	k    string
	v    interface{}
	next *listNode
	pre  *listNode
}

func NewListRoot() *List {
	root := new(List)
	return root
}

func NewListRootWithInit(key string, v interface{}) *List {
	root := new(List)
	node := newListNode(key, v)
	root.head = node
	root.tail = node

	return root
}

func newListNode(key string, v interface{}) *listNode {
	return &listNode{
		k:    key,
		v:    v,
		next: nil,
		pre:  nil,
	}
}

func (l *List) init(key string, v interface{}) {
	node := newListNode(key, v)
	l.tail = node
	l.head = node
	l.root = node
}

func (l *List) Set(key string, v interface{}) {
	if l.head == nil {
		l.init(key, v)
		return
	}

	node, ok := l.Search(key)
	if ok {
		node.v = v
		return
	}

	node = newListNode(key, v)
	l.tail.next = node
	node.pre = l.tail
	l.tail = node
}

func (l *List) Delete(key string) {
	nextNode := l.root
	var deleteNode *listNode

	for nextNode != nil {
		if nextNode.k == key {
			deleteNode = nextNode
			break
		}
		nextNode = nextNode.next
	}

	// 找不到 return 掉
	if deleteNode == nil {
		return
	}

	// 头尾相等，找到的就是 头
	if l.head == l.tail {
		l.root = nil
		l.tail = nil
		l.head = nil
		return
	}

	// 如果是尾巴，把尾巴的前一个节点的next置为空
	if l.tail == deleteNode {
		deleteNode.pre.next = nil
		l.tail = deleteNode.pre
		return
	}

	if l.head == deleteNode {
		l.head = deleteNode.next
		l.root = deleteNode.next
		return
	}

	// 这边删的是中间的，非头，非尾
	deleteNext := deleteNode.next
	deleteNode.pre.next = deleteNext
	deleteNext.pre = deleteNode.pre
	deleteNode.next = nil

	return
}

func (l *List) Search(key string) (*listNode, bool) {
	nextNode := l.root

	for nextNode != nil {
		if nextNode.k == key {
			return nextNode, true
		}
		nextNode = nextNode.next
	}

	return nil, false
}

func (l *List) String() string {
	if l.head == nil {
		return "empty List"
	}

	var listString = ""

	nextNode := l.root

	for nextNode != nil {
		listString += fmt.Sprintf("-->[%v]", nextNode.v)
		nextNode = nextNode.next
	}

	return listString
}

func (l *List) Get(key string) (interface{}, bool) {
	node, ok := l.Search(key)
	if !ok {
		return nil, false
	}
	return node.v, ok
}