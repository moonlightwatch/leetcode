package leetcode

import "fmt"

// https://leetcode-cn.com/problems/palindrome-number/

// 9. 回文数
// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//     例如，121 是回文，而 123 不是。

func isPalindrome(x int) bool {
	s := fmt.Sprintf("%d", x)
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l += 1
		r -= 1
	}
	return true
}

func Test9() {
	fmt.Println(isPalindrome(121))  // true
	fmt.Println(isPalindrome(-121)) // false
	fmt.Println(isPalindrome(10))   // false
}
