package L_0905

import "fmt"

// https://leetcode-cn.com/problems/sort-array-by-parity/

// 905. 按奇偶排序数组

// 给你一个整数数组 nums，将 nums 中的的所有偶数元素移动到数组的前面，后跟所有奇数元素。

// 返回满足此条件的 任一数组 作为答案。

// 双指针
func sortArrayByParity(nums []int) []int {
	a := 0
	b := len(nums) - 1
	for a < b {
		for a < b && nums[a]%2 == 0 {
			a++
		}
		nums[a], nums[b] = nums[b], nums[a]
		b--
	}
	return nums
}

// // 冒泡（贼慢）
// func sortArrayByParity(nums []int) []int {
// 	l := len(nums)
// 	for i := 0; i < l; i++ {
// 		if nums[i]%2 == 0 {
// 			continue
// 		}
// 		for j := i + 1; j < l; j++ {
// 			if nums[j]%2 == 0 {
// 				nums[i], nums[j] = nums[j], nums[i]
// 				break
// 			}
// 		}
// 	}
// 	return nums
// }

func Test905() {
	fmt.Println(sortArrayByParity([]int{0, 2}))       // 0,2
	fmt.Println(sortArrayByParity([]int{0, 1}))       // 0,1
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4})) // 2,4,3,1
	fmt.Println(sortArrayByParity([]int{0}))          // 0
}
