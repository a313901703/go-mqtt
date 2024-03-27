package api

type Resp struct {
	Code    int         `json:"code" xml:"code"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
}

type RespErr struct {
	Code    int    `json:"code" xml:"code"`
	Message string `json:"message" xml:"message"`
}
