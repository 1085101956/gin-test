package algorithm

import (
	"errors"
	"fmt"
	"os"
)

type SingleQueue struct {
	MaxSize int
	Array   [4]int //数组=>模拟队列
	Front   int    //表示指向队列首
	Rear    int    //表示指向队列的尾部
}

// 添加数据到队列
func (this *SingleQueue) AddSingleQueue(value int) (err error) {
	//先判断队列是否已经满了
	if this.Rear == this.MaxSize-1 { //重要提示，rear 是队列尾部（含最后元素）
		return errors.New("队列已经满了，无法写入")
	}
	//添加队列
	this.Rear++
	this.Array[this.Rear] = value
	return
}

// 显示队列，找到队首，然后遍历到队尾
func (this *SingleQueue) ShowSingleQueue() {
	//this.front 不包含队首的元素
	for i := this.Front + 1; i <= this.Rear; i++ {
		fmt.Printf("array[%d]=%d\n", i, this.Array[i])
	}
}

// 获取队列
func (this *SingleQueue) GetSingleQueue() (val int, err error) {
	//先判断队列是否为空
	if this.Front == this.Rear {
		return -1, errors.New("队列为空，已经取不出来了")
	}
	this.Front++
	return this.Array[this.Front], err
}

// 展示
func UseSingleQueue() {
	SingleQueue := SingleQueue{
		MaxSize: 4,
		Front:   -1,
		Rear:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("1.输入add 表示添加数据到队列")
		fmt.Println("2.输入get 表示从队列获取数据")
		fmt.Println("3.输入show 表示显示队列数据")
		fmt.Println("4.输入exit 表示退出")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列的数字")
			fmt.Scanln(&val)
			err := SingleQueue.AddSingleQueue(val)
			if err == nil {
				fmt.Println("加入队列成功！")
			} else {
				fmt.Println(err.Error())
			}
			break
		case "get":
			value, err := SingleQueue.GetSingleQueue()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("数据已经拿到：%d \n", value)
		case "show":
			fmt.Println("队列当前情况：")
			SingleQueue.ShowSingleQueue()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("您的输入有误，请重新输入")
		}
	}

}
