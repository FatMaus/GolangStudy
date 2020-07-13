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

// 合并升序链表，形成一个新的升序链表
func mergeLists(l1 *ChainListNode, l2 *ChainListNode) *ChainListNode {
	var (
		// 在首节点可能被替换时，使用哑节点缓冲
		dummy *ChainListNode = &ChainListNode{value: 0}
		head  *ChainListNode = dummy
	)
	// 比较元素大小，逐个合并
	for l1 != nil && l2 != nil {
		if l1.value > l2.value {
			head.nextEle = l2
			l2 = l2.nextEle
		} else {
			head.nextEle = l1
			l1 = l1.nextEle
		}
		head = head.nextEle
	}
	// 处理剩余节点
	for l1 != nil {
		head.nextEle = l1
		l1 = l1.nextEle
		head = head.nextEle
	}
	for l2 != nil {
		head.nextEle = l2
		l2 = l2.nextEle
		head = head.nextEle
	}
	return dummy.nextEle
}

// 给定一个值x，使得链表中小于x的节点都位于链表前端(左)，大于x的节点都位于链表后端(右)
func partition(head *ChainListNode, x int) *ChainListNode {
	var (
		dummyHead *ChainListNode = &ChainListNode{value: 0}
		dummy     *ChainListNode = &ChainListNode{value: 0}
		tail      *ChainListNode
		temp      *ChainListNode
	)
	if head == nil {
		return head
	}
	tail = dummy // 确定新链表起点
	dummyHead.nextEle = head
	head = dummyHead
	for head.nextEle != nil {
		if head.nextEle.value < x {
			head = head.nextEle
		} else {
			// 从原链表中移除
			temp = head.nextEle
			head.nextEle = head.nextEle.nextEle
			// 加入新链表
			tail.nextEle = temp
			tail = tail.nextEle
		}
	}
	// 链表对接
	tail.nextEle = nil
	head.nextEle = dummy.nextEle
	return dummyHead.nextEle
}

// ChainListMergeSort 链表归并排序
func ChainListMergeSort(head *ChainListNode) *ChainListNode {
	var (
		mid   *ChainListNode
		tail  *ChainListNode
		left  *ChainListNode
		right *ChainListNode
		ret   *ChainListNode
	)
	if head == nil || head.nextEle == nil {
		return head
	}
	mid = getMid(head)
	tail = mid.nextEle
	mid.nextEle = nil
	left = ChainListMergeSort(head)
	right = ChainListMergeSort(tail)
	ret = mergeLists(left, right) // 利用两个有序链表拼接的方法进行对接
	return ret
}

func getMid(head *ChainListNode) *ChainListNode {
	var (
		mid *ChainListNode
		end *ChainListNode
	)
	mid = head
	end = head.nextEle
	for end != nil && end.nextEle != nil {
		end = end.nextEle.nextEle
		mid = mid.nextEle
	}
	return mid
}
