package controllers

import (
	"github.com/kataras/iris/v12"
	"ucenter/models/vo"
	"ucenter/services"
)

type UserController struct {
	Service services.UserService
	Ctx     iris.Context
}

// @Tags 用户管理
// @Summary 创建用户信息
// @Description 创建用户信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param body body vo.UserParamsVO true "request body"
// @Success 200 {object} vo.UserVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /users [post]
func (c *UserController) Post() (vo.UserVO, error) {
	return vo.UserVO{}, nil
}

// @Tags 用户管理
// @Summary 更新用户信息
// @Description 更新用户信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param id query integer true "用户id"
// @Param body body vo.UserParamsVO true "request body"
// @Success 200 {object} vo.UserVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /users [put]
func (c *UserController) Put() (vo.UserVO, error) {
	return vo.UserVO{}, nil
}

// @Tags 用户管理
// @Summary 删除用户信息
// @Description 删除用户信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param id query integer true "用户id"
// @Param body body vo.UserDeletedVO true "request body"
// @Success 200 {object} vo.UserVO "success"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /users [delete]
func (c *UserController) Delete() error {
	return nil
}

// @Tags 用户管理
// @Summary 设置用户禁用状态
// @Description 设置用户禁用状态
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param id query integer true "用户id"
// @Param body body vo.UserDisabledVO true "request body"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /users/disabled [put]
func (c *UserController) PutDisabled() error {
	return nil
}

// @Tags 用户管理
// @Summary 修改登陆用户密码
// @Description 修改登陆用户密码
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param id query integer true "用户id"
// @Param body body vo.PasswordVO true "request body"
// @Success 200 {string} string	"success"
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
// @Param token header string true "token认证"
// @Param id query integer false "用户id"
// @Param uuid query string false "用户uuid"
// @Success 200 {object} vo.UserVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /users [get]
func (c *UserController) Get() (vo.UserVO, error) {
	return vo.UserVO{}, nil
}

// @Tags 用户管理
// @Summary 获取用户分页列表
// @Description 获取全部用户分页列表
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param body body vo.PageVO true "request body"
// @Success 200 {object} vo.UserPageVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /users/list [get]
func (c *UserController) GetList() (vo.UserPageVO, error) {
	return vo.UserPageVO{}, nil
}

// @Tags 用户管理
// @Summary 获取用户及关联应用角色信息
// @Description 通过用户ID获取用户及关联应用角色信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param id query integer true "用户id"
// @Success 200 {object} vo.UserAppRolesVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /users/userapprole [get]
func (c *UserController) GetUserapprole() (vo.UserAppRolesVO, error) {
	return vo.UserAppRolesVO{}, nil
}
