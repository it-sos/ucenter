package controllers

import "github.com/kataras/iris/v12"

func Login(ctx iris.Context) {
	ctx.JSON(iris.Map{"code":"200", "message":"hello world."})
}
