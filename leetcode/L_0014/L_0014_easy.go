package L_0014

import "fmt"

// https://leetcode-cn.com/problems/longest-common-prefix/

// 14. 最长公共前缀

// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""。

func longestCommonPrefix(strs []string) string {
	// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
	// 内存消耗：2.2 MB, 在所有 Go 提交中击败了84.13% 的用户
	perfix := ""
	for i := 0; ; i++ {
		if i >= len(strs[0]) {
			return perfix
		}
		m := strs[0][i]
		for _, str := range strs {
			if i >= len(str) || m != str[i] {
				return perfix
			}
		}
		perfix += string(m)
	}
}

func Test14() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""
}
