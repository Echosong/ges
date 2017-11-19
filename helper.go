package ges

import (
	"github.com/go-ini/ini"
	"log"
	"os"
	"time"
)

type Helper struct {

} 

/**
get es config
 */
func (h *Helper) GetConfig(section string, key string)  string {
	cfg,err := ini.InsensitiveLoad("config.ini");
	if err != nil{
		return "";
	}
	return cfg.Section(section).Key(key).Value()
}

/**
write es log
 */
func (h *Helper) Log(message string, level string)  {
	fileName := "tmp/"+ level+"_" + time.Now().Format("2006-01-02")+"_.log"
	logFile,err  := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	logInfo := log.New(logFile,"",log.LstdFlags)
	logInfo.Println(message);
}