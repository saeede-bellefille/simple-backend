package product

import (
	"github.com/saeede-bellefille/simple-backend/internal/domain"
	"github.com/saeede-bellefille/simple-backend/internal/repository"
)

type Service struct {
	repo *repository.ProductRepo
}

func New(repo *repository.ProductRepo) *Service {
	return &Service{repo: repo}
}

func (p *Service) List() ([]domain.Product, error) {
	return p.repo.List()
}

func (p *Service) Get(id uint) (*domain.Product, error) {
	return p.repo.Get(id)
}
func (p *Service) Create(product *domain.Product) error {
	return p.repo.Create(product)
}

func (p *Service) Update(id uint, product *domain.Product) error {
	return p.repo.Update(id, product)
}

func (p *Service) Delete(id uint) error {
	return p.repo.Delete(id)
}
