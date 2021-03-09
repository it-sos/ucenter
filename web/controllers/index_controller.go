package controllers

import (
	"errors"
	"ucenter/datamodels"
	"ucenter/services"
)

type IndexController struct {
	Service services.UserService
}

/**
 * 获取打卡记录
 */
func (c *IndexController) Get() (datamodels.App, error) {
	//data := c.Service.GetByDate()
	return datamodels.App{}, errors.New("hello is error.")
}
