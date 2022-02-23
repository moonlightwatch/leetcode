package leetcode

import (
	"fmt"
	"math"
	"sort"
)

// https://leetcode-cn.com/problems/the-number-of-good-subsets/

// 1994. 好子集的数目

// 给你一个整数数组 nums 。如果 nums 的一个子集中，所有元素的乘积可以表示为一个或多个 互不相同的质数 的乘积，那么我们称它为 好子集 。

//     比方说，如果 nums = [1, 2, 3, 4] ：
//         [2, 3] ，[1, 2, 3] 和 [1, 3] 是 好 子集，乘积分别为 6 = 2*3 ，6 = 2*3 和 3 = 3 。
//         [1, 4] 和 [4] 不是 好 子集，因为乘积分别为 4 = 2*2 和 4 = 2*2 。

// 请你返回 nums 中不同的 好 子集的数目对 10⁹ + 7 取余 的结果。

// nums 中的 子集 是通过删除 nums 中一些（可能一个都不删除，也可能全部都删除）元素后剩余元素组成的数组。如果两个子集删除的下标不同，那么它们被视为不同的子集。

// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 30

func numberOfGoodSubsets(nums []int) int {
	targetMap := map[int][]int{
		2:  {2},
		3:  {3},
		5:  {5},
		6:  {2, 3},
		7:  {7},
		10: {2, 5},
		11: {11},
		13: {13},
		14: {2, 7},
		15: {3, 5},
		17: {17},
		19: {19},
		21: {3, 7},
		22: {2, 11},
		23: {23},
		26: {2, 13},
		29: {29},
		30: {2, 3, 5},
	}

	oneCount := 0 // 原数组包含 1 的数量
	targetList := [][]int{}
	for _, num := range nums {
		if num == 1 {
			oneCount++
		}
		if primes, ok := targetMap[num]; ok {
			targetList = append(targetList, primes)
		}
	}
	res := [][]int{{}}
	for _, primes := range targetList {
		l := len(res)
		for i := 0; i < l; i++ {
			newList := append(res[i], primes...)
			// 确认没有重复
			mark := false
			sort.Ints(newList)
			for j := 1; j < len(newList); j++ {
				if newList[j-1] == newList[j] {
					mark = true
					break
				}
			}
			if !mark {
				res = append(res, newList)
				// fmt.Println(newList)
			}
		}
	}
	// fmt.Println(res)
	return int(math.Pow(2, float64(oneCount))) * (len(res) - 1) % int(math.Pow10(9)+7)
}

func Test1994() {
	fmt.Println(numberOfGoodSubsets([]int{5, 10, 1, 26, 24, 21, 24, 23, 11, 12, 27, 4, 17, 16, 2, 6, 1, 1, 6, 8, 13, 30, 24, 20, 2, 19})) // 5368
	fmt.Println(numberOfGoodSubsets([]int{6, 8, 1, 8, 6, 5, 6, 11, 17}))                                                                  // 62
	// fmt.Println(numberOfGoodSubsets([]int{18, 28, 2, 17, 29, 30, 15, 9, 12}))                                                             // 19
	// fmt.Println(numberOfGoodSubsets([]int{1, 2, 3, 4}))                                                                                   // 6
	// fmt.Println(numberOfGoodSubsets([]int{4, 2, 3, 15}))                                                                                  // 5
	// fmt.Println(numberOfGoodSubsets([]int{1, 2}))                                                                                         // 2
	// fmt.Println(numberOfGoodSubsets([]int{1, 1, 1, 1, 1, 1}))                                                                             // 0
	// fmt.Println(numberOfGoodSubsets([]int{1, 2, 2}))                                                                                      // 4
	// fmt.Println(numberOfGoodSubsets([]int{2}))                                                                                            // 1
}
