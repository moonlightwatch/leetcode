package L_1981

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/minimize-the-difference-between-target-and-chosen-elements/

// 1981. 最小化目标值与所选元素的差

// 给你一个大小为 m x n 的整数矩阵 mat 和一个整数 target 。

// 从矩阵的 每一行 中选择一个整数，你的目标是 最小化 所有选中元素之 和 与目标值 target 的 绝对差 。

// 返回 最小的绝对差 。

// a 和 b 两数字的 绝对差 是 a - b 的绝对值。

// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 70
// 1 <= mat[i][j] <= 70
// 1 <= target <= 800

// 动态规划
// 逐行计算每一个元素与上一行的差
// 记录到最后一行，即所有可能的结果
// 筛选合适的结果即可

func minimizeTheDifference(mat [][]int, target int) int {
	m := len(mat)
	n := len(mat[0])
	lastCache := []int{target} // 记录上一行内容，初始填入target，即计算与target的差
	for i := 0; i < m; i++ {
		curCache := make(map[int]bool, n*len(lastCache)) // 计算当前行的内容，使用map，减少重复元素的反复计算
		for j := 0; j < n; j++ {
			for _, v := range lastCache { // 循环上一行内容，与当前元素求差
				curCache[v-mat[i][j]] = true
			}
		}
		lastCache = []int{}
		temp := -10000 // 设定一个负数的最大值，用于剔除较小的负值。因为较小的负值在后续计算中没有用。再求差的话，绝对差肯定越来越大
		for v, _ := range curCache {
			if v >= 0 { // 大于 0 的，加到上一行内容中
				lastCache = append(lastCache, v)
			} else if v > temp { // 小于 0 的，寻找最大值
				temp = v
			}
		}
		lastCache = append(lastCache, temp) // 将小于0的最大值填入上一行内容
	}

	// 最后求出最小绝对值即可
	var min float64 = 10000
	for _, item := range lastCache {
		v := math.Abs(float64(item))
		if v < min {
			min = v
		}
	}
	return int(min)
}

func Test1981() {
	fmt.Println(minimizeTheDifference([][]int{
		{10, 3, 7, 7, 9, 6, 9, 8, 9, 5},
		{1, 1, 6, 8, 6, 7, 7, 9, 3, 9},
		{3, 4, 4, 1, 3, 6, 3, 3, 9, 9},
		{6, 9, 9, 3, 8, 7, 9, 6, 10, 6},
	}, 5)) // 3
	fmt.Println(minimizeTheDifference([][]int{{3, 5}, {4, 4}, {1, 7}, {1, 1}}, 67)) // 50

	fmt.Println(minimizeTheDifference([][]int{{1, 2, 9, 8, 7}}, 6)) // 1

	fmt.Println(minimizeTheDifference([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 13)) // 0

	fmt.Println(minimizeTheDifference([][]int{{1}, {2}, {3}}, 100)) // 94

}
