package errors

import (
	"encoding/json"
	"errors"
)

type Error struct {
	ErrorCode    int
	ErrorMessage string
}

var errResponse = map[string]Error{
	"param_err": {
		ErrorCode:    4001001,
		ErrorMessage: "参数异常",
	},
}

func Errors(key string) error {
	ret, _ := json.Marshal(errResponse[key])
	return errors.New(string(ret))
}
