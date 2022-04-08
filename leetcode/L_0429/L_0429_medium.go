package L_0429

import "fmt"

// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/

// 429. N 叉树的层序遍历

// 给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

// 树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

//  Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	result = append(result, []int{root.Val})
	curLevelNodes := []*Node{root}
	level := 0
	for {
		nextLevelNodes := []*Node{}
		level++
		for i := 0; i < len(curLevelNodes); i++ {
			nextLevelNodes = append(nextLevelNodes, curLevelNodes[i].Children...)
		}
		l := len(nextLevelNodes)
		if l == 0 {
			break
		}
		result = append(result, make([]int, l))
		for i := 0; i < l; i++ {
			result[level][i] = nextLevelNodes[i].Val
		}
		curLevelNodes = nextLevelNodes
	}
	return result
}

func Test429() {
	fmt.Println(levelOrder(tree01())) // [[1],[3,2,4],[5,6]]
}

func tree01() *Node {
	return &Node{
		Val: 1,
		Children: []*Node{
			{Val: 2},
			{Val: 3, Children: []*Node{
				{Val: 5},
				{Val: 6},
			}},
			{Val: 4},
		},
	}
}
