package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	name := c.GetString("name") // query param
	if name == "" {
		name = "World"
	}

	c.Ctx.WriteString("Hello " + name)
}

func (c *MainController) GetWithName() {
	name := c.Ctx.Input.Param(":name") // path param
	if name == "" {
		name = "World"
	}

	c.Ctx.WriteString("Hello " + name)
}
