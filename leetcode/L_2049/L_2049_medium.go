package L_2049

import "fmt"

// https://leetcode-cn.com/problems/count-nodes-with-the-highest-score/

// 2049. 统计最高分的节点数目

// 给你一棵根节点为 0 的 二叉树 ，它总共有 n 个节点，节点编号为 0 到 n - 1 。
// 同时给你一个下标从 0 开始的整数数组 parents 表示这棵树，其中 parents[i] 是节点 i 的父节点。由于节点 0 是根，所以 parents[0] == -1 。
// 一个子树的 大小 为这个子树内节点的数目。每个节点都有一个与之关联的 分数 。
// 求出某个节点分数的方法是，将这个节点和与它相连的边全部 删除 ，剩余部分是若干个 非空 子树，这个节点的 分数 为所有这些子树 大小的乘积 。
// 请你返回有 最高得分 节点的 数目 。

// n == parents.length
// 2 <= n <= 10⁵
// parents[0] == -1
// 对于 i != 0 ，有 0 <= parents[i] <= n - 1
// parents 表示一棵二叉树。

var n int
var childs [][]int

// 思路：深度优先遍历，边遍历边统计
//   节点得分 =（节点总数 - 这个节点的子树Size）* 节点左子树Size * 节点右子树Size
//   需要注意的是，避免空子树Size为0，导致的计算得分计算错误

// countTree 递归。
// 输入：
//    节点数量，用于统计节点得分
//    节点和子节点关系，用于深度优先遍历
//    当前节点
// 返回：
//    节点Size
//    节点的得分
//    节点子树最高得分数量统计
func countTree(i int) (int, int, int) {
	curSize := 1  // 节点Size，默认1
	curScore := 1 // 节点得分，因为后面是乘法，所以记为1
	maxCount := 0 // 节点最高分数量
	maxScore := 0 // 节点最高分分数
	// 遍历子节点
	for _, c := range childs[i] {
		// 递归计算子节点数据
		childSize, childScore, childMaxCount := countTree(c)
		curSize += childSize       // Size直接累加即可
		curScore *= childSize      // 得分是（左子树Size * 右子树Size），直接累乘即可
		if childScore > maxScore { // 记录左右子节点的较高得分，和较高得分数量
			maxScore = childScore
			maxCount = childMaxCount
		} else if maxScore != 0 && childScore == maxScore { // 处理两个子节点得分相等的情况
			maxCount += childMaxCount
		}
	}
	// 根节点的得分是左右子树的得分乘积，需要特殊处理
	if n != curSize {
		curScore *= n - curSize
	}
	// 如果当前节点得分高于左右子树中的最高分
	if curScore > maxScore {
		maxScore = curScore // 记录最高分
		maxCount = 1        // 重置最高分数量统计
	} else if curScore == maxScore { // 如果当前节点得分等于左右子树中的最高分
		maxCount++ // // 最高分数量统计 +1
	}
	return curSize, maxScore, maxCount // 返回当前节点数据
}

func countHighestScoreNodes(parents []int) int {
	n = len(parents)
	childs = make([][]int, n) // 保存节点的左右子节点，用于遍历
	// 左右子节点，根节点不处理
	for i := 1; i < n; i++ {
		// 向父节点的字节点内添加自己
		childs[parents[i]] = append(childs[parents[i]], i)
	}
	// 递归便利，直接返回结果
	_, _, result := countTree(0)
	return result
}

func Test2049() {
	fmt.Println(countHighestScoreNodes([]int{-1, 2, 4, 0, 0, 4, 2})) // 5
	fmt.Println(countHighestScoreNodes([]int{-1, 2, 0, 2, 0}))       // 3
	fmt.Println(countHighestScoreNodes([]int{-1, 2, 0}))             // 2
}
