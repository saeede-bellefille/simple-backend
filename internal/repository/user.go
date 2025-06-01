package repository

import (
	"github.com/saeede-bellefille/simple-backend/internal/domain"
	"github.com/saeede-bellefille/simple-backend/internal/repository/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *domain.User, password string) error {
	u := models.User{
		Username: user.Username,
		Password: password,
		Email:    user.Email,
		Name:     user.Name,
		Age:      user.Age,
	}
	return r.db.Create(&u).Error
}

func (r *UserRepo) Get(username string) (*domain.User, error) {
	u := models.User{}
	if err := r.db.Where("username = ?", username).First(&u).Error; err != nil {
		return nil, err
	}
	return &domain.User{
		Username: u.Username,
		Email:    u.Email,
		Name:     u.Name,
		Age:      u.Age,
	}, nil
}

func (r *UserRepo) GetByUsernamePassword(username, password string) (*domain.User, error) {
	u := models.User{}
	if err := r.db.Where("username = ? AND password = ?", username, password).First(&u).Error; err != nil {
		return nil, err
	}
	return &domain.User{
		Username: u.Username,
		Email:    u.Email,
		Name:     u.Name,
		Age:      u.Age,
	}, nil
}

func (r *UserRepo) Update(username string, user *domain.User) error {
	u := models.User{
		Email: user.Email,
		Name:  user.Name,
		Age:   user.Age,
	}
	return r.db.Where("username = ?", username).Updates(&u).Error
}

func (r *UserRepo) ChangePassword(username, password string) error {
	return r.db.Where("username = ?", username).Update("password", password).Error
}
