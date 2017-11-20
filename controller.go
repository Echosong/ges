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

}


// Begin runs  before request function execution.
func (c *Controller) Begin()  {

}

// After runs after request function execution.
func (c *Controller) After() {

}

//Display sends the response with rendered template bytes as text/html type.
func (c *Controller) Display()  {

}

