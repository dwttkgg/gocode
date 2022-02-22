package monster

import (
	"fmt"
	"testing"
)

func TestMonster(t *testing.T) {
	monster := &Monster{
		Name:  "孙悟空",
		Age:   500,
		Skill: "七十二变",
	}
	fmt.Println(monster)
	res := monster.Store()
	if !res {
		t.Fatal("测试错误", res)
	}
	t.Log("测试为真", res)
}

func TestUnMonster(t *testing.T) {
	var monster Monster
	err := monster.ReStore()
	if !err {
		t.Fatal("", err)
	}

	t.Log("反序列化成功", monster)
}
