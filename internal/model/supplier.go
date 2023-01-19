package model

import (
	"github.com/thanh-vt/splash-inventory-service/internal/constant"
	"gorm.io/gorm"
	"time"
)

type Supplier struct {
	Code        *string          `gorm:"primaryKey" json:"code"`
	Name        *string          `gorm:"not null" json:"name"`
	CreatedDate *time.Time       `json:"createdDate"`
	UpdatedDate *time.Time       `json:"updatedDate"`
	Status      *constant.Status `gorm:"not null" json:"status"`
}

func (supplier *Supplier) TableName() string {
	return "supplier"
}

func (supplier *Supplier) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	supplier.CreatedDate = &now
	return
}

func (supplier *Supplier) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	supplier.UpdatedDate = &now
	return
}
