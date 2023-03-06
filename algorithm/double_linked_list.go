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
	temp.next = newNode
	newNode.pre = temp

	newNode.next = temp.next
	temp.pre = newNode.pre
}

// 删除双向链表
func DeleteDoubleLink() {

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
	//从尾部插入链表
	//algorithm.InsertDoubleLink(&afterbody,&head)
	//algorithm.InsertDoubleLink(&afterbody,&head1)
	//algorithm.InsertDoubleLink(&afterbody,&head2)
	//按照no排序插入链表
	WellInstallDoubleLink(&afterbody1, &head1)
	WellInstallDoubleLink(&afterbody1, &head2)
	WellInstallDoubleLink(&afterbody1, &head)
	//algorithm.ShowSingleLinkList(&afterbody)
	ShowDoubleLink(&afterbody1)
	//DeleteSingleLinkList(&afterbody1,56)
	//ShowSingleLinkList(&afterbody1)
}
