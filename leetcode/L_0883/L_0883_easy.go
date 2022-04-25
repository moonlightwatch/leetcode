package L_0883

import "fmt"

// https://leetcode-cn.com/problems/projection-area-of-3d-shapes/

// 883. 三维形体投影面积

// 在 n x n 的网格 grid 中，我们放置了一些与 x，y，z 三轴对齐的 1 x 1 x 1 立方体。

// 每个值 v = grid[i][j] 表示 v 个正方体叠放在单元格 (i, j) 上。

// 现在，我们查看这些立方体在 xy 、yz 和 zx 平面上的投影。

// 投影 就像影子，将 三维 形体映射到一个 二维 平面上。从顶部、前面和侧面看立方体时，我们会看到“影子”。

// 返回 所有三个投影的总面积 。

func projectionArea(grid [][]int) int {
	n := len(grid)
	result := 0

	xz := make([]int, n)
	yz := make([]int, n)

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if grid[x][y] > 0 {
				result += 1 // (x, y) 平面投影
			}
			if xz[x] < grid[x][y] {
				xz[x] = grid[x][y] // (x, z) 平面投影，取最大值
			}
			if yz[y] < grid[x][y] {
				yz[y] = grid[x][y] // (y, z) 平面投影，取最大值
			}
		}
	}
	// 累计投影面积
	for i := 0; i < n; i++ {
		result += xz[i] + yz[i]
	}
	return result
}

func Test883() {
	fmt.Println(projectionArea([][]int{{1, 2}, {3, 4}})) // 17
	fmt.Println(projectionArea([][]int{{2}}))            // 5
	fmt.Println(projectionArea([][]int{{1, 0}, {0, 2}})) // 8
}
