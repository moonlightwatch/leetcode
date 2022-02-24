package leetcode

import "fmt"

// https://leetcode-cn.com/problems/generate-parentheses/

// 22. 括号生成

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

func makeString(l, r int, s string) []string {
	result := []string{}
	if l == 0 && r == 0 {
		result = append(result, s)
	}
	if l > 0 {
		result = append(result, makeString(l-1, r, s+"(")...)
	}
	if r > l {
		result = append(result, makeString(l, r-1, s+")")...)
	}
	return result
}

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	return makeString(n, n, "")
}

func Test22() {
	fmt.Println(generateParenthesis(3)) // ["((()))","(()())","(())()","()(())","()()()"]
	fmt.Println(generateParenthesis(1)) // ["()"]
}
