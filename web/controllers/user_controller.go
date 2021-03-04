package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"ucenter/services"
)

type UserController struct {
	Service services.UserService
	Ctx     iris.Context
}

// @Summary 创建用户
// @Description 创建用户
// @Accept  json
// @Produce  json
// @Param account      query string    true		"账号"
// @Param password     query string    true		"密码"
// @Param nickname     query string    true     "昵称"
// @Param phone        query string    false    "手机号"
// @Param expired      query string    false 	"有效期，0=永久/指定过期时间，默认:否"
// @Param is_disabled  query uint8     false 	"是否禁用状态1=是;0=否，默认:否"
// @Success 200 {string} string	"ok"
// @Router /user [post]
func (c *UserController) Post() mvc.Result {
	return mvc.Response{
		//Err: err,
		Path: "/",
	}
}
