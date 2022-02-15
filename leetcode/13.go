package leetcode

import "fmt"

// 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

// 字符          数值
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

// 例如， 罗马数字 2 写做 II ，即为两个并列的 1 。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

// 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

//     I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//     X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
//     C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

// 给定一个罗马数字，将其转换成整数。

func romanToInt(s string) int {
	// 执行用时：4 ms, 在所有 Go 提交中击败了92.01% 的用户
	// 内存消耗：2.7 MB, 在所有 Go 提交中击败了96.85% 的用户
	result := 0
	last := '~'
	for i := len(s) - 1; i >= 0; i-- {
		if (last == 'D' || last == 'M') && s[i] == 'C' {
			result -= 100
		} else if (last == 'C' || last == 'L') && s[i] == 'X' {
			result -= 10
		} else if (last == 'V' || last == 'X') && s[i] == 'I' {
			result -= 1
		} else {
			switch s[i] {
			case 'I':
				result += 1
			case 'V':
				result += 5
			case 'X':
				result += 10
			case 'L':
				result += 50
			case 'C':
				result += 100
			case 'D':
				result += 500
			case 'M':
				result += 1000
			}
		}
		last = rune(s[i])
	}
	return result
}

func Test13() {
	fmt.Println(romanToInt("III"))     // 3
	fmt.Println(romanToInt("IV"))      // 4
	fmt.Println(romanToInt("IX"))      // 9
	fmt.Println(romanToInt("LVIII"))   // 58
	fmt.Println(romanToInt("MCMXCIV")) // 1994

}
