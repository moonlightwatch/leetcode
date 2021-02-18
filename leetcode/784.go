package leetcode

import (
	"fmt"
	"strings"
)

// 给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。

func letterCasePermutation(S string) []string {
	result := []string{}
	if len(S) == 1 {
		if S[0] >= 65 && S[0] <= 90 {
			return []string{string(S[0]), string(S[0] + 32)}
		} else if S[0] >= 97 && S[0] <= 122 {
			return []string{string(S[0]), string(S[0] - 32)}
		} else {
			return []string{S}
		}
	}
	first := []byte(S)[0]
	if first >= 65 && first <= 90 {
		for _, item := range letterCasePermutation(S[1:]) {
			result = append(result, strings.Join([]string{string(first), item}, ""), strings.Join([]string{string(first + 32), item}, ""))
		}

	} else if first >= 97 && first <= 122 {
		for _, item := range letterCasePermutation(S[1:]) {
			result = append(result, strings.Join([]string{string(first), item}, ""), strings.Join([]string{string(first - 32), item}, ""))
		}
	} else {
		for _, item := range letterCasePermutation(S[1:]) {
			result = append(result, strings.Join([]string{string(first), item}, ""))
		}
	}
	return result
}

// 执行用时：8 ms, 在所有 Go 提交中击败了97.33% 的用户
// 内存消耗：6.5 MB, 在所有 Go 提交中击败了17.33% 的用户

func Test784() {
	fmt.Printf("%v\n", letterCasePermutation("a"))
	fmt.Printf("%v\n", letterCasePermutation("3z4"))
	fmt.Printf("%v\n", letterCasePermutation("a1b2"))
	fmt.Printf("%v\n", letterCasePermutation("123456"))
}
