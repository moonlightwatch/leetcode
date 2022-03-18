package L_0589

// https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/

// 589. N 叉树的前序遍历

// 给定一个 n 叉树的根节点  root ，返回 其节点值的 前序遍历 。

// n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

func preorder(root *Node) []int {
	result := []int{}
	if root == nil {
		return result
	}
	cache := []*Node{root}
	for len(cache) > 0 {
		node := cache[len(cache)-1]
		cache = cache[:len(cache)-1]
		result = append(result, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			cache = append(cache, node.Children[i])
		}
	}
	return result
}

type Node struct {
	Val      int
	Children []*Node
}
