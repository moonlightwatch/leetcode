package leetcode

import "fmt"

// 今天，书店老板有一家店打算试营业 customers.length 分钟。每分钟都有一些顾客（customers[i]）会进入书店，所有这些顾客都会在那一分钟结束后离开。
// 在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，那么 grumpy[i] = 1，否则 grumpy[i] = 0。 当书店老板生气时，那一分钟的顾客就会不满意，不生气则他们是满意的。
// 书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 X 分钟不生气，但却只能使用一次。
// 请你返回这一天营业下来，最多有多少客户能够感到满意的数量。

/*

思路：

max  ： 因克制情绪导致的客户满意数量的最大值
temp ： X 天内因克制情绪导致的客户满意数量
count： 满意客户总数

逐日求取上述值。
- 当找到新的 max 时，则 count 需减去旧的 max， 加上新的 max


*/

func maxSatisfied(customers []int, grumpy []int, X int) int {
	count := 0
	max := 0
	temp := 0
	i := 0
	l := len(customers)
	for i < X {
		count += customers[i]
		if grumpy[i] == 1 {
			max += customers[i]
			temp += customers[i]
		}
		i++
	}
	for i < l {
		temp -= customers[i-X] * grumpy[i-X]
		if grumpy[i] == 0 {
			count += customers[i]
		} else {
			temp += customers[i]
			if max < temp {
				count -= max
				count += temp
				max = temp
			}
		}
		i++
	}
	return count
}

// 执行用时：40 ms, 在所有 Go 提交中击败了71.19% 的用户
// 内存消耗：6.6 MB, 在所有 Go 提交中击败了55.93% 的用户

func Test1052() {
	fmt.Printf("%v\n\n", maxSatisfied( // 16
		[]int{1, 0, 1, 2, 1, 1, 7, 5},
		[]int{0, 1, 0, 1, 0, 1, 0, 1}, 3))
	fmt.Printf("%v\n\n", maxSatisfied( // 19
		[]int{9, 10, 4, 5},
		[]int{1, 0, 1, 1}, 1))
	fmt.Printf("%v\n\n", maxSatisfied( // 17
		[]int{2, 6, 6, 9},
		[]int{0, 0, 1, 1}, 1))
	fmt.Printf("%v\n\n", maxSatisfied( // 23
		[]int{3, 8, 8, 7, 1},
		[]int{1, 1, 1, 1, 1}, 3))
}
