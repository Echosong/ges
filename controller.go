package ges

import (
	"strings"
	"html/template"
	"os"
)

type ControllerInterface interface {
	Init()
	Begin()
	After()
	Display()
}

type Routes struct {
	m string
	c string
	a string
}

type Controller struct {
	Ctx     *Context
	ViewDir string
	Layout  string
	Routes  *Routes
	Data    map[interface{}]interface{}
}

// Init generates default values of controller operations.
func (c *Controller) Init() {
	c.Layout = "Layout.html"
	c.Ctx = Cx
	//init m,c,a
	urlPath := Cx.Request.URL.Path
	paths := strings.Split(urlPath, "/")
	c.Routes.m = App.Helper.GetConfig("controller", "m")
	if c.Routes.m == "" {
		c.Routes.m = "web"
	}
	c.Routes.c = App.Helper.GetConfig("controller", "c")
	if c.Routes.c == "" {
		c.Routes.c = "default"
	}
	if len(paths) == 4 {
		c.Routes.m = paths[1];
		c.Routes.c = paths[2];
		c.Routes.a = paths[3]
	} else if len(paths) == 3 {
		c.Routes.c = paths[1];
		c.Routes.a = paths[2]
	} else if len(paths) == 2 {
		c.Routes.a = paths[1]
	}
	c.ViewDir = App.DirCurrent + "/src/views/" + c.Routes.m + "/"
}

// Begin runs  before request function execution.
func (c *Controller) Begin() {

}

// After runs after request function execution.
func (c *Controller) After() {

}

//Display sends the response with rendered template bytes as text/html type.
func (c *Controller) Display(tplName string) {
	t := template.Must(template.New(c.Routes.a).ParseFiles(c.ViewDir + c.Layout))
	t = template.Must(t.ParseFiles(c.ViewDir + tplName))
	t.Execute(os.Stdout, c.Data)
}
