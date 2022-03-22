package L_1184

import "fmt"

// https://leetcode-cn.com/problems/distance-between-bus-stops/

// 1184. 公交站间的距离

// 环形公交路线上有 n 个站，按次序从 0 到 n - 1 进行编号。我们已知每一对相邻公交站之间的距离，distance[i] 表示编号为 i 的车站和编号为 (i + 1) % n 的车站之间的距离。

// 环线上的公交车都可以按顺时针和逆时针的方向行驶。

// 返回乘客从出发点 start 到目的地 destination 之间的最短距离。

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	d1, d2 := 0, 0
	tail := len(distance) - 1
	for p := start; p != destination; {
		if p == 0 {
			d1 += distance[tail]
			p = tail
		} else {
			d1 += distance[p-1]
			p--
		}
	}

	for p := start; p != destination; {
		if p == tail {
			d2 += distance[tail]
			p = 0
		} else {
			d2 += distance[p]
			p++
		}
	}

	if d1 < d2 {
		return d1
	} else {
		return d2
	}
}

func Test1184() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1)) // 1
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2)) // 3
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3)) // 4
}
