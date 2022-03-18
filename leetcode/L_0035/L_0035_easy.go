package L_0035

import "fmt"

// https://leetcode-cn.com/problems/search-insert-position/

// 35. 搜索插入位置

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

// 请必须使用时间复杂度为 O(log n) 的算法。

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	if nums[0] >= target {
		return 0
	} else if nums[r] < target {
		return r + 1
	}
	for {
		if target == nums[l] {
			return l
		} else if target == nums[r] || r-l == 1 {
			return r
		}

		mid := (l + r) / 2
		if target < nums[mid] {
			r = mid
		} else {
			l = mid
		}
	}
}

func Test35() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5)) // 2
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2)) // 1
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7)) // 4
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0)) // 0
	fmt.Println(searchInsert([]int{1}, 0))          // 0
}
