package L_0868

import "fmt"

// https://leetcode-cn.com/problems/binary-gap/

// 868. 二进制间距

// 给定一个正整数 n，找到并返回 n 的二进制表示中两个 相邻 1 之间的 最长距离 。如果不存在两个相邻的 1，返回 0 。

// 如果只有 0 将两个 1 分隔开（可能不存在 0 ），则认为这两个 1 彼此 相邻 。
// 两个 1 之间的距离是它们的二进制表示中位置的绝对差。例如，"1001" 中的两个 1 的距离为 3 。

// 逐位右移，判断二进制最右侧两次出现1之间的间隔即可。

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了53.85% 的用户

func binaryGap(n int) int {
	maxGap := 0
	lastIndex := -1 // 记录上一个1的位置
	i := 0          // 位置计数
	for n > 0 {     // 循环直到 n<=0
		if n&1 == 1 { // 遇到1时
			if lastIndex != -1 { // 若已有1的位置记录
				// 计算距离
				if i-lastIndex > maxGap {
					maxGap = i - lastIndex
				}
			}
			// 重置上一个1的位置
			lastIndex = i
		}
		n = n >> 1 // 右移一位
		i++        // 位置计数+1
	}
	return maxGap
}

func Test868() {
	fmt.Println(binaryGap(15)) // 1
	fmt.Println(binaryGap(22)) // 2
	fmt.Println(binaryGap(8))  // 0
}
