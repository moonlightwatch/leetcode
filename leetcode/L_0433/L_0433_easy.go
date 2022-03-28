package L_0433

import "fmt"

// https://leetcode-cn.com/problems/binary-number-with-alternating-bits/

// 693. 交替位二进制数

// 给定一个正整数，检查它的二进制表示是否总是 0、1 交替出现：换句话说，就是二进制表示中相邻两位的数字永不相同。

func hasAlternatingBits(n int) bool {
	// 异或判断
	// t := n ^ (n >> 1)
	// return t&(t+1) == 0

	// 逐位判断
	for n > 0 {
		if ((n >> 1) & 1) == n&1 {
			return false
		}
		n = n >> 1
	}
	return true
}

func Test433() {
	fmt.Println(hasAlternatingBits(5))  // true
	fmt.Println(hasAlternatingBits(6))  // false
	fmt.Println(hasAlternatingBits(7))  // false
	fmt.Println(hasAlternatingBits(11)) // false

}
