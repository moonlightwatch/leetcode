package L_0417

import "fmt"

// https://leetcode-cn.com/problems/pacific-atlantic-water-flow/

// 417. 太平洋大西洋水流问题

// 有一个 m × n 的矩形岛屿，与 太平洋 和 大西洋 相邻。 “太平洋” 处于大陆的左边界和上边界，而 “大西洋” 处于大陆的右边界和下边界。

// 这个岛被分割成一个由若干方形单元格组成的网格。给定一个 m x n 的整数矩阵 heights ， heights[r][c] 表示坐标 (r, c) 上单元格 高于海平面的高度 。

// 岛上雨水较多，如果相邻单元格的高度 小于或等于 当前单元格的高度，雨水可以直接向北、南、东、西流向相邻单元格。水可以从海洋附近的任何单元格流入海洋。

// 返回 网格坐标 result 的 2D列表 ，其中 result[i] = [ri, ci] 表示雨水可以从单元格 (ri, ci) 流向 太平洋和大西洋 。

// 从外至内，分别找出 能流向太平洋的单元格 和 能流向大西洋的单元格
// 取交集
// 对能流入上述交集内的单元格进行广度优先搜索

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func pacificAtlantic(heights [][]int) (ans [][]int) {
	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	bfs := func(x, y int, ocean [][]bool) {
		if ocean[x][y] {
			return
		}
		ocean[x][y] = true
		q := []pair{{x, y}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			for _, d := range dirs {
				if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && !ocean[x][y] && heights[x][y] >= heights[p.x][p.y] {
					ocean[x][y] = true
					q = append(q, pair{x, y})
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		bfs(i, 0, pacific)
	}
	for j := 1; j < n; j++ {
		bfs(0, j, pacific)
	}
	for i := 0; i < m; i++ {
		bfs(i, n-1, atlantic)
	}
	for j := 0; j < n-1; j++ {
		bfs(m-1, j, atlantic)
	}

	for i, row := range pacific {
		for j, ok := range row {
			if ok && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return
}

func Test417() {
	fmt.Println(pacificAtlantic([][]int{{3, 3, 3, 3, 3, 3}, {3, 0, 3, 3, 0, 3}, {3, 3, 3, 3, 3, 3}}))
	// [[0,2],[1,0],[1,1],[1,2],[2,0],[2,1],[2,2]]
	fmt.Println(pacificAtlantic([][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}))
	// [[0,0],[0,1],[1,0],[1,1]]
	fmt.Println(pacificAtlantic([][]int{{2, 1}, {1, 2}}))
	// [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]

	fmt.Println(pacificAtlantic([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}))
}
