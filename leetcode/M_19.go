package leetcode

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 将链表各个节点存入数组
	// 根据输入 n 确定删除位置
	// 执行链表删除操作
	// *
	// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
	// 内存消耗：2.2 MB, 在所有 Go 提交中击败了74.82% 的用户

	cache := []*ListNode{} // 节点数组

	// 遍历链表，将节点存入数组
	next := head
	for {
		cache = append(cache, next)
		if next.Next == nil {
			break
		}
		next = next.Next

	}

	// 确认删除位置
	index := len(cache) - n

	// 若删除位置是 head，则直接返回 head 的 Next
	if index == 0 {
		return cache[0].Next
	}

	// 执行链表删除
	cache[index-1].Next = cache[index].Next

	// 返回 head
	return head
}

func Test19() {
	_printListNode(removeNthFromEnd(_makeListNode([]int{1, 2, 3, 4, 5}), 2)) // 1 2 3 5
	_printListNode(removeNthFromEnd(_makeListNode([]int{1}), 1))             //
	_printListNode(removeNthFromEnd(_makeListNode([]int{1, 2}), 1))          // 1
}
