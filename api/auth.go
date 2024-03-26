package api

type (
	// session key request
	SessionKeyReq struct {
		Sn   string `json:"sn" validate:"required"`
		Rand string `json:"rand" validate:"required"`
	}

	SessionKeyResp struct {
		Sta int    `json:"sta"`
		Msg string `json:"msg"`
		Dat string `json:"data"`
	}
)
