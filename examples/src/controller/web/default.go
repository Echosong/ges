package web

import "github.com/Echosong/ges"

type DefaultController struct {
	ges.Controller
}

func (c *DefaultController) GetIndex()  {
	c.Data["hi"] = "hello world!"
	c.Display("")
}
