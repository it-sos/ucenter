package controllers

import (
	"github.com/kataras/iris/v12"
	"ucenter/models/vo"
	"ucenter/services"
)

type UserPermissionController struct {
	Service services.UserService
	Ctx     iris.Context
}

// @Tags 权限管理
// @Summary 更新用户权限
// @Description 更新用户权限
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param body body vo.UserPermissionParamsVO true "request body"
// @Success 200 {object} vo.UserPermissionVO	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /userpermissions [put]
func (c *UserPermissionController) Put() (vo.UserPermissionVO, error) {
	return vo.UserPermissionVO{}, nil
}

// @Tags 权限管理
// @Summary 获取用户权限列表
// @Description 通过用户id和应用id获取权限列表
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param appid query integer true "应用id"
// @Param uid query integer true "用户id"
// @Success 200 {object} vo.UserPermissionListVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /userpermission/listbyuidandappid [get]
func (c *UserPermissionController) GetListbyuidandappid() (vo.UserPermissionListVO, error) {
	return vo.UserPermissionListVO{}, nil
}

// @Tags 权限管理
// @Summary 获取用户已授权的权限列表
// @Description 通过用户id和应用id获取已授权的权限列表
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param appid query integer true "应用id"
// @Param uid query integer true "用户id"
// @Success 200 {object} vo.UserPermissionListVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /userpermission/allowlistbyuidandappid [get]
func (c *UserPermissionController) GetAllowlistbyuidandappid() (vo.UserPermissionListVO, error) {
	return vo.UserPermissionListVO{}, nil
}
