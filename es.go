package ges

import (
	"net/http"
	"reflect"
)

type Application struct {
	Static string
	Helper Helper
}

var App *Application

func autoRoute(w http.ResponseWriter, r *http.Request)  {
	requestPath := r.URL.Path
	w.Write([]byte(requestPath))
}

func static(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/"):];
	http.ServeFile(w, r, file)
}

func  Run()  {
	App := &Application{}
	App.Helper = Helper{}
	App.Static = App.Helper.GetConfig("server","static")
	http.HandleFunc("/static/", static)
	http.HandleFunc("/" , autoRoute)
	http.ListenAndServe(App.Helper.GetConfig("server","address"),nil)
}

