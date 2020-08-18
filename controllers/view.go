package controllers

import (
	"strconv"

	"github.com/SakaiTaka23/BlogSystemGo/models"

	"github.com/astaxie/beego"
)

// ViewController struct
type ViewController struct {
	beego.Controller
}

func (this *ViewController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	this.Data["Post"] = models.GetBlog(id)
	this.Layout = "layout.tpl"
	this.TplName = "view.tpl"
}
