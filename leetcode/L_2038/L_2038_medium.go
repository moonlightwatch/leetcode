package L_2038

import "fmt"

// https://leetcode-cn.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/

// 2038. 如果相邻两个颜色均相同则删除当前颜色

// 总共有 n 个颜色片段排成一列，每个颜色片段要么是 'A' 要么是 'B' 。给你一个长度为 n 的字符串 colors ，其中 colors[i] 表示第 i 个颜色片段的颜色。

// Alice 和 Bob 在玩一个游戏，他们 轮流 从这个字符串中删除颜色。Alice 先手 。

//     如果一个颜色片段为 'A' 且 相邻两个颜色 都是颜色 'A' ，那么 Alice 可以删除该颜色片段。Alice 不可以 删除任何颜色 'B' 片段。
//     如果一个颜色片段为 'B' 且 相邻两个颜色 都是颜色 'B' ，那么 Bob 可以删除该颜色片段。Bob 不可以 删除任何颜色 'A' 片段。
//     Alice 和 Bob 不能 从字符串两端删除颜色片段。
//     如果其中一人无法继续操作，则该玩家 输 掉游戏且另一玩家 获胜 。

// 假设 Alice 和 Bob 都采用最优策略，如果 Alice 获胜，请返回 true，否则 Bob 获胜，返回 false。

func winnerOfGame(colors string) bool {
	// 执行用时：8 ms, 在所有 Go 提交中击败了100.00% 的用户
	// 内存消耗：6.2 MB, 在所有 Go 提交中击败了72.73% 的用户

	// 遍历colors，对连续的 A 或者 B 进行计数
	// 连续出现 3 次，则选手增加一回合
	// 连续出现 3 次以上 每多一次，增加一回合
	// 最后统计回合数，由于 Alice 先手，故 当Alice的回合数大于Bob的回合数时，Alice胜，否则Bob胜

	countA, roundA := 0, 0 // 连续出现计数
	countB, roundB := 0, 0 // 回合计数
	curColor := colors[0]  // 当前计数的颜色
	if curColor == 'A' {   // 第一个颜色计数
		countA++
	} else {
		countB++
	}
	// 遍历所有颜色
	for i := 1; i < len(colors); i++ {
		if curColor == colors[i] { // 若颜色与当前计数的颜色相同
			if curColor == 'A' { // 且颜色为 A
				countA++        // 则A的连续计数加一
				if countA > 2 { // 若连续3次或者以上
					roundA++ // 则为Alice增加回合数
				}
			} else { // 若颜色为 B
				countB++        //则B的连续计数加一
				if countB > 2 { // 若连续3次或者以上
					roundB++ // 则为Bob增加回合数
				}
			}
		} else { // 若颜色与当前计数颜色不同
			countA, countB = 0, 0 // 则清空计数
			curColor = colors[i]  // 重设当前计数颜色
			if curColor == 'A' {  // 根据当前计数颜色，增加连续计数
				countA++
			} else {
				countB++
			}
		}
	}
	return roundA > roundB // 若 Alice 的回合数 大于 Bob 的回合数，则 Alice 胜，否则 Bob 胜
}

func Test2038() {
	fmt.Println(winnerOfGame("AAABABB"))     // true
	fmt.Println(winnerOfGame("AA"))          // false
	fmt.Println(winnerOfGame("ABBBBBBBAAA")) // false
}
