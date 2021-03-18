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
// @Param body body vo.RoleParamsVO true "request body"
// @Success 200 {object} vo.RoleVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles [post]
func (c *RoleController) Post() (vo.RoleVO, error) {
	return vo.RoleVO{}, nil
}

// @Tags 角色管理
// @Summary 更新角色信息
// @Description 更新角色信息
// @Accept json
// @Produce json
// @Param id query integer true "角色id"
// @Param body body vo.RoleParamsVO true "request body"
// @Success 200 {object} vo.RoleVO "success"
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
// @Param id query integer true "角色id"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles [delete]
func (c *RoleController) Delete() error {
	return nil
}

// @Tags 角色管理
// @Summary 获取角色信息
// @Description 通过角色ID获取角色信息
// @Accept json
// @Produce json
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
// @Param body body vo.PageVO true "request body"
// @Success 200 {object} vo.RolePageVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles/list [get]
func (c *RoleController) GetList() (vo.RolePageVO, error) {
	return vo.RolePageVO{}, nil
}
