package L_0440

import (
	"fmt"
)

// https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/

// 440. 字典序的第K小数字

// 给定整数 n 和 k，返回  [1, n] 中字典序第 k 小的数字。

// 1 <= k <= n <= 10^9

// 计算小于 n 的子孙节点数量
func countNodes(num, n int) int {
	// 根据规律：
	// 同层的最小值是 (上一层最小值)×10，即  [1，10, 100, 1000, 10000...]
	// 同层的最大值是 (上一层最大值×10)+9，即  [9 ，99, 999, 9999, 99999...]
	// 每层节点数量就是  (同层的最大值 - 同层的最小值 + 1)
	// 我们只取值小于 n 的节点

	if num > n { // 若节点值大于n，直接返回 0（没有合适的子孙节点）
		return 0
	}
	count := 0      // 节点计数
	left := num     // 本层最小值，即最左侧的节点
	right := num    // 本层最大值，即最右侧的节点
	for left <= n { // 循环层数直到当前层的最小值比 n 大，则之后的节点都不合适
		if right > n { // 我们只取值小于 n 的节点，所以当最大值比 n 大的时候，我们只当本层最大值是 n
			count += (n - left + 1) // 累加本层节点数量
		} else { // 最大值比 n 大的时候，我们取整层的数量
			count += (right - left + 1) // 累加本层节点数量
		}
		left = left * 10     // 按照规律计算下一层最小值
		right = right*10 + 9 // 按照规律计算下一层最大值
	}
	return count // 返回累加结果
}

func findKthNumber(n int, k int) int {
	num := 1 // 第一个节点是 1
	k--      // 由于第一个节点已知，所以计数减一

	// 构建数字0-9有序组成的10叉树，前序遍历树，即是字典序
	// 通过计算子孙节点数量，得到遍历子树所消耗的计数，可以避免遍历所带来的消耗
	// 若节点的子孙节点数量小于计数，则认为遍历完其子树后，仍未到达目标值，继续遍历同层的下一个子树。
	// 若节点的子孙节点数量大于计数，则认为目标值在其子树内，开始按照上述方法遍历其子树。
	for k > 0 {
		count := countNodes(num, n)
		if count <= k { // 如果子孙节点的数量小于 k，则目标值不在子树中
			num++      // 计算同层下一个节点的子孙节点数量
			k -= count // 消耗掉子节点数量的计数，相当与遍历完子孙节点
		} else { // 如果子孙节点的数量大于k，说明目标在当前节点的子树内
			num *= 10 // 进入当前节点的子树，（×10，即是第一个子树）
			k--       // 消耗一个计数
		}
	}
	return num
}

func Test440() {
	fmt.Println(findKthNumber(957747794, 424238336))
	fmt.Println(findKthNumber(10000, 10000)) // 9999
	fmt.Println(findKthNumber(1000, 1000))   // 999
	fmt.Println(findKthNumber(100, 10))      // 17
	fmt.Println(findKthNumber(10, 3))        // 2
	fmt.Println(findKthNumber(2, 2))         // 2
	fmt.Println(findKthNumber(13, 2))        // 10
	fmt.Println(findKthNumber(1, 1))         // 1
}
