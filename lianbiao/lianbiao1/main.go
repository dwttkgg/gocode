package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode // 指向下一个节点
}

// 第一种插入方式，直接在队列尾插入
func InsertHeroNode(head *HeroNode, newHero *HeroNode) {
	// 先找到链表最后的节点
	// 创建一个辅助节点来操作，头节点不能改动
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next //移动到下一个节点
	}
	temp.next = newHero
}

// 第二种插入方式 ，按no大小插入
func InsertHeroNode2(head *HeroNode, newHero *HeroNode) {
	// 先找到链表最后的节点
	// 创建一个辅助节点来操作，头节点不能改动
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHero.no {
			break
		} else if temp.next.no == newHero.no {
			flag = false
		}
		temp = temp.next //移动到下一个节点
	}
	if !flag {
		fmt.Println("已存在该no")
		return
	} else {
		newHero.next = temp.next
		temp.next = newHero
	}
}

// 根据no 删除对应的链表
func DeleteNode(head *HeroNode, no int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == no {
			flag = true
			break
		}
		temp = temp.next //移动到下一个节点
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("未找到数据")

	}
}
func ListHero(head *HeroNode) {
	// 创建辅助节点，并判断链表是否为空
	temp := head
	if temp.next == nil {
		return
	}
	for {
		fmt.Printf("%d  %s  %s\n", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}
func main() {
	// 创建一个头节点
	head := &HeroNode{}

	// 创建一个新的节点
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	hero4 := &HeroNode{
		no:       4,
		name:     "吴用",
		nickname: "智多星",
	}
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero4)

	DeleteNode(head, 2)
	ListHero(head)
}
