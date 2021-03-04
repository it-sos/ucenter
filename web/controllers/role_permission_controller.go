package controllers

import (
	"github.com/kataras/iris/v12"
	"ucenter/services"
)

type RolePermissionController struct {
	Service services.UserService
	Ctx     iris.Context
}
