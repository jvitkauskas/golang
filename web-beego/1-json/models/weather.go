package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Weather struct {
	City        string `orm:"pk"`
	Temperature float64
}

func init() {
	// Register model
	orm.RegisterModel(new(Weather))
}
