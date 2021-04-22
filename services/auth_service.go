package services

import (
	"ucenter/caches"
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
	authToken := vo.AuthTokenVO{}
	if err := s.validCaptcha(auth.Account, auth.Captcha); err != nil {
		return authToken, err
	}
	user, has := s.user.GetByAccount(auth.Account)
	if !has || !s.validatePassword(auth.Password, user.Salt, user.Password) {
		return authToken, errors.Error("login_auth_err")
	}
	caches.NewAuthCache().Key(auth.Account).Clear()
	// todo 生成token与refresh_token
	return authToken, nil
}

// 超过3次需要验证码
const validCaptchaLimit = 3

// 输入3次错误密码则要求验证码验证
func (s *authService) validCaptcha(account, captcha string) error {
	valid := caches.NewAuthCache().Key(account)
	validTimes := valid.Get()
	if validTimes <= validCaptchaLimit {
		validTimes = valid.Incr()
	}
	if validTimes > validCaptchaLimit {
		if len(captcha) <= 6 {
			return errors.Error("login_captcha_err")
		}
	}
	// todo 验证码验证模块
	return nil
}
