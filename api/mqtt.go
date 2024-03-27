package api

type PublishReq struct {
	Topic   string `json:"topic" validate:"required"`
	Payload string `json:"payload" validate:"required"`
	Qos     int    `json:"qos"`
	Retain  bool   `json:"retain"`
}
