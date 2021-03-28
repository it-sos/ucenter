package controllers

import (
	"github.com/kataras/iris/v12"
	"ucenter/models/vo"
	"ucenter/services"
)

type AppController struct {
	Service services.UserService
	Ctx     iris.Context
}

// @Tags 应用管理
// @Summary 创建应用信息
// @Description 创建应用信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param body body vo.AppParamsVO true "request body"
// @Success 200 {object} vo.AppVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /apps [post]
func (c *AppController) Post() (vo.AppVO, error) {
	return vo.AppVO{}, nil
}

// @Tags 应用管理
// @Summary 上传图标文件
// @Description 上传图标文件
// @Accept  multipart/form-data
// @Produce json
// @Param token header string true "token认证"
// @Param id query integer true "应用id"
// @Param file formData file true "request file data"
// @Success 200 {object} vo.AppIconVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /apps/icon [post]
func (c *AppController) PostIcon() (vo.AppIconVO, error) {
	return vo.AppIconVO{}, nil
}

// @Tags 应用管理
// @Summary 更新应用信息
// @Description 更新应用信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param id query integer true "应用id"
// @Param body body vo.AppParamsVO true "request body"
// @Success 200 {object} vo.AppVO	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /apps [put]
func (c *AppController) Put() (vo.AppVO, error) {
	return vo.AppVO{}, nil
}

// @Tags 应用管理
// @Summary 删除应用信息
// @Description 删除应用信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param id query integer true "应用id"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /apps [delete]
func (c *AppController) Delete() error {
	return nil
}

// @Tags 应用管理
// @Summary 获取应用信息
// @Description 通过应用ID获取应用信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param icon query string true "图标地址"
// @Success 200 {string} string "图片流"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /apps/icon [get]
func (c *AppController) GetIcon() error {
	return nil
}

// @Tags 应用管理
// @Summary 获取应用信息
// @Description 通过应用ID获取应用信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param id query integer true "应用id"
// @Success 200 {object} vo.AppVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /apps [get]
func (c *AppController) Get() (vo.AppVO, error) {
	return vo.AppVO{}, nil
}

// @Tags 应用管理
// @Summary 获取应用分页列表
// @Description 获取全部应用分页列表
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param body body vo.PageVO true "request body"
// @Success 200 {object} vo.AppPageVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /apps/list [get]
func (c *AppController) GetList() (vo.AppPageVO, error) {
	return vo.AppPageVO{}, nil
}

// @Tags 应用管理
// @Summary 通过appid获取应用信息
// @Description 通过应用appid获取应用信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param appid query integer true "应用appid"
// @Success 200 {object} vo.AppVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /apps/byappid [get]
func (c *AppController) GetByappid() (vo.AppVO, error) {
	return vo.AppVO{}, nil
}
