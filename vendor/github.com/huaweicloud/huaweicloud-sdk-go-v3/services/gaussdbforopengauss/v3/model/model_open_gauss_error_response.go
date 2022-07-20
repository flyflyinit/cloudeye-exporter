package model

import (
	"encoding/json"

	"strings"
)

type OpenGaussErrorResponse struct {
	// 错误码

	ErrorCode string `json:"error_code"`
	// 错误消息。

	ErrorMsg string `json:"error_msg"`
}

func (o OpenGaussErrorResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OpenGaussErrorResponse struct{}"
	}

	return strings.Join([]string{"OpenGaussErrorResponse", string(data)}, " ")
}