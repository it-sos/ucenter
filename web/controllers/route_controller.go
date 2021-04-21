package controllers

import (
	"github.com/kataras/iris/v12"
	"ucenter/models/vo"
	"ucenter/services"
)

type RouteController struct {
	Service services.UserService
	Ctx     iris.Context
}

// @Tags 路由管理
// @Summary 创建路由信息
// @Description 创建路由信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param body body vo.RouteParamsVO true "request body"
// @Success 200 {object} vo.RouteVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /testapp [post]
func (c *RouteController) Post() (vo.RouteVO, error) {
	return vo.RouteVO{}, nil
}

// @Tags 路由管理
// @Summary 更新路由信息
// @Description 更新路由信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param id query integer true "路由id"
// @Param body body vo.RouteParamsVO true "request body"
// @Success 200 {object} vo.RouteVO	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /testapp [put]
func (c *RouteController) Put() (vo.RouteVO, error) {
	return vo.RouteVO{}, nil
}

// @Tags 路由管理
// @Summary 删除路由信息
// @Description 通过应用id批量删除路由信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param appid query integer true "应用id"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /testapp/byappid [delete]
func (c *RouteController) DeleteByappid() error {
	return nil
}

// @Tags 路由管理
// @Summary 删除路由信息
// @Description 通过路由id删除路由信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param id query integer true "路由id"
// @Success 200 {string} string	"success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /testapp [delete]
func (c *RouteController) Delete() error {
	return nil
}

// @Tags 路由管理
// @Summary 获取路由信息
// @Description 通过路由ID获取路由信息
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param id query integer true "路由id"
// @Success 200 {object} vo.RouteVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /testapp [get]
func (c *RouteController) Get() (vo.RouteVO, error) {
	return vo.RouteVO{}, nil
}

// @Tags 路由管理
// @Summary 获取路由分页列表
// @Description 获取全部路由分页列表
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param body body vo.PageVO true "request body"
// @Success 200 {object} vo.RoutePageVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /testapp/list [get]
func (c *RouteController) GetList() (vo.RoutePageVO, error) {
	return vo.RoutePageVO{}, nil
}

// @Tags 路由管理
// @Summary 通过应用id获取路由列表
// @Description 通过应用id获取路由列表
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Param appid query integer true "路由appid"
// @Success 200 {object} vo.RoutePageVO "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /testapp/listbyappid [get]
func (c *RouteController) GetListByappid() (vo.RoutePageVO, error) {
	return vo.RoutePageVO{}, nil
}
