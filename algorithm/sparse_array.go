package algorithm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ValNode struct {
	Row int
	Col int
	Val int
}

// 一个稀疏数组算法
func SparseArray() {
	//先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	var sparseArr []ValNode
	//标准的稀疏数组应该有一个 记录元素的二维数组的规模(行和列)
	InitialValNode := ValNode{
		Row: 11,
		Col: 11,
		Val: 0,
	}
	sparseArr = append(sparseArr, InitialValNode)

	for k, v := range chessMap {
		for k2, v2 := range v {
			if v2 != 0 {
				//创建一个节点 值节点
				ValNode := ValNode{
					Row: k,
					Col: k2,
					Val: v2,
				}
				sparseArr = append(sparseArr, ValNode)
			}
		}
	}

	//输出稀疏数组
	for i, v := range sparseArr {
		fmt.Printf("%d:row:%d col:%d val:%d \n", i, v.Row, v.Col, v.Val)
	}
	//存盘操作
	sparseArrJson, _ := json.Marshal(sparseArr)
	fmt.Println(string(sparseArrJson))
	file := "d:/golang/2023/project/MyGinChat/data/sparse_arr.log"
	err := ioutil.WriteFile(file, sparseArrJson, 0666)
	if err != nil {
		fmt.Println("read file err = %v", err)
	}

}
