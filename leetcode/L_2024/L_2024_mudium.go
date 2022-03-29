package L_2024

import "fmt"

// https://leetcode-cn.com/problems/maximize-the-confusion-of-an-exam/

// 2024. 考试的最大困扰度

// 一位老师正在出一场由 n 道判断题构成的考试，每道题的答案为 true （用 'T' 表示）或者 false （用 'F' 表示）。
// 老师想增加学生对自己做出答案的不确定性，方法是 最大化 有 连续相同 结果的题数。（也就是连续出现 true 或者连续出现 false）。

// 给你一个字符串 answerKey ，其中 answerKey[i] 是第 i 个问题的正确结果。除此以外，还给你一个整数 k ，表示你能进行以下操作的最多次数：

//     每次操作中，将问题的正确答案改为 'T' 或者 'F' （也就是将 answerKey[i] 改为 'T' 或者 'F' ）。

// 请你返回在不超过 k 次操作的情况下，最大 连续 'T' 或者 'F' 的数目。

// n == answerKey.length
// 1 <= n <= 5 * 10^4
// answerKey[i] 要么是 'T' ，要么是 'F'
// 1 <= k <= n

func maxConsecutiveAnswers(answerKey string, k int) int {
	n := len(answerKey)
	if n == 1 {
		return 1
	}
	keys := []byte{}
	keyCount := []int{}
	temp := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			keys = append(keys, answerKey[i])
			temp++
			continue
		}
		if answerKey[i] == keys[len(keys)-1] {
			temp++
		} else {
			keyCount = append(keyCount, temp)
			temp = 1
			keys = append(keys, answerKey[i])
		}
		if i == n-1 {
			keyCount = append(keyCount, temp)
		}
	}

	// 全都相等的情况
	if len(keyCount) == 1 {
		return keyCount[0]
	}

	longest := 0 // 最长结果计数

	// 无法连接的情况
	for i := 0; i < len(keyCount); i++ {
		if longest < keyCount[i]+k {
			longest = keyCount[i] + k
		}
	}

	for i := 0; i < len(keyCount)-1; i++ { // 循环统计结果
		t := keyCount[i]        // 计数
		if keyCount[i+1] <= k { // 如果每到最后一个，且下一个长度小于k
			// t += keyCount[i+1] // 则还要加上下一个
			tempK := k
			for j := 1; j < len(keyCount)-i; j += 2 {
				if keyCount[i+j] <= tempK {
					t += keyCount[i+j]
					tempK -= keyCount[i+j]
					if i+j+1 < len(keyCount) {
						t += keyCount[i+j+1]
					}
				} else {
					break
				}
			}
			t += tempK
		} else { // 若下一个结果比k长，则加上k即可
			t += k
		}
		if longest < t {
			longest = t
		}
		if longest >= n {
			return n
		}
	}
	return longest
}

func Test2024() {
	fmt.Println(maxConsecutiveAnswers("TFF", 1))        // 3
	fmt.Println(maxConsecutiveAnswers("FFFTTFTTFT", 3)) // 8
	fmt.Println(maxConsecutiveAnswers("TTFTTTTTFT", 1)) // 8
	fmt.Println(maxConsecutiveAnswers("TTTTTFTFFT", 2)) // 8
	fmt.Println(maxConsecutiveAnswers("TTFF", 2))       // 4
	fmt.Println(maxConsecutiveAnswers("TFFT", 1))       // 3
	fmt.Println(maxConsecutiveAnswers("TTFTTFTT", 1))   // 5
}
