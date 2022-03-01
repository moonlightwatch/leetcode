package leetcode

// https://leetcode-cn.com/problems/merge-k-sorted-lists/

// 23. 合并K个升序链表

// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。

// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4

func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	var head *ListNode
	var tail *ListNode
	for {
		min := 100000
		markIndex := -1
		for i := 0; i < k; i++ {
			if lists[i] != nil {
				if lists[i].Val < min {
					min = lists[i].Val
					markIndex = i
				}
			}
		}
		if markIndex != -1 {
			lists[markIndex] = lists[markIndex].Next
		} else {
			break
		}
		if head == nil {
			head = &ListNode{Val: min}
			tail = head
		} else {
			tail.Next = &ListNode{Val: min}
			tail = tail.Next
		}
	}
	return head
}

func Test23() {
	_printListNode(mergeKLists([]*ListNode{ // [1,1,2,3,4,4,5,6]
		_makeListNode([]int{1, 4, 5}),
		_makeListNode([]int{1, 3, 4}),
		_makeListNode([]int{2, 6})}))
	_printListNode(mergeKLists([]*ListNode{}))                       // [ ]
	_printListNode(mergeKLists([]*ListNode{_makeListNode([]int{})})) // [ ]
}
