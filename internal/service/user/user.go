package user

import (
	"errors"

	"github.com/saeede-bellefille/simple-backend/internal/domain"
	"github.com/saeede-bellefille/simple-backend/internal/repository"
)

type Service struct {
	repo *repository.UserRepo
}

func New(repo *repository.UserRepo) *Service {
	return &Service{repo: repo}
}

func (u *Service) Test() string {
	return "Hello World"
}

func (u *Service) Read(username string) (*domain.User, error) {
	return u.repo.Get(username)
}

func (u *Service) Register(user *domain.User, password, repeat string) error {
	if password != repeat {
		return errors.New("Password do not match!")
	}
	if len(password) < 6 {
		return errors.New("Password is too short")
	}
	return u.repo.Create(user, password)
}
