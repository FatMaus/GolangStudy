package datastructure

import "math"

// PriQueueNode 优先队列节点
type PriQueueNode struct {
	value    int
	priority int
	nextEle  *PriQueueNode
}

// newPriQueueNode 构造函数，构造一个新的优先队列节点
func newPriQueueNode(value, priority int) *PriQueueNode {
	return &PriQueueNode{
		value:    value,
		priority: priority,
		nextEle:  nil,
	}
}

// PriorityQueue 优先队列结构
type PriorityQueue struct {
	head *PriQueueNode
}

// NewPriorityQueue 构造函数，构造一个新的优先队列
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		head: nil,
	}
}

// Push 向优先队列中插入新元素，复杂度O(N)
func (p *PriorityQueue) Push(value, priority int) {
	var newNode *PriQueueNode = newPriQueueNode(value, priority)
	// 对于空队列，直接填入头节点
	if p.head == nil {
		p.head = newNode
		return
	}
	var cur *PriQueueNode = p.head
	if p.head.priority < priority {
		// 对于非空队列，头节点满足条件则进行替换，否则逐一对比
		newNode.nextEle = p.head
		p.head = newNode
	} else {
		// 定位到最后一个优先级高于新节点的节点，进行元素插入
		for cur.nextEle != nil && cur.nextEle.priority > priority {
			cur = cur.nextEle
		}
		newNode.nextEle = cur.nextEle
		cur.nextEle = newNode
	}
}

// Pop 移除优先级最高的节点，节点值作为返回值，空节点返回MinInt
func (p *PriorityQueue) Pop() int {
	if p.head == nil {
		return int(math.MinInt64)
	}
	var retValue int = p.head.value
	p.head = p.head.nextEle
	return retValue
}

// Peek 查看优先级最高的节点的值
func (p *PriorityQueue) Peek() int {
	return p.head.value
}

// IsEmpty 判断队列是否为空
func (p *PriorityQueue) IsEmpty() bool {
	return p.head == nil
}
