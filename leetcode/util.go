package leetcode

import "fmt"

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
