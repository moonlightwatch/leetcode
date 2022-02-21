package leetcode

import "fmt"

// https://leetcode-cn.com/problems/two-sum

// 1. 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。

func twoSum(nums []int, target int) []int {
	i := 0
	j := 1
	for {
		if nums[i]+nums[j] == target {
			return []int{i, j}
		}
		j++
		if j == len(nums) {
			i++
			j = i + 1
		}
	}
}

func Test1() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0,1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1,2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // [0,1]
}
