package L_0037

import "fmt"

// https://leetcode-cn.com/problems/sudoku-solver/

// 37. 解数独

// 编写一个程序，通过填充空格来解决数独问题。

// 数独的解法需 遵循如下规则：

//     数字 1-9 在每一行只能出现一次。
//     数字 1-9 在每一列只能出现一次。
//     数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）

// 数独部分空格内已填入了数字，空白格用 '.' 表示。

func trySetNum(board [][]byte) bool {
	// 遍历所有位置
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			// 如果已经填了值，则跳过
			if board[x][y] != '.' {
				continue
			}
			// 尝试1-9所有值
			for v := '1'; v <= '9'; v++ {
				if isValid(board, x, y, byte(v)) { // 如果值合法
					board[x][y] = byte(v) // 则向此位置设置值
					if trySetNum(board) { // 进行下一次
						return true // 递归没有问题，则确认找到解
					}
					board[x][y] = '.' // 如果有问题，则回退此次设置，进行下一次
				}
			}
			return false // 若尝试9个值都不能填，则返回错误。
		}
	}
	return true // 遍历完成，则认为找到解
}

// 检查位置上的值是否合适。r (row), c (column), v (value)
func isValid(board [][]byte, r int, c int, v byte) bool {
	for i := 0; i < 9; i++ {
		if board[r][i] == v {
			return false
		}
		if board[i][c] == v {
			return false
		}
	}
	rowIndex := (r / 3) * 3
	colIndex := (c / 3) * 3
	for x := rowIndex; x < rowIndex+3; x++ {
		for y := colIndex; y < colIndex+3; y++ {
			if board[x][y] == v {
				return false
			}
		}
	}
	return true
}

func solveSudoku(board [][]byte) {
	trySetNum(board)
}

func print(b [][]byte) {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			fmt.Printf(" %c", b[x][y])
		}
		fmt.Println()
	}
	fmt.Println()
}

func Test37() {
	fmt.Println("%d", '9')
	v := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	print(v)
	solveSudoku(v)
	print(v)

	v = [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'}}
	print(v)
	solveSudoku(v)
	print(v)

}
