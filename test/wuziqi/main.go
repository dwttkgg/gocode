package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	// 创建原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2
	// 显示原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	// 将原始数组转换成稀疏数组
	// 遍历chessMap 如果值不为0 创建一个node结构体
	// 然后放到对应的切片中
	var sparseArr []ValNode
	//标准的稀疏数组应该有一个 记录元素的二维数组的规模(行，列，默认值)
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			//创建一个值节点
			if v2 != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	// 输出稀疏数组
	fmt.Println("当前的稀疏数组是：：：：")
	for i, v := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, v.row, v.col, v.val)
	}

	// 将稀疏数组持久化
	filename := "d:/chessMap.data"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//写入带缓存的
	writer := bufio.NewWriter(file)
	for _, v := range sparseArr {
		fmt.Println(v)
		data, err := json.Marshal(v)
		fmt.Println(string(data))
		if err != nil {
			fmt.Println(err)
		}
		writer.WriteString(string(data))
		fmt.Println(string(data))
	}

	// 将稀疏数组还原程二维数组
	var chessMap2 [11][11]int

	for i, v := range sparseArr {
		if i != 0 {
			chessMap2[v.row][v.col] = v.val
		}
	}

}
