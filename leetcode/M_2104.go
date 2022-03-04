package leetcode

import (
	"fmt"
)

// https://leetcode-cn.com/problems/sum-of-subarray-ranges/

// 2104. 子数组范围和

// 给你一个整数数组 nums 。nums 中，子数组的 范围 是子数组中最大元素和最小元素的差值。
// 返回 nums 中 所有 子数组范围的 和 。
// 子数组是数组中一个连续 非空 的元素序列。

// 1 <= nums.length <= 1000
// -10⁹  <= nums[i] <= 10⁹

func subArrayRanges(nums []int) int64 {
	// 直接遍历取子数组，但不用记录子数组，记录最小值最大值即可
	// 每动一次，判断并记录最小值、最大值，同时收集其差值
	// 遍历完即获得结果
	var result int64
	l := len(nums)
	for i := 0; i < l; i++ {
		min := nums[i]
		max := nums[i]
		for j := i + 1; j < l; j++ {
			if nums[j] < min {
				min = nums[j]
			}
			if nums[j] > max {
				max = nums[j]
			}
			result += int64(max - min)
		}
	}
	return result
}

func Test2104() {
	fmt.Println(subArrayRanges([]int{1, 3, 3}))         // 4
	fmt.Println(subArrayRanges([]int{4, -2, -3, 4, 1})) // 59
}
