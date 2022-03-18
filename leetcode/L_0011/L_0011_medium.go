package L_0011

import (
	"fmt"
)

// https://leetcode-cn.com/problems/container-with-most-water/

// 11. 盛最多水的容器

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

// 返回容器可以储存的最大水量。

// 说明：你不能倾斜容器。

func maxArea(height []int) int {
	max := 0
	leftIndex := 0
	rightIndex := len(height) - 1
	for leftIndex < rightIndex {
		lowHeight := height[leftIndex]
		if height[rightIndex] < height[leftIndex] {
			lowHeight = height[rightIndex]
			rightIndex -= 1
		} else {
			leftIndex += 1
		}
		area := (rightIndex - leftIndex + 1) * lowHeight
		if area > max {
			max = area
		}
	}
	return max
}

func Test11() {
	fmt.Println(maxArea([]int{1, 1}))                      // 1
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})) // 49
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))             // 16
	fmt.Println(maxArea([]int{1, 2, 1}))                   // 2
}
