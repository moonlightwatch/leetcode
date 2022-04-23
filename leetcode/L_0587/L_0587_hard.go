package L_0587

import "fmt"

//  https://leetcode-cn.com/problems/erect-the-fence/

// 587. 安装栅栏

// 在一个二维的花园中，有一些用 (x, y) 坐标表示的树。
// 由于安装费用十分昂贵，你的任务是先用最短的绳子围起所有的树。
// 只有当所有的树都被绳子包围时，花园才能围好栅栏。
// 你需要找到正好位于栅栏边界上的树的坐标。

// 思路：
//     栅栏边界上，每条边界都具备特性：所有其他的树，都在边界一侧或者边界线上
//     因此，遍历所有树的连线，若其他树都在连线其同一侧，则记录连线上的树。
//         所有记录到的树即为边界上的树。

// side 计算 点 p3 与 线 (p1,p2) 的位置关系，等于 0 则共线
func side(p1 []int, p2 []int, p3 []int) int {
	return (p1[0]-p3[0])*(p2[1]-p3[1]) - (p1[1]-p3[1])*(p2[0]-p3[0])
}

func outerTrees(trees [][]int) [][]int {
	l := len(trees)
	if l < 4 { // 小于三条边，直接就返回原数据
		return trees
	}
	resultIndex := map[int]bool{} // 记录结果

	// 先找到最右的一个点，此点一定是边界上的点
	curPointIndex := 0
	for i := 0; i < l; i++ {
		if trees[i][0] > trees[curPointIndex][0] {
			curPointIndex = i
		}
	}
	resultIndex[curPointIndex] = true
	firstPoint := curPointIndex
	for { // 循环直到没有发现新的点
		nextIndex := (curPointIndex + 1) % l
		// 遍历其他点，若有点 处于线 (curPointIndex, nextIndex) 小于0一侧，则将nextIndex设为此点
		// 完成后，所有的点都在 线 (curPointIndex, nextIndex) 的大于0一侧或者共线
		for i := 0; i < l; i++ {
			if side(trees[curPointIndex], trees[nextIndex], trees[i]) < 0 {
				nextIndex = i
			}
		}

		resultIndex[nextIndex] = true

		// 找出所有共线的点，也记入结果
		for i := 0; i < l; i++ {
			if _, ok := resultIndex[i]; ok {
				continue
			}
			if side(trees[curPointIndex], trees[nextIndex], trees[i]) == 0 {
				resultIndex[i] = true
			}
		}
		// 若找到第一个点，则搜索结束
		if nextIndex == firstPoint {
			result := [][]int{}
			for i, _ := range resultIndex {
				result = append(result, trees[i])
				i++
			}
			return result
		}
		curPointIndex = nextIndex
	}
}

func Test587() {
	// [[7,4],[3,0],[1,2],[2,5],[5,5],[4,5],[1,4],[2,1],[3,5],[0,3],[6,5],[7,2],[7,3],[4,0],[5,0],[6,1]]
	fmt.Println(outerTrees([][]int{{3, 0}, {4, 0}, {5, 0}, {6, 1}, {7, 2}, {7, 3}, {7, 4}, {6, 5}, {5, 5}, {4, 5}, {3, 5}, {2, 5}, {1, 4}, {1, 3}, {1, 2}, {2, 1}, {4, 2}, {0, 3}}))
	fmt.Println(outerTrees([][]int{{1, 2}, {2, 2}, {4, 2}}))                         // 1,2 2,2 4,2
	fmt.Println(outerTrees([][]int{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}})) // [[1,1],[2,0],[4,2],[3,3],[2,4]]
}
