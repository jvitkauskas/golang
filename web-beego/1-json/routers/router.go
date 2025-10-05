package routers

import (
	"1-json/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/weather/:city", &controllers.MainController{})
}
