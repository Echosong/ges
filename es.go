package ges

import (
	"net/http"
	"strings"
	"reflect"
)

type Application struct {
	Static string
	Helper Helper
}

type ControllerInterface interface{
	Begin()
	After()
	Display()
}

var App *Application
var RouterMaps  map[string] *ControllerInterface

func autoRoute(w http.ResponseWriter, r *http.Request)  {
	requestPath := r.URL.Path
	paths := strings.Split(requestPath, "/");
	rootPath := paths[0]+"/"+paths[1];
	controller :=RouterMaps[rootPath]
	if(controller != nil){
		//reflect.TypeOf();
		reflect.ValueOf(controller).MethodByName(paths[2]).Call([]reflect.Value{})
	}else{
		http.NotFoundHandler()
	}
}

func static(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/"):];
	http.ServeFile(w, r, file)
}

func Router(rootpath string, c *ControllerInterface){
	RouterMaps = make(map[string] *ControllerInterface)
	RouterMaps[rootpath] = c
}

func  Run()  {
	App := &Application{}
	App.Helper = Helper{}
	App.Static = App.Helper.GetConfig("server","static")
	http.HandleFunc("/"+App.Static+"/", static)
	http.HandleFunc("/" , autoRoute)
	http.ListenAndServe(App.Helper.GetConfig("server","address"),nil)
}

