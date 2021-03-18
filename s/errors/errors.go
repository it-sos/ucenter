package errors

import (
	"encoding/json"
	"errors"
)

type Errors struct {
	ErrCode int    `json:"errCode"`
	Message string `json:"message"`
}

var errorResponse = map[string]Errors{
	"param_err": {4001001, "参数异常"},
}

func Error(key string) error {
	ret, _ := json.Marshal(errorResponse[key])
	return errors.New(string(ret))
}

func (e *Errors) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

func (e *Errors) SetErrCode(errCode int) {
	e.ErrCode = errCode
}

func (e *Errors) SetMessage(message string) {
	e.Message = message
}
