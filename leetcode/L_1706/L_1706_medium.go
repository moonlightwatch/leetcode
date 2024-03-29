package L_1706

import "fmt"

// https://leetcode-cn.com/problems/where-will-the-ball-fall/

// 1706. 球会落何处

// 用一个大小为 m x n 的二维网格 grid 表示一个箱子。你有 n 颗球。箱子的顶部和底部都是开着的。

// 箱子中的每个单元格都有一个对角线挡板，跨过单元格的两个角，可以将球导向左侧或者右侧。

//     将球导向右侧的挡板跨过左上角和右下角，在网格中用 1 表示。
//     将球导向左侧的挡板跨过右上角和左下角，在网格中用 -1 表示。

// 在箱子每一列的顶端各放一颗球。每颗球都可能卡在箱子里或从底部掉出来。如果球恰好卡在两块挡板之间的 "V" 形图案，或者被一块挡导向到箱子的任意一侧边上，就会卡住。

// 返回一个大小为 n 的数组 answer ，其中 answer[i] 是球放在顶部的第 i 列后从底部掉出来的那一列对应的下标，如果球卡在盒子里，则返回 -1 。

// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 100
// grid[i][j] 为 1 或 -1

func findBall(grid [][]int) []int {
	// 逐一逐行判断下落情况，收集下落结果并返回
	result := make([]int, len(grid[0]))
	for n := 0; n < len(grid[0]); n++ {
		i := n
		m := 0
		for ; m < len(grid); m++ {
			if grid[m][i] == 1 {
				if i == len(grid[0])-1 || grid[m][i+1] == -1 {
					break
				}
				i++
			}
			if grid[m][i] == -1 {
				if i == 0 || grid[m][i-1] == 1 {
					break
				}
				i--
			}
		}
		if m == len(grid) {
			result[n] = i
		} else {
			result[n] = -1
		}
	}
	return result
}
func findBall_(grid [][]int) []int {
	// 逐行判断结果

	result := make([]int, len(grid[0]))
	// 初始化状态
	for i := 0; i < len(grid[0]); i++ {
		result[i] = i
	}
	for _, line := range grid {
		for n, i := range result {
			if result[n] == -1 {
				continue
			}
			if line[i] == 1 {
				if i == len(grid[0])-1 || line[i+1] == -1 {
					result[n] = -1
					break
				}
				result[n]++
			}
			if line[i] == -1 {
				if i == 0 || line[i-1] == 1 {
					result[n] = -1
					break
				}
				result[n]--
			}
		}
	}
	return result

}

func Test1706() {
	fmt.Println(findBall_([][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}})) // [1,-1,-1,-1,-1]
	fmt.Println(findBall_([][]int{{-1}}))                                                                                             // [-1]
	fmt.Println(findBall_([][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}))       // [0,1,2,3,4,-1]
}
