package L_0006

import (
	"fmt"
)

// https://leetcode-cn.com/problems/zigzag-conversion/

// 6. Z 字形变换

// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

// P   A   H   N
// A P L S I I G
// Y   I   R

// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

func convert(s string, numRows int) string {
	// 按题目顺序摆就行，题目中说是Z字形。但实际是趴着的 Z（可以看作 N，但笔顺不同）
	// 第一列从上摆到下，然后从倒数第二列逐个v摆到第二列，然后从第一列重复上述过程，即得答案
	// 最后整理成所需的输出形式即可

	// 用一个指定行数的矩阵，缓存摆出来的字符
	cache := make([][]byte, numRows)

	// 循环所有字符，但是索引的推进在循环体内进行
	for i := 0; i < len(s); {

		// 从上倒下逐行添加字符
		for l := 0; l < numRows && i < len(s); l++ {
			cache[l] = append(cache[l], s[i])
			i += 1
		}

		// 从下到上逐行添加字符，注意只添加中间的行，第一行和倒数第一行不处理（由上一个循环处理）
		for l := numRows - 2; l > 0 && i < len(s); l-- {
			cache[l] = append(cache[l], s[i])
			i += 1
		}
	}

	// 逐行拼接，按要求返回数据
	tempBytes := []byte{}
	for _, item := range cache {
		tempBytes = append(tempBytes, item...)
	}
	return string(tempBytes)
}

func Test6() {
	fmt.Println(convert("PAYPALISHIRING", 3)) // PAHNAPLSIIGYIR
	fmt.Println(convert("PAYPALISHIRING", 4)) // PINALSIGYAHRPI
	fmt.Println(convert("A", 1))              // A

}
