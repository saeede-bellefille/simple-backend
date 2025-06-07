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
		Role:     string(user.Role),
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
		Role:     domain.Role(u.Role),
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
		Role:     domain.Role(u.Role),
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
func (r *UserRepo) UpdateRole(username string, role domain.Role) error {
	return r.db.Where("username = ?", username).Update("role", string(role)).Error
}

func (r *UserRepo) ChangePassword(username, password string) error {
	return r.db.Where("username = ?", username).Update("password", password).Error
}

func (r *UserRepo) List() ([]domain.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	var domainUsers []domain.User
	for _, u := range users {
		domainUsers = append(domainUsers, domain.User{
			Username: u.Username,
			Email:    u.Email,
			Name:     u.Name,
			Age:      u.Age,
			Role:     domain.Role(u.Role),
		})
	}
	return domainUsers, nil
}
