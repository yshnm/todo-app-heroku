package config

import (
	"log"

	"todo-app/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port string
	SQLDriver string
	DbName string
	LogFile string
	Static string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	ctg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port: ctg.Section("web").Key("port").MustString("8080"),
		SQLDriver: ctg.Section("db").Key("driver").String(),
		DbName: ctg.Section("db").Key("name").String(),
		LogFile: ctg.Section("web").Key("logfile").String(),
		Static: ctg.Section("web").Key("static").String(),
	}
}