package leetcode

import (
	"fmt"
)

// https://leetcode-cn.com/problems/valid-parentheses/

// 20. 有效的括号

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：

//     左括号必须用相同类型的右括号闭合。
//     左括号必须以正确的顺序闭合。

func isValid(s string) bool {
	// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
	// 内存消耗：1.9 MB, 在所有 Go 提交中击败了80.11% 的用户

	// 整一个栈，左括号入栈，匹配栈顶的右括号出栈，不匹配返回 false。结束后栈非空，返回 false 。
	if len(s) <= 1 {
		return false
	}
	cache := []rune{}
	for _, c := range s {
		l := len(cache)
		switch c {
		case '(', '[', '{':
			cache = append(cache, c)
		case ')':
			if l > 0 && cache[l-1] == '(' {
				cache = cache[:l-1]
			} else {
				return false
			}
		case ']':
			if l > 0 && cache[l-1] == '[' {
				cache = cache[:l-1]
			} else {
				return false
			}
		case '}':
			if l > 0 && cache[l-1] == '{' {
				cache = cache[:l-1]
			} else {
				return false
			}
		}
	}
	return len(cache) == 0
}

func Test20() {
	fmt.Println(isValid(""))       // false
	fmt.Println(isValid(")"))      // false
	fmt.Println(isValid("]]"))     // false
	fmt.Println(isValid("["))      // false
	fmt.Println(isValid("(("))     // false
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true
}
