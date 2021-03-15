package controllers

import (
	"github.com/kataras/iris/v12"
	"ucenter/datamodels"
	"ucenter/s/errors"
	"ucenter/services"
)

type UserController struct {
	Service services.UserService
	Ctx     iris.Context
}

// @Tags 用户管理
// @Summary 创建用户
// @Description 创建用户
// @Accept json
// @Produce json
// @Param body body parameter.User true "request body"
// @Success 200 {object} datamodels.User "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /users [post]
func (c *UserController) Post() (datamodels.User, error) {
	return datamodels.User{}, nil
}

// @Tags 用户管理
// @Summary 更新用户
// @Description 更新用户
// @Accept json
// @Produce json
// @Param body body parameter.User true "body"
// @Success 200 {object} datamodels.User "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /users [put]
func (c *UserController) Put() (datamodels.User, error) {
	return datamodels.User{}, nil
}

// @Tags 用户管理
// @Summary 删除用户
// @Description 删除用户
// @Accept json
// @Produce json
// @Param body body parameter.Deleted true "request body"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /users [delete]
func (c *UserController) Delete() error {
	return errors.Error("param_err")
}

// @Tags 用户管理
// @Summary 设置用户禁用状态
// @Description 设置用户禁用状态
// @Accept json
// @Produce json
// @Param body body parameter.Disabled true "request body"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /users/disabled [put]
func (c *UserController) PutDisabled() error {
	return nil
}

// @Tags 用户管理
// @Summary 修改当前用户密码
// @Description 修改当前用户密码
// @Accept json
// @Produce json
// @Param body body parameter.Password true "request body"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /users/password [put]
func (c *UserController) PutPassword() error {
	return nil
}

// @Tags 用户管理
// @Summary 获取用户信息
// @Description 通过用户ID获取用户信息
// @Accept json
// @Produce json
// @Param id query integer true "用户id"
// @Success 200 {object} datamodels.User "ok"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /users [get]
func (c *UserController) Get() datamodels.User {
	return datamodels.User{}
}

// @Tags 用户管理
// @Summary 获取用户信息
// @Description 通过用户ID获取用户信息
// @Accept json
// @Produce json
// @Param uuid query integer true "用户uuid"
// @Success 200 {object} datamodels.User "ok"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /users/uuid [get]
func (c *UserController) GetUuid() error {
	return errors.Error("param_err")
}
