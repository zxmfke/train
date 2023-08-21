package list

import "fmt"

type LinkedList struct {
	head *LNode
	tail *LNode
	root *LNode
}

type LNode struct {
	k    int
	v    interface{}
	next *LNode
	pre  *LNode
}

func NewLinkedListRoot() *LinkedList {
	root := new(LinkedList)
	return root
}

func NewLinkedListRootWithInit(key int, v interface{}) *LinkedList {
	l := new(LinkedList)
	node := newLinkedListNode(key, v)
	l.head = node
	l.tail = node
	l.root = node

	return l
}

func newLinkedListNode(key int, v interface{}) *LNode {
	return &LNode{
		k:    key,
		v:    v,
		next: nil,
		pre:  nil,
	}
}

func (l *LinkedList) init(key int, v interface{}) {
	node := newLinkedListNode(key, v)
	l.tail = node
	l.head = node
	l.root = node
}

func (l *LinkedList) Get(key int) (interface{}, bool) {
	node, ok := l.search(key)
	if !ok {
		return nil, false
	}
	return node.v, ok
}

func (l *LinkedList) Set(key int, v interface{}) {
	if l.head == nil {
		l.init(key, v)
		return
	}

	node, ok := l.search(key)
	if ok {
		node.v = v
		return
	}

	node = newLinkedListNode(key, v)
	l.tail.next = node
	node.pre = l.tail
	l.tail = node
}

func (l *LinkedList) Delete(key int) {
	nextNode := l.root
	var deleteNode *LNode

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
		l.DeleteTail()
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

// IsEmpty 是否为空
func (l *LinkedList) IsEmpty() bool {
	return l.head == nil
}

// DeleteTail 删除链表尾
func (l *LinkedList) DeleteTail() (interface{}, bool) {

	if l.IsEmpty() {
		return nil, false
	}

	var tailNodeValue = l.tail.v

	if l.tail == l.head {
		l.head = nil
		l.root = nil
		l.tail = nil
		return tailNodeValue, true
	}

	l.tail.pre.next = nil
	l.tail = l.tail.pre
	return tailNodeValue, true
}

// DeleteHead 删除链表头
func (l *LinkedList) DeleteHead() (interface{}, bool) {

	if l.IsEmpty() {
		return nil, false
	}

	var headNodeValue = l.head.v

	if l.tail == l.head {
		l.head = nil
		l.root = nil
		l.tail = nil
		return headNodeValue, true
	}

	l.head = l.head.next
	l.root = l.head.next
	return headNodeValue, true
}

func (l *LinkedList) String() string {
	if l.IsEmpty() {
		return "empty LinkedList"
	}

	var listString = ""

	nextNode := l.root

	for nextNode != nil {
		listString += fmt.Sprintf("-->[%v]", nextNode.v)
		nextNode = nextNode.next
	}

	return listString
}

func (l *LinkedList) search(key int) (*LNode, bool) {
	nextNode := l.root

	for nextNode != nil {
		if nextNode.k == key {
			return nextNode, true
		}
		nextNode = nextNode.next
	}

	return nil, false
}

func (l *LinkedList) Search(key int) (interface{}, bool) {
	node, has := l.search(key)
	if !has {
		return nil, false
	}

	return node.v, true
}
