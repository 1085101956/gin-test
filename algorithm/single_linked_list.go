package algorithm

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

// 从链表的最后加入
func InsterHeroNode(head, node *HeroNode) {
	//思路，先找到该链表的最后这个节点
	temp := head //创建一个辅助节点
	for {
		//如果
		if temp.next == nil { //表示找到了链表的最后节点了
			break
		}
		temp = head.next
	}
	temp.next = node //把新的节点 追加到后面节点上
}
func ShowSingleLinkList() {

}
