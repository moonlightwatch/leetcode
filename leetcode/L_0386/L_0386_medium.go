package L_0386

import (
	"fmt"
)

// https://leetcode-cn.com/problems/lexicographical-numbers/

// 386. 字典序排数

// 给你一个整数 n ，按字典序返回范围 [1, n] 内所有整数。

// 你必须设计一个时间复杂度为 O(n) 且使用 O(1) 额外空间的算法。

// 1 <= n <= 5 * 10^4

func lexicalOrder(n int) []int {
	result := []int{}
	for i := 1; i < 10; i++ {
		dfs(n, i, &result)
	}
	return result
}

func dfs(n int, cur int, result *[]int) {
	if cur > n {
		return
	}
	*result = append(*result, cur)
	for i := 0; i < 10; i++ {
		dfs(n, (cur*10)+i, result)
	}
}

func Test386() {
	fmt.Println(lexicalOrder(13)) // [1,10,11,12,13,2,3,4,5,6,7,8,9]
	fmt.Println(lexicalOrder(2))  // [1,2]
}
