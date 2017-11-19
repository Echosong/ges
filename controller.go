package ges

type Routs struct {
	m string
	c string
	a string
}

type Controller struct {
	Layout string
	Routs *Routs
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

