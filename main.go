package main

import (
	"github.com/SakaiTaka23/BlogSystemGo/controllers"

	"github.com/astaxie/beego"
)

func main() {
	//ブログのトップページを表示
	beego.Router("/", &controllers.IndexController{})
	//ブログの詳細な情報を検索
	beego.Router("/view/:id([0-9]+)", &controllers.ViewController{})
	//ブログの文章を作成
	beego.Router("/new", &controllers.NewController{})
	//ブログの削除
	beego.Router("/delete/:id([0-9]+)", &controllers.DeleteController{})
	//ブログの編集
	beego.Router("/edit/:id([0-9]+)", &controllers.EditController{})

	beego.Run(":3000")
}
