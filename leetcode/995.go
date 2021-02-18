package leetcode

import "fmt"

// 在仅包含 0 和 1 的数组 A 中，一次 K 位翻转包括选择一个长度为 K 的（连续）子数组，同时将子数组中的每个 0 更改为 1，而每个 1 更改为 0。
// 返回所需的 K 位翻转的最小次数，以便数组没有值为 0 的元素。如果不可能，返回 -1。

func minKBitFlips(A []int, K int) int {
	count := 0
	for len(A) >= K {
		if A[0] == 0 {
			for i := 1; i < K; i++ {
				A[i] = A[i] ^ 1
			}
			count++
		}
		A = A[1:]
	}
	for _, x := range A {
		if x == 0 {
			return -1
		}
	}
	return count
}

// 执行用时：1068 ms, 在所有 Go 提交中击败了8.70% 的用户
// 内存消耗：7.1 MB, 在所有 Go 提交中击败了73.91% 的用户

func Test995() {
	fmt.Printf("%v\n", minKBitFlips([]int{0, 1, 0}, 1))
	fmt.Printf("%v\n", minKBitFlips([]int{1, 1, 0}, 2))
	fmt.Printf("%v\n", minKBitFlips([]int{0, 0, 0, 1, 0, 1, 1, 0}, 3))
}
