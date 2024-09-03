package services

import (
	"new-shout-golang/models"
	"new-shout-golang/repositories"
)

type SellerService interface {
	CreateSeller(seller *models.Seller) (*models.Seller, error)
	GetAllSellers() ([]models.Seller, error)
	GetSellerByID(id uint) (*models.Seller, error)
	UpdateSeller(seller *models.Seller) (*models.Seller, error)
	DeleteSeller(id uint) error
}

type sellerService struct {
	repo repositories.SellerRepository
}

func NewSellerService(repo repositories.SellerRepository) SellerService {
	return &sellerService{repo: repo}
}

func (s *sellerService) CreateSeller(seller *models.Seller) (*models.Seller, error) {
	return s.repo.CreateSeller(seller)
}

func (s *sellerService) GetAllSellers() ([]models.Seller, error) {
	return s.repo.GetAllSellers()
}

func (s *sellerService) GetSellerByID(id uint) (*models.Seller, error) {
	return s.repo.GetSellerByID(id)
}

func (s *sellerService) UpdateSeller(seller *models.Seller) (*models.Seller, error) {
	return s.repo.UpdateSeller(seller)
}

func (s *sellerService) DeleteSeller(id uint) error {
	return s.repo.DeleteSeller(id)
}
