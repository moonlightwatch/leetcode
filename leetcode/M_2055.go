package leetcode

import (
	"fmt"
)

// https://leetcode-cn.com/problems/plates-between-candles/

// 2055. 蜡烛之间的盘子

// 给你一个长桌子，桌子上盘子和蜡烛排成一列。给你一个下标从 0 开始的字符串 s ，它只包含字符 '*' 和 '|' ，其中 '*' 表示一个 盘子 ，'|' 表示一支 蜡烛 。

// 同时给你一个下标从 0 开始的二维整数数组 queries ，其中 queries[i] = [lefti, righti] 表示 子字符串 s[lefti...righti] （包含左右端点的字符）。对于每个查询，你需要找到 子字符串中 在 两支蜡烛之间 的盘子的 数目 。如果一个盘子在 子字符串中 左边和右边 都 至少有一支蜡烛，那么这个盘子满足在 两支蜡烛之间 。

//     比方说，s = "||**||**|*" ，查询 [3, 8] ，表示的是子字符串 "*||**|" 。子字符串中在两支蜡烛之间的盘子数目为 2 ，子字符串中右边两个盘子在它们左边和右边 都 至少有一支蜡烛。

// 请你返回一个整数数组 answer ，其中 answer[i] 是第 i 个查询的答案。

func platesBetweenCandles(s string, queries [][]int) []int {
	l := len(s)
	// 先做统计工作
	leftCandleMap := make([]int, l)       // 记录各个位置上，左侧最近的蜡烛位置
	rightCandleMap := make([]int, l)      // 记录各个位置上，右侧最近的蜡烛位置
	pansBeforeCandleMap := make([]int, l) // 记录各个位置上，前面有几个盘子
	// i 从左到右，j 从右到左，进行循环
	for i, j := 0, l-1; i < len(s) && j >= 0; i, j = i+1, j-1 {
		if s[i] == '|' { // 左侧遇到蜡烛
			leftCandleMap[i] = i // 则当前位置左侧最近的蜡烛是自己
			if i > 0 {           // 当前位置前面的盘子数量与前一个位置相同
				pansBeforeCandleMap[i] = pansBeforeCandleMap[i-1]
			}
		} else if i > 0 { // 左侧遇到盘子
			leftCandleMap[i] = leftCandleMap[i-1]                 // 则当前位置左侧最近的蜡烛与前一个位置相同
			pansBeforeCandleMap[i] = pansBeforeCandleMap[i-1] + 1 // 当前位置的盘子数量比前一个位置多一个（即自己）
		}
		if s[j] == '|' { // 右侧遇到蜡烛
			rightCandleMap[j] = j // 则当前位置右侧最近的蜡烛是自己
		} else if j < l-1 { // 右侧遇到盘子
			rightCandleMap[j] = rightCandleMap[j+1] // 则当前位置右侧最近的蜡烛，与前一个位置相同
		}
	}
	// 然后计算数量
	result := make([]int, len(queries))
	for i, q := range queries {
		if q[1]-q[0] < 2 { // 符合条件的盘子至少需要三个位置，即 "|*|"。小于三个的字串，直接填 0
			result[i] = 0
		} else if rightCandleMap[q[0]] != -1 && leftCandleMap[q[1]] != -1 && rightCandleMap[q[0]] < leftCandleMap[q[1]] {
			// 此处条件：1. 区间左端位置的右侧有蜡烛；2. 区间右端的左侧有蜡烛；3. 左端的蜡烛 在 右端蜡烛 的 左侧
			// 符合条件的：区间内有效的盘子数量 = 右端蜡烛位置的盘子数量 - 左端蜡烛位置的盘子数量
			result[i] = pansBeforeCandleMap[leftCandleMap[q[1]]] - pansBeforeCandleMap[rightCandleMap[q[0]]]
		}
	}
	return result
}

func Test2055() {
	fmt.Println(platesBetweenCandles("|*|*", [][]int{{3, 3}}))                                                        // 0
	fmt.Println(platesBetweenCandles("**|**|***|", [][]int{{2, 5}, {5, 9}}))                                          // 2 3
	fmt.Println(platesBetweenCandles("***|**|*****|**||**|*", [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}})) // 9 0 0 0 0
}
