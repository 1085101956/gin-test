package main

import (
	"NyGinChat/algorithm"
	_ "fmt"
)

func Test[T any](a, b T) T {
	return a
}

func Add[T string | int | int64 | float64](a, b T) T {
	return a + b
}

func main() {
	algorithm.SparseArr()

	//arr := []int{99, 25, 34, 11, 66, 77, 44, 22, 41, 10}
	//algorithm.BubbleSort(arr)
	//s := []int{65, 25, 34, 11, 32, 77, 44, 105, 41, 10}
	//algorithm.TestSort(s)
	//res := Test(1.1, 2.2)
	//fmt.Println(res)
	//稀疏数组
	//algorithm.SparseArray()

	//singlequeue 单队列展示
	//algorithm.UseSingleQueue()

	//环形队列示例
	//algorithm.UseCircleQueue()

	//链表的使用 单项链表
	algorithm.UseHeroNode()
	//双相链表的使用
	algorithm.UseDoubleLink()
}
