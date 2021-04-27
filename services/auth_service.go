package services

import (
	"github.com/kataras/iris/v12/middleware/jwt"
	"time"
	"ucenter/caches"
	"ucenter/models/vo"
	"ucenter/s/config"
	"ucenter/s/crypt"
	"ucenter/s/errors"
	"ucenter/s/utils"
	"ucenter/s/utils/captcha"
)

const (
	accessTokenMaxAge  = 10 * time.Minute
	refreshTokenMaxAge = time.Hour
)

var (
	privateKey, publicKey = jwt.MustLoadRSA(config.C.GetFile("pem/rsa_private_key.pem"), config.C.GetFile("rsa_public_key.pem"))

	signer   = jwt.NewSigner(jwt.RS256, privateKey, accessTokenMaxAge)
	verifier = jwt.NewVerifier(jwt.RS256, publicKey)
)

func NewAuthService(user UserService) AuthService {
	return &authService{
		user: user,
	}
}

type AuthService interface {
	Login(auth vo.AuthVO) (vo.AuthTokenVO, error)
	GenerateCaptcha(account string) string
	generalPassword(password string) (passwordShal, salt string)
	validatePassword(password, salt, passwordShal string) bool
	validCaptcha(account, captcha string) error
}

type authService struct {
	user UserService
}

func (s *authService) GenerateCaptcha(account string) string {
	id, b64s := captcha.Generate()
	caches.NAuthCaptcha.Key(account).Set(id)
	return b64s
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
	caches.NAuthTimes.Key(auth.Account).Clear()
	// todo 生成token与refresh_token
	return authToken, nil
}

// 超过3次需要验证码
const validCaptchaLimit = 3

// 输入3次错误密码则要求验证码验证
func (s *authService) validCaptcha(account, captchas string) error {
	valid := caches.NAuthTimes.Key(account)
	validTimes := valid.Get()
	if validTimes < validCaptchaLimit {
		valid.Incr()
		return nil
	}

	id := caches.NAuthCaptcha.Key(account).Get()
	if !captcha.Verify(id, captchas) {
		return errors.Error("login_captcha_err")
	}
	return nil
}
