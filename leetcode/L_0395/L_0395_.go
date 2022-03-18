package L_0395

import (
	"fmt"
	"strings"
)

// 给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。
// 返回这一子串的长度。

func longestSubstring(s string, k int) int {
	temps := map[rune]int{}
	for _, item := range s {
		if _, ok := temps[item]; !ok {
			temps[item] = 0
		}
		temps[item]++
	}
	for item, c := range temps {
		if c < k {
			max := 0
			for _, sub := range strings.Split(s, string(item)) {
				t := longestSubstring(sub, k)
				if max < t {
					max = t
				}
			}
			return max
		}
	}
	return len(s)
}

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
// 内存消耗：2.3 MB, 在所有 Go 提交中击败了25.92% 的用户

func Test395() {
	fmt.Printf("%v\n", longestSubstring("aaabb", 3))
	fmt.Printf("%v\n", longestSubstring("ababbc", 2))
}
