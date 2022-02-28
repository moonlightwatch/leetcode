package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/maximum-number-of-achievable-transfer-requests/

// 1601. 最多可达成的换楼请求数目

// 我们有 n 栋楼，编号从 0 到 n - 1 。每栋楼有若干员工。由于现在是换楼的季节，部分员工想要换一栋楼居住。
// 给你一个数组 requests ，其中 requests[i] = [fromi, toi] ，表示一个员工请求从编号为 fromi 的楼搬到编号为 toi 的楼。
// 一开始 所有楼都是满的，所以从请求列表中选出的若干个请求是可行的需要满足 每栋楼员工净变化为 0 。意思是每栋楼 离开 的员工数目 等于 该楼 搬入 的员工数数目。比方说 n = 3 且两个员工要离开楼 0 ，一个员工要离开楼 1 ，一个员工要离开楼 2 ，如果该请求列表可行，应该要有两个员工搬入楼 0 ，一个员工搬入楼 1 ，一个员工搬入楼 2 。
// 请你从原请求列表中选出若干个请求，使得它们是一个可行的请求列表，并返回所有可行列表中最大请求数目。

// 1 <= n <= 20
// 1 <= requests.length <= 16
// requests[i].length == 2
// 0 <= fromi, toi < n

func maximumRequests(n int, requests [][]int) int {
	// 穷举
	cacheIndex := []string{""} // 用字符串保存索引
	// 枚举出所有组合
	for i := 0; i < len(requests); i++ {
		l := len(cacheIndex)
		for j := 0; j < l; j++ {
			cacheIndex = append(cacheIndex, fmt.Sprintf("%s %d", cacheIndex[j], i))
		}
	}
	max := 0 // 记录最大值

	// 对所有组合进行判断
	for _, items := range cacheIndex {
		// 空组合直接跳过
		if items == "" {
			continue
		}
		// 记录器
		counter := make([]int, n)
		// 分离所有索引
		p := strings.Split(strings.TrimSpace(items), " ")
		// 若长度小于已记录的最大值，则跳过。其必不可能是答案
		if len(p) < max {
			continue
		}
		// 按照索引，逐个安排调换
		for _, r := range p {
			i, _ := strconv.Atoi(r)
			counter[requests[i][0]]--
			counter[requests[i][1]]++
		}
		// 判断是否满足要求
		good := true
		for _, temp := range counter {
			if temp != 0 {
				good = false
				break
			}
		}
		// 如满足要求，且比已记录的数量大，则更新记录
		if good && len(p) > max {
			max = len(p)
		}
	}
	return max
}

func Test1601() {
	fmt.Println(maximumRequests(5, [][]int{{2, 4}, {2, 3}}))                                                                                                         // 0
	fmt.Println(maximumRequests(2, [][]int{{1, 0}, {0, 0}, {1, 0}, {0, 1}, {0, 1}, {1, 1}, {0, 1}, {0, 0}, {0, 0}, {0, 1}, {1, 0}, {0, 0}, {0, 1}, {1, 1}, {1, 1}})) // 13
	fmt.Println(maximumRequests(2, [][]int{{1, 1}, {1, 0}, {0, 1}, {0, 0}, {0, 0}, {0, 1}, {0, 1}, {1, 0}, {1, 0}, {1, 1}, {0, 0}, {1, 0}}))                         // 11
	fmt.Println(maximumRequests(3, [][]int{{1, 2}, {1, 2}, {2, 2}, {0, 2}, {2, 1}, {1, 1}, {1, 2}}))                                                                 // 4
	fmt.Println(maximumRequests(5, [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}))                                                                         // 5
	fmt.Println(maximumRequests(3, [][]int{{0, 0}, {1, 2}, {2, 1}}))                                                                                                 // 3
	fmt.Println(maximumRequests(4, [][]int{{0, 3}, {3, 1}, {1, 2}, {2, 0}}))                                                                                         // 4
}
