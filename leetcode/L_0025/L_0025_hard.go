package L_0025

import "fmt"

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

// 25. K 个一组翻转链表

// 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

// 进阶：
//     你可以设计一个只使用常数额外空间的算法来解决此问题吗？
//     你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

// 列表中节点的数量在范围 sz 内
// 1 <= sz <= 5000
// 0 <= Node.val <= 1000
// 1 <= k <= sz

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 遍历链表，将节点存入缓存
	// 直到 k 个节点进入缓存
	// 将缓存中的节点翻转
	// 继续遍历，直至结束

	if k == 1 { // 特殊情况判断
		return head
	}
	pointer := head     // 遍历游标
	var front *ListNode // 记录当前翻转节点的前一个节点
	for {
		cache := make([]*ListNode, k) // 节点缓存
		for i := 0; i < k; i++ {      // 将k个节点放入缓存
			if pointer == nil {
				return head
			}
			cache[i] = pointer
			pointer = pointer.Next // 游标后移
		}
		// 反转缓存中的节点
		for i := k - 1; i > 0; i-- {
			cache[i].Next = cache[i-1]
		}
		cache[0].Next = pointer // 续接尾部
		if front == nil {       // 若是头部
			head = cache[k-1] // 则更新head
		} else {
			front.Next = cache[k-1] // 非头部。则续接前一个节点
		}
		front = cache[0] // 更新 “前一个节点”
	}
}

func Test25() {
	_printListNode(reverseKGroup(_makeListNode([]int{1, 2, 3, 4, 5}), 4)) // [4,3,2,1,5]
	_printListNode(reverseKGroup(_makeListNode([]int{1, 2, 3, 4, 5}), 2)) // [2,1,4,3,5]
	_printListNode(reverseKGroup(_makeListNode([]int{1, 2, 3, 4, 5}), 3)) // [3,2,1,4,5]
	_printListNode(reverseKGroup(_makeListNode([]int{1, 2, 3, 4, 5}), 1)) // [1,2,3,4,5]
	_printListNode(reverseKGroup(_makeListNode([]int{1}), 1))             // [1]
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
