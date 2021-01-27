package services

import (
	"ucenter/repositories"
)

type UserService interface {
	Test() bool
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
}

func (s *userService) Test() {
	s.repo.Test()
}
