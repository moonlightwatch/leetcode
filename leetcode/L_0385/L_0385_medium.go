package L_0385

import (
	"fmt"
	"strconv"
)

// https://leetcode-cn.com/problems/mini-parser/

// 385. 迷你语法分析器

// 给定一个字符串 s 表示一个整数嵌套列表，实现一个解析它的语法分析器并返回解析的结果 NestedInteger 。

// 列表中的每个元素只可能是整数或整数嵌套列表

// 思路：
//     遍历字符串，遇见 '[' 则创建新的 NestedInteger ，将其加入栈顶的 NestedInteger 中 并入栈，遇见 ']' 则出栈，遇见 ',' 跳过
//          其他情况，则是数字部分，向后提取到 ',' 或者 ']'，转换为数字。为此数字创建 NestedInteger ，将其加入栈顶的 NestedInteger 中
//     遍历完就行

func deserialize(s string) *NestedInteger {
	if s[0] != '[' { // 特殊情况：字符串由单个数字组成
		v, _ := strconv.Atoi(s)
		n := NestedInteger{}
		n.SetInteger(v)
		return &n
	}
	l := len(s)
	root := &NestedInteger{}
	cache := []*NestedInteger{root}
	for i := 0; i < l; i++ {
		cacheLen := len(cache)
		switch s[i] {
		case '[': // 创建
			cache[cacheLen-1].Add(NestedInteger{}) // 新建 NestedInteger 将其加入栈顶的 NestedInteger 中
			temp := cache[cacheLen-1].GetList()
			cache = append(cache, temp[len(temp)-1]) // 将新建的 NestedInteger 入栈
		case ']':
			cache = cache[:cacheLen-1] // 出栈
		case ',':
			continue // 跳过
		default:
			n := 1
			if s[i] == '-' {
				n = -1
				i++
			}
			num := int(s[i] - byte('0'))
			i++
			for ; s[i] != ',' && s[i] != ']'; i++ {
				num = num*10 + int(s[i]-byte('0'))
			}
			cache[cacheLen-1].Add(NestedInteger{})
			temp := cache[cacheLen-1].GetList()
			temp[len(temp)-1].SetInteger(n * num)
			if s[i] == ']' {
				cache = cache[:len(cache)-1]
			}
		}
	}
	return root.GetList()[0]
}

func Test385() {
	// fmt.Printf("%+v", deserialize("324"))
	// a := deserialize("[123,[456,[789]]]")
	a := deserialize("[-16,[456,[789]]]")
	fmt.Printf("%+v", a)
}

type NestedInteger struct {
	Items  []*NestedInteger
	Value  int
	Single bool
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return n.Single
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return n.Value
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.Single = true
	n.Value = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	n.Single = false
	n.Value = 0
	if n.Items == nil {
		n.Items = []*NestedInteger{}
	}
	n.Items = append(n.Items, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n *NestedInteger) GetList() []*NestedInteger {
	if n.Items == nil {
		n.Items = []*NestedInteger{}
	}
	if n.Single {
		return []*NestedInteger{}
	}
	return n.Items
}
