package main

import (
	_ "uuid/routers"
	"github.com/astaxie/beego"
	"uuid/services"
)

func main() {
	services.Init()
	beego.Run()
}


