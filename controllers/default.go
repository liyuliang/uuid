package controllers

import (
	"github.com/astaxie/beego"
	"github.com/liyuliang/utils"
	"uuid/services"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	num, _ := c.GetInt("n")

	if num > 0 && num <= 100 {

		var ids []string
		for i := 0; i < num; i++ {
			id, err := services.Get()
			if err == nil {

				ids = append(ids, format.Int64ToStr(id))
			}
		}
		c.Ctx.WriteString(format.StructToStr(ids))

	} else {

		id, err := services.Get()
		if err != nil {

			c.Ctx.WriteString(format.ErrorToStr(err))
		} else {
			c.Ctx.WriteString(format.Int64ToStr(id))
		}
	}
}

