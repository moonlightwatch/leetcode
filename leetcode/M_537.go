package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/complex-number-multiplication/

// 537. 复数乘法

// 复数 可以用字符串表示，遵循 "实部+虚部i" 的形式，并满足下述条件：

//     实部 是一个整数，取值范围是 [-100, 100]
//     虚部 也是一个整数，取值范围是 [-100, 100]
//     i² == -1

// 给你两个字符串表示的复数 num1 和 num2 ，请你遵循复数表示形式，返回表示它们乘积的字符串。

func complexNumberMultiply(num1 string, num2 string) string {
	// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
	// 内存消耗：1.9 MB, 在所有 Go 提交中击败了78.95% 的用户

	// 负数乘法：(a+bi)(c+di)=(ac-bd)+(bc+ad)i
	num1Parts := strings.SplitN(num1, "+", 2)
	a, _ := strconv.Atoi(num1Parts[0])
	b, _ := strconv.Atoi(strings.Replace(num1Parts[1], "i", "", 1))

	num2Parts := strings.SplitN(num2, "+", 2)
	c, _ := strconv.Atoi(num2Parts[0])
	d, _ := strconv.Atoi(strings.Replace(num2Parts[1], "i", "", 1))

	return fmt.Sprintf("%d+%di", a*c-b*d, b*c+a*d)
}

func Test537() {
	fmt.Println(complexNumberMultiply("1+1i", "1+1i"))   // 0+2i
	fmt.Println(complexNumberMultiply("1+-1i", "1+-1i")) // 0+-2i
}
