package L_1606

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/find-servers-that-handled-most-number-of-requests/

// 1606. 找到处理最多请求的服务器

// 你有 k 个服务器，编号为 0 到 k-1 ，它们可以同时处理多个请求组。每个服务器有无穷的计算能力但是 不能同时处理超过一个请求 。请求分配到服务器的规则如下：

//     第 i （序号从 0 开始）个请求到达。
//     如果所有服务器都已被占据，那么该请求被舍弃（完全不处理）。
//     如果第 (i % k) 个服务器空闲，那么对应服务器会处理该请求。
//     否则，将请求安排给下一个空闲的服务器（服务器构成一个环，必要的话可能从第 0 个服务器开始继续找下一个空闲的服务器）。
//     比方说，如果第 i 个服务器在忙，那么会查看第 (i+1) 个服务器，第 (i+2) 个服务器等等。

// 给你一个 严格递增 的正整数数组 arrival ，表示第 i 个任务的到达时间，和另一个数组 load ，其中 load[i] 表示第 i 个请求的工作量（也就是服务器完成它所需要的时间）。
// 你的任务是找到 最繁忙的服务器 。最繁忙定义为一个服务器处理的请求数是所有服务器里最多的。

// 请你返回包含所有 最繁忙服务器 序号的列表，你可以以任意顺序返回这个列表。

type hi struct{ sort.IntSlice }

func (h *hi) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hi) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type pair struct{ end, id int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func busiestServers(k int, arrival, load []int) (ans []int) {
	available := hi{make([]int, k)}
	for i := 0; i < k; i++ {
		available.IntSlice[i] = i
	}
	busy := hp{}
	requests := make([]int, k)
	maxRequest := 0
	for i, t := range arrival {
		for len(busy) > 0 && busy[0].end <= t {
			heap.Push(&available, i+((busy[0].id-i)%k+k)%k) // 保证得到的是一个不小于 i 的且与 id 同余的数
			heap.Pop(&busy)
		}
		if available.Len() == 0 {
			continue
		}
		id := heap.Pop(&available).(int) % k
		requests[id]++
		if requests[id] > maxRequest {
			maxRequest = requests[id]
			ans = []int{id}
		} else if requests[id] == maxRequest {
			ans = append(ans, id)
		}
		heap.Push(&busy, pair{t + load[i], id})
	}
	return
}

func Test1606() {
	fmt.Println(busiestServers(2, []int{2, 3, 4, 8}, []int{3, 2, 4, 3})) // 1
	// fmt.Println(busiestServers(3, []int{1, 2, 3, 4, 5}, []int{5, 2, 3, 3, 3}))               // 1
	// fmt.Println(busiestServers(3, []int{1, 2, 3}, []int{10, 12, 11}))                        // 0 1 2
	// fmt.Println(busiestServers(3, []int{1, 2, 3, 4, 8, 9, 10}, []int{5, 2, 10, 3, 1, 2, 2})) // 1
	// fmt.Println(busiestServers(1, []int{1}, []int{1}))                                       // 0
}

// 暴力解法，会超时
// func busiestServers(k int, arrival []int, load []int) []int {
//     max := 0
// 	results := []int{}
// 	count := make([]int, k)
// 	serverStatus := make([]int, k)
// 	lastTask := 0
// 	for i := 0; i < len(arrival); i++ {
// 		// fmt.Printf("server status[%d]:%v\n", i, serverStatus)
// 		temp := arrival[i] - lastTask
// 		for j := 0; j < k; j++ {
// 			if temp < serverStatus[j] {
// 				serverStatus[j] -= temp
// 			} else {
// 				serverStatus[j] = 0
// 			}
// 		}
// 		lastTask = arrival[i]
// 		serverIndex := i % k
// 		// 如果服务器正在处理任务，则搜索下一个服务器
// 		if serverStatus[serverIndex] != 0 {
// 			for j := (serverIndex + 1); ; j++ {
// 				if j == k {
// 					j = 0
// 				}
//                 if j == serverIndex {
// 					break
// 				}
// 				if serverStatus[j] == 0 {
// 					serverIndex = j
// 					break
// 				}
// 			}
// 			// 没找到空闲服务器
// 			if serverIndex == i%k {
// 				continue
// 			}
// 		}
// 		// 分配任务
// 		serverStatus[serverIndex] = load[i]
// 		count[serverIndex]++
// 		if count[serverIndex] > max {
// 			max = count[serverIndex]
// 			results = []int{serverIndex}
// 		} else if count[serverIndex] == max {
// 			results = append(results, serverIndex)
// 		}
// 	}
// 	// fmt.Println(count)
// 	return results
// }
