package datastructure

import "math"

// MaxHeap max堆结构，也是一种平衡二叉树结构，这里使用数组实现
type MaxHeap struct {
	CAP   int
	size  int
	Array []int
}

// NewMaxHeap 构造函数，构造一个新的max堆
func NewMaxHeap(CAP int) *MaxHeap {
	var array []int = make([]int, CAP)
	return &MaxHeap{
		CAP:   CAP,
		size:  0,
		Array: array,
	}
}

// Insert 将新元素加入堆，并放到合适的位置
func (m *MaxHeap) Insert(elem int) {
	// 查看数组是否已满，满则扩容
	if m.size == m.CAP {
		var newArray []int = make([]int, m.CAP*2)
		for i := 0; i < m.CAP; i++ {
			newArray[i] = m.Array[i]
		}
		m.CAP *= 2
		m.Array = newArray
	}
	// 新元素放在末位再逐次判断升位
	m.Array[m.size] = elem
	m.size++
	m.heapifyUP()
}

// Poll 移除根节点并作为返回值返回，同时设定新的根节点，空堆则返回MinInt
func (m *MaxHeap) Poll() int {
	if m.size == 0 {
		return int(math.MinInt64)
	}
	var elem int = m.Array[0]
	m.Array[0] = m.Array[m.size-1]
	m.size--
	m.Array[m.size] = 0
	m.heapifyDown()
	return elem
}

// Peek 查看根节点，若为空堆则返回MinInt
func (m *MaxHeap) Peek() int {
	if m.size == 0 {
		return int(math.MinInt64)
	}
	return m.Array[0]
}

// swap 辅助方法，交换元素的位置
func (m *MaxHeap) swap(index1, index2 int) {
	var temp = m.Array[index1]
	m.Array[index1] = m.Array[index2]
	m.Array[index2] = temp
}

// heapifyUP 辅助方法，末位判断机制，末位升位到比父节点小为止
func (m *MaxHeap) heapifyUP() {
	var index int = m.size - 1
	for m.hasParent(index) && m.parent(index) < m.Array[index] {
		m.swap(getParentIndex(index), index)
		index = getParentIndex(index)
	}
}

// heapifyDown 辅助方法，首位判断机制，首位降位到比所有子节点大为止
func (m *MaxHeap) heapifyDown() {
	var index int = 0
	for m.hasLeftChild(index) {
		var largerChildIndex int = getLeftChildIndex(index)
		if m.hasRightChild(index) && m.rightChild(index) > m.leftChild(index) {
			largerChildIndex = getRightChildIndex(index)
		}
		// 将较大子节点与节点比较，子节点较大则更换位置，子节点小则完成降位
		if m.Array[index] < m.Array[largerChildIndex] {
			m.swap(index, largerChildIndex)
		} else {
			break
		}
		index = largerChildIndex
	}
}

// hasLeftChild 辅助方法，判断是否存在左子节点
func (m *MaxHeap) hasLeftChild(index int) bool {
	return getLeftChildIndex(index) < m.size
}

// hasRightChild 辅助方法，判断是否存在右子节点
func (m *MaxHeap) hasRightChild(index int) bool {
	return getRightChildIndex(index) < m.size
}

// hasParent 辅助方法，判断是否存在父节点
func (m *MaxHeap) hasParent(index int) bool {
	return getParentIndex(index) >= 0
}

// leftChild 辅助方法，获取左子节点的值
func (m *MaxHeap) leftChild(parentIndex int) int {
	return m.Array[getLeftChildIndex(parentIndex)]
}

// rightChild 辅助方法，获取右子节点的值
func (m *MaxHeap) rightChild(parentIndex int) int {
	return m.Array[getRightChildIndex(parentIndex)]
}

// parent 辅助方法，获取父节点的值
func (m *MaxHeap) parent(childIndex int) int {
	return m.Array[getParentIndex(childIndex)]
}

// getLeftChildIndex 辅助函数，获取左子节点在数组中的索引
func getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

// getRightChildIndex 辅助函数，获取右子节点在数组中的索引
func getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

// getParentIndex 辅助函数，获取父节点在数组中的索引
func getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}
