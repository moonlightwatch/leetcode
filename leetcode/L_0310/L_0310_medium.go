package L_0310

import (
	"fmt"
)

// https://leetcode-cn.com/problems/minimum-height-trees/

// 310. 最小高度树

// 树是一个无向图，其中任何两个顶点只通过一条路径连接。 换句话说，一个任何没有简单环路的连通图都是一棵树。

// 给你一棵包含 n 个节点的树，标记为 0 到 n - 1 。
// 给定数字 n 和一个有 n - 1 条无向边的 edges 列表（每一个边都是一对标签），其中 edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间存在一条无向边。

// 可选择树中任何一个节点作为根。当选择节点 x 作为根节点时，设结果树的高度为 h 。在所有可能的树中，具有最小高度的树（即，min(h)）被称为 最小高度树 。

// 请你找到所有的 最小高度树 并按 任意顺序 返回它们的根节点标签列表。
// 树的 高度 是指根节点和叶子节点之间最长向下路径上边的数量。

// 1 <= n <= 2 * 10^4
// edges.length == n - 1
// 0 <= ai, bi < n
// ai != bi
// 所有 (ai, bi) 互不相同
// 给定的输入 保证 是一棵树，并且 不会有重复的边

func findMinHeightTrees(n int, edges [][]int) []int {
	nodeMap := make([][]int, n)
	for i := 0; i < n; i++ {
		nodeMap[i] = []int{}
	}
	for i := 0; i < n-1; i++ {
		nodeMap[edges[i][0]] = append(nodeMap[edges[i][0]], edges[i][1])
		nodeMap[edges[i][1]] = append(nodeMap[edges[i][1]], edges[i][0])
	}

	parents := make([]int, n)

	bfs := func(node int) int {
		cache := []int{node}
		passed := map[int]bool{node: true}
		n := -1
		for len(cache) > 0 {
			n, cache = cache[0], cache[1:]
			for _, v := range nodeMap[n] {
				if !passed[v] {
					passed[v] = true
					parents[v] = n
					cache = append(cache, v)
				}
			}
		}
		return n
	}
	node1 := bfs(0)
	node2 := bfs(node1)
	parents[node1] = -1
	pathNodes := []int{}
	for node2 != -1 {
		pathNodes = append(pathNodes, node2)
		node2 = parents[node2]
	}
	l := len(pathNodes)
	if l%2 == 0 {
		return []int{pathNodes[l/2-1], pathNodes[l/2]}
	}
	return []int{pathNodes[l/2]}
}

func Test310() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))                 // 1
	fmt.Println(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}})) // 3 4
	fmt.Println(findMinHeightTrees(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {3, 4}, {4, 5}})) // 3
}
