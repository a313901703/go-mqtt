package libs

import (
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	var err error
	user := viper.GetString("mysql.user")
	pass := viper.GetString("mysql.password")
	dsn := user + ":" + pass + "@tcp(" + viper.GetString("mysql.host") + ")/" + viper.GetString("mysql.dbname") + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Println("Connected to database")
}
