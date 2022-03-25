package L_0172

import "fmt"

// https://leetcode-cn.com/problems/factorial-trailing-zeroes/

// 172. 阶乘后的零

// 给定一个整数 n ，返回 n! 结果中尾随零的数量。

// 提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
// 0 <= n <= 10^4

// 能乘出末尾 0 的，只有 5 的倍数
// 即 5的倍数 与 2的倍数 相乘，则尾随的 0 数量加一（1*5 = 10）
// 由于2的倍数明显比5的倍数多，所以 5的倍数 一定有一个 2的倍数与 之相乘
// 又通过统计得知：
// 25的倍数 与 4 的倍数相乘，则尾随的 0 数量加二（4*25 = 100）
// 125的倍数 与 8 的倍数相乘，则尾随的 0 数量加三（8*125 = 1000）
// 625的倍数 与 16 的倍数相乘，则尾随的 0 数量加四（16*625 = 10000）
// 3125的倍数 与 32 的倍数相乘，则尾随的 0 数量加五（32*3125 = 100000）
// （规律是：5的幂数 与其 幂次相同的2的幂数 相乘，会产生与幂次相等数量的0）
// 题目限制 数字最大到 10000，后续不再统计
// 因而，我们可以直接找到上述数字，累加其对应的0的数量即可

func trailingZeroes(n int) int {
	zeroes := 0
	// 高效解法
	// for n > 0 {
	// 	zeroes += n / 5
	// 	n = n / 5
	// }
	for i := 1; i <= n; i++ {
		if i%3125 == 0 {
			zeroes += 5
		} else if i%625 == 0 {
			zeroes += 4
		} else if i%125 == 0 {
			zeroes += 3
		} else if i%25 == 0 {
			zeroes += 2
		} else if i%5 == 0 {
			zeroes += 1
		}
	}
	return zeroes
}

func Test172() {

	fmt.Println(trailingZeroes(625)) // 156
	fmt.Println(trailingZeroes(200)) // 49
	fmt.Println(trailingZeroes(30))  // 7
	fmt.Println(trailingZeroes(0))   // 0
	fmt.Println(trailingZeroes(3))   // 0
	fmt.Println(trailingZeroes(5))   // 1
	fmt.Println(trailingZeroes(10))  // 2
	fmt.Println(trailingZeroes(14))  // 2
	fmt.Println(trailingZeroes(15))  // 3
	fmt.Println(trailingZeroes(16))  // 3
	fmt.Println(trailingZeroes(19))  // 3
	fmt.Println(trailingZeroes(20))  // 4

}
