package errors

import (
	"encoding/json"
	"errors"
)

type Errors struct {
	ErrCode int    `json:"errCode" example:"4000000"`
	Message string `json:"message"`
}

var errCodeList = map[string]Errors{
	"param_err":         {4001001, "参数异常"},
	"login_auth_err":    {4001002, "用户名或密码错误"},
	"login_captcha_err": {4001003, "验证码不正确"},
}

func Error(key string) error {
	ret, _ := json.Marshal(errCodeList[key])
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

func (e *Errors) GetErrCode() int {
	return e.ErrCode
}

func (e *Errors) GetMessage() string {
	return e.Message
}
