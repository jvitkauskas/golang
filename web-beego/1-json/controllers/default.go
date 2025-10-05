package controllers

import (
	"1-json/models"
	"1-json/types"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	city := c.Ctx.Input.Param(":city")

	o := orm.NewOrm()
	weather := models.Weather{City: city}
	err := o.Read(&weather)
	if err == orm.ErrNoRows {
		c.CustomAbort(400, "No such city")
		return
	} else if err != nil {
		c.CustomAbort(500, "Error reading data")
		return
	}

	c.Data["json"] = weather
	c.ServeJSON()
}

func (c *MainController) Put() {
	saveWeather(c, true)
}

func (c *MainController) Post() {
	saveWeather(c, false)
}

func saveWeather(c *MainController, checkExisting bool) {
	var req types.WeatherUpdateRequest
	c.BindJSON(&req)

	if req.Temperature == nil {
		c.CustomAbort(400, "Temperature not provided")
		return
	}

	city := c.Ctx.Input.Param(":city")

	o := orm.NewOrm()
	weather := models.Weather{City: city}

	if checkExisting {
		err := o.Read(&weather)
		if err == orm.ErrNoRows {
			c.CustomAbort(400, "No such city")
			return
		} else if err != nil {
			c.CustomAbort(500, "Error reading data")
			return
		}
	}

	weather.Temperature = *req.Temperature

	o.Insert(&weather) // InsertOrUpdate is not supported for sqlite, so just try inserting
	o.Update(&weather)

	c.Data["json"] = weather
	c.ServeJSON()
}
