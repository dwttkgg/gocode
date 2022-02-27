package main

import (
	"fmt"
	"html/template"
	"net/http"
)
func sayHello(w http.ResponseWriter,r *http.Request){
	// 解析模板
	t,err:=template.ParseFiles("./hello.tmpl")
	if err!=nil{
		fmt.Println("Parse template failed ,err:%v",err)
		return
	}
	//渲染模板
	err=t.Execute(w,"ganyu")
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
