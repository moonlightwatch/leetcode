package L_0032

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/longest-valid-parentheses/

// 32. 最长有效括号

// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

// 统计所有有效括号，并记录其索引
// 将所有索引排序后
// 计算统计最长连续的索引长度即可
func longestValidParentheses(s string) int {

	l := len(s)
	// 长度小于2的，没有有效括号
	if l < 2 {
		return 0
	}
	maxLen := 0                  // 记录最长有效括号长度
	cache := []int{}             // 保存有效括号索引
	indexStack := make([]int, l) // 统计有效括号使用的栈
	top := 0                     // 栈顶
	for i := 0; i < l; i++ {
		if s[i] == '(' { // 左括号入栈
			indexStack[top] = i
			top++
		} else if top > 0 && s[indexStack[top-1]] == '(' { // 右括号，判断栈顶后，出栈（top-1）。并且记录有效括号索引
			cache = append(cache, i)
			cache = append(cache, indexStack[top-1])
			top--
		}
	}

	// 排序
	sort.Slice(cache, func(i, j int) bool {
		return cache[i] < cache[j]
	})

	tempLen := 1
	// 统计最长连续索引
	for i := 1; i < len(cache); i++ {
		if cache[i-1] == cache[i]-1 {
			tempLen++
		} else {
			tempLen = 1
		}
		if tempLen > maxLen {
			maxLen = tempLen
		}
	}
	return maxLen
}

func Test32() {

	fmt.Println(longestValidParentheses("))))((()((")) // 2
	fmt.Println(longestValidParentheses("())"))        // 2
	fmt.Println(longestValidParentheses("()(()"))      // 2
	fmt.Println(longestValidParentheses("("))          // 0
	fmt.Println(longestValidParentheses("(()"))        // 2
	fmt.Println(longestValidParentheses(""))           // 0
	fmt.Println(longestValidParentheses(")()())"))     // 4
	fmt.Println(longestValidParentheses("(()())"))     // 6
}
