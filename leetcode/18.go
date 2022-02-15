package leetcode

import (
	"fmt"
	"sort"
)

// 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//     0 <= a, b, c, d < n
//     a、b、c 和 d 互不相同
//     nums[a] + nums[b] + nums[c] + nums[d] == target
// 你可以按 任意顺序 返回答案 。
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)
	for a := 0; a < n; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < n; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			for c := b + 1; c < n; c++ {
				if c > b+1 && nums[c] == nums[c-1] {
					continue
				}
				for d := c + 1; d < n; d++ {
					if d > c+1 && nums[d] == nums[d-1] {
						continue
					}
					if nums[a]+nums[b]+nums[c]+nums[d] == target {
						result = append(result, []int{nums[a], nums[b], nums[c], nums[d]})
					}
				}
			}
		}
	}
	return result
}
func Test18() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))        // [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))             // [[2,2,2,2]]
	fmt.Println(fourSum([]int{-5, 5, 4, -3, 0, 0, 4, -2}, 4)) // [[-5,0,4,5],[-3,-2,4,5]]
}
