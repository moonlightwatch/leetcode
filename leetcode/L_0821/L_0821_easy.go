package L_0821

import (
	"fmt"
)

// https://leetcode-cn.com/problems/shortest-distance-to-a-character/

// 821. 字符的最短距离

// 给你一个字符串 s 和一个字符 c ，且 c 是 s 中出现过的字符。

// 返回一个整数数组 answer ，其中 answer.length == s.length 且 answer[i] 是 s 中从下标 i 到离它 最近 的字符 c 的 距离 。

// 两个下标 i 和 j 之间的 距离 为 abs(i - j) ，其中 abs 是绝对值函数。

// 1 <= s.length <= 104
// s[i] 和 c 均为小写英文字母
// 题目数据保证 c 在 s 中至少出现一次

// 思路：
// 字符串分三部分，分别对应不同的距离规则：
//   1. 从开始到第一个 c : 距离是从第一个c逆推到开始位置，距离逐步 +1
//   2. 两个 c 之间 : 距离是从两个 c 的位置，双向同时向内遍历，距离逐步 +1
//   3. 最后一个 c 到末尾： 距离是从c向后遍历到末尾，距离逐步 +1

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
// 内存消耗：2.2 MB, 在所有 Go 提交中击败了100.00% 的用户

func shortestToChar(s string, c byte) []int {
	l := len(s)
	answer := make([]int, l)
	lastIndex := -1 // 记录上一个 c 的索引
	for i := 0; i < l; i++ {
		if s[i] == c {
			if lastIndex == -1 { // 从开始到第一个 c : 距离是从第一个c逆推到开始位置，距离逐步 +1
				for j := i - 1; j >= 0; j-- {
					answer[j] = answer[j+1] + 1
				}
			} else { // 两个 c 之间 : 距离是从两个 c 的位置，双向同时向内遍历，距离逐步 +1
				for j := 1; lastIndex+j <= i-j; j++ {
					answer[lastIndex+j] = j
					answer[i-j] = j
				}
			}
			lastIndex = i
		}
	}
	// 最后一个 c 到末尾： 距离是从c向后遍历到末尾，距离逐步 +1
	for j := lastIndex + 1; j < l; j++ {
		answer[j] = answer[j-1] + 1
	}
	return answer
}

func Test821() {
	fmt.Println(shortestToChar("loveleetcode", 'e')) // 3,2,1,0,1,0,0,1,2,2,1,0
	fmt.Println(shortestToChar("aaab", 'b'))         // 3,2,1,0
	fmt.Println(shortestToChar("zxabaaa", 'b'))      // 3,2,1,0,1,2,3
}
