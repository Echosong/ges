package ges

type Routs struct {
	m string
	c string
	a string
}

type ControllerInterface interface{
	Begin()
	After()
	Display()
}

type Controller struct {
	Ctx  *Context
}


// Begin runs  before request function execution.
func (c *Controller) Begin()  {
	c.Ctx = Cx;
}

// After runs after request function execution.
func (c *Controller) After() {

}

//Display sends the response with rendered template bytes as text/html type.
func (c *Controller) Display()  {

}

