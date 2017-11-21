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


//get es config
func (h *Helper) GetConfig(section string, key string) string {
	cfg, err := ini.InsensitiveLoad(h.GetCurrentDirectory() + "/config.ini");
	if err != nil {
		return ""
	}
	sec, err := cfg.GetSection(section)
	if err != nil {
		return ""
	}
	if sec.HasKey(key) {
		return cfg.Section(section).Key(key).Value()
	} else {
		return ""
	}
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
	routes.m = App.Helper.GetConfig("controller", "m")
	if routes.m == "" {
		routes.m = "web"
	}
	routes.c = App.Helper.GetConfig("controller", "c")
	if routes.c == "" {
		routes.c = "default"
	}
	routes.a = App.Helper.GetConfig("controller", "a")
	if routes.a == "" {
		routes.a = "index"
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
