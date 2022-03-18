package L_1380

import "fmt"

// 给你一个 m * n 的矩阵，矩阵中的数字 各不相同 。请你按 任意 顺序返回矩阵中的所有幸运数。
// 幸运数是指矩阵中满足同时下列两个条件的元素：
//     在同一行的所有元素中最小
//     在同一列的所有元素中最大

func luckyNumbers(matrix [][]int) []int {
	// 执行用时：12 ms, 在所有 Go 提交中击败了100.00% 的用户
	// 内存消耗：6.4 MB, 在所有 Go 提交中击败了95.45% 的用户
	// 通过测试用例：104 / 104
	result := []int{}
	for _, nums := range matrix { // 循环各行
		lucky := 100000 // 当前行的最小值
		minX := 0       // 最小值所在列索引

		// 求当前行的最小值
		for x, num := range nums {
			if num < lucky {
				lucky = num
				minX = x
			}
		}

		// 判断当前行的最小值是否所在列的最大值
		for i := 0; i < len(matrix); i++ {
			if matrix[i][minX] > lucky { // 列中有更大的值，则置为0并跳出
				lucky = 0
				break
			}
		}

		// 若非0则是幸运数
		if lucky != 0 {
			result = append(result, lucky)
		}
	}
	return result
}

func Test1380() {
	fmt.Println(luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}))          // [15]
	fmt.Println(luckyNumbers([][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}})) // [12]
	fmt.Println(luckyNumbers([][]int{{7, 8}, {1, 2}}))                                // [7]
}
