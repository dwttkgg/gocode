package UtilDemo

import (
	"testing"
)

// testing框架查找 *_test.go文件，然后_test.go文件去查找 Test*的函数
// 测试用例的文件名必须以 _test.go结尾 ，如sub_test.go
// 测试用例函数必须以Test开头  一般为Test加需要测试的函数名
// 形参必须是 *testing.T
// 测试命令 go test 运行正确无日志，错误输出日志
//         go test -v 正确错误都将输出日志
// 出现错误时 可以使用 t.Fatalf来格式化输出信息，并退出程序
// t.Logf可以输出对应日志
// 测试单个文件 go test sub_test.go cat.go
// 测试单个方法 go test -v -test.run TestSub

func TestSub(t *testing.T) {
	res2 := Subdc(9, 5)
	if res2 != 5 {
		t.Fatalf("执行错误预期值%v ,实际值%v", 5, res2)
	}

	t.Logf("Subdc执行正确")
}
