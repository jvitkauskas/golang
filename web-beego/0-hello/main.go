package main

import (
	_ "0-hello/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
