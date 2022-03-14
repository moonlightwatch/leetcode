package leetcode

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/next-permutation/

// 31. 下一个排列

// 整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

//     例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。

// 整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

//     例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
//     类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
//     而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。

// 给你一个整数数组 nums ，找出 nums 的下一个排列。

// 必须 原地 修改，只允许使用额外常数空间。

func nextPermutation(nums []int) {
	signIndex := 101 // 记录第一个小于右侧元素的数
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			signIndex = i - 1
			break
		}
	}
	if signIndex == 101 { // 如果没有，那就是最后一个排列，直接逆序输出
		sort.Slice(nums, func(i, j int) bool {
			return nums[i] < nums[j]
		})
		return
	}
	// 从右至左，找到第一个比上述标记位置元素大的元素，交换其位置，将其后元素从小到大排序
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[signIndex] {
			nums[i], nums[signIndex] = nums[signIndex], nums[i]
			// 排序
			for j := signIndex + 1; j < len(nums); j++ {
				for k := j; k < len(nums); k++ {
					if nums[k] < nums[j] {
						nums[j], nums[k] = nums[k], nums[j]
					}
				}
			}
			break
		}
	}
}

func Test31() {
	items := []int{1, 3, 2}
	nextPermutation(items) // 2 1 3
	fmt.Println(items)

	items = []int{1, 2, 3}
	nextPermutation(items) // 1 3 2
	fmt.Println(items)

	items = []int{3, 2, 1}
	nextPermutation(items) // 1 2 3
	fmt.Println(items)

	items = []int{1, 1, 5}
	nextPermutation(items) // 1 5 1
	fmt.Println(items)

}
