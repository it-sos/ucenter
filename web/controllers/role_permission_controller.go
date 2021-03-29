package controllers

import (
	"github.com/kataras/iris/v12"
	"ucenter/models/vo"
)

type RolePermissionController struct {
	Ctx iris.Context
}

// @Tags 权限管理
// @Summary 更新角色权限
// @Description 更新角色权限
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param body body vo.RolePermissionParamsVO true "request body"
// @Success 200 {object} vo.RolePermissionVO	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /rolepermissions [put]
func (c *RolePermissionController) Put() (vo.RolePermissionVO, error) {
	return vo.RolePermissionVO{}, nil
}

// @Tags 权限管理
// @Summary 获取角色权限列表
// @Description 通过角色id和应用id获取权限列表
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param appid query integer true "应用id"
// @Param uid query integer true "角色id"
// @Success 200 {object} vo.RolePermissionListVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /rolepermission/listbyuidandappid [get]
func (c *RolePermissionController) GetListbyuidandappid() (vo.RolePermissionListVO, error) {
	return vo.RolePermissionListVO{}, nil
}

// @Tags 权限管理
// @Summary 获取角色已授权的权限列表
// @Description 通过角色id和应用id获取已授权的权限列表
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param appid query integer true "应用id"
// @Param uid query integer true "角色id"
// @Success 200 {object} vo.RolePermissionListVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /rolepermission/allowlistbyuidandappid [get]
func (c *RolePermissionController) GetAllowlistbyuidandappid() (vo.RolePermissionListVO, error) {
	return vo.RolePermissionListVO{}, nil
}
