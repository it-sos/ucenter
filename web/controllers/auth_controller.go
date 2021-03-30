package controllers

import (
	"github.com/kataras/iris/v12"
	"ucenter/models/vo"
	"ucenter/services"
)

type AuthController struct {
	Service services.UserService
	Ctx     iris.Context
}

// @Tags 认证管理
// @Summary 换取新的token认证
// @Description 通过refresh_token换取token
// @Accept json
// @Produce json
// @Param body body vo.AuthRefreshTokenVO true "request body"
// @Success 200 {object} vo.AuthTokenVO	"success"
// @Failure 400 {object} errors.Errors "error"
// @Router /auth/refresh [put]
func (c *AuthController) PostRefresh() (vo.AuthTokenVO, error) {
	return vo.AuthTokenVO{}, nil
}

// @Tags 认证管理
// @Summary
// @Description 登陆认证
// @Accept json
// @Produce json
// @Param body body vo.AuthVO true "request body"
// @Success 200 {object} vo.AuthTokenVO	"success"
// @Failure 400 {object} errors.Errors "error"
// @Router /auth/login [put]
func (c *AuthController) PostLogin() (vo.AuthTokenVO, error) {
	return vo.AuthTokenVO{}, nil
}

// @Tags 认证管理
// @Summary 获取验证码图片
// @Description 获取验证码图片
// @Accept json
// @Produce json
// @Param account query string true "用户帐号"
// @Success 200 {string} string	"图片流"
// @Failure 400 {object} errors.Errors "error"
// @Router /auth/captcha [get]
func (c *AuthController) GetCaptcha() error {
	return nil
}

// @Tags 认证管理
// @Summary 退出登陆
// @Description 退出登陆状态
// @Accept json
// @Produce json
// @Param token header string true "token认证"
// @Success 200 {string} string "success"
// @Failure 400 {object} errors.Errors "error"
// @Security token[read]
// @Router /auth/logout [get]
func (c *AuthController) GetLogout() error {
	return nil
}
