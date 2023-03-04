package algorithm

import (
	"errors"
	"fmt"
	"os"
)

// 实现环形队列
// 使用一个结构体管理环形队列
type CircleQueue struct {
	MaxSize int    //数组最多存放值
	Array   [5]int //数组
	Head    int    //队列的第一个值
	Tail    int    //队列的最后一个值
}

// 加入环形队列
func (this *CircleQueue) PushQueue(value int) (err error) {
	if this.IsFill() {
		return errors.New("队列已经满了！")
	}
	//先把值给尾部
	this.Array[this.Tail] = value
	this.Tail = (this.Tail + 1) % this.MaxSize
	return
}

// 出队列
func (this *CircleQueue) PopQueue() (val int, err error) {
	if this.IsEmpey() {
		return -1, errors.New("队列里面不存在数据")
	}
	//取出head，是指向队首，并且包含队首的元素
	val = this.Array[this.Head]
	this.Head = (this.Head + 1) % this.MaxSize
	return
}
func (this *CircleQueue) ShowQueue() {
	fmt.Println("当前队列里面的值如下:")
	size := this.Size()
	if size == 0 {
		fmt.Println("环形队列里面还没有任何数值！")
		return
	}
	TempHead := this.Head
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d", TempHead, this.Array[TempHead])
		TempHead = (TempHead + 1) % this.MaxSize
	}
}

// 判断环形队列是否满了
func (this *CircleQueue) IsFill() bool {
	fmt.Println(this.Tail, this.Head)
	return (this.Tail+1)%this.MaxSize == this.Head
}

// 判断环形队列是否为空
func (this *CircleQueue) IsEmpey() bool {
	return this.Tail == this.Head
}

// 获取环形队列里面有多少个元素
func (this *CircleQueue) Size() int {
	fmt.Println((this.Tail + this.MaxSize - this.Head) % this.MaxSize)
	return (this.Tail + this.MaxSize - this.Head) % this.MaxSize
}

func UseCircleQueue() {
	SingleQueue := &CircleQueue{
		MaxSize: 5,
		Head:    0,
		Tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1.输入add 表示添加数据到队列")
		fmt.Println("2.输入get 表示从队列获取数据")
		fmt.Println("3.输入show 表示显示队列数据")
		fmt.Println("4.输入size 表示查看队列里面存在多少数据")
		fmt.Println("5.输入exit 表示退出")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列的数字")
			fmt.Scanln(&val)
			err := SingleQueue.PushQueue(val)
			if err == nil {
				fmt.Println("加入队列成功！")
			} else {
				fmt.Println(err.Error())
			}
			break
		case "get":
			value, err := SingleQueue.PopQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("数据已经拿到：%d \n", value)
			}
		case "show":
			fmt.Println("队列当前情况：")
			SingleQueue.ShowQueue()
		case "size":
			fmt.Println("当前队列里面存在：", SingleQueue.Size(), "个数据")
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("您的输入有误，请重新输入")
		}
	}

}
