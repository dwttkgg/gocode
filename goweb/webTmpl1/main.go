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
	// 解析模板
	t,err:=template.ParseFiles("./hello.tmpl")
	if err!=nil{
		fmt.Println("Parse template failed ,err:%v",err)
		return
	}
	//渲染模板
	user:=&User{
		Name: "ganyu",
		gender: "nv",
		Age: 10088,
	}
	map1 :=map[string]interface{}{
		"name":"神里",
		"age":18,
		"gender":"nv",
	}
	hobbyList:=[]string{
		"篮球",
		"羽毛球",
		"双色球",
	}
	err=t.Execute(w,map[string]interface{}{
		"user":user,
		"map1":map1,
		"hobby":hobbyList,
	})

	if err!=nil{
		fmt.Println("render template failed ,err:%v",err)
		return
	}

}

func main() {
	http.HandleFunc("/sayHello",sayHello)
	err:=http.ListenAndServe(":8088",nil)
	if err!=nil{
		fmt.Println("http service start failed ,err：%v ",err)
		return
	}

}
