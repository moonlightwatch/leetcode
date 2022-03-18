package L_0720

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/longest-word-in-dictionary/

// 720. 词典中最长的单词

// 给出一个字符串数组 words 组成的一本英语词典。返回 words 中最长的一个单词，该单词是由 words 词典中其他单词逐步添加一个字母组成。

// 若其中有多个可行的答案，则返回答案中字典序最小的单词。若无答案，则返回空字符串。

// 递归解法
// 执行用时：56 ms, 在所有 Go 提交中击败了5.22% 的用户
// 内存消耗：6 MB, 在所有 Go 提交中击败了100.00% 的用户

func searchLongestWordByPrefix(prefix string, words []string) string {
	longestWord := prefix
	for _, word := range words {
		if len(word) == len(prefix)+1 && strings.HasPrefix(word, prefix) {
			w := searchLongestWordByPrefix(word, words)
			if len(w) > len(longestWord) {
				longestWord = w
			} else if len(w) == len(longestWord) {
				for i := 0; i < len(w); i++ {
					if w[i] != longestWord[i] {
						if w[i] < longestWord[i] {
							longestWord = w
						} else {
							break
						}
					}
				}
			}
		}
	}
	return longestWord
}

func longestWord(words []string) string {
	longestWord := ""
	for _, word := range words {
		if len(word) == 1 {
			w := searchLongestWordByPrefix(word, words)
			if len(w) > len(longestWord) {
				longestWord = w
			} else if len(w) == len(longestWord) {
				for i := 0; i < len(w); i++ {
					if w[i] != longestWord[i] {
						if w[i] < longestWord[i] {
							longestWord = w
						} else {
							break
						}
					}
				}
			}
		}
	}
	return longestWord
}

// 循环解法

// func longestWord(words []string) string {
// 	sort.Slice(words, func(i, j int) bool {
// 		if len(words[i]) == len(words[j]) {
// 			for k := 0; k < len(words[i]); k++ {
// 				if words[i][k] != words[j][k] {
// 					return words[i][k] < words[j][k]
// 				}
// 			}
// 		}
// 		return len(words[i]) < len(words[j])
// 	})
// 	longest := ""
// 	cache := map[string]bool{}
// 	for i := 0; i < len(words); i++ {
// 		_, ok := cache[words[i][:len(words[i])-1]]
// 		if len(words[i]) == 1 || ok {
// 			if len(words[i]) > len(longest) {
// 				longest = words[i]
// 			}
// 			cache[words[i]] = true
// 		}
// 	}
// 	return longest
// }

func Test720() {
	fmt.Println(longestWord([]string{"rac", "rs", "ra", "on", "r", "otif", "o", "onpdu", "rsf", "rs", "ot", "oti", "racy", "onpd"}))
	fmt.Println(longestWord([]string{"b", "br", "bre", "brea", "break", "breakf", "breakfa", "breakfas", "breakfast", "l", "lu", "lun", "lunc", "lunch", "d", "di", "din", "dinn", "dinne", "dinner"})) // breakfast
	fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))                                                                                                                               // world
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))                                                                                                            // apple
}
