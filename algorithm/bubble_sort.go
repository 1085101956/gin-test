package algorithm

import "fmt"

// 冒泡排序
func BubbleSort(arr []int) {
	lenght := len(arr)
	//a := []int{4,2,8,10,30,66,1}
	for i := 0; i < lenght-1; i++ {
		for j := 0; j < lenght-1-i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
}

// 冒泡排序
func TestSort(arr []int) {
	lenght := len(arr) - 1
	var temp int
	for i := 0; i < lenght; i++ {
		for j := 0; j < lenght-i; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
}

func BubbleSortOfficial(values []int) {
	flag := true
	for i := 0; i < len(values)-1; i++ {
		flag = true
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}
