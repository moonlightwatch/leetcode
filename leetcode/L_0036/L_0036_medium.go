package L_0036

// https://leetcode-cn.com/problems/valid-sudoku/

// 36. 有效的数独

// 请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

//     数字 1-9 在每一行只能出现一次。
//     数字 1-9 在每一列只能出现一次。
//     数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）

// 注意：

//     一个有效的数独（部分已被填充）不一定是可解的。
//     只需要根据以上规则，验证已经填入的数字是否有效即可。
//     空白格用 '.' 表示。

// check 判断参数内没有重复元素，忽略 '.'
func check(items ...byte) bool {
	cache := map[byte]bool{}
	for _, item := range items {
		if item == '.' {
			continue
		}
		if _, ok := cache[item]; !ok {
			cache[item] = true
		} else {
			return false
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	// 检查行和列
	for i := 0; i < 9; i++ {
		if !check(
			board[i][0], // 行
			board[i][1],
			board[i][2],
			board[i][3],
			board[i][4],
			board[i][5],
			board[i][6],
			board[i][7],
			board[i][8],
		) || !check( // 列
			board[0][i],
			board[1][i],
			board[2][i],
			board[3][i],
			board[4][i],
			board[5][i],
			board[6][i],
			board[7][i],
			board[8][i],
		) {
			return false // 如有重复，则返回false
		}
	}
	// 检查九个 3x3 宫
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !check(
				board[i][j],
				board[i][j+1],
				board[i][j+2],
				board[i+1][j],
				board[i+1][j+1],
				board[i+1][j+2],
				board[i+2][j],
				board[i+2][j+1],
				board[i+2][j+2],
			) {
				return false // 如有重复，则返回false
			}
		}
	}
	return true
}
