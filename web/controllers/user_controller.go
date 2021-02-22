package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"strconv"
	"ucenter/services"
)

type UserController struct {
	Service services.UserService
	Ctx     iris.Context
}

/**
 * 提交打卡
 */
func (c *UserController) PostBy(id int64) mvc.Result {
	var (
		morning, _ = strconv.ParseUint(c.Ctx.FormValue("morning"), 10, 8)
		noon, _    = strconv.ParseUint(c.Ctx.FormValue("noon"), 10, 8)
		night, _   = strconv.ParseUint(c.Ctx.FormValue("night"), 10, 8)
	)

	c.Service.UpdatePosterAndGenreByID(id, uint8(morning), uint8(noon), uint8(night))
	return mvc.Response{
		//Err: err,
		Path: "/",
	}
}
