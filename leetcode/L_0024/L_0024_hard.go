package L_0024

import "fmt"

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/

// 24. 两两交换链表中的节点

// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100

func swapPairs(head *ListNode) *ListNode {
	// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
	// 内存消耗：1.9 MB, 在所有 Go 提交中击败了85.32% 的用户

	// 就从头到尾，两两交换
	// 通过三个标记，标记 “前一个节点”、“交换节点1”、“交换节点2”
	// 直接交换节点：
	//   frontNode.Next = node2
	//   node1.Next = node2.Next
	//   node2.Next = node1
	// 需要注意：头节点交换，不需要“前一个节点”。尾节点交换，需要处理奇数个节点的情况。

	if head == nil || head.Next == nil { // 特殊情况
		return head
	}
	var frontNode *ListNode // 前一个节点
	node1 := head           // 交换节点1
	node2 := head.Next      // 交换节点2
	for {
		// 两两交换
		if frontNode == nil { // 处理头部节点交换的特殊情况
			node1.Next = node2.Next
			node2.Next = node1
			head = node2 // 头节点变化了，需要重新赋值
		} else {
			// 交换节点
			frontNode.Next = node2
			node1.Next = node2.Next
			node2.Next = node1
		}
		if node1.Next == nil { // 如果交换后的后面一个节点没有下一个节点，则交换完成，退出循环
			break
		}
		// 下一组
		frontNode = node1  // 下一次的 “前一个节点” 是本次的最后一个节点，即交换后的node1
		node1 = node1.Next // 下一次的 交换节点1
		node2 = node1.Next // 下一次的 交换节点2
		if node2 == nil {  // 若交换节点2是空的，即链表元素是奇数个，直接跳出，最后一个元素不必交换
			break
		}
	}
	return head
}

func Test24() {
	_printListNode(swapPairs(_makeListNode([]int{1, 2, 3})))    // 2, 1, 4, 3
	_printListNode(swapPairs(_makeListNode([]int{1, 2, 3, 4}))) // 2, 1, 4, 3
	_printListNode(swapPairs(_makeListNode([]int{})))           //
	_printListNode(swapPairs(_makeListNode([]int{1})))          // 1
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
