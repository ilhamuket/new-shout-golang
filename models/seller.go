package models

import (
	"time"
)

type Seller struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	UserID       uint   `json:"user_id"`
	StoreName    string `json:"store_name"`
	Description  string `json:"description"`
	StoreAddress string `json:"store_address"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
