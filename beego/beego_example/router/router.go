package router

import (
	"beego_example/controller/indexController"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &indexController.IndexController{}, "*:Index")
}
