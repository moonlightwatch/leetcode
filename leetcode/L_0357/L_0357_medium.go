package L_0357

import (
	"fmt"
)

// https://leetcode-cn.com/problems/count-numbers-with-unique-digits/

// 357. 统计各位数字都不同的数字个数
// 给你一个整数 n ，统计并返回各位数字都不同的数字 x 的个数，其中 0 <= x < 10^n 。

// 思路：
//   即从 0-9 十个数中取 n 个 进行组合，求所有组合的数量
//   直接应用组合公式：C(10,n) , （要注意第一位不能是 0）

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了86.36% 的用户

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	c := 1
	for i := 0; i < n; i++ {
		if i == 0 { // 第一位不能是 0，所以用 9 乘
			c *= 9
		} else { // 其他位可以选 10 个数
			c *= 10 - i
		}
	}
	// 因为组合只统计了 n 位时的情况
	// 所以要加上低于 n 位时的数量
	return c + countNumbersWithUniqueDigits(n-1)

	// 有限集合内的快速解：
	// return []int{1, 10, 91, 739, 5275, 32491, 168571, 712891, 2345851}[n] // 穷举答案，一行搞定……
}

func Test357() {

	fmt.Println(countNumbersWithUniqueDigits(0)) // 1
	fmt.Println(countNumbersWithUniqueDigits(1)) // 10
	fmt.Println(countNumbersWithUniqueDigits(2)) // 91
	fmt.Println(countNumbersWithUniqueDigits(3)) // 739
	fmt.Println(countNumbersWithUniqueDigits(4)) // 5275
	fmt.Println(countNumbersWithUniqueDigits(5)) // 32491
	fmt.Println(countNumbersWithUniqueDigits(6)) // 168571
	fmt.Println(countNumbersWithUniqueDigits(7)) // 712891
	fmt.Println(countNumbersWithUniqueDigits(8)) // 2345851
}
