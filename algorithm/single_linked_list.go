package algorithm

import "fmt"

// 单链表
// 定义一个 HeroNode
type HeroNode struct {
	No       int
	Name     string
	NickName string
	next     *HeroNode //指向下一个节点
}

func Init() HeroNode {
	//创建一个头节点。
	head := HeroNode{}
	return head
}

// 使用链表
func UseHeroNode() {
	//afterbody := algorithm.HeroNode{}
	afterbody1 := HeroNode{}

	head := HeroNode{
		No:       56,
		Name:     "林冲",
		NickName: "test",
	}
	head1 := HeroNode{
		No:       57,
		Name:     "林2冲",
		NickName: "tesxxt",
	}
	head2 := HeroNode{
		No:       58,
		Name:     "林w冲",
		NickName: "teszzt",
	}
	head3 := HeroNode{
		No:       1,
		Name:     "周玲",
		NickName: "哈哈哈",
	}
	//从尾部插入链表
	//algorithm.InsterHeroNode(&afterbody,&head)
	//algorithm.InsterHeroNode(&afterbody,&head1)
	//algorithm.InsterHeroNode(&afterbody,&head2)
	//按照no排序插入链表
	InsterHeroNode(&afterbody1, &head1)
	InsterHeroNode(&afterbody1, &head2)
	InsterHeroNode(&afterbody1, &head)
	InsterHeroNode(&afterbody1, &head3)

	//algorithm.ShowSingleLinkList(&afterbody)
	ShowSingleLinkList(&afterbody1)
	DeleteSingleLinkList(&afterbody1, 56)
	ShowSingleLinkList(&afterbody1)
}

// 从链表的最后加入
func InsterAfterHeroNode(head, node *HeroNode) {
	//思路，先找到该链表的最后这个节点
	temp := head //创建一个辅助节点
	for {
		//如果
		if temp.next == nil { //表示找到了链表的最后节点了
			break
		}
		temp = temp.next
	}
	temp.next = node //把新的节点 追加到后面节点上
}

// 根据 no 从小到大的顺序插入
func InsterHeroNode(head, node *HeroNode) {
	//思路，先找到该链表的最后这个节点
	temp := head //创建一个辅助节点
	for {
		if temp.next == nil { //说明已经找到了链表的最后了
			break
		} else if temp.next.No >= node.No {
			//说明我们的链表就应该插入到这个temp后面
			break
		}
		temp = temp.next
	}
	//
	node.next = temp.next
	temp.next = node //把新的节点 追加到后面节点上
}

// 显示链表所有的节点信息
func ShowSingleLinkList(head *HeroNode) {
	//1.创建一个辅助节点
	temp := head
	//先判断这个链表是不是空链表
	if temp.next == nil {
		fmt.Println("该链表是一个空链表")
		return
	}
	//2.遍历这个链表
	for {
		fmt.Printf("[%d , %s , %s ] ==>", temp.next.No, temp.next.Name, temp.next.NickName)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
	fmt.Println()
}

// 删除一个节点
func DeleteSingleLinkList(head *HeroNode, no int) {
	temp := head
	flag := true
	for {
		if temp.next == nil { //说明找到了链表的最后没有找到需要删除的ID
			flag = false
			break
		} else if temp.next.No == no {
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Printf("未找到需要删除的no->%d", no)
	} else {
		temp.next = temp.next.next
	}

}
