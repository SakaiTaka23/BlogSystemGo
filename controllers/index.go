package controllers

import (
	"github.com/SakaiTaka23/BlogSystemGo/models"

	"github.com/astaxie/beego"
)

// IndexController struct
type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Data["blogs"] = models.GetAll()
	this.Layout = "layout.tpl"
	this.TplName = "index.tpl"
}
