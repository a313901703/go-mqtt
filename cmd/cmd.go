package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func Start() {
	e := echo.New()

	//加载组件
	registerCore()
	commands()
	register()
	setRouters(e)
	port := viper.GetString("serverPort")
	s := http.Server{
		Addr:    ":" + port,
		Handler: e,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}
	fmt.Println("Server is running on port: " + port)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
