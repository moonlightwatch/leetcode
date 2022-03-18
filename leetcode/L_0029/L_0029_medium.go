package L_0029

import (
	"fmt"
	"math"
)

// https://leetcode-cn.com/problems/divide-two-integers/

// 29. 两数相除

// 给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

// 返回被除数 dividend 除以除数 divisor 得到的商。

// 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

func divide(dividend int, divisor int) int {
	a := dividend
	if a < 0 {
		a = 0 - a
	}
	b := divisor
	if b < 0 {
		b = 0 - b
	}
	if a == 0 || a < b {
		return 0
	}
	i := 0
	for a >= b {
		b = b << 1 // 不断左移，直到大于被除数。左移一次相当于与2相乘
		i++
	}
	result := (1 << (i - 1)) + divide(a-(b>>1), int(math.Abs(float64(divisor))))
	if dividend^divisor < 0 { // 判断符号是否相等
		result = 0 - result
	}
	if result < -2147483648 || result > 2147483647 {
		return 2147483647
	}
	return result
}

func Test29() {
	fmt.Println(divide(-2147483648, -1)) // 2147483647
	fmt.Println(divide(-2, -1))          // 2
	fmt.Println(divide(-10, -3))         // 3
	fmt.Println(divide(10, 3))           // 3
	fmt.Println(divide(7, -3))           // -2
}
