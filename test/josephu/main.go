package main

import "fmt"

// 人员的结构体
type Boy struct {
	no   int
	next *Boy
}

// 构建环形单向链表
// num 添加小孩个数
// 返回环形链表的第一个指针
func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}
	if num < 1 {
		fmt.Println("num的值不正确")
		return first
	}

	for i := 1; i <= num; i++ {
		boy := &Boy{
			no: i,
		}
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.next = first
		} else {
			curBoy.next = boy
			curBoy = boy
			curBoy.next = first
		}
	}
	return first
}

// 显示单向环形列表
func ShowBoy(first *Boy) {
	// 判断是否为空链表
	if first.next == nil {
		fmt.Println("链表为空，无数据")
		return
	}
	// 创建辅助指针帮助遍历
	curBoy := first
	for {
		fmt.Println("小孩的编号：", curBoy.no)
		// 退出条件
		if curBoy.next == first {
			break
		}
		curBoy = curBoy.next
	}
}

// n个人围成一圈，从m开始数起，约定编号为K的人出列，下一位从新开始从1报数
// 1<=k<=n 直到留下最后一个人
func PlayGame() {

}
func main() {
	first := AddBoy(5)
	ShowBoy(first)
}
