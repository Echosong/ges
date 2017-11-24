package ges

import (
	"net/http"
	"strings"
	"reflect"
	"fmt"
)

type Application struct {
	DirCurrent string
	Helper     Helper
	Config  map[string] interface{}
	Context *Context
}

type Context struct {
	Response http.ResponseWriter
	Request *http.Request
}

var App = &Application{}
var routerMaps = make(map[string]ControllerInterface)

func autoRoute(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(w, "action not fund or Error:[ %S ]",err)
			App.Helper.Log(err, "error")
		}
	}()

	routes := App.Helper.GetRoutes(r.URL.Path)
	var controller ControllerInterface
	for r,v := range routerMaps{
		if strings.ToLower(r) == strings.ToLower(routes.c) || strings.ToLower(r) == strings.ToLower(routes.m+"/"+routes.c) {
			controller = v
			break;
		}
	}

	if controller != nil {
		methodName := routes.a;
		rs := []rune(methodName)
		methodName = strings.ToUpper(string(rs[0:1])) + string(rs[1:])
		if r.Method == "GET" {
			methodName = "Get"+methodName
		}else if r.Method == "POST"{
			methodName = "Post"+methodName
		}
		cxt := &Context{}
		cxt.Request = r
		cxt.Response = w
		App.Context = cxt
		instance := reflect.ValueOf(controller)
		instance.MethodByName("Init").Call([]reflect.Value{})
		instance.MethodByName("Begin").Call([]reflect.Value{})
		instance.MethodByName(methodName).Call([]reflect.Value{})
		instance.MethodByName("After").Call([]reflect.Value{})
	} else {
		w.Write([]byte("404 page not found!"))
	}
}

func static(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/"):];
	http.ServeFile(w, r, file)
}

func Router(rootPath string, c ControllerInterface) {
	routerMaps[rootPath] = c
}

func Run() {
	App.Helper = Helper{}
	App.DirCurrent = App.Helper.GetCurrentDirectory()
	App.Config = App.Helper.InitConfig()
	var staticPath = App.Config["server.staticPath"]
	if staticPath == nil{
		staticPath = "res"
	}
	http.HandleFunc("/"+staticPath.(string)+"/", static)
	http.HandleFunc("/", autoRoute)
	http.ListenAndServe(App.Config["server.address"].(string), nil)
}
