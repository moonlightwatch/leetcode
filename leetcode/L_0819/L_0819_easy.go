package L_0819

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/most-common-word/

// 819. 最常见的单词

// 给定一个段落 (paragraph) 和一个禁用单词列表 (banned)。返回出现次数最多，同时不在禁用列表中的单词。

// 题目保证至少有一个词不在禁用列表中，而且答案唯一。

// 禁用列表中的单词用小写字母表示，不含标点符号。段落中的单词不区分大小写。答案都是小写字母。

// 1 <= 段落长度 <= 1000
// 0 <= 禁用单词个数 <= 100
// 1 <= 禁用单词长度 <= 10
// 答案是唯一的, 且都是小写字母 (即使在 paragraph 里是大写的，即使是一些特定的名词，答案都是小写的。)
// paragraph 只包含字母、空格和下列标点符号!?',;.
// 不存在没有连字符或者带有连字符的单词。
// 单词里只包含字母，不会出现省略号或者其他标点符号。

func mostCommonWord(paragraph string, banned []string) string {
	// 移除标点符号
	paragraph = strings.ReplaceAll(paragraph, "!", "  ")
	paragraph = strings.ReplaceAll(paragraph, "?", "  ")
	paragraph = strings.ReplaceAll(paragraph, "'", "  ")
	paragraph = strings.ReplaceAll(paragraph, ",", "  ")
	paragraph = strings.ReplaceAll(paragraph, ";", "  ")
	paragraph = strings.ReplaceAll(paragraph, ".", "  ")
	paragraph = strings.ReplaceAll(paragraph, " ", "  ")

	// 转小写，两端加空格，用于下一步移除禁用词
	paragraph = fmt.Sprintf(" %s ", strings.ToLower(paragraph))

	// 移除禁用词
	for _, w := range banned {
		r := fmt.Sprintf(" %s ", strings.ToLower(w))
		paragraph = strings.ReplaceAll(paragraph, r, "  ")
	}

	// 单词列表
	words := strings.Split(paragraph, " ")
	maxCount := 0
	maxWord := words[0]
	cache := map[string]int{}
	// 循环单词
	for i := 0; i < len(words); i++ {
		if words[i] == "" { // 跳过空单词
			continue
		}
		if _, ok := cache[words[i]]; !ok {
			cache[words[i]] = 1
		} else {
			cache[words[i]]++
		}
		if cache[words[i]] > maxCount {
			maxCount = cache[words[i]]
			maxWord = words[i]
		}
	}
	return maxWord
}

func Test819() {
	fmt.Println(mostCommonWord("abc abc? abcd the jeff!", []string{"abc", "abcd", "jeff"}))
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"})) // ball
}
