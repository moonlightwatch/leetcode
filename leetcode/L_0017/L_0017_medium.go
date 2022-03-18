package L_0017

import "fmt"

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

// 17. 电话号码的字母组合

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

func letterCombinations(digits string) []string {
	// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
	// 内存消耗：1.9 MB, 在所有 Go 提交中击败了93.69% 的用户

	// 解法：本题可以看作遍历矩阵，将便利路径记录输出即可
	// （亦可看作若干矩阵求积）

	l := len(digits) // 获取长度
	if l == 0 {      // 非法长度直接返回
		return []string{}
	}
	result := []string{} // 存储结果

	// 构造数组映射
	codeMap := map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	// 用于记录遍历过程中，每一位对应的索引
	indexList := make([]int, l)

	for indexList[0] < len(codeMap[rune(digits[0])]) {
		tempString := "" // 临时存储字符串

		// 遍历索引列表，按其中索引分别提取各位的字母，拼接到临时字符转中
		for i, index := range indexList {
			tempString += string(codeMap[rune(digits[i])][index])
		}

		// 将拼好的字符串加入结果
		result = append(result, tempString)

		// 更新最后一位的索引
		indexList[l-1] += 1

		// 从后往前，逐位判断索引是否溢出。溢出的话，向前一位加一，并将此位置为0
		// （这里的操作类似于加法进位，为的是遍历矩阵）
		for i := l - 1; i > 0; i-- {
			if indexList[i] >= len(codeMap[rune(digits[i])]) {
				indexList[i] = 0
				indexList[i-1] += 1
			}
		}
	}
	return result
}

func Test17() {
	fmt.Println(letterCombinations("23")) // ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	fmt.Println(letterCombinations("2"))  // ["a","b","c"]
	fmt.Println(letterCombinations(""))   // []
}
