package leetcode

import (
	"fmt"
)

// n 张多米诺骨牌排成一行，将每张多米诺骨牌垂直竖立。在开始时，同时把一些多米诺骨牌向左或向右推。
// 每过一秒，倒向左边的多米诺骨牌会推动其左侧相邻的多米诺骨牌。同样地，倒向右边的多米诺骨牌也会推动竖立在其右侧的相邻多米诺骨牌。
// 如果一张垂直竖立的多米诺骨牌的两侧同时有多米诺骨牌倒下时，由于受力平衡， 该骨牌仍然保持不变。
// 就这个问题而言，我们会认为一张正在倒下的多米诺骨牌不会对其它正在倒下或已经倒下的多米诺骨牌施加额外的力。
// 给你一个字符串 dominoes 表示这一行多米诺骨牌的初始状态，其中：

//     dominoes[i] = 'L'，表示第 i 张多米诺骨牌被推向左侧，
//     dominoes[i] = 'R'，表示第 i 张多米诺骨牌被推向右侧，
//     dominoes[i] = '.'，表示没有推动第 i 张多米诺骨牌。

// 返回表示最终状态的字符串。

func pushDominoes(dominoes string) string {
	// 循环所有字符
	// 遇到 'L' 或者 'R' 则处理当前字符到上一次记录字符之间的牌的状态，并记录位置与字符；遇到 '.' 则跳过

	runes := []rune(dominoes) // rune 列表，主要对此对象进行操作
	cacheIndex := 0           // 记录位置
	cacheRune := 'L'          // 记录方向，默认为 'L' 是因为不对后续判断产生影响
	for i, c := range runes {
		switch c {
		case 'L': // 若当前字符是 L
			if cacheRune == 'L' { // 若当前字符是 L, 上一次记录的字符是 L， 则从上一次记录位置到当前位置的所有 '.' 都更新为 'L'
				for j := cacheIndex; j < i; j++ {
					runes[j] = 'L'
				}
			}
			if cacheRune == 'R' { // 若当前字符是 L, 上一次记录的字符是 R, 则左侧一半改为 'R', 右侧一半改为 'L'
				// 记录左侧索引和右侧索引
				// 开始循环
				// 将左侧索引位置设为 'R'；右侧索引位置设为 'L'
				// 左侧索引向右推进一位，右侧索引向左推进一位
				// 直到左右索引相遇，则结束循环
				l := cacheIndex + 1
				r := i - 1
				for l < r { // 此处可以保证左右两侧向中间倒时，中间的不动
					runes[l] = 'R'
					runes[r] = 'L'
					l += 1
					r -= 1
				}
			}
			cacheRune = c
			cacheIndex = i
		case 'R': // 若当前字符是 R
			if cacheRune == 'R' { // 若当前字符是 R, 上一次记录的字符是 R， 则从上一次记录位置到当前位置的所有 '.' 都更新为 'R'
				for j := cacheIndex + 1; j < i; j++ {
					runes[j] = 'R'
				}
			}
			// 若上次一记录为 'L'，则不操作，仅记录当前字符
			cacheRune = c
			cacheIndex = i
		}
	}
	// 需要处理尾部的 '.' 字符，当上一个记录字符为 'R' 的情况下，尾部的 '.' 都改为 'R'。其他情况则不处理。
	if cacheIndex < len(runes) && cacheRune == 'R' {
		for i := cacheIndex; i < len(runes); i++ {
			runes[i] = 'R'
		}
	}

	return string(runes)
}

func Test838() {
	fmt.Println(pushDominoes("RR.L"))           // "RR.L"
	fmt.Println(pushDominoes(".L.R...LR..L..")) // "LL.RR.LLRRLL.."
	fmt.Println(pushDominoes(".L.R."))          // "LL.RR"
}
