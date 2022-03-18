package L_0504

import (
	"fmt"
)

// https://leetcode-cn.com/problems/base-7/

// 504. 七进制数

// 给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。

func convertToBase7(num int) string {
	// return strconv.FormatInt(int64(num), 7)
	if num == 0 {
		return "0"
	}
	result := []rune{}
	sign := ""
	if num < 0 {
		sign = "-"
		num = num * -1
	}
	for num > 0 {
		result = append(result, rune(num%7+48))
		num = num / 7
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return fmt.Sprintf("%s%s", sign, string(result))
}

func Test504() {
	fmt.Println(convertToBase7(100)) // 202
	fmt.Println(convertToBase7(-7))  // -10
}
