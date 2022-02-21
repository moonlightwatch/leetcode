package leetcode

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

// 4. 寻找两个正序数组的中位数
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
// 算法的时间复杂度应该为 O(log (m+n)) 。

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 执行用时：12 ms, 在所有 Go 提交中击败了85.31% 的用户
	// 内存消耗：5.5 MB, 在所有 Go 提交中击败了36.52% 的用户
	nums := append(nums1, nums2...) // 合并
	sort.Ints(nums)                 // 排序
	// 取中位数
	if len(nums)%2 == 0 { // 偶数个
		i := len(nums) / 2
		return float64(nums[i]+nums[i-1]) / 2.0
	} else { // 奇数个
		return float64(nums[len(nums)/2])
	}
}

func Test4() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))    // 2.00000
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // 2.50000
}
