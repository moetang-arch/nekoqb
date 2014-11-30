package main

import (
	_ "github.com/moetang-arch/nekoqb/docs"
	_ "github.com/moetang-arch/nekoqb/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
