package L_0003

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

// 3. 无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	max := 1
	t := ""
	for _, r := range s {
		if strings.ContainsRune(t, r) {
			t = t[strings.Index(t, string(r))+1:]
		}
		t = t + string(r)
		if max < len(t) {
			max = len(t)
		}
	}
	return max
}

func Test3() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3

}
