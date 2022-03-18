package L_0007

import (
	"fmt"
)

// https://leetcode-cn.com/problems/reverse-integer/

// 7. 整数反转
// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
// 如果反转后整数超过 32 位的有符号整数的范围 [−2³¹,  2³¹ − 1] ，就返回 0。
// 假设环境不允许存储 64 位整数（有符号或无符号）。

func reverse(x int) int {
	cache := []int8{}
	sign := 1
	if x < 0 {
		sign = -1
		x = x * -1
	}
	n := 10
	for x > 0 {
		cache = append(cache, int8(x%n))
		x = x / n
	}

	// fmt.Println(cache)
	n = 1
	result := 0
	for i := len(cache); i > 0; i-- {
		result += int(cache[i-1]) * n
		n *= 10
	}
	if 2147483647 < result {
		return 0
	}
	return result * sign

}

func Test7() {
	fmt.Println(reverse(123))  // 321
	fmt.Println(reverse(-123)) // -321
	fmt.Println(reverse(120))  // 21
	fmt.Println(reverse(0))    // 0
}
