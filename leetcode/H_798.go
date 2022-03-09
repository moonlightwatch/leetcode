package leetcode

import (
	"fmt"
)

// https://leetcode-cn.com/problems/smallest-rotation-with-highest-score/

// 798. 得分最高的最小轮调

// 给你一个数组 nums，我们可以将它按一个非负整数 k 进行轮调，这样可以使数组变为 [nums[k], nums[k + 1], ... nums[nums.length - 1], nums[0], nums[1], ..., nums[k-1]] 的形式。此后，任何值小于或等于其索引的项都可以记作一分。
//     例如，数组为 nums = [2,4,1,3,0]，我们按 k = 2 进行轮调后，它将变成 [1,3,0,2,4]。这将记为 3 分，因为 1 > 0 [不计分]、3 > 1 [不计分]、0 <= 2 [计 1 分]、2 <= 3 [计 1 分]，4 <= 4 [计 1 分]。
// 在所有可能的轮调中，返回我们所能得到的最高分数对应的轮调下标 k 。如果有多个答案，返回满足条件的最小的下标 k 。

func bestRotation(nums []int) int {
	// 思路：主要是避免嵌套循环导致的超时。所以将计算过程分为几次循环统计操作：
	//      因为每次轮调，第一个数调到最后，必定会变成记分项。而其他的位置，轮调后可能会变成不记分。
	//      从计分变为不记分的情况只有一种，即 nums[i]==i， 这种情况在下次轮调中，会变为不记分。
	//      我们可以通过计算，获得一个数，在几次轮调后，会处于nums[i]==i的状态。据此，只要知道未轮调的得分数量，即可逐次计算得分：
	//      本次轮调得分 = 上次轮调得分 - 上次轮调中处于nums[i]==i的状态的数量 + 1
	// 1. 先统计未轮调（即k=0）时的得分数量
	// 2. 统计每次轮调中 处于nums[i]==i的状态的数量
	// 3. 根据步骤 1 的结果逐次计算论调结果
	// 4. 统计得分最高的最小轮调

	l := len(nums)
	list := make([]int, l) // 保存第 k 轮，恰好得分的分数，即 nums[i]==i 的数量
	score := 0             // 未轮调的得分数量
	// 统计初始得分
	for i := 0; i < l; i++ {
		if nums[i] <= i {
			score++
		}
	}
	// 统计各个数字在第几轮恰好得分
	for i := 0; i < l; i++ {
		// 直接计算移动几次可以实现 nums[i]==i，即为所需的K
		k := (l + (i - nums[i])) % l
		list[k] += 1 // 在统计数组中记录
	}
	maxK := 0
	maxScore := score

	// 计算每次轮调得分
	for i := 1; i < l; i++ {
		// 因为恰好得分的位置，轮调后，必定不得分
		// 本次得分 = 上次得分 减 上次恰好得分的数量 加上 头部移动到尾部这个必定得分的一个
		score = score - list[i-1] + 1
		// 顺手统计最高得分
		if maxScore < score {
			maxScore = score
			maxK = i
		}
	}
	return maxK
}

func Test798() {
	fmt.Println(bestRotation([]int{2, 3, 1, 4, 0})) // 3
	fmt.Println(bestRotation([]int{1, 3, 0, 2, 4})) // 0
}
