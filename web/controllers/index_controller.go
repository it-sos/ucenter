package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"ucenter/services"
)

type IndexController struct {
	Service services.UserService
}

/**
 * 获取打卡记录
 */
func (c *IndexController) Get() mvc.Result {
	data := c.Service.GetByDate()
	return mvc.View{
		Name: "index.html",
		Data: data,
	}
}
