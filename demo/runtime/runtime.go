package runtime

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

func Getinfo() (funcname, filename string, line int) {
	// 向上查找几次,如果需要查看当前的函数则是0,一次调用则是1,两次调用则是2
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("未获取到数据")
	}
	funcname = runtime.FuncForPC(pc).Name()
	filename = path.Base(file)
	funcname = strings.Split(funcname, ".")[3]
	//fmt.Println(funcname)
	//fmt.Println(file)
	//fmt.Println(path.Base(file))
	//fmt.Println(line)
	return
}
