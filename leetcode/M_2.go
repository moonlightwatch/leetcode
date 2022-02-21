package leetcode

// https://leetcode-cn.com/problems/add-two-numbers/

// 2. 两数相加
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := l1
	for {
		l1.Val += l2.Val
		if l1.Val > 9 {
			l1.Val = l1.Val % 10
			if l1.Next == nil {
				l1.Next = &ListNode{Val: 1, Next: nil}
			} else {
				l1.Next.Val++
			}
		}

		if l1.Next == nil && l2.Next == nil {
			break
		}

		if l1.Next == nil {
			l1.Next = &ListNode{Val: 0, Next: nil}
		}
		l1 = l1.Next
		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2 = &ListNode{Val: 0, Next: nil}
		}
	}
	return r
}
func Test2() {
	_printListNode(addTwoNumbers(_makeListNode([]int{2, 4, 3}), _makeListNode([]int{5, 6, 4})))                // [7,0,8]
	_printListNode(addTwoNumbers(_makeListNode([]int{0}), _makeListNode([]int{0})))                            // [0]
	_printListNode(addTwoNumbers(_makeListNode([]int{9, 9, 9, 9, 9, 9, 9}), _makeListNode([]int{9, 9, 9, 9}))) // [8,9,9,9,0,0,0,1]
}
