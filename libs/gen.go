package libs

import (
	"fmt"
	"mqtt/help"

	"gorm.io/gen"
	"gorm.io/gorm"
)

func GenModel(tableName string, gormdb *gorm.DB) {
	fmt.Println("start gen model", tableName)

	g := gen.NewGenerator(gen.Config{
		OutPath: "./model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions

	g.ApplyBasic(
		g.GenerateModelAs(tableName, help.ToCamelCase(tableName)),
	)
	// g.ApplyBasic(
	// 	// Generate structs from all tables of current database
	// 	g.GenerateAllTable()...,
	// )
	// Generate the code
	g.Execute()
}
