package L_0954

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/array-of-doubled-pairs/

// 954. 二倍数对数组

// 给定一个长度为偶数的整数数组 arr，只有对 arr 进行重组后可以满足 “对于每个 0 <= i < len(arr) / 2，都有 arr[2 * i + 1] = 2 * arr[2 * i]” 时，返回 true；否则，返回 false。

// 题目意思是：每个奇数索引的值都是其前一个值的2倍

func canReorderDoubled(arr []int) bool {
	l := len(arr)
	zeroCount := 0 // 0 的数量
	numCount := make(map[int]int, l)
	// 统计一波
	for i := 0; i < l; i++ {
		if arr[i] == 0 {
			zeroCount++
		} else {
			if _, ok := numCount[arr[i]]; !ok {
				numCount[arr[i]] = 0
			}
			numCount[arr[i]]++
		}
	}
	// 如果有奇数个0，表示必定有一个 0 不符合 arr[2 * i + 1] = 2 * arr[2 * i]
	if zeroCount%2 != 0 {
		return false
	}
	// 处理数字
	keys := make([]int, 0, len(numCount))
	for k, v := range numCount {
		if v == 0 {
			continue
		}
		keys = append(keys, k)
	}
	// 排个序，保证从小到大处理
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, num := range keys {
		count := numCount[num]
		if count <= 0 { // 不处理已经处理完的数据
			continue
		}
		if _, ok := numCount[num*2]; ok { // 若有对应的，二倍数字，则减掉较小的数量
			if numCount[num*2] < numCount[num] {
				numCount[num] -= numCount[num*2]
				numCount[num*2] = 0
			} else {
				numCount[num*2] -= numCount[num]
				numCount[num] = 0
			}
		}
	}
	// 若有没处理完的数据，则不符合要求
	for _, c := range numCount {
		if c != 0 {
			return false
		}
	}
	return true
}

func Test954() {
	fmt.Println(canReorderDoubled([]int{-62, 86, 96, -18, 43, -24, -76, 13, -31, -26, -88, -13, 96, -44, 9, -20, -42, 100, -44, -24, -36, 28, -32, 58, -72, 20, 48, -36, -45, 14, 24, -64, -90, -44, -16, 86, -33, 48, 26, 29, -84, 10, 46, 50, -66, -8, -38, 92, -19, 43, 48, -38, -22, 18, -32, -48, -64, -10, -22, -48})) // true

	fmt.Println(canReorderDoubled([]int{10, 20, 40, 80}))                   // true
	fmt.Println(canReorderDoubled([]int{2, 1, 2, 6}))                       // false
	fmt.Println(canReorderDoubled([]int{2, 1, 1, 4, 8, 8}))                 // false
	fmt.Println(canReorderDoubled([]int{1, 2, 1, -8, 8, -4, 4, -4, 2, -2})) // true
	fmt.Println(canReorderDoubled([]int{2, 4, 0, 0, 8, 1}))                 // true
	fmt.Println(canReorderDoubled([]int{3, 1, 3, 6}))                       // false
	fmt.Println(canReorderDoubled([]int{2, 1, 2, 6}))                       // false
	fmt.Println(canReorderDoubled([]int{4, -2, 2, -4}))                     // true
}
