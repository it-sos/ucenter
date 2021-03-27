package controllers

import (
	"github.com/kataras/iris/v12"
	"ucenter/models/vo"
)

type RoleController struct {
	//Service services.RoleService
	Ctx iris.Context
}

// @Tags 角色管理
// @Summary 创建角色信息
// @Description 创建角色信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param body body vo.RoleParamsVO true "request body"
// @Success 200 {object} vo.RoleVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles [post]
func (c *RoleController) Post() (vo.RoleVO, error) {
	return vo.RoleVO{}, nil
}

// @Tags 角色管理
// @Summary 关联用户
// @Description 用户加入到应用角色
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param body body vo.RoleAppUserVO true "request body"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles/appuser [post]
func (c *RoleController) PostAppuser() (vo.RoleVO, error) {
	return vo.RoleVO{}, nil
}

// @Tags 角色管理
// @Summary 关联应用
// @Description 角色关联应用
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param body body vo.RoleAppVO true "request body"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles/app [post]
func (c *RoleController) PostApp() (vo.RoleVO, error) {
	return vo.RoleVO{}, nil
}

// @Tags 角色管理
// @Summary 更新角色信息
// @Description 更新角色信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param id query integer true "角色id"
// @Param body body vo.RoleParamsVO true "request body"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles [put]
func (c *RoleController) Put() (vo.RoleVO, error) {
	return vo.RoleVO{}, nil
}

// @Tags 角色管理
// @Summary 删除角色信息
// @Description 删除角色信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param id query integer true "角色id"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles [delete]
func (c *RoleController) Delete() error {
	return nil
}

// @Tags 角色管理
// @Summary 删除应用信息
// @Description 删除角色与应用关联信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param role_id query integer true "角色id"
// @Param app_id query integer true "应用id"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles [delete]
func (c *RoleController) DeleteApp() error {
	return nil
}

// @Tags 角色管理
// @Summary 删除用户信息
// @Description 删除角色应用用户关联信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param role_id query integer true "角色id"
// @Param user_id query integer true "用户id"
// @Param app_id query integer true "应用id"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles [delete]
func (c *RoleController) DeleteUserApp() error {
	return nil
}

// @Tags 角色管理
// @Summary 获取角色信息
// @Description 通过角色ID获取角色信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param id query integer true "角色id"
// @Success 200 {object} vo.RoleVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles [get]
func (c *RoleController) Get() (vo.RoleVO, error) {
	return vo.RoleVO{}, nil
}

// @Tags 角色管理
// @Summary 获取角色分页列表
// @Description 获取全部角色分页列表
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param body body vo.PageVO true "request body"
// @Success 200 {object} vo.RolePageVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles/list [get]
func (c *RoleController) GetList() (vo.RolePageVO, error) {
	return vo.RolePageVO{}, nil
}

// @Tags 角色管理
// @Summary 获取用户列表
// @Description 通过角色ID和应用ID获取用户列表
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param role_id query integer true "角色id"
// @Param app_id query integer true "应用id"
// @Success 200 {object} vo.UserPageVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles/users [get]
func (c *RoleController) GetUsers() (vo.UserPageVO, error) {
	return vo.UserPageVO{}, nil
}

// @Tags 角色管理
// @Summary 获取应用角色列表
// @Description 通过应用ID获取角色列表
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param role_id query integer true "角色id"
// @Param app_id query integer true "应用id"
// @Success 200 {object} vo.RolePageVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles/listbyappid [get]
func (c *RoleController) GetListbyappid() (vo.RolePageVO, error) {
	return vo.RolePageVO{}, nil
}
