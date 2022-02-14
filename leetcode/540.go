package leetcode

import (
	"fmt"
)

// 给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
// 请你找出并返回只出现一次的那个数。
// 你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。

func singleNonDuplicate(nums []int) int {
	i := 0
	for i = 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			break
		}
	}
	return nums[i]
}

func Test540() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8})) // 2
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))    // 10
	fmt.Println(singleNonDuplicate([]int{6}))                         // 6
	fmt.Println(singleNonDuplicate([]int{1, 1, 2}))                   // 2
	fmt.Println(singleNonDuplicate([]int{1, 2, 2}))                   // 1
}
