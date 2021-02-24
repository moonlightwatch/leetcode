package leetcode

import (
	"fmt"
)

// 给你一个二维整数数组 matrix， 返回 matrix 的 转置矩阵 。
// 矩阵的 转置 是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。

/*

思路：

就硬遍历

*/

func transpose(matrix [][]int) [][]int {
	result := [][]int{}
	w := len(matrix[0])
	h := len(matrix)
	for j := 0; j < w; j++ {
		result = append(result, make([]int, h))
		for i := 0; i < h; i++ {
			result[j][i] = matrix[i][j]
		}
	}
	return result
}

// 执行用时：12 ms, 在所有 Go 提交中击败了59.77% 的用户
// 内存消耗：6.1 MB, 在所有 Go 提交中击败了100.00% 的用户

// func transpose(matrix [][]int) [][]int {
// 	result := [][]int{}
// 	for j := 0; j < len(matrix[0]); j++ {
// 		temp := []int{}
// 		for i := 0; i < len(matrix); i++ {
// 			temp = append(temp, matrix[i][j])
// 		}
// 		result = append(result, temp)
// 	}
// 	return result
// }

//执行用时：8 ms, 在所有 Go 提交中击败了88.72% 的用户
// 内存消耗：6.2 MB, 在所有 Go 提交中击败了19.92% 的用户

// // 这个方法就是遍历，硬算。比较慢
// func transpose(matrix [][]int) [][]int {
// 	linesCount := len(matrix)
// 	itemsCount := len(matrix[0])
// 	result := make([][]int, itemsCount)
// 	for i := 0; i < itemsCount; i++ {
// 		result[i] = make([]int, linesCount)
// 		for j := 0; j < linesCount; j++ {
// 			result[i][j] = matrix[j][i]
// 		}
// 	}
// 	return result
// }

func Test867() {
	fmt.Printf("%v\n", transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Printf("%v\n", transpose([][]int{{1, 2, 3}, {4, 5, 6}}))
}
