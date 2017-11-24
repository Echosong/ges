package ges

import (
	"github.com/go-ini/ini"
	"log"
	"os"
	"time"
	"path/filepath"
	"strings"
)

type Helper struct {
}

func (h *Helper) InitConfig() map[string]interface{} {
	cfg, err := ini.InsensitiveLoad(h.GetCurrentDirectory() + "/config.ini");
	if err != nil {
		return nil
	}
	 config:= make( map[string] interface{})
	for _,name:= range cfg.SectionStrings(){
		sec, err := cfg.GetSection(name)
		if err != nil {
			continue
		}
		for _,keyName := range sec.KeyStrings(){
			config[name+"."+keyName] = sec.Key(keyName).Value()
		}
	}
	return config;
}

//write es log
func (h *Helper) Log(message interface{}, level string) {
	fileName := h.GetCurrentDirectory() + "/tmp/" + level + "_" + time.Now().Format("2006-01-02") + "_.log"
	logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	logInfo := log.New(logFile, "", log.LstdFlags)
	logInfo.Println(message);
}

func (h *Helper) GetCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}
	if dir == "" {
		dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//init global routes
func (h *Helper) GetRoutes(urlPath string) Routes {
	routes := Routes{}
	m := App.Config["controller.m"]
	if m == nil {
		routes.m = "web"
	}else {
		routes.m = m.(string)
	}
	c := App.Config["controller.c"]
	if c == nil {
		routes.c = "default"
	}else{
		routes.c = c.(string)
	}
	a := App.Config["controller.a"]
	if a == nil {
		routes.a = "index"
	}else {
		routes.a = a.(string)
	}
	if urlPath == "" || urlPath == "/" {
		urlPath = "/" + routes.m + "/" + routes.c + "/" + routes.a
	}
	paths := strings.Split(urlPath, "/")

	if len(paths) == 4 {
		routes.m = paths[1]
		routes.c = paths[2]
		routes.a = paths[3]
	} else if len(paths) == 3 {
		routes.c = paths[1]
		routes.a = paths[2]
	} else if len(paths) == 2 {
		routes.a = paths[1]
	}
	return routes
}
