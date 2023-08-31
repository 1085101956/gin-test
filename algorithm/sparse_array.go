package algorithm

//稀疏数组算法
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ValNode struct {
	Row int //行
	Col int //列
	Val int //值
}

// 一个稀疏数组算法
func SparseArray() {
	//先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子
	chessMap[2][4] = 1 //白子

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
	addFile(sparseArrJson)
	resumeData()
}

// 将稀疏后的数组进行存盘
func addFile(sparseArrJson []byte) {
	file := "d:/golang/2023/project/MyGinChat/data/sparse_arr.log"
	err := ioutil.WriteFile(file, sparseArrJson, 0666)
	if err != nil {
		fmt.Println("write file err = %v", err)
	}
}

// 将稀疏后的数组恢复
func resumeData() {
	var chessMap [11][11]int
	file := "d:/golang/2023/project/MyGinChat/data/sparse_arr.log"
	str, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read file err = %v", err)
	}
	ValNode := []ValNode{}
	err = json.Unmarshal(str, &ValNode)
	if err != nil {
		fmt.Println("json 解析失败")
		return
	}
	for k, v := range ValNode {
		if k == 0 {
			continue
		}
		chessMap[v.Row][v.Col] = v.Val
	}
	for _, v := range chessMap {
		for _, value := range v {
			fmt.Printf("%d\t", value)
		}
		fmt.Println()
	}
}

func SparseArr() {
	//创建一个原始数组
	var sparseArr [11][11]int
	sparseArr[1][2] = 2 //2是黑子
	sparseArr[2][3] = 1 //1是白子
	for _, v := range sparseArr {
		for _, val := range v {
			fmt.Printf("%d \t", val)
		}
		fmt.Println(v)
	}
	sparseArrays := []ValNode{}
	originalArr := ValNode{
		Row: 11,
		Col: 11,
		Val: 0,
	}
	sparseArrays = append(sparseArrays, originalArr)
	for x, y := range sparseArr {
		for j, vv := range y {
			if vv != 0 {
				array := ValNode{
					Row: x,
					Col: j,
					Val: vv,
				}
				sparseArrays = append(sparseArrays, array)
			}
		}
	}
	fmt.Println(sparseArrays)
	//遍历稀疏数组，稀疏数组恢复
	arrMap := [11][11]int{}
	for z, s := range sparseArrays {
		if z == 0 { //第一行过滤掉
			continue
		}
		//赋值给数组
		arrMap[s.Row][s.Col] = s.Val
	}
	fmt.Println(arrMap)
}
