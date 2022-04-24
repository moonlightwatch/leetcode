package L_0398

import "fmt"

// https://leetcode-cn.com/problems/random-pick-index/

// 398. 随机数索引

// 给定一个可能含有重复元素的整数数组，要求随机输出给定的数字的索引。 您可以假设给定的数字一定存在于数组中。

// 注意：
// 数组大小可能非常大。 使用太多额外空间的解决方案将不会通过测试。

type Solution struct {
	cache map[int][]int // 数字和索引列表的map
}

func Constructor(nums []int) Solution {
	s := Solution{
		cache: make(map[int][]int, len(nums)),
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := s.cache[nums[i]]; !ok {
			s.cache[nums[i]] = []int{i}
		} else {
			s.cache[nums[i]] = append(s.cache[nums[i]], i)
		}
	}
	return s
}

func (this *Solution) Pick(target int) int {
	// 先将对应数字的索引列表第一位倒到最后一位
	this.cache[target] = append(this.cache[target][1:], this.cache[target][0])
	// 返回第一位，以保证各个索引出现次数概率相同
	return this.cache[target][0]
}

func Test398() {
	obj := Constructor([]int{1, 2, 3, 3, 3})
	fmt.Println(obj.Pick(1)) // 0
	fmt.Println(obj.Pick(3)) // 2/3/4
	fmt.Println(obj.Pick(3)) // 2/3/4
	fmt.Println(obj.Pick(3)) // 2/3/4
	fmt.Println(obj.Pick(3)) // 2/3/4
	fmt.Println(obj.Pick(3)) // 2/3/4
	fmt.Println(obj.Pick(3)) // 2/3/4
	fmt.Println(obj.Pick(3)) // 2/3/4
	fmt.Println(obj.Pick(3)) // 2/3/4
}
