package services

import (
	"ucenter/models/vo"
	"ucenter/s/crypt"
	"ucenter/s/errors"
	"ucenter/s/utils"
)

func NewAuthService(user UserService) AuthService {
	return &authService{
		user: user,
	}
}

type AuthService interface {
	Login(auth vo.AuthVO) (vo.AuthTokenVO, error)
	generalPassword(password string) (passwordShal, salt string)
	validatePassword(password, salt, passwordShal string) bool
}

type authService struct {
	user UserService
}

func (s *authService) generalPassword(password string) (passwordShal, salt string) {
	salt = utils.Rand(32)
	passwordShal = crypt.Sha1(password + salt)
	return
}

func (s *authService) validatePassword(password, salt, passwordShal string) bool {
	return crypt.Sha1(password+salt) == passwordShal
}

func (s *authService) Login(auth vo.AuthVO) (vo.AuthTokenVO, error) {
	// todo 三次错误，启动验证码机制
	authToken := vo.AuthTokenVO{}
	user, has := s.user.GetByAccount(auth.Account)
	if !has || !s.validatePassword(auth.Password, user.Salt, user.Password) {
		return authToken, errors.Error("login_auth_err")
	}
	// todo 生成token与refresh_token
	return authToken, nil
}
