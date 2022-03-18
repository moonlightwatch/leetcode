package L_0697

import "fmt"

// https://leetcode-cn.com/problems/degree-of-an-array/

// 697. 数组的度

// 给定一个非空且只包含非负数的整数数组 nums，数组的 度 的定义是指数组里任一元素出现频数的最大值。

// 你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。

/*

思路：

遍历数组，记录每种数字 出现的次数 、最开始的位置 和 最后的位置
找到出现次数最多的数字，分别计算其 最开始的位置 和 最后的位置 之间的跨度
跨度最小的一个，其值就是结果

*/

func findShortestSubArray(nums []int) int {
	tempCount := map[int][]int{} // 计数缓存
	maxCount := 0                // nums 的度
	for index, item := range nums {
		if _, ok := tempCount[item]; !ok {
			tempCount[item] = []int{0, index, index}
		}
		tempCount[item][0]++
		tempCount[item][2] = index
		if tempCount[item][0] > maxCount {
			maxCount = tempCount[item][0]
		}
	}
	minLen := len(nums)
	for _, item := range tempCount {
		if item[0] == maxCount && item[2]-item[1]+1 < minLen {
			minLen = item[2] - item[1] + 1
		}
	}
	return minLen
}

// 执行用时：36 ms, 在所有 Go 提交中击败了68.55% 的用户
// 内存消耗：6.6 MB, 在所有 Go 提交中击败了70.44% 的用户

func Test697() {
	fmt.Printf("%v\n", findShortestSubArray([]int{1, 2, 2, 3, 1}))                         // 2
	fmt.Printf("%v\n", findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))                   // 6
	fmt.Printf("%v\n", findShortestSubArray([]int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2})) // 7
}
