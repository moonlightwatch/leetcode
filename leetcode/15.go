package leetcode

import (
	"fmt"
	"sort"
)

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。

func threeSum(nums []int) [][]int {
	l := len(nums)
	if l < 3 {
		return [][]int{}
	}
	sort.Ints(nums) // 排个序
	result := [][]int{}

	for i := 0; i < l && nums[i] <= 0; i++ {
		Lindex := i + 1
		Rindex := l - 1

		if nums[i] > 0 {
			return result
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for Lindex < Rindex {
			sum := nums[i] + nums[Lindex] + nums[Rindex]

			if sum > 0 {
				Rindex -= 1
			} else if sum < 0 {
				Lindex += 1
			} else {
				result = append(result, []int{nums[i], nums[Lindex], nums[Rindex]})
				for Lindex <= Rindex && nums[Rindex] == nums[Rindex-1] {
					Rindex -= 1
				}
				for Lindex <= Rindex && nums[Lindex] == nums[Lindex+1] {
					Lindex += 1
				}
				Rindex -= 1
				Lindex += 1
			}
		}
	}

	return result
}

func Test15() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))                  // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))                       // [[-2,0,2],[-2,1,1]]
	fmt.Println(threeSum([]int{0, 0, 0}))                              // [[0,0,0]]
	fmt.Println(threeSum([]int{0, 0, 0, 0}))                           // [[0,0,0]]
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4})) // [[-4,0,4],[-4,1,3],[-3,-1,4],[-3,0,3],[-3,1,2],[-2,-1,3],[-2,0,2],[-1,-1,2],[-1,0,1]]
}
