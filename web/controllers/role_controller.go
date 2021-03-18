package controllers

import (
	"github.com/kataras/iris/v12"
	"ucenter/datamodels"
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
// @Param body body parameter.Role true "request body"
// @Success 200 {object} datamodels.Role "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles [post]
func (c *RoleController) Post() (datamodels.Role, error) {
	return datamodels.Role{}, nil
}

// @Tags 角色管理
// @Summary 更新角色信息
// @Description 更新角色信息
// @Accept json
// @Produce json
// @Param body body parameter.Role true "request body"
// @Success 200 {object} datamodels.Role "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles [put]
func (c *RoleController) Put() (datamodels.Role, error) {
	return datamodels.Role{}, nil
}

// @Tags 角色管理
// @Summary 删除角色信息
// @Description 删除角色信息
// @Accept json
// @Produce json
// @Param body body parameter.RoleDeleted true "request body"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles [delete]
func (c *RoleController) Delete() error {
	return nil
}

// @Tags 角色管理
// @Summary 设置角色禁用状态
// @Description 设置角色禁用状态
// @Accept json
// @Produce json
// @Param body body parameter.RoleDisabled true "request body"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles/disabled [put]
func (c *RoleController) PutDisabled() error {
	return nil
}

// @Tags 角色管理
// @Summary 修改登陆角色密码
// @Description 修改登陆角色密码
// @Accept json
// @Produce json
// @Param body body parameter.Password true "request body"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles/password [put]
func (c *RoleController) PutPassword() error {
	return nil
}

// @Tags 角色管理
// @Summary 获取角色信息
// @Description 通过角色ID获取角色信息
// @Accept json
// @Produce json
// @Param id query integer true "角色id"
// @Success 200 {object} datamodels.Role "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles [get]
func (c *RoleController) Get() (datamodels.Role, error) {
	return datamodels.Role{}, nil
}

// @Tags 角色管理
// @Summary 获取角色信息
// @Description 通过角色uuid获取角色信息
// @Accept json
// @Produce json
// @Param uuid query integer true "角色uuid"
// @Success 200 {object} datamodels.Role "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles/byuuid [get]
func (c *RoleController) GetByUuid() (datamodels.Role, error) {
	return datamodels.Role{}, nil
}

// @Tags 角色管理
// @Summary 获取角色分页列表
// @Description 获取全部角色分页列表
// @Accept json
// @Produce json
// @Param body body parameter.Page true "request body"
// @Success 200 {object} vo.RoleVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /roles/list [get]
func (c *RoleController) GetList() (datamodels.Role, error) {
	return datamodels.Role{}, nil
}
