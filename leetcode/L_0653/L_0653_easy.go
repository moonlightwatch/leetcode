package L_0653

import "fmt"

// https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/

// 653. 两数之和 IV - 输入 BST

// 给定一个二叉搜索树 root 和一个目标结果 k，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 搜索。root=当前节点，targets=其他节点与目标的差，k=目标值
func find(root *TreeNode, targets map[int]bool, k int) bool {
	if root == nil { // 空节点则直接返回false
		return false
	}
	if _, ok := targets[root.Val]; ok { // 如果当前节点的值，是其他节点与目标的差
		return true // 则存在所找的情况（存在两个元素且它们的和等于给定的目标结果）
	} else { // 如果当前节点不是其他节点与目标的差
		targets[k-root.Val] = true // 则将当前节点的值与目标的差加入集合
	}
	// 搜索两个字节点
	return find(root.Left, targets, k) || find(root.Right, targets, k)
}

func findTarget(root *TreeNode, k int) bool {
	return find(root, map[int]bool{}, k)
}

func Test653() {
	fmt.Println(findTarget(_makeTree([]int{5, 3, 6, 2, 4, 100000, 7}), 9))  // true
	fmt.Println(findTarget(_makeTree([]int{5, 3, 6, 2, 4, 100000, 7}), 28)) // false
}

func _makeTree(root []int) *TreeNode {
	if len(root) == 0 || root[0] == 100000 {
		return nil
	}
	result := &TreeNode{Val: root[0]}
	cache := []*TreeNode{result}
	l := len(root)
	for i := 1; i < l; i++ {
		newcache := []*TreeNode{result}
		for _, item := range cache {
			if i < l && root[i] != 100000 {
				item.Left = &TreeNode{Val: root[i]}
				newcache = append(newcache, item.Left)
			}
			i++
			if i < l && root[i] != 100000 {
				item.Right = &TreeNode{Val: root[i]}
				newcache = append(newcache, item.Right)
			}
			i++
		}
		cache = newcache
	}
	return result
}
