package datastructure

import (
	"hash/crc32"
	"math"
)

// HashNode 哈希节点
type HashNode struct {
	key     string
	value   int
	nextEle *HashNode
}

// newHashNode 构造函数，创建新节点
func newHashNode(key string, value int) *HashNode {
	return &HashNode{
		key:     key,
		value:   value,
		nextEle: nil,
	}
}

// HashTable 哈希表结构
type HashTable struct {
	CAP   int
	size  int
	Array []*HashNode
}

// NewHashTable 构造函数，创建一个新哈希表
func NewHashTable(CAP int) *HashTable {
	// 创建新数组，默认值nil
	var Array = make([]*HashNode, CAP)
	return &HashTable{
		CAP:   CAP,
		size:  0,
		Array: Array,
	}
}

// getIndex 辅助方法 根据键计算哈希值，并返回存储索引
func (h *HashTable) getIndex(key string) (index int) {
	// 计算哈希值，并保证是无符号整型
	var hashCode = int(crc32.ChecksumIEEE([]byte(key)))
	if hashCode < 0 {
		hashCode = -hashCode
	}
	index = hashCode % h.CAP
	return index
}

// Add 向哈希表增加元素
func (h *HashTable) Add(key string, value int) {
	// 获取存储索引
	var ArrIndex = h.getIndex(key)
	var head = h.Array[ArrIndex]
	// 查找键是否已存在
	for head != nil {
		// 若键已存在则替换值
		if head.key == key {
			head.value = value
			return
		}
		head = head.nextEle
	}
	// 表中不存在该键则进行添加
	head = h.Array[ArrIndex]
	var newNode = newHashNode(key, value)
	newNode.nextEle = head
	h.Array[ArrIndex] = newNode
	h.size++
	// 为优化查找速率，存储量过多时应扩容
	if h.size-h.CAP*4/5 >= 0 {
		h.expansion()
	}
}

// 辅助方法，给哈希表内置数组扩容
func (h *HashTable) expansion() {
	var backup = h.Array
	h.CAP *= 2
	h.size = 0
	h.Array = make([]*HashNode, h.CAP)
	for _, node := range backup {
		for node != nil {
			h.Add(node.key, node.value)
			node = node.nextEle
		}
	}
}

// Get 根据键找出对应的值，若键不存在则返回一个MinInt
func (h *HashTable) Get(key string) int {
	var ArrIndex = h.getIndex(key)
	var head = h.Array[ArrIndex]
	for head != nil {
		if head.key == key {
			return head.value
		}
		head = head.nextEle
	}
	return int(math.MinInt64)
}

// Remove 根据键移除键值对，若不存在该键则返回MinInt
func (h *HashTable) Remove(key string) int {
	var ArrIndex = h.getIndex(key)
	var head = h.Array[ArrIndex]
	var prev *HashNode = nil
	for head != nil {
		if head.key == key {
			break
		}
		prev = head
		head = head.nextEle
	}
	// 若未找到该键，返回MinInt
	if head == nil {
		return int(math.MinInt64)
	}
	// 找到了对应的键，移除键值对
	h.size--
	if prev != nil {
		prev.nextEle = head.nextEle
	} else {
		// prev为nil则表示首位就是目标键值对，下一位移入数组顶替位置
		h.Array[ArrIndex] = head.nextEle
	}
	return head.value
}

// GetSize 返回哈希表元素个数
func (h *HashTable) GetSize() int {
	return h.size
}

// IsEmpty 查看哈希表是否为空
func (h *HashTable) IsEmpty() bool {
	return h.size == 0
}
