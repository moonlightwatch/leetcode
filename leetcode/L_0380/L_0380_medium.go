package L_0380

import (
	"fmt"
	"math/rand"
)

// https://leetcode-cn.com/problems/insert-delete-getrandom-o1/

// 380. O(1) 时间插入、删除和获取随机元素

// 实现RandomizedSet 类：

//     RandomizedSet() 初始化 RandomizedSet 对象
//     bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
//     bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
//     int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。

// 你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。
type RandomizedSet struct {
	m   map[int]int // 字典存储数据，用于去重
	l   []int       // 列表存储数据，用于检索
	end int         // 当前容量
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:   map[int]int{},
		l:   []int{0},
		end: 0,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.l = append(append(this.l[:this.end], val), this.l[this.end:]...)
	this.m[val] = this.end
	this.end++

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; ok {
		this.l[this.m[val]], this.l[this.end-1] = this.l[this.end-1], this.l[this.m[val]]
		this.m[this.l[this.m[val]]] = this.m[val]
		delete(this.m, val)
		this.end--
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.l[rand.Intn(this.end)]
}

func Test380() {
	r := Constructor()
	// fmt.Println(r.Insert(1), r.Remove(2), r.Insert(2), r.GetRandom(), r.Remove(1), r.Insert(2), r.GetRandom())

	// r = Constructor()
	// fmt.Println(r.Remove(0), r.Remove(0), r.Insert(0), r.GetRandom(), r.Remove(0), r.Insert(0))

	r = Constructor()
	fmt.Println(r.Insert(0), r.Insert(1), r.Remove(0), r.Insert(2), r.Remove(1), r.Insert(0), r.GetRandom()) // [null,true,true,true,true,true,2]

}
