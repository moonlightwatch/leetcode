package L_0008

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/string-to-integer-atoi/

// 8. 字符串转换整数 (atoi)

// 请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。

// 函数 myAtoi(string s) 的算法如下：

//     读入字符串并丢弃无用的前导空格
//     检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
//     读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
//     将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
//     如果整数数超过 32 位有符号整数范围 [−2³¹,  2³¹ − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −2³¹ 的整数应该被固定为 −2³¹ ，大于 2³¹ − 1 的整数应该被固定为 2³¹ − 1 。
//     返回整数作为最终结果。

// 注意：

//     本题中的空白字符只包括空格字符 ' ' 。
//     除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。

func myAtoi(s string) int {
	// 就按题目描述直接操作即可
	// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
	// 内存消耗：2.2 MB, 在所有 Go 提交中击败了81.87% 的用户

	s = strings.TrimLeft(s, " ") // 移除前导空格
	if s == "" {                 // 空字串直接返回
		return 0
	}

	// 符号判断
	sign := 0
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		sign = 1
		s = s[1:]
	} else { // 默认为正
		sign = 1
	}
	s = strings.TrimLeft(s, "0") // 移除前导 0
	if s == "" {                 // 空字串直接返回
		return 0
	}

	cache := []rune{}
	for _, c := range s {
		if c >= '0' && c <= '9' {
			cache = append(cache, c)
		} else {
			break
		}
	}

	if len(cache) > 10 {
		if sign == -1 {
			return -2147483648
		}
		if sign == 1 {
			return 2147483647
		}
	}
	n := 1
	var result int64
	for i := len(cache); i > 0; i-- {
		result += int64(int(cache[i-1]-48) * n)
		n *= 10
	}
	if sign == -1 && result > 2147483648 {
		return -2147483648
	}
	if sign == 1 && result > 2147483647 {
		return 2147483647
	}
	return int(result) * sign
}

func Test8() {
	fmt.Println(myAtoi("+-12"))                  // 0
	fmt.Println(myAtoi("42"))                    // 42
	fmt.Println(myAtoi("   -42"))                // -42
	fmt.Println(myAtoi("4193 with words"))       // 4193
	fmt.Println(myAtoi("words and 987"))         // 0
	fmt.Println(myAtoi("-91283472332"))          // -2147483648
	fmt.Println(myAtoi("9223372036854775808"))   // 2147483647
	fmt.Println(myAtoi("  0000000000012345678")) // 12345678
}
