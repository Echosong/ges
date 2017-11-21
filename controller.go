package ges

import (
	"html/template"
)

type ControllerInterface interface {
	Init()
	Begin()
	After()
	Display(tplName string)
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
	Routes  Routes
	Data    map[interface{}]interface{}
}

// Init generates default values of controller operations.
func (c *Controller) Init() {
	c.Layout = "Layout.html"
	c.Ctx = Cx
	c.Data = make(map[interface{}]interface{})
	c.Routes = App.Helper.GetRoutes(c.Ctx.Request.URL.Path)
	c.ViewDir = App.DirCurrent + "/src/view/" + c.Routes.m + "/"
}

// Begin runs  before request function execution.
func (c *Controller) Begin() {

}

// After runs after request function execution.
func (c *Controller) After() {

}

//Display sends the response with rendered template bytes as text/html type.
func (c *Controller) Display(tplName string) {
	if tplName == "" {
		tplName = c.Routes.a+".html"
	}
	//t := template.Must(template.New(c.Routes.a).ParseFiles(c.ViewDir + c.Layout))
	t := template.Must(template.ParseFiles(c.ViewDir + tplName))
	t.Execute(c.Ctx.Response, c.Data)
}