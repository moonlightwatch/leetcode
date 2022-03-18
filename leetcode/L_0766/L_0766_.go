package L_0766

import "fmt"

// 给你一个 m x n 的矩阵 matrix 。如果这个矩阵是托普利茨矩阵，返回 true ；否则，返回 false 。
// 如果矩阵上每一条由左上到右下的对角线上的元素都相同，那么这个矩阵是 托普利茨矩阵 。

/*

思路：

除了第一行和第一列，每个元素都和其左上元素比较一下。若有不同则返回 false

*/

func isToeplitzMatrix(matrix [][]int) bool {
	for x, line := range matrix {
		if x == 0 {
			continue
		}
		for y, item := range line {
			if y == 0 {
				continue
			}
			if item != matrix[x-1][y-1] {
				return false
			}
		}
	}
	return true
}

// 执行用时：12 ms, 在所有 Go 提交中击败了68.54% 的用户
// 内存消耗：4.4 MB, 在所有 Go 提交中击败了67.42% 的用户

func Test766() {
	fmt.Printf("%v\n", isToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}))
	fmt.Printf("%v\n", isToeplitzMatrix([][]int{{1, 2}, {2, 2}}))
}
