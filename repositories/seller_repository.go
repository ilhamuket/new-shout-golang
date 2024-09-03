package repositories

import (
	"gorm.io/gorm"
	"new-shout-golang/models"
)

type SellerRepository interface {
	CreateSeller(seller *models.Seller) (*models.Seller, error)
	GetAllSellers() ([]models.Seller, error)
	GetSellerByID(id uint) (*models.Seller, error)
	UpdateSeller(seller *models.Seller) (*models.Seller, error)
	DeleteSeller(id uint) error
}

type sellerRepository struct {
	db *gorm.DB
}

func NewSellerRepository(db *gorm.DB) SellerRepository {
	return &sellerRepository{db: db}
}

func (r *sellerRepository) CreateSeller(seller *models.Seller) (*models.Seller, error) {
	if err := r.db.Create(seller).Error; err != nil {
		return nil, err
	}
	return seller, nil
}

func (r *sellerRepository) GetAllSellers() ([]models.Seller, error) {
	var sellers []models.Seller
	if err := r.db.Find(&sellers).Error; err != nil {
		return nil, err
	}
	return sellers, nil
}

func (r *sellerRepository) GetSellerByID(id uint) (*models.Seller, error) {
	var seller models.Seller
	if err := r.db.First(&seller, id).Error; err != nil {
		return nil, err
	}
	return &seller, nil
}

func (r *sellerRepository) UpdateSeller(seller *models.Seller) (*models.Seller, error) {
	if err := r.db.Save(seller).Error; err != nil {
		return nil, err
	}
	return seller, nil
}

func (r *sellerRepository) DeleteSeller(id uint) error {
	if err := r.db.Delete(&models.Seller{}, id).Error; err != nil {
		return err
	}
	return nil
}
