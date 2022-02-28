package main

import (
	"fmt"
	"html/template"
	"net/http"
)
func tempBC(w http.ResponseWriter,r *http.Request){
	t,err:=template.New("index.tmpl").Delims("{[","]}").ParseFiles("./index.tmpl")
	if err!=nil{
		fmt.Printf("template new parsefile failed ,err:%v",err)
	}
	name:="神里"
	err=t.Execute(w,name)
	if err != nil {
		fmt.Printf("template execute failed ,err:%v",err)
	}
}
func xss(w http.ResponseWriter,r *http.Request){
	// 使用 funcs自定义函数 返回js转换后的html值
	t,err:=template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string)template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.tmpl")
	if err!=nil{
		fmt.Printf("template new parsefile failed ,err:%v",err)
	}
	str1:="<script>alert(123)</script>"
	str2:="<a href='http://www.baidu.com'>百度</a>"
	err=t.Execute(w,map[string]string{
		"str1":str1,
		"str2":str2,
	})
	if err != nil {
		fmt.Printf("template execute failed ,err:%v",err)
	}
}

func main() {
	http.HandleFunc("/tempBC",tempBC)
	http.HandleFunc("/xss",xss)
	err:=http.ListenAndServe(":8088",nil)
	if err!=nil{
		fmt.Printf("listenAndServe start failed ,err:%v",err)
	}
}
