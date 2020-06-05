package datastructure

import "math"

// ArrQueue 数组队列结构
type ArrQueue struct {
	CAP   int
	front int
	rear  int
	size  int
	Array []int
}

// NewArrQueue 构造函数，创建一个新数组队列
func NewArrQueue(CAP int) *ArrQueue {
	var array []int = make([]int, CAP)
	return &ArrQueue{
		CAP:   CAP,
		front: 0,
		rear:  0,
		size:  0,
		Array: array,
	}
}

// IsFull 检查队列是否装满
func (a *ArrQueue) IsFull() bool {
	return a.size == a.CAP
}

// IsEmpty 检查队列是否为空
func (a *ArrQueue) IsEmpty() bool {
	return a.size == 0
}

// Enqueue 向队列末位增加元素，若队列已满则返回false
func (a *ArrQueue) Enqueue(elem int) bool {
	if a.IsFull() {
		return false
	}
	a.Array[a.rear] = elem
	// 取余操作实现循环队列
	a.rear = (a.rear + 1) % a.CAP
	a.size++
	return true
}

// Dequeue 从队列首位提取元素，取出的元素不再存放于队列。空队列则返回MinInt
func (a *ArrQueue) Dequeue() int {
	if a.IsEmpty() {
		return int(math.MinInt64)
	}
	var ret int = a.Array[a.front]
	// 取余操作实现循环队列
	a.front = (a.front + 1) % a.CAP
	a.size--
	return ret
}

// Peek 查看队列首位元素，若为空队列则返回MinInt
func (a *ArrQueue) Peek() int {
	if a.IsEmpty() {
		return int(math.MinInt64)
	}
	return a.Array[a.front]
}

// QueueNode 队列节点，用于链表队列
type QueueNode struct {
	value   int
	nextEle *QueueNode
}

// newQueueNode 构造函数，构造一个队列节点
func newQueueNode(value int) *QueueNode {
	return &QueueNode{
		value:   value,
		nextEle: nil,
	}
}

// LinkQueue 链表队列结构
type LinkQueue struct {
	front *QueueNode
	rear  *QueueNode
}

// NewLinkQueue 构造函数，构造一个新的链表队列
func NewLinkQueue() *LinkQueue {
	return &LinkQueue{
		front: nil,
		rear:  nil,
	}
}

// Enqueue 向链表队列后端增加元素
func (l *LinkQueue) Enqueue(value int) {
	var newNode *QueueNode = newQueueNode(value)
	// 若队列为空时注意与前端取出位关联
	if l.rear == nil {
		l.rear = newNode
		l.front = l.rear
	} else {
		// 非空队列直接加入
		l.rear.nextEle = newNode
		l.rear = newNode
	}
}

// Dequeue 提取链表队列前端元素，提取后的元素不再存于队列中,队列为空返回MinInt
func (l *LinkQueue) Dequeue() int {
	if l.front == nil {
		return int(math.MinInt64)
	}
	var frontNode *QueueNode = l.front
	l.front = frontNode.nextEle
	// 若将队列提取空，则需要首尾一致
	if l.front == nil {
		l.rear = nil
	}
	return frontNode.value
}

// Peek 查看链表队列前端元素，队列为空则返回MinInt
func (l *LinkQueue) Peek() int {
	if l.front == nil {
		return int(math.MinInt64)
	}
	return l.front.value
}
