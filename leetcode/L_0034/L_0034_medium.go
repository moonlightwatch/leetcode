package L_0034

import "fmt"

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

// 34. 在排序数组中查找元素的第一个和最后一个位置

// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

// 如果数组中不存在目标值 target，返回 [-1, -1]。

// 进阶：

//     你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？

func searchRange(nums []int, target int) []int {
	// 使用两次二分法分别查找左边界和右边界
	// 执行用时：4 ms, 在所有 Go 提交中击败了96.19% 的用户
	// 内存消耗：3.8 MB, 在所有 Go 提交中击败了76.56% 的用户
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}
	left := 0
	right := len(nums) - 1
	for {
		if nums[left] == target && (left == 0 || nums[left-1] < target) {
			result[0] = left
			break
		}
		if nums[right] == target && (right == 0 || nums[right-1] < target) {
			result[0] = right
			break
		}
		if right-left <= 1 {
			break
		}
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid
		} else {
			right = mid
		}
	}
	left = 0
	right = len(nums) - 1
	for {
		if nums[right] == target && (right == len(nums)-1 || nums[right+1] > target) {
			result[1] = right
			break
		}
		if nums[left] == target && (left == len(nums)-1 || nums[left+1] > target) {
			result[1] = left
			break
		}
		if right-left <= 1 {
			break
		}
		mid := (left + right) / 2
		if target < nums[mid] {
			right = mid
		} else {
			left = mid
		}
	}
	return result
}

func Test34() {
	fmt.Println(searchRange([]int{1}, 1))                 // 0 0
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // 3 4
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6)) // -1 -1
	fmt.Println(searchRange([]int{}, 0))                  // -1 -1
}
