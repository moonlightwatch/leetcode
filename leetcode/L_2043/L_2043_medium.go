package L_2043

import "fmt"

// https://leetcode-cn.com/problems/simple-bank-system/

// 2043. 简易银行系统

// 你的任务是为一个很受欢迎的银行设计一款程序，以自动化执行所有传入的交易（转账，存款和取款）。银行共有 n 个账户，编号从 1 到 n 。
// 每个账号的初始余额存储在一个下标从 0 开始的整数数组 balance 中，其中第 (i + 1) 个账户的初始余额是 balance[i] 。

// 请你执行所有 有效的 交易。如果满足下面全部条件，则交易 有效 ：

//     指定的账户数量在 1 和 n 之间，且
//     取款或者转账需要的钱的总数 小于或者等于 账户余额。

// 实现 Bank 类：

//     Bank(long[] balance) 使用下标从 0 开始的整数数组 balance 初始化该对象。
//     boolean transfer(int account1, int account2, long money) 从编号为 account1 的账户向编号为 account2 的账户转帐 money 美元。如果交易成功，返回 true ，否则，返回 false 。
//     boolean deposit(int account, long money) 向编号为 account 的账户存款 money 美元。如果交易成功，返回 true ；否则，返回 false 。
//     boolean withdraw(int account, long money) 从编号为 account 的账户取款 money 美元。如果交易成功，返回 true ；否则，返回 false 。

type Bank struct {
	balance []int64
	count   int
}

func Constructor(balance []int64) Bank {
	return Bank{
		balance: balance,
		count:   len(balance),
	}
}

func (this *Bank) checkAccount(account int) bool {
	return account > 0 && account <= this.count
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if this.checkAccount(account1) && this.checkAccount(account2) && this.balance[account1-1] >= money {
		this.balance[account1-1] -= money
		this.balance[account2-1] += money
		return true
	}
	return false
}

func (this *Bank) Deposit(account int, money int64) bool {
	if this.checkAccount(account) {
		this.balance[account-1] += money
		return true
	}
	return false
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if this.checkAccount(account) && this.balance[account-1] >= money {
		this.balance[account-1] -= money
		return true
	}
	return false
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */

func Test2043() {
	// ["Bank","withdraw","transfer","deposit","transfer","withdraw"]
	// [[[10,100,20,50,30]],[3,10],[5,1,20],[5,20],[3,4,15],[10,50]]

	obj := Constructor([]int64{10, 100, 20, 50, 30})
	fmt.Println(obj.Withdraw(3, 10))    // true
	fmt.Println(obj.Transfer(5, 1, 20)) // true
	fmt.Println(obj.Deposit(5, 20))     // true
	fmt.Println(obj.Transfer(3, 4, 15)) // false
	fmt.Println(obj.Withdraw(10, 50))   // false
}
