package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"mqtt/help"
	"mqtt/libs"
	"mqtt/model"
	"runtime/debug"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

type sRemoteCommand struct {
}

var RemoteCommandService sRemoteCommand

type commandResp struct {
	Sn  string // 设备唯一标识
	Dat interface{}
}

type commandData struct {
	Seq int
	Sta int
	Cmd int
}

// {"sn":"123123131","dat":{"seq":"5","sta":1,"cmd":"102"}}
func (s *sRemoteCommand) ReciveResp() (respHandler mqtt.MessageHandler) {

	respHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MSG: %s\n", msg.Payload())

		defer func() {
			if err := recover(); err != nil {
				s := string(debug.Stack())
				help.PanicLog(fmt.Sprintf("ReciveResp.respHandler panic: %v, payload: %s, stack=%s\n", err, msg.Payload(), s))
			}
		}()
		var err error
		var payload commandResp

		if err = json.Unmarshal([]byte(msg.Payload()), &payload); err != nil {
			help.ErrorLog(fmt.Sprintf("json.Unmarshal is err:%v", err))
			return
		}
		help.WriteLog(fmt.Sprintf("recive mqtt resp data %+v", payload), "mqtt_recive")
		// log.Infof("deal mqtt recive data %+v", payload)
		var commandModel model.MqttRemoteCommand
		var commandData commandData
		var sn string
		if viper.GetBool("debug") {
			commandData.Seq = int(payload.Dat.(map[string]interface{})["seq"].(float64))
			commandData.Sta = int(payload.Dat.(map[string]interface{})["sta"].(float64))
			commandData.Cmd = int(payload.Dat.(map[string]interface{})["cmd"].(float64))
		} else {
			var bodyByte []byte
			// 解密数据
			if bodyByte, err = dectyptData(payload.Dat.(string), payload.Sn); err != nil {
				help.ErrorLog(err.Error())
				return
			}
			if err = json.Unmarshal(bodyByte, &commandData); err != nil {
				help.ErrorLog(fmt.Sprintf("json.Unmarshal is err, sn:%s, data:%s", sn, string(bodyByte)))
				return
			}
		}
		cModel, err := commandModel.One(commandData.Seq)
		if err != nil {
			help.ErrorLog(fmt.Sprintf("command not found,command id: %d", commandData.Seq))
			return
		}
		fmt.Println(cModel)
		// TODO 具体的业务

	}
	return respHandler
}

func dectyptData(dataStr string, sn string) (bodyByte []byte, err error) {
	bodyByte, _ = base64.StdEncoding.DecodeString(dataStr)
	ctx := context.Background()
	var sessionKey string
	if sessionKey, err = libs.Redis.Get(ctx, "sessionKey:"+sn).Result(); err != nil {
		return nil, errors.New("get session key failed, sn: " + sn)
	}
	bodyByte, err = libs.TdesDescrypt([]byte(sessionKey), bodyByte)
	if err != nil {
		return nil, errors.New("decrypt paload failed: " + sn)
	}

	return
}
