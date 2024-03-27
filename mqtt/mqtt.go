package mqtt

import (
	"fmt"
	"mqtt/service"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

var MqttClient mqtt.Client

func Start() {
	opts := mqtt.NewClientOptions().AddBroker(viper.GetString("mqtt.host")).SetClientID("golang-mqtt-12345")

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)
	opts.AutoReconnect = true

	opts.OnConnect = func(client mqtt.Client) {
		println("连接成功")
		subscribe(client)
	}
	opts.Username = viper.GetString("mqtt.username")
	opts.Password = viper.GetString("mqtt.password")

	MqttClient = mqtt.NewClient(opts)
	if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

}

func subscribe(c mqtt.Client) {
	// 订阅主题
	if token := c.Subscribe("resp", 0, service.RemoteCommandService.ReciveResp()); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error(), "Subscribe Error")
	}
}
