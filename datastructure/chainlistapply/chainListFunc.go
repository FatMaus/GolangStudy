package chainlistapply

// 循环链表的循环起点
func listCirclePoint(head *ChainListNode) *ChainListNode {
	var (
		fastPointer *ChainListNode
		slowPointer *ChainListNode
	)
	if head == nil {
		return head
	}
	fastPointer = head.nextEle
	slowPointer = head
	for fastPointer != nil && fastPointer.nextEle != nil {
		fastPointer = fastPointer.nextEle.nextEle
		slowPointer = slowPointer.nextEle
		if fastPointer == slowPointer {
			fastPointer = head
			slowPointer = slowPointer.nextEle
			for fastPointer != slowPointer {
				fastPointer = fastPointer.nextEle
				slowPointer = slowPointer.nextEle
			}
			return slowPointer
		}
	}
	return nil
}
