package L_1791

import "fmt"

// 有一个无向的 星型 图，由 n 个编号从 1 到 n 的节点组成。星型图有一个 中心 节点，并且恰有 n - 1 条边将中心节点与其他每个节点连接起来。
// 给你一个二维整数数组 edges ，其中 edges[i] = [ui, vi] 表示在节点 ui 和 vi 之间存在一条边。请你找出并返回 edges 所表示星型图的中心节点。

func findCenter(edges [][]int) int {
	if len(edges) == 1 {
		return 0
	}
	for i := 1; i < len(edges); i++ {
		if i == 0 {
			continue
		}
		if edges[0][0] == edges[i][0] || edges[0][0] == edges[i][1] {
			return edges[0][0]
		}
		if edges[0][1] == edges[i][0] || edges[0][1] == edges[i][1] {
			return edges[0][1]
		}
	}

	return 0
}

func Test1791() {

	fmt.Println(findCenter([][]int{{1, 2}, {2, 3}, {4, 2}}))         // 2
	fmt.Println(findCenter([][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}})) // 1

}
