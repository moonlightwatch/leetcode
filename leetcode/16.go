package leetcode

import "fmt"

// 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
// 返回这三个数的和。
// 假定每组输入只存在恰好一个解。

func threeSumClosest(nums []int, target int) int {
	result := 0
	c := 100000
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				sum := nums[i] + nums[j] + nums[k]
				if sum == result {
					continue
				}
				if sum > target {
					if (sum - target) < c {
						result = sum
						c = sum - target
					}
				} else {
					if (target - sum) < c {
						result = sum
						c = target - sum
					}
				}
			}
		}
	}
	return result
}

func Test16() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1)) // 2
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))      // 0
}
