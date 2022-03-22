package L_1269

import (
	"fmt"
)

// https://leetcode-cn.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/

// 1269. 停在原地的方案数

// 有一个长度为 arrLen 的数组，开始有一个指针在索引 0 处。

// 每一步操作中，你可以将指针向左或向右移动 1 步，或者停在原地（指针不能被移动到数组范围外）。

// 给你两个整数 steps 和 arrLen ，请你计算并返回：在恰好执行 steps 次操作以后，指针仍然指向索引 0 处的方案数。

// 由于答案可能会很大，请返回方案数 模 10^9 + 7 后的结果。

func numWays(steps int, arrLen int) int {
	if steps/2+1 < arrLen {
		arrLen = steps/2 + 1
	}
	cache := [][]int{make([]int, arrLen)}
	cache[0][0] = 1
	if arrLen > 1 {
		cache[0][1] = 1
	}
	for i := 1; i < steps; i++ {
		cache = append(cache, make([]int, arrLen))
		for j := 0; j < arrLen; j++ {
			cache[i][j] = cache[i-1][j]
			if j > 0 {
				cache[i][j] += cache[i-1][j-1]
			}
			if j < arrLen-1 {
				cache[i][j] += cache[i-1][j+1]
			}
			cache[i][j] = cache[i][j] % 1000000007

		}
	}

	return cache[steps-1][0]
}

func Test1269() {
	fmt.Println(numWays(3, 1))   //
	fmt.Println(numWays(47, 38)) // 318671228
	fmt.Println(numWays(3, 2))   // 4
	fmt.Println(numWays(2, 4))   // 2
	fmt.Println(numWays(4, 2))   // 8
	fmt.Println(numWays(27, 7))  //
}
