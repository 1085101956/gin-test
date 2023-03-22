package algorithm

import "fmt"

// 双向链表
// 定义一个 DoubleHeroNode
type DoubleHeroNode struct {
	No       int
	Name     string
	NickName string
	pre      *DoubleHeroNode //指向前一个节点
	next     *DoubleHeroNode //指向下一个节点
}

// 插入双向链表 尾部
func InsertDoubleLink(head, newNode *DoubleHeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
	newNode.pre = temp
}

// 根据顺序插入双向链表中
func WellInstallDoubleLink(head, newNode *DoubleHeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		} else if temp.next.No >= newNode.No {
			break
		}
		temp = temp.next
	}
	//先关联后面的节点，不然会断掉
	//新节点的newNode.next = temp.next
	newNode.next = temp.next
	newNode.pre = temp
	if temp.next != nil {
		//双向列表 temp 的下一个节点 = newNode
		temp.next.pre = newNode
	}
	temp.next = newNode
}

// 删除双向链表中的一个节点
func DeleteDoubleLink(head *DoubleHeroNode, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil { //说明到了双向链表尾部，找不到了
			break
		} else if temp.next.No == id {
			//说明找到了
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
		//这儿很绕，此时temp.next 已经等于 temp.next.next 了
		//需要做一个判断，如果 next.pre != nil 那么才链接
		if temp.next != nil {
			temp.next.pre = temp
		}
	}
}

// 双向链表 倒叙展示
func ShowDoubleLinkOrderDesc(head *DoubleHeroNode) {
	//创建一个辅助节点
	temp := head
	//先判断这个链表是不是空链表
	if temp.next == nil {
		fmt.Println("该链表是一个空链表")
		return
	}

	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	for {
		fmt.Printf("[%d , %s , %s ] ==>", temp.No, temp.Name, temp.NickName)
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
	fmt.Println()
}

// 展示双向链表
func ShowDoubleLink(head *DoubleHeroNode) {
	//1.创建一个辅助节点
	temp := head
	//先判断这个链表是不是空链表
	if temp.next == nil {
		fmt.Println("该链表是一个空链表")
		return
	}
	for {
		fmt.Printf("[%d , %s , %s ] ==>", temp.next.No, temp.next.Name, temp.next.NickName)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
	fmt.Println()
}
func UseDoubleLink() {
	//afterbody := algorithm.HeroNode{}
	afterbody1 := DoubleHeroNode{}

	head := DoubleHeroNode{
		No:       56,
		Name:     "林冲",
		NickName: "test",
	}
	head1 := DoubleHeroNode{
		No:       57,
		Name:     "林2冲",
		NickName: "tesxxt",
	}
	head2 := DoubleHeroNode{
		No:       58,
		Name:     "林w冲",
		NickName: "teszzt",
	}
	head3 := DoubleHeroNode{
		No:       2,
		Name:     "林w冲",
		NickName: "teszzt",
	}
	//从尾部插入链表
	//InsertDoubleLink(&afterbody1,&head)
	//InsertDoubleLink(&afterbody1,&head1)
	//InsertDoubleLink(&afterbody1,&head2)
	//按照no排序插入链表
	WellInstallDoubleLink(&afterbody1, &head1)
	WellInstallDoubleLink(&afterbody1, &head2)
	WellInstallDoubleLink(&afterbody1, &head3)
	WellInstallDoubleLink(&afterbody1, &head)
	//algorithm.ShowSingleLinkList(&afterbody)
	ShowDoubleLink(&afterbody1)
	ShowDoubleLinkOrderDesc(&afterbody1)

	DeleteDoubleLink(&afterbody1, 57)

	ShowDoubleLinkOrderDesc(&afterbody1)

	//DeleteSingleLinkList(&afterbody1,56)
	//ShowSingleLinkList(&afterbody1)
}
