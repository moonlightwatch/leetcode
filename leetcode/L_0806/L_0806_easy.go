package L_0806

import "fmt"

// https://leetcode-cn.com/problems/number-of-lines-to-write-string/

// 806. 写字符串需要的行数

// 我们要把给定的字符串 S 从左到右写到每一行上，每一行的最大宽度为100个单位，如果我们在写某个字母的时候会使这行超过了100 个单位，那么我们应该把这个字母写到下一行。
// 我们给定了一个数组 widths ，这个数组 widths[0] 代表 'a' 需要的单位， widths[1] 代表 'b' 需要的单位，...， widths[25] 代表 'z' 需要的单位。

// 现在回答两个问题：至少多少行能放下S，以及最后一行使用的宽度是多少个单位？将你的答案作为长度为2的整数列表返回。

// 字符串 S 的长度在 [1, 1000] 的范围。
// S 只包含小写字母。
// widths 是长度为 26的数组。
// widths[i] 值的范围在 [2, 10]。

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
// 内存消耗：1.8 MB, 在所有 Go 提交中击败了100.00% 的用户

func numberOfLines(widths []int, s string) []int {
	result := []int{1, 0}
	for _, c := range s { // 循环所有字母
		result[1] += widths[c-'a'] // 给当前的最后一行，加上当前字母的宽度
		if result[1] > 100 {       // 如果当前的最后一行大于100了
			result[0]++               // 则行数 +1
			result[1] = widths[c-'a'] // 将当前字母加在下一行
		}
	}
	return result
}

func Test806() {
	// 3 60
	fmt.Println(numberOfLines([]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "abcdefghijklmnopqrstuvwxyz"))
	// 2 4
	fmt.Println(numberOfLines([]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa"))

}
