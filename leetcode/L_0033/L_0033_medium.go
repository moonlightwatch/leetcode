package L_0033

import "fmt"

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

// 33. 搜索旋转排序数组

// 整数数组 nums 按升序排列，数组中的值 互不相同 。

// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
// 例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

func search(nums []int, target int) int {
	// 二分法查找
	// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
	// 内存消耗：2.4 MB, 在所有 Go 提交中击败了99.87% 的用户

	left := 0              // 搜索起点
	right := len(nums) - 1 // 搜索终点
	for {
		// 判断是否命中目标
		if nums[left] == target {
			return left
		} else if nums[right] == target {
			return right
		} else if right-left < 2 {
			return -1
		}

		// 确定中间位置
		mid := (left + right) / 2
		if nums[0] < nums[mid] { // 若前半部分有序
			if target > nums[0] && target < nums[mid] { // 判断target在不在前半部分
				right = mid // 若在前半部分，则重设搜索终点
			} else {
				left = mid // 若在后半部分，则重设搜索起点
			}
		} else { // 若后半部分有序
			if target > nums[mid] && target < nums[right] { // target 在后半部分
				left = mid // 若在后半部分，则重设搜索起点
			} else {
				right = mid // 若在前半部分，则重设搜索终点
			}
		}
	}
}

func Test33() {
	fmt.Println(search([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8)) // 4
	fmt.Println(search([]int{1, 3, 5}, 3))                // 1
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))    // 4
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))    // -1
	fmt.Println(search([]int{1}, 0))                      // -1
}
