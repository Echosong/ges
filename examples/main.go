package main

import (
	"github.com/Echosong/ges"
	"github.com/Echosong/ges/examples/src/controller/web"
)

func main()  {
	ges.Router("default", &web.DefaultController{})
	ges.Run()
}
