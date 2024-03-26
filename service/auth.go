package service

import (
	"context"
	"errors"
	"mqtt/help"
	"mqtt/libs"
	"time"
)

type sAuth struct {
	BaseService
}

var Auth sAuth

func (s *sAuth) SessionKey(sn string, clientKey string) (string, error) {
	ctx := context.Background()
	if len(clientKey) != 16 {
		return "", errors.New("the length of the client key must be 16 ")
	}
	serverKey := help.GenerateRandomString(16)
	sessionKey := serverKey + clientKey
	// 存入redis 24小时过期时间
	if err := libs.Redis.Set(ctx, "sessionKey:"+sn, sessionKey, 24*time.Hour).Err(); err != nil {
		return "", errors.New("set redis failed")
	}
	return sessionKey, nil
}
