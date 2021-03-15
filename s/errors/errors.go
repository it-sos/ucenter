package errors

import (
	"encoding/json"
	"errors"
)

type Errors struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

var errorResponse = map[string]Errors{
	"param_err": {4001001, "参数异常"},
}

func Error(key string) error {
	ret, _ := json.Marshal(errorResponse[key])
	return errors.New(string(ret))
}
