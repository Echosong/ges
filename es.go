package ges

import (
	"net/http"
	"strings"
	"reflect"
	"fmt"
)

type Application struct {
	Static string
	DirCurrent string
	Helper Helper
}


type Context struct {
	W http.ResponseWriter
	R *http.Request
}

var App  = &Application{}
var RouterMaps  map[string] ControllerInterface
var Cx  = &Context{}

func autoRoute(w http.ResponseWriter, r *http.Request)  {
	defer func(){
		if err:=recover();err!=nil{
			w.Write([]byte("internal error"))
			fmt.Println(err)
		}
	}()
	requestPath := r.URL.Path
	paths := strings.Split(requestPath, "/");
	rootPath := paths[1]+"/"+paths[2];
	controller := RouterMaps[rootPath]
	Cx.R = r
	Cx.W = w
	if(controller != nil){
		methodName := strings.Trim(paths[3]," ");
		rs :=[]rune(methodName)
		methodName = strings.ToUpper(string(rs[0:1]))+string(rs[1:])
		reflect.ValueOf(controller).MethodByName(methodName).Call([]reflect.Value{})
	}else{
		w.Write([]byte("404 page not found!"))
	}
}

func static(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/"):];
	http.ServeFile(w, r, file)
}

func Router(rootpath string, c ControllerInterface){
	RouterMaps = make(map[string] ControllerInterface)
	RouterMaps[rootpath] = c
}

func  Run()  {
	App.Helper = Helper{}
	App.DirCurrent = App.Helper.GetCurrentDirectory()
	App.Static = App.Helper.GetConfig("server","staticPath")
	http.HandleFunc("/"+App.Static+"/", static)
	http.HandleFunc("/" , autoRoute)
	http.ListenAndServe(App.Helper.GetConfig("server","address"),nil)
}

