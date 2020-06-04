package datastructure

import (
	"fmt"
	"math"
)

// ArrStack 数组堆栈
type ArrStack struct {
	top   int
	CAP   int
	stack []int
}

// NewArrStack 构造函数，构造一个新数组堆栈
func NewArrStack(CAP int) *ArrStack {
	var arr = make([]int, CAP)
	return &ArrStack{
		top:   -1,
		CAP:   CAP,
		stack: arr,
	}
}

// Push 向堆栈添加元素，满溢则报错
func (a *ArrStack) Push(value int) bool {
	if a.top >= a.CAP-1 {
		fmt.Println("overflow")
		return false
	}
	a.top++
	a.stack[a.top] = value
	return true
}

// Pop 从堆栈中移除元素并作为返回值，若无元素则报错并返回MinInt
func (a *ArrStack) Pop() int {
	if a.top < 0 {
		fmt.Println("empty")
		return int(math.MinInt64)
	}
	var ret int = a.stack[a.top]
	a.top--
	return ret
}

// Peek 读取堆栈顶部元素，空栈则报错并返回MinInt
func (a *ArrStack) Peek() int {
	if a.top < 0 {
		fmt.Println("empty")
		return int(math.MinInt64)
	}
	return a.stack[a.top]
}

// IsEmpty 检查堆栈是否为空
func (a *ArrStack) IsEmpty() bool {
	return a.top < 0
}

// StackNode 堆栈节点，用于链表堆栈
type StackNode struct {
	value   int
	nextEle *StackNode
}

// newStackNode 构造函数，构造一个新堆栈节点
func newStackNode(value int) *StackNode {
	return &StackNode{
		value:   value,
		nextEle: nil,
	}
}

// LinkStack 链表堆栈结构
type LinkStack struct {
	top *StackNode
}

// NewLinkStack 构造函数，构造一个新链表堆栈
func NewLinkStack() *LinkStack {
	return &LinkStack{
		top: nil,
	}
}

// Push 向链表堆栈中装入元素
func (l *LinkStack) Push(value int) {
	var newNode *StackNode = newStackNode(value)
	// 空堆栈直接装入
	if l.top == nil {
		l.top = newNode
	} else {
		// 非空堆栈则进行顶元素替换
		var temp = l.top
		l.top = newNode
		newNode.nextEle = temp
	}
}

// Pop 从堆栈中移除元素并作为返回值显示，若堆栈为空则返回MinInt
func (l *LinkStack) Pop() int {
	if l.top == nil {
		fmt.Println("empty")
		return int(math.MinInt64)
	}
	var popped int = l.top.value
	l.top = l.top.nextEle
	return popped
}

// Peek 读取堆栈顶部元素，若堆栈为空则返回MinInt
func (l *LinkStack) Peek() int {
	if l.top == nil {
		fmt.Println("empty")
		return int(math.MinInt64)
	}
	return l.top.value
}

// IsEmpty 判断堆栈是否为空
func (l *LinkStack) IsEmpty() bool {
	return l.top == nil
}
