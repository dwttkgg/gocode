package main

import (
	"fmt"
	"html/template"
	"net/http"
)
type User struct {
	Name string
	gender string
	Age int
}
func (u *User) GetUsergender() string{
	return u.gender
}
func sayHello(w http.ResponseWriter,r *http.Request){
	// 定义模板
	t:=template.New("f.tmpl")

	// 自定义函数,要么只有一个返回值,要么两个返回值  第二个返回值必须是error
	k:= func(name string) (string,error){
		return name +"年前又帅气！！！",nil
	}
	// 告诉模板引擎 多了一个自定义函数
	t.Funcs(template.FuncMap{
		"kua":k,
	})
	// 解析模板
	_,err:=t.ParseFiles("./f.tmpl")
	if err!=nil{
		fmt.Println("parse template faild ,err:=%v\n",err)
		return
	}
	name:="甘雨"
	//渲染模板
	t.Execute(w,name)
}
func demo1(w http.ResponseWriter,r *http.Request){
	// 定义模板
	// 解析模板
	t,err:=template.ParseFiles("./t.tmpl","./ul.tmpl")
	if err!=nil{
		fmt.Println("parse template faild ,err:=%v\n",err)
		return
	}
	name:="甘雨"
	//渲染模板
	t.Execute(w,name)
}

func main() {
	http.HandleFunc("/sayHello",sayHello)
	http.HandleFunc("/demo1",demo1)
	err:=http.ListenAndServe(":8088",nil)
	if err!=nil{
		fmt.Println("http service start failed ,err：%v ",err)
		return
	}

}
