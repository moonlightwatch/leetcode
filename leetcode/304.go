package leetcode

import "fmt"

// 给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。

type NumMatrix struct {
	m [][]int
}

func Constructor304(matrix [][]int) NumMatrix {
	return NumMatrix{m: matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for r := row1; r <= row2; r++ {
		for c := col1; c <= col2; c++ {
			sum += this.m[r][c]
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
