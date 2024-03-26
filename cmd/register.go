package cmd

import (
	"fmt"
	"log"
	"mqtt/libs"
	"mqtt/mqtt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func register() {
	mqtt.Start()
}

func registerCore() {
	initConfig()
	libs.InitMysql()
	libs.InitRedis()
}

func initConfig() {
	viper.SetConfigFile("config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig()failed,err:%v\n", err)
		return
	}
	//监听修改
	viper.WatchConfig()
	//为配置修改增加一个回调函数
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
	})
}

func commands() {
	if len(os.Args) > 2 && os.Args[1] == "run" {
		// 执行脚本
		switch os.Args[2] {
		case "gen:model":
			var db *gorm.DB
			if len(os.Args) < 4 {
				log.Fatal("not match params for gen")
			}
			if len(os.Args) < 5 {
				db = libs.DB
			} else {
				db = libs.DB
			}

			libs.GenModel(os.Args[3], db)
		default:
			log.Fatal("unknown command:", os.Args[2])
		}
	}
}
