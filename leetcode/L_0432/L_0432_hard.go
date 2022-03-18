package L_0432

import (
	"fmt"
)

// https://leetcode-cn.com/problems/all-oone-data-structure/

// 432. 全 O(1) 的数据结构

// 请你设计一个用于存储字符串计数的数据结构，并能够返回计数最小和最大的字符串。

// 实现 AllOne 类：

//     AllOne() 初始化数据结构的对象。
//     inc(String key) 字符串 key 的计数增加 1 。如果数据结构中尚不存在 key ，那么插入计数为 1 的 key 。
//     dec(String key) 字符串 key 的计数减少 1 。如果 key 的计数在减少后为 0 ，那么需要将这个 key 从数据结构中删除。测试用例保证：在减少计数前，key 存在于数据结构中。
//     getMaxKey() 返回任意一个计数最大的字符串。如果没有元素存在，返回一个空字符串 "" 。
//     getMinKey() 返回任意一个计数最小的字符串。如果没有元素存在，返回一个空字符串 "" 。

// 使用有序的双向链表，加一个字典保存字符串和链表节点的关系

// 双向链表节点
type SNode struct {
	S     string
	Count int
	Next  *SNode
	Front *SNode
}

type AllOne struct {
	tab  map[string]*SNode // 保存字符串和链表节点的关系
	head *SNode            // 头节点，即最小计数字符串所在的节点
	tail *SNode            // 尾节点，即最大计数字符串所在节点
}

func Constructor() AllOne {
	return AllOne{
		tab: map[string]*SNode{}, // 初始化字典
	}
}

func (this *AllOne) Inc(key string) {
	// 增加计数
	// 先从字典中找到对应的节点，增加计数
	// 然后调整增加计数后，此节点的位置

	if _, ok := this.tab[key]; !ok { // 若是新的字符串，则添加到字典中
		this.tab[key] = &SNode{S: key, Count: 1}
		// 若非空链表，则先把节点放到链表末尾，后续再挪
		if this.tail != nil {
			this.tail.Next = this.tab[key]
			this.tab[key].Front = this.tail
			this.tail = this.tab[key]
		}
	} else {
		// 若是已经存在的字符串，则增加计数
		this.tab[key].Count++
	}
	if this.head == nil { // 链表还是空的情况下，设置第一个节点
		this.head = this.tab[key]
		this.tail = this.tab[key]
	} else if this.tab[key].Front != nil && this.tab[key].Front.Count > this.tab[key].Count { // 当前节点计数比前一个节点小，应调整位置
		tempNode := this.tab[key]
		// 将节点往前交换（仅交换节点的内容），直到达到有序链表的合理位置，即前一个节点的计数小于当前节点
		for tempNode.Front != nil && tempNode.Front.Count >= tempNode.Count {
			tempNode.Front.S, tempNode.S = tempNode.S, tempNode.Front.S
			tempNode.Front.Count, tempNode.Count = tempNode.Count, tempNode.Front.Count
			this.tab[tempNode.S] = tempNode
			this.tab[tempNode.Front.S] = tempNode.Front
			tempNode = tempNode.Front
		}
		// 如果挪到头了，则更新头节点
		if tempNode.Front == nil {
			this.head = tempNode
		}
	} else if this.tab[key].Next != nil && this.tab[key].Next.Count < this.tab[key].Count { // 当前节点计数比后一个节点大，应调整位置
		tempNode := this.tab[key]
		// 将节点往后交换（仅交换节点的内容），直到达到有序链表的合理位置，即后一个节点的计数大于当前节点
		for tempNode.Next != nil && tempNode.Next.Count <= tempNode.Count {
			tempNode.Next.S, tempNode.S = tempNode.S, tempNode.Next.S
			tempNode.Next.Count, tempNode.Count = tempNode.Count, tempNode.Next.Count
			this.tab[tempNode.S] = tempNode
			this.tab[tempNode.Next.S] = tempNode.Next
			tempNode = tempNode.Next
		}
		// 若挪到了尾部，则更新尾节点
		if tempNode.Next == nil {
			this.tail = tempNode
		}
	}
}

func (this *AllOne) Dec(key string) {
	// 减少计数
	// 先从字典中找到对应的节点，减少计数
	// 然后调整减少计数后，此节点的位置

	if _, ok := this.tab[key]; !ok { // 不存在的字符串，不操作
		return
	} else {
		this.tab[key].Count-- // 减少计数
	}
	if this.tab[key].Count == 0 { // 如果减到0了，就删除这个节点
		if this.tab[key].Front != nil { // 如果不是头节点，则更新前一个节点的Next
			this.tab[key].Front.Next = this.tab[key].Next
		} else { // 如果是头节点，直接更新头节点
			this.head = this.tab[key].Next
		}
		if this.tab[key].Next != nil { // 如果不是尾节点，则更新后一个节点的Front
			this.tab[key].Next.Front = this.tab[key].Front
		} else { // 如果是尾节点，直接更新尾节点
			this.tail = this.tab[key].Front
		}
		// delete(this.tab, key) // 此处可以不删除，不影响结果，甚至会更快
	} else if this.tab[key].Front != nil && this.tab[key].Front.Count > this.tab[key].Count { // 当前节点计数比前一个节点小，应调整位置
		tempNode := this.tab[key]
		for tempNode.Front != nil && tempNode.Front.Count >= tempNode.Count {
			tempNode.Front.S, tempNode.S = tempNode.S, tempNode.Front.S
			tempNode.Front.Count, tempNode.Count = tempNode.Count, tempNode.Front.Count
			this.tab[tempNode.S] = tempNode
			this.tab[tempNode.Front.S] = tempNode.Front
			tempNode = tempNode.Front
		}
		if tempNode.Front == nil {
			this.head = tempNode
		}
	}

}

func (this *AllOne) GetMaxKey() string {
	if this.tail == nil {
		return ""
	}
	return this.tail.S
}

func (this *AllOne) GetMinKey() string {
	if this.head == nil {
		return ""
	}
	return this.head.S
}

func Test432() {
	allOne := Constructor()
	allOne.Inc("hello")
	allOne.Inc("world")
	allOne.Inc("leet")
	allOne.Inc("code")
	allOne.Inc("ds")
	allOne.Inc("leet")
	fmt.Println(allOne.GetMaxKey())
	allOne.Inc("ds")
	allOne.Dec("leet")
	fmt.Println(allOne.GetMaxKey())
	// fmt.Println(allOne.GetMaxKey())
	// fmt.Println(allOne.GetMinKey())
	// allOne.Dec("a")
	// allOne.Inc("leet")
	// fmt.Println(allOne.GetMaxKey())
	// fmt.Println(allOne.GetMinKey())

}
