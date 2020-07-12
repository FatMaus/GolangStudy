package chainlistapply

// ChainListNode nodes for chainlist
type ChainListNode struct {
	value   int
	nextEle *ChainListNode
}

// ChainList a single chain list
type ChainList struct {
	head *ChainListNode
	tail *ChainListNode
	size int
}

// 在一个已排序的链表中，删除重复的元素，使每个元素只出现一次
func deleteDuplicates(head *ChainListNode) *ChainListNode {
	var cur *ChainListNode = head
	for cur != nil {
		if cur.nextEle != nil && cur.value == cur.nextEle.value {
			cur.nextEle = cur.nextEle.nextEle
		} else {
			cur = cur.nextEle
		}
	}
	return head
}

// 在一个已排序的链表中，删除所有重复的元素，只保留未重复过的元素
func removeDuplicates(head *ChainListNode) *ChainListNode {
	var (
		prevDel *ChainListNode
		cur     *ChainListNode
		delVal  int
	)
	if head == nil {
		return head
	}
	prevDel = &ChainListNode{value: 0}
	prevDel.nextEle = head
	cur = prevDel
	// 此处必须从prev开始，否则输出无改变
	for cur.nextEle != nil && cur.nextEle.nextEle != nil {
		if cur.nextEle.value == cur.nextEle.nextEle.value {
			delVal = cur.nextEle.value
			for cur.nextEle != nil && cur.nextEle.value == delVal {
				cur.nextEle = cur.nextEle.nextEle
			}
		} else {
			cur = cur.nextEle
		}
	}
	return prevDel.nextEle
}

// 反转单链表
func reverseChainList(head *ChainListNode) *ChainListNode {
	var (
		temp *ChainListNode
		prev *ChainListNode
		cur  *ChainListNode
	)
	cur = head
	for cur != nil {
		temp = cur.nextEle
		cur.nextEle = prev
		prev = cur
		cur = temp
	}
	return prev
}

// 在第m个元素到第n个元素之间反转链表
func reverseChainListBetween(head *ChainListNode, m int, n int) *ChainListNode {
	var (
		prev     *ChainListNode
		start    *ChainListNode
		temp     *ChainListNode
		linkNode *ChainListNode
		nextNode *ChainListNode
	)
	n++
	if head == nil {
		return head
	}
	start = &ChainListNode{value: 0}
	start.nextEle = head
	head = start
	for i := 0; i < m; i++ {
		prev = head
		head = head.nextEle
	}
	linkNode = head
	for j := m; j < n; j++ {
		if head == nil {
			break
		}
		temp = head.nextEle
		head.nextEle = nextNode
		nextNode = head
		head = temp
	}
	prev.nextEle = nextNode
	linkNode.nextEle = head
	return start.nextEle
}
