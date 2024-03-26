package controller

import (
	"mqtt/api"
	"mqtt/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type cAuth struct {
	BaseController
}

var AuthController = cAuth{}

func (auth *cAuth) GetSessionKey(c echo.Context) (err error) {
	req := new(api.SessionKeyReq)
	if err := c.Bind(req); err != nil {
		log.Error(err)
	}
	if err = auth.Validate(req); err != nil {
		return RespError(c, http.StatusBadRequest, err.Error())
	}

	// 验证通过，生成sessionKey
	var key string
	if key, err = service.Auth.SessionKey(req.Sn, req.Rand); err != nil {
		return RespError(c, http.StatusBadRequest, err.Error())
	}
	return RespSuccess(c, key)
}
