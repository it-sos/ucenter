package controllers

import (
	"github.com/kataras/iris/v12"
	"ucenter/services"
)

type UserPermissionController struct {
	Service services.UserService
	Ctx     iris.Context
}
