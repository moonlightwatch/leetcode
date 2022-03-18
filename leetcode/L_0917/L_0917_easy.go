package L_0917

import "fmt"

// https://leetcode-cn.com/problems/reverse-only-letters/

// 917. 仅仅反转字母

// 给你一个字符串 s ，根据下述规则反转字符串：

//     所有非英文字母保留在原有位置。
//     所有英文字母（小写或大写）位置反转。

// 返回反转后的 s 。

func reverseOnlyLetters(s string) string {
	// 执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
	// 内存消耗：1.8 MB, 在所有 Go 提交中击败了75.36% 的用户

	// 左右两端指针相向搜索。遇到字母，则相互交换内容，直至指针相遇。

	cache := []rune(s)
	length := len(s)
	l, r := 0, length-1
	for l < r && l < length && r > 0 {
		for (l < length) && (cache[l] < 'a' || cache[l] > 'z') && (cache[l] < 'A' || cache[l] > 'Z') {
			l++
		}
		for (r > 0) && (cache[r] < 'a' || cache[r] > 'z') && (cache[r] < 'A' || cache[r] > 'Z') {
			r--
		}
		if l < r && l < length && r > 0 {
			cache[l], cache[r] = cache[r], cache[l]
			l++
			r--
		}
	}
	return string(cache)
}

func Test917() {

	fmt.Println(reverseOnlyLetters("?6C40E"))               // ?6E40C
	fmt.Println(reverseOnlyLetters("7_28]"))                // 7_28]
	fmt.Println(reverseOnlyLetters("ab-cd"))                // dc-ba
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))        // j-Ih-gfE-dCba
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!")) // Qedo1ct-eeLg=ntse-T!
}
