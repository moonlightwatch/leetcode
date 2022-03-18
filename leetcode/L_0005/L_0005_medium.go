package L_0005

import "fmt"

// https://leetcode-cn.com/problems/longest-palindromic-substring/

// 5. 最长回文子串
// 给你一个字符串 s，找到 s 中最长的回文子串。

func longestPalindrome(s string) string {
	length := len(s)
	if length == 1 {
		return s
	}
	if length == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	}
	longestPalindrome := s[0:1]
	runes := []rune(s)
	l, r := 0, 0
	for i := 1; i < length; i++ {
		if i+1 < length && runes[i-1] == runes[i+1] {
			l, r = i-1, i+1
			for l > 0 && r+1 < length && runes[l-1] == runes[r+1] {
				l -= 1
				r += 1
			}
			temp := runes[l : r+1]
			if len(longestPalindrome) < len(temp) {
				longestPalindrome = string(temp)
			}
		}
		if runes[i] == runes[i-1] {
			l, r = i-1, i
			for l > 0 && r+1 < length && runes[l-1] == runes[r+1] {
				l -= 1
				r += 1
			}
			temp := runes[l : r+1]
			if len(longestPalindrome) < len(temp) {
				longestPalindrome = string(temp)
			}
		}
	}
	return longestPalindrome
}

func Test5() {
	fmt.Println(longestPalindrome("abcda")) // a
	fmt.Println(longestPalindrome("eabcb")) // bcb
	fmt.Println(longestPalindrome("aaaa"))  // aaaa
	fmt.Println(longestPalindrome("abb"))   // bb
	fmt.Println(longestPalindrome("babad")) // bab
	fmt.Println(longestPalindrome("cbbd"))  // bb
	fmt.Println(longestPalindrome("aa"))    // aa
	fmt.Println(longestPalindrome("ccc"))   // ccc
}
