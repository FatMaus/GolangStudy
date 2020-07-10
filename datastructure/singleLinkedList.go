package datastructure

import (
	"fmt"
	"reflect"
)

// SingleLinkNode 单链表节点
type SingleLinkNode struct {
	value   int
	nextEle *SingleLinkNode
}

// 构造函数，构造一个单链表节点
func newSingleLinkNode(value int) *SingleLinkNode {
	return &SingleLinkNode{
		value:   value,
		nextEle: nil,
	}
}

// SingleLinkList 单链表结构
type SingleLinkList struct {
	head *SingleLinkNode
	tail *SingleLinkNode
	size int
}

// NewSingleLinkList 构造函数，构造一个单链表
func NewSingleLinkList() *SingleLinkList {
	return &SingleLinkList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Append 向最后一位增加元素
func (s *SingleLinkList) Append(num int) {
	var newNode = newSingleLinkNode(num)
	// 对于空链表，直接加入。非空链表加到最后一位并建立链接
	if s.size == 0 {
		s.head = newNode
		s.tail = s.head.nextEle
	} else if s.tail == nil {
		s.head.nextEle = newNode
		s.tail = newNode
	} else {
		s.tail.nextEle = newNode
		s.tail = newNode
	}
	s.size++
}

// Insert 按索引插入元素
func (s *SingleLinkList) Insert(position, num int) {
	// 若索引超出范围，直接返回
	if position > s.size || position < 0 {
		fmt.Println("out range")
		return
	}
	var newNode = newSingleLinkNode(num)
	// 首位添加
	if position == 0 {
		newNode.nextEle = s.head
		s.head = newNode
		if s.tail == nil {
			s.tail = newNode.nextEle
		}
		s.size++
	} else if position == s.size {
		// 末尾添加
		s.Append(num)
	} else {
		// 其余位置添加
		var prev = s.head
		for i := 0; i < position-1; i++ {
			prev = prev.nextEle
		}
		newNode.nextEle = prev.nextEle
		prev.nextEle = newNode
		s.size++
	}
}

// GetIndex 按值查找索引
func (s *SingleLinkList) GetIndex(num int) int {
	var cur = s.head
	var ret int = -1
	for i := 0; i < s.size; i++ {
		if cur.value == num {
			ret = i
			break
		} else {
			cur = cur.nextEle
		}
	}
	// 若未找到则返回-1
	return ret
}

// Delete 按值删除元素
func (s *SingleLinkList) Delete(num int) {
	if s.head.value == num {
		// 移除头元素
		s.head = s.head.nextEle
		s.size--
		// 若成为空链表，则保证头尾相同
		if s.size == 0 {
			s.tail = s.head
		}
	} else {
		// 非头元素的移除
		var cur = s.head
		var prev = s.head
		// 逐一查找
		for prev != nil && cur != nil {
			if cur.value == num {
				// 尾元素直接断尾
				if cur == s.tail {
					s.tail = prev
				}
				// 非尾元素注意链接不断
				prev.nextEle = cur.nextEle
				s.size--
				return
			}
			// 若不满足条件则继续逐一匹配
			prev = cur
			cur = cur.nextEle
		}
	}
}

// Update 更新链的值
func (s *SingleLinkList) Update(oldNum, newNum int) int {
	// 从头开始，逐一匹配
	var cur = s.head
	var ret int = -1
	for i := 0; i < s.size; i++ {
		// 匹配到更新值，返回索引
		if cur.value == oldNum {
			cur.value = newNum
			ret = i
			break
		} else {
			cur = cur.nextEle
		}
	}
	// 无对应元素则返回-1
	return ret
}

// Display 遍历元素并打印
func (s *SingleLinkList) Display() {
	var cur = s.head
	for i := 0; i < s.size; i++ {
		fmt.Print(cur.value, ", ")
		cur = cur.nextEle
	}
}

// IsCircle 判断链表是否是循环链表，设定两个变量，行进速度不同，若某一刻等同，则闭环
func (s *SingleLinkList) IsCircle() bool {
	var node1 = s.head
	var node2 = s.head
	var isCircle bool = false
	for i := 0; i < s.size; i++ {
		if reflect.DeepEqual(node1, node2) {
			isCircle = true
			break
		} else if node2 == nil {
			break
		} else {
			node1 = node1.nextEle
			node2 = node2.nextEle.nextEle
		}
	}
	return isCircle
}
