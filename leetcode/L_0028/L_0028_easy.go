package L_0028

import (
	"fmt"
)

// https://leetcode-cn.com/problems/implement-strstr/

// 28. 实现 strStr()

// 实现 strStr() 函数。

// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。

// 说明：

// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	t := len(haystack) - len(needle)
	if t < 0 {
		return -1
	}
	for i := 0; i < t+1; i++ {
		found := true
		for n := 0; n < len(needle); n++ {
			if needle[n] != haystack[i+n] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}

	return -1
}

func Test28() {
	fmt.Println(strStr("hello", "ll"))  // 2
	fmt.Println(strStr("aaaaa", "bba")) // -1
	fmt.Println(strStr("a", "a"))       // 0
	fmt.Println(strStr("", ""))         // 0
}
