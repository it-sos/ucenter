package errors

import (
	"encoding/json"
	"errors"
)

type Error struct {
	errorCode    int
	errorMessage string
}

var errorResponse = map[string]Error{
	"param_err": {4001001, "参数异常"},
}

func Errors(key string) error {
	ret, _ := json.Marshal(errorResponse[key])
	return errors.New(string(ret))
}
