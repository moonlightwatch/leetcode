package L_0011

import (
	"fmt"
)

// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

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
