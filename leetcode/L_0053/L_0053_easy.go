package L_0053

import "fmt"

// https://leetcode-cn.com/problems/maximum-subarray/

// 53. 最大子数组和

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

// 子数组 是数组中的一个连续部分。

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func Test53() {
	fmt.Println(maxSubArray([]int{-2, 1}))                         // 1
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(maxSubArray([]int{1}))                             // 1
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))                // 23
}
