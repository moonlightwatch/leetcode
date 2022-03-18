package L_0304

import "fmt"

// https://leetcode-cn.com/problems/range-sum-query-2d-immutable/

// 304. 二维区域和检索 - 矩阵不可变

// 给定一个二维矩阵 matrix，以下类型的多个请求：

//     计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1, col1) ，右下角 为 (row2, col2) 。

// 实现 NumMatrix 类：

//     NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
//     int sumRegion(int row1, int col1, int row2, int col2) 返回 左上角 (row1, col1) 、右下角 (row2, col2) 所描述的子矩阵的元素 总和 。

type NumMatrix struct {
	m [][]int
}

func Constructor304(matrix [][]int) NumMatrix {
	return NumMatrix{m: matrix}
}

func (M *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for r := row1; r <= row2; r++ {
		for c := col1; c <= col2; c++ {
			sum += M.m[r][c]
		}
	}
	return sum
}

// 执行用时：132 ms, 在所有 Go 提交中击败了7.37% 的用户
// 内存消耗：8.2 MB, 在所有 Go 提交中击败了97.89% 的用户

func Test304() {
	c := Constructor304([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5}})
	fmt.Printf("%v\n", c.SumRegion(2, 1, 4, 3)) // 8
	fmt.Printf("%v\n", c.SumRegion(1, 1, 2, 2)) // 11
	fmt.Printf("%v\n", c.SumRegion(1, 2, 2, 4)) // 12
}
