package L_0021

import "fmt"

// https://leetcode-cn.com/problems/merge-two-sorted-lists/

// 21. 合并两个有序链表

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// mergeTwoLists 循环解
func mergeTwoLists_(list1 *ListNode, list2 *ListNode) *ListNode {
	// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
	// 内存消耗：2.4 MB, 在所有 Go 提交中击败了99.98% 的用户

	// 将 list2 按顺序插入到 list1, 返回list1。

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	node1 := list1
	var frontNode *ListNode
	for list2 != nil {
		if node1 == nil {
			frontNode.Next = list2
			break
		}
		if node1.Val > list2.Val { // 需要将list2，插入到node1前面
			next2 := list2.Next
			if frontNode == nil {
				list2.Next = node1
				list1 = list2
				node1 = list2.Next
				frontNode = list2
			} else {
				frontNode.Next = list2
				list2.Next = node1
				frontNode = list2
			}

			list2 = next2
		} else {
			frontNode = node1
			node1 = node1.Next
		}
	}
	return list1
}

// mergeTwoLists 递归解
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func Test21() {
	_printListNode(mergeTwoLists(_makeListNode([]int{1, 2, 4}), _makeListNode([]int{1, 3, 4}))) // [1,1,2,3,4,4]
	_printListNode(mergeTwoLists(_makeListNode([]int{}), _makeListNode([]int{})))               // []
	_printListNode(mergeTwoLists(_makeListNode([]int{}), _makeListNode([]int{0})))              // [0]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func _makeListNode(items []int) *ListNode {
	if len(items) == 0 {
		return nil
	}
	head := &ListNode{Val: items[0], Next: nil}
	tail := head
	cache := head
	for _, v := range items {
		cache.Val = v
		cache.Next = &ListNode{}
		tail = cache
		cache = cache.Next
	}
	tail.Next = nil
	return head
}
func _printListNode(head *ListNode) {
	cache := head
	for cache != nil {
		fmt.Printf("%d ", cache.Val)
		cache = cache.Next
	}
	fmt.Println()
}
