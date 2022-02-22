package main

import (
	"fmt"
	"sort"
)

func main() {
	// map 的声明和注意事项 map创建的三种方式
	var a map[int]string
	// 在使用map之前，需要先make 分配空间
	a = make(map[int]string, 10)
	a[1] = "宋江"
	a[2] = "无用"
	a[2] = "智多星"
	fmt.Println(a)

	var cli map[string]string = make(map[string]string)
	cli["11"] = "sss"
	cli["12"] = "sss"
	cli["31"] = "sss"
	fmt.Println(cli)

	var sli map[string]string = map[string]string{
		"ss":    "11",
		"sss":   "1231",
		"sssss": "dadfa",
	}

	fmt.Println(sli)

	ms := make(map[string]map[string]string)
	ms["01"] = make(map[string]string, 10)
	ms["01"]["name"] = "张三"
	ms["01"]["age"] = "33"
	ms["01"]["sex"] = "men"
	ms["02"] = make(map[string]string, 10)
	ms["02"]["name"] = "李四"
	ms["02"]["age"] = "44"
	ms["02"]["sex"] = "women"

	fmt.Println(ms)

	// 删除操作
	delete(ms, "01")
	fmt.Println(ms)

	// 删除所有的map 方法两种
	// 1 遍历所有的key  挨个删除
	// 2 重新make空间 GC回收之前的map

	for i, v := range ms {
		fmt.Printf("key %v ,值 %v", i, v)
	}

	// map 切片
	// 声明一个map切片
	var maps []map[string]string
	maps = make([]map[string]string, 2)
	if maps[0] == nil {
		maps[0] = make(map[string]string, 2)
		maps[0]["name"] = "sss22"
		maps[0]["age"] = "22"

	}
	if maps[1] == nil {
		maps[1] = make(map[string]string, 2)
		maps[1]["name"] = "lisi"
		maps[1]["age"] = "222"

	}
	// 不能直接使用赋值的方式增加超过容量的数据，需要使用append才能让slice动态增长
	// if maps[2] == nil {
	// 	maps[2] = make(map[string]string, 2)
	// 	maps[2]["name"] = "wangmazi"
	// 	maps[2]["age"] = "111"

	// }
	cons := map[string]string{
		"name": "wangmazi",
		"age":  "22",
	}

	maps = append(maps, cons)

	fmt.Println(maps)

	// map的排序
	// 先将map的key放入切片中，对切片排序，然后遍历key取对应的值
	var keyd = make(map[int]int, 10)
	var keys []int
	keyd[0] = 10
	keyd[9] = 9
	keyd[4] = 92
	keyd[95] = 5
	for k, _ := range keyd {
		keys = append(keys, k)
	}
	fmt.Println(keys)

	sort.Ints(keys)
	fmt.Println(keys)
	for _, k := range keys {
		fmt.Printf("keyd[%v]=%v", k, keyd[k])
	}

}
