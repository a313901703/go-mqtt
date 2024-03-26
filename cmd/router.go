package cmd

import (
	"encoding/json"
	"fmt"
	"mqtt/controller"
	"mqtt/help"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func setRouters(e *echo.Echo) {
	middlewares(e)
	e.POST("mqtt/publish", controller.MqttController.Publish)
	agroup := e.Group("/auth")
	agroup.POST("/session-key", controller.AuthController.GetSessionKey)
	routers, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(routers))
}

type ErrResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func middlewares(e *echo.Echo) {
	// CORS
	e.Use(middleware.CORS())
	// Rate Limiter  每秒限制 1000
	reteLimit(e)
	e.Use(middleware.Decompress())
	e.Use(middleware.RequestID())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:       1 << 10, // 1 KB
		DisableStackAll: true,
		LogErrorFunc:    recoverF,
	}))

	if viper.GetBool("debug") {
		// dump
		e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			fmt.Println("BodyDump", string(reqBody), string(resBody))
		}))

	}
	f, err := os.OpenFile("./logs/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening file: %v", err))
	}
	defer f.Close()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: f,
	}))
}

func reteLimit(e *echo.Echo) {
	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 1000},
		),
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}

	e.Use(middleware.RateLimiterWithConfig(config))
}

func recoverF(c echo.Context, err error, stack []byte) error {
	msg := fmt.Sprintf("[PANIC RECOVER] %v %s\n", err, stack)
	c.Logger().Error(msg)
	help.PanicLog(msg)
	return err
}
