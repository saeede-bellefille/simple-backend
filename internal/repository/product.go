package repository

import (
	"github.com/saeede-bellefille/simple-backend/internal/domain"
	"github.com/saeede-bellefille/simple-backend/internal/repository/models"
	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (p *ProductRepo) List() ([]domain.Product, error) {
	var products []domain.Product

	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepo) Get(id uint) (*domain.Product, error) {
	product := models.Product{}
	if err := p.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &domain.Product{
		Id:    product.ID,
		Name:  product.Name,
		Group: product.Group,
		Price: product.Price,
	}, nil
}
func (p *ProductRepo) Create(product *domain.Product) error {
	pro := models.Product{
		Name:  product.Name,
		Group: product.Group,
		Price: product.Price,
	}
	return p.db.Create(&pro).Error
}
func (p *ProductRepo) Update(id uint, product *domain.Product) error {
	pro := models.Product{
		Name:  product.Name,
		Group: product.Group,
		Price: product.Price,
	}
	return p.db.Where("id = ?", id).Updates(&pro).Error
}

func (p *ProductRepo) Delete(id uint) error {
	return p.db.Delete(&domain.Product{}, id).Error
}
