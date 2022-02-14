package leetcode

import (
	"fmt"
	"strings"
)

// 罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

// 字符          数值
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

// 例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

// 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

//     I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//     X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
//     C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

// 给你一个整数，将其转为罗马数字。

func intToRoman(num int) string {
	roman := ""
	if num >= 1000 {
		roman = strings.Repeat("M", num/1000)
		num = num % 1000
	}
	if num >= 900 {
		roman += strings.Repeat("CM", num/900)
		num = num % 900
	}
	if num >= 500 {
		roman += strings.Repeat("D", num/500)
		num = num % 500
	}
	if num >= 400 {
		roman += strings.Repeat("CD", num/400)
		num = num % 400
	}
	if num >= 100 {
		roman += strings.Repeat("C", num/100)
		num = num % 100
	}
	if num >= 90 {
		roman += strings.Repeat("XC", num/90)
		num = num % 90
	}
	if num >= 50 {
		roman += strings.Repeat("L", num/50)
		num = num % 50
	}
	if num >= 40 {
		roman += strings.Repeat("XL", num/40)
		num = num % 40
	}
	if num >= 10 {
		roman += strings.Repeat("X", num/10)
		num = num % 10
	}
	if num >= 9 {
		roman += strings.Repeat("IX", num/9)
		num = num % 9
	}
	if num >= 5 {
		roman += strings.Repeat("V", num/5)
		num = num % 5
	}
	if num >= 4 {
		roman += strings.Repeat("IV", num/4)
		num = num % 4
	}
	roman += strings.Repeat("I", num/1)

	return roman
}

func Test12() {
	fmt.Println(intToRoman(3))    // III
	fmt.Println(intToRoman(4))    // IV
	fmt.Println(intToRoman(9))    // IX
	fmt.Println(intToRoman(58))   // LVIII
	fmt.Println(intToRoman(1994)) //MCMXCIV
	fmt.Println(intToRoman(400))  //CD
}
