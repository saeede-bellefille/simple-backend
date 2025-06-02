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

func (u *Service) Login(username, password string) (*domain.User, error) {
	return u.repo.GetByUsernamePassword(username, password)
}
func (u *Service) UpdateProfile(username string, user *domain.User) error {
	return u.repo.Update(username, user)
}

func (u *Service) ChangePassword(username, currentPassword, newPassword, repeatPassword string) error {
	_, err := u.repo.GetByUsernamePassword(username, currentPassword)
	if err != nil {
		return errors.New("current password is incorrect")
	}

	if newPassword != repeatPassword {
		return errors.New("new passwords do not match")
	}

	if len(newPassword) < 6 {
		return errors.New("new password is too short (minimum 6 characters)")
	}

	return u.repo.ChangePassword(username, newPassword)
}
