package ges

import (
	"net/http"
	"strings"
	"reflect"
)

type Application struct {
	Static     string
	DirCurrent string
	Helper     Helper
}

type Context struct {
	Response http.ResponseWriter
	Request *http.Request
}

var App = &Application{}
var routerMaps = make(map[string]ControllerInterface)
var Cx = &Context{}

func autoRoute(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.Write([]byte("internal error"))
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
		Cx.Request = r
		Cx.Response = w
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
	App.Static = App.Helper.GetConfig("server", "staticPath")
	http.HandleFunc("/"+App.Static+"/", static)
	http.HandleFunc("/", autoRoute)
	http.ListenAndServe(App.Helper.GetConfig("server", "address"), nil)
}
