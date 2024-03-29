package controllers

import (
	"github.com/SakaiTaka23/BlogSystemGo/models"
	"strconv"

	"github.com/astaxie/beego"
)

// DeleteController struct
type DeleteController struct {
	beego.Controller
}

func (this *DeleteController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	blog := models.GetBlog(id)
	this.Data["Post"] = blog
	models.DelBlog(blog)
	this.Ctx.Redirect(302, "/")
}
