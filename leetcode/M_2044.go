package leetcode

import "fmt"

// https://leetcode-cn.com/problems/count-number-of-maximum-bitwise-or-subsets/

// 2044. 统计按位或能得到最大值的子集数目

// 给你一个整数数组 nums ，请你找出 nums 子集 按位或 可能得到的 最大值 ，并返回按位或能得到最大值的 不同非空子集的数目 。
// 如果数组 a 可以由数组 b 删除一些元素（或不删除）得到，则认为数组 a 是数组 b 的一个 子集 。如果选中的元素下标位置不一样，则认为两个子集 不同 。
// 对数组 a 执行 按位或 ，结果等于 a[0] OR a[1] OR ... OR a[a.length - 1]（下标从 0 开始）。

func countMaxOrSubsets(nums []int) int {
	// 按照遍历子集的方式累计计算或的结果，计算过程中统计最大值的数量
	cache := []int{0}
	max := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		l := len(cache)
		for j := 0; j < l; j++ {
			o := cache[j] | nums[i]
			cache = append(cache, o)
			if o > max {
				max = o
				count = 1
			} else if o == max {
				count++
			}
		}
	}
	return count

}

func Test2044() {
	fmt.Println(countMaxOrSubsets([]int{3, 1}))       // 2
	fmt.Println(countMaxOrSubsets([]int{2, 2, 2}))    // 7
	fmt.Println(countMaxOrSubsets([]int{3, 2, 1, 5})) // 6
}
