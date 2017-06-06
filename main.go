package main

import (
	_ "github.com/edot92/api-simple-gocoverage/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.CopyRequestBody = true
	// if beego.BConfig.RunMode == "dev" {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/docs"] = "swagger"
	// }
	beego.Run()
}
