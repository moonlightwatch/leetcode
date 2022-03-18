package L_0030

import (
	"fmt"
)

// https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/

// 30. 串联所有单词的子串

// 给定一个字符串 s 和一些 长度相同 的单词 words 。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑 words 中单词串联的顺序。

// 1 <= s.length <= 104
// s 由小写英文字母组成
// 1 <= words.length <= 5000
// 1 <= words[i].length <= 30
// words[i] 由小写英文字母组成

func findSubstring(s string, words []string) []int {
	// 可以设置一个长度为所有单词长度综合的滑动窗口。逐位从前到后进行处理。
	// 由于所有单词都一样长
	// 所以滑动窗口内的字符串，可以直接按长度分割为若干份
	// 比较 滑动窗口内的字符串分割出的数组 和 words 数组是否一一对应
	// 若一一对应，则记录当前滑动窗口的头部索引，即为答案

	wl := len(words[0])  // 每个单词的长度
	l := wl * len(words) // 统计所有单词长度
	if len(s) < l {      // 若给定字符串长度小于所有单词长度之和，则直接返回空
		return []int{}
	}
	result := []int{} // 存储结果
	// 逐位向后检索
	for i := 0; i < len(s)-l+1; i++ {
		tempS := string(s[i : i+l])  // 提取滑动窗口内的字符串
		parts := map[string]int{}    // 这个map存储字符串按长度分割后，个部分出现的次数，用于和words一一对应
		for n := 0; n < l; n += wl { // 这个循环用于构建上述map
			x := tempS[n : n+wl]
			if _, ok := parts[x]; !ok {
				parts[x] = 1
			} else {
				parts[x]++
			}

		}
		// 遍历words
		for _, word := range words {
			if _, ok := parts[word]; ok { // 判断word在map中是否存在
				parts[word]--         // 存在的话次数 -1
				if parts[word] == 0 { // 若降低至0，则删除此项
					delete(parts, word)
				}
			} else { // 若 word 在 map 中不存在，直接返回
				break
			}
		}
		// 如果map的内容被删干净了，那么说明words和map一一对应，记录当前索引
		if len(parts) == 0 {
			result = append(result, i)
		}
	}
	return result
}

func Test30() {
	fmt.Println(findSubstring("ababaab", []string{"ab", "ba", "ba"}))                                // 1
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))                         // 0 9
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})) //
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))            // 6 9 12

}
