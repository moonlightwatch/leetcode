package leetcode

import "fmt"

// https://leetcode-cn.com/problems/add-digits/

// 258. 各位相加

// 给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。返回这个结果。

func addDigits(num int) int {
	// 观察规律，获得规律：最终结果是 num 对 9 取余数，特别的是 9 的倍数返回 9
	if num == 0 {
		return 0
	}
	n := num % 9
	if n == 0 {
		return 9
	}
	return n
}

// 更优秀的解法，666！
// func addDigits(num int) int {
// 	return (num-1)%9+1
// }

func Test258() {
	fmt.Println(addDigits(1))
	fmt.Println(addDigits(2))
	fmt.Println(addDigits(3))
	fmt.Println(addDigits(4))
	fmt.Println(addDigits(10))
	fmt.Println(addDigits(11))
	fmt.Println(addDigits(19))
	fmt.Println(addDigits(38)) // 2
	fmt.Println(addDigits(0))  // 0
}
