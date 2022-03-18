package L_0303

import "fmt"

// 给定一个整数数组  nums，求出数组从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点。
// 实现 NumArray 类：
//     NumArray(int[] nums) 使用数组 nums 初始化对象
//     int sumRange(int i, int j) 返回数组 nums 从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点（也就是 sum(nums[i], nums[i + 1], ... , nums[j])）

type NumArray struct {
	array []int
}

func Constructor303(nums []int) NumArray {
	return NumArray{array: nums}
}

func (A *NumArray) SumRange(i int, j int) int {
	result := 0
	for index := i; index <= j; index++ {
		result += A.array[index]
	}
	return result
}

// 执行用时：92 ms, 在所有 Go 提交中击败了23.27% 的用户
// 内存消耗：9.4 MB, 在所有 Go 提交中击败了64.08% 的用户

func Test303() {
	nums := Constructor303([]int{-2, 0, 3, -5, 2, -1})
	fmt.Printf("%v\n", nums.SumRange(0, 2))
	fmt.Printf("%v\n", nums.SumRange(2, 5))
	fmt.Printf("%v\n", nums.SumRange(0, 5))
}
