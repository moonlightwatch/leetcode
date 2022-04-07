package L_0796

import (
	"fmt"
	"strings"
)

// https://leetcode-cn.com/problems/rotate-string/

// 796. 旋转字符串

// 给定两个字符串, s 和 goal。如果在若干次旋转操作之后，s 能变成 goal ，那么返回 true 。

// s 的 旋转操作 就是将 s 最左边的字符移动到最右边。

//     例如, 若 s = 'abcde'，在旋转一次之后结果就是'bcdea' 。

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	var b strings.Builder
	b.WriteString(goal)
	b.WriteString(goal)
	return strings.Contains(b.String(), s)
}

func Test796() {
	fmt.Println(rotateString("abcde", "cdeab")) // true
	fmt.Println(rotateString("abcde", "abced")) // false
}
