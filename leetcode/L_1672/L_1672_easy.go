package L_1672

import (
	"fmt"
)

// https://leetcode-cn.com/problems/richest-customer-wealth/

// 1672. 最富有客户的资产总量

// 给你一个 m x n 的整数网格 accounts ，其中 accounts[i][j] 是第 i​​​​​​​​​​​​ 位客户在第 j 家银行托管的资产数量。返回最富有客户所拥有的 资产总量 。

// 客户的 资产总量 就是他们在各家银行托管的资产数量之和。最富有客户就是 资产总量 最大的客户。

// m == accounts.length
// n == accounts[i].length
// 1 <= m, n <= 50
// 1 <= accounts[i][j] <= 100

func maximumWealth(accounts [][]int) int {
	cache := make([]int, len(accounts))
	for i, b := range accounts {
		for _, a := range b {
			cache[i] += a
		}
	}
	max := 0
	for _, v := range cache {
		if max < v {
			max = v
		}
	}
	return max
}

func Test1672() {
	fmt.Println(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))   // 6
	fmt.Println(maximumWealth([][]int{{1, 5}, {7, 3}, {3, 5}})) // 10
}
