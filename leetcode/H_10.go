package leetcode

import "fmt"

// https://leetcode-cn.com/problems/regular-expression-matching/

// 10. 正则表达式匹配

// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

//     '.' 匹配任意单个字符
//     '*' 匹配零个或多个前面的那一个元素

// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

func isMatch(s string, p string) bool {
	if len(p) == 0 { // p 长度为 0
		return len(s) == 0 // 若解析完毕则返回true
	} else if len(p) == 1 { // p 长度为 1
		return len(s) == 1 && (s[0] == p[0] || p[0] == '.') // 若最后一个字符能匹配，则返回true
	} else if p[1] == '*' { // p 至少长度为 2，且第二个字符是 *
		if isMatch(s, p[2:]) { // 如能直接匹配后续子串，则返回投true
			return true
		}
		if len(s) > 0 && (s[0] == p[0] || p[0] == '.') {
			return isMatch(s[1:], p) // 若能匹配一个字符，则继续匹配后续子串
		} else {
			return isMatch(s, p[2:]) // 若不能匹配，则认为此 * 匹配 0 次，继续匹配后续内容
		}
	} else { // p 的第二个字符不是 *
		if len(s) > 0 && (s[0] == p[0] || p[0] == '.') { // 如果当前字符能匹配，则匹配后续子串
			return isMatch(s[1:], p[1:])
		}
		return false // 若当前字符不能匹配，则认为匹配失败
	}
}

func Test10() {
	fmt.Println(isMatch("aaa", "a.a"))                // true
	fmt.Println(isMatch("bbaa", "a..."))              // false
	fmt.Println(isMatch("a", ".*..a*"))               // false
	fmt.Println(isMatch("mississippi", "mis*is*p*.")) // false
	fmt.Println(isMatch("a", "ab*"))                  // true
	fmt.Println(isMatch("aaa", "a*a"))                // true
	fmt.Println(isMatch("ab", ".*c"))                 // false
	fmt.Println(isMatch("abcd", "d*"))                // false
	fmt.Println(isMatch("aab", "c*a*b"))              // true
	fmt.Println(isMatch("aa", "a"))                   // false
	fmt.Println(isMatch("aa", "a*"))                  // true
	fmt.Println(isMatch("ab", ".*"))                  // true
}
