package main

import (
	"fmt"
	"html/template"
	"net/http"
)
func index(w http.ResponseWriter,r *http.Request){
	t,err:=template.ParseFiles("./index.tmpl")
	if err!=nil{
		fmt.Printf("parse file failed, err:%v",err)
		return
	}
	msg:="shenli"
	t.Execute(w,msg)
}
func home(w http.ResponseWriter,r *http.Request){
	t,err:=template.ParseFiles("./home.tmpl")
	if err!=nil{
		fmt.Printf("parse file failed, err:%v",err)
		return
	}
	msg:="ganyu"
	t.Execute(w,msg)
}
func index2(w http.ResponseWriter,r *http.Request){
	t,err:=template.ParseFiles("./template/base.tmpl","./template/index2.tmpl")
	if err!=nil{
		fmt.Printf("parse file failed, err:%v",err)
		return
	}
	msg:="shenli"
	t.ExecuteTemplate(w,"index2.tmpl",msg)
}
func home2(w http.ResponseWriter,r *http.Request){
	t,err:=template.ParseFiles("./template/base.tmpl","./template/home2.tmpl")
	if err!=nil{
		fmt.Printf("parse file failed, err:%v",err)
		return
	}
	msg:="ganyu"
	t.ExecuteTemplate(w,"home2.tmpl",msg)
}
func main() {
	http.HandleFunc("/index",index)
	http.HandleFunc("/home",home)
	http.HandleFunc("/index2",index2)
	http.HandleFunc("/home2",home2)
	err:=http.ListenAndServe(":8000",nil)
	if err!=nil{
		fmt.Println("http service start failed ,err：%v ",err)
		return
	}

}
