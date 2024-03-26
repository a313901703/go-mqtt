package model

import (
	"mqtt/libs"
	"time"
)

type MqttSessionKey struct {
	Id         uint   `gorm:"primarykey"`
	Sn         string `json:"sn"`
	SessionKey string `json:"session_key"` // 设备会话密钥
	VendorId   uint   `json:"vendor_id"`
	CreatedAt  time.Time
}

// TableName 表名称
func (m *MqttSessionKey) TableName() string {
	return "mqtt_session_key"
}

func (MqttSessionKey) One(sn string) (MqttSessionKey, error) {
	var data MqttSessionKey
	ret := libs.DB.Where("sn = ?", sn).First(&data)
	if ret.Error != nil {
		return data, ret.Error
	}
	return data, nil
}

func (m MqttSessionKey) FindOrCreate(sn string) {
	libs.DB.FirstOrInit(&m, MqttSessionKey{Sn: sn})
}

func (m MqttSessionKey) Save() (err error) {
	if ret := libs.DB.Save(&m); ret.Error != nil {
		return ret.Error
	}
	return nil
}
