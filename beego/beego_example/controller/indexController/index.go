package indexController

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (p *IndexController) Index() {
	p.TplName = "index/index.html"
	m := make(map[string]interface{})
	m["code"] = 200
	m["status"] = "success"
	p.Data["json"] = m

	//p.ServeJSON()
}
