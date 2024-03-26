package controller

import (
	"mqtt/api"
	"mqtt/mqtt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type cMqtt struct {
	BaseController
}

var MqttController = cMqtt{}

func (c *cMqtt) Publish(ctx echo.Context) (err error) {
	req := new(api.PublishReq)
	if err := ctx.Bind(req); err != nil {
		log.Error(err)
	}
	if err = c.Validate(req); err != nil {
		return RespError(ctx, http.StatusBadRequest, err.Error())
	}

	mqtt.MqttClient.Publish(req.Topic, byte(req.Qos), req.Retain, req.Payload)
	return RespSuccess(ctx, "")
}
