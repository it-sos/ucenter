package services

import (
	"ucenter/s/crypt"
	"ucenter/s/utils"
)

type authService struct {
}

func (s *authService) GeneralPassword(password string) (passwordShal string, salt string) {
	salt = utils.Rand(32)
	passwordShal = crypt.Sha1(password + salt)
	return
}

func (s *authService) ValidatePassword(passwordShal, password, salt string) bool {
	return crypt.Sha1(password+salt) == passwordShal
}
