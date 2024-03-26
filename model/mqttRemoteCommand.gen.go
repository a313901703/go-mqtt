package model

import (
	"mqtt/libs"
	"time"
)

// MqttRemoteCommand undefined
type MqttRemoteCommand struct {
	ID           int64     `json:"id" gorm:"id"`
	Uln          string    `json:"uln" gorm:"uln"`
	RoomId       int64     `json:"room_id" gorm:"room_id"`
	Sn           string    `json:"sn" gorm:"sn"`         // SN #
	Action       string    `json:"action" gorm:"action"` // remote action name
	Comments     string    `json:"comments" gorm:"comments"`
	Payload      string    `json:"payload" gorm:"payload"`
	SendStatus   int8      `json:"send_status" gorm:"send_status"` // 1:success  0:false
	Error        string    `json:"error" gorm:"error"`             // mqtt err message
	MqttId       string    `json:"mqtt_id" gorm:"mqtt_id"`
	ExecStatus   int8      `json:"exec_status" gorm:"exec_status"` // The execution status of commands. 0: wati process  1:processing  2:success 3:failed
	IssueTime    time.Time `json:"issue_time" gorm:"issue_time"`   // Command execution time
	CreatedBy    int64     `json:"created_by" gorm:"created_by"`
	CreatedName  string    `json:"created_name" gorm:"created_name"`
	CreatedEmail string    `json:"created_email" gorm:"created_email"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
}

// TableName 表名称
func (m *MqttRemoteCommand) TableName() string {
	return "mqtt_remote_command"
}

func (m MqttRemoteCommand) Update(cond map[string]interface{}) (int64, error) {
	ret := libs.DB.Model(&m).Updates(cond)
	if ret.Error != nil {
		return 0, ret.Error
	}
	return ret.RowsAffected, nil
}

func (MqttRemoteCommand) One(id int) (MqttRemoteCommand, error) {
	var m MqttRemoteCommand
	ret := libs.DB.First(&m, id)
	if ret.Error != nil {
		return m, ret.Error
	}
	return m, nil
}
