package L_0784

import (
	"fmt"
	"strings"
)

// 给定一个字符串S，通过将字符串S中的每个字母转变大小写，我们可以获得一个新的字符串。返回所有可能得到的字符串集合。

/*

思路： （递归）

若输入只有一位，则根据是否为字母，返回大小写集合，或者其本身。
否则，通过本函数获取除第一位外的子串的大小写全集。
然后，判断第一位是否字母。若是，则在字串大小写全集前分别加上第一位的大小写，全部写入结果集；若否，则仅将第一位分别添加到字串大小写全集各元素前，写入结果集。
返回结果集

*/

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
	firstStr := string(first)
	if first >= 65 && first <= 90 {
		for _, item := range letterCasePermutation(S[1:]) {
			result = append(result, strings.Join([]string{firstStr, item}, ""), strings.Join([]string{string(first + 32), item}, ""))
		}

	} else if first >= 97 && first <= 122 {
		for _, item := range letterCasePermutation(S[1:]) {
			result = append(result, strings.Join([]string{firstStr, item}, ""), strings.Join([]string{string(first - 32), item}, ""))
		}
	} else {
		for _, item := range letterCasePermutation(S[1:]) {
			result = append(result, strings.Join([]string{firstStr, item}, ""))
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
