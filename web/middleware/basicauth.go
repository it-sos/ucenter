package middleware

import "github.com/kataras/iris/v12/middleware/basicauth"

var BasicAuth = basicauth.Default(map[string]string{
	"feiyu": "11118888",
})
