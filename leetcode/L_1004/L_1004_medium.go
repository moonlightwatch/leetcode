package L_1004

import (
	"fmt"
)

// https://leetcode-cn.com/problems/max-consecutive-ones-iii/

// 1004. 最大连续1的个数 III

// 给定一个二进制数组 nums 和一个整数 k ，如果可以翻转最多k 个 0 ，则返回 数组中连续 1 的最大个数 。

/*

思路：

在数组 A 上设置滑动窗口，head 为起点，tail 为终点
逐步后移tail
当移动到 0 值上，则 K 减一；
	当 K 不足一时，记录窗口长度（与历史长度取最大值）
	将 head 后移到下一个 0 值处
记录窗口长度（与历史长度取最大值）
即得最长结果

上述窗口在滑动时，始终保持窗口内有 K （或小于 K） 个 0。最终记录窗口最大值即可。

*/

func longestOnes(A []int, K int) int {
	head := 0
	tail := 0
	max := 0
	l := len(A)
	for tail < l {
		if A[tail] == 0 {
			if K > 0 {
				K--
			} else {
				if max < tail-head {
					max = tail - head
				}
				for A[head] == 1 {
					head++
				}
				head++
			}
		}
		tail++
		// fmt.Printf("head:%d tail:%d K:%d max:%d\n", head, tail, K, max)
	}
	if max < tail-head {
		max = tail - head
	}
	return max
}

// 执行用时：72 ms, 在所有 Go 提交中击败了80.42% 的用户
// 内存消耗：6.9 MB, 在所有 Go 提交中击败了9.79% 的用户

func Test1004() {
	fmt.Printf("%v\n", longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Printf("%v\n", longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
	fmt.Printf("%v\n", longestOnes([]int{0, 0, 0, 1}, 4))
}
