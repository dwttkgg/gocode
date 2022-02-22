package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// 序列化
func (this *Monster) Store() bool {
	date, err := json.Marshal(this)
	if err != nil {
		fmt.Println("序列化错误", err)
		return false
	}
	filename := "d:/monsterStore.txt"
	err = ioutil.WriteFile(filename, date, 0666)
	if err != nil {
		fmt.Println("写文件错误", err)
		return false

	}
	return true
}

// 反序列化
func (this *Monster) ReStore() bool {
	filename := "d:/monsterStore.txt"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("文件读取错误", err)
		return false
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return false
	}
	return true
}
