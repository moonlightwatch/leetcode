package L_1305

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/all-elements-in-two-binary-search-trees/

// 1305. 两棵二叉搜索树中的所有元素

// 给你 root1 和 root2 这两棵二叉搜索树。请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	result := []int{} // 定义结果集

	// 遍历，同时搜索两颗树
	cache := []*TreeNode{root1, root2}
	i := 0
	for len(cache) > i {
		if cache[i] != nil { // 不处理nil
			result = append(result, cache[i].Val) // 将遍历到的节点值加入结果集
			cache = append(cache, cache[i].Left)
			cache = append(cache, cache[i].Right)
		}
		i++
	}

	// 排个序
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}

func Test1305() {
	root1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(getAllElements(root1, root2)) // [0,1,1,2,3,4]
}
