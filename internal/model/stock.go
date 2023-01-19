package model

import (
	"github.com/oklog/ulid/v2"
	"github.com/thanh-vt/splash-inventory-service/internal/constant"
	"gorm.io/gorm"
	"time"
)

type Stock struct {
	Id                   *string          `json:"id"`
	ProductId            *string          `json:"productId"`
	ProductName          *string          `json:"productName"`
	Sku                  *string          `json:"sku"`
	Classification1Name  *string          `json:"classification1Name"`
	Classification1Value *string          `json:"classification1Value"`
	Classification2Name  *string          `json:"classification2Name"`
	Classification2Value *string          `json:"classification2Value"`
	CreatedDate          *time.Time       `json:"createdDate"`
	UpdatedDate          *time.Time       `json:"updatedDate"`
	Status               *constant.Status `json:"status"`
}

func (stock *Stock) TableName() string {
	return "stock"
}

func (stock *Stock) BeforeCreate(tx *gorm.DB) (err error) {
	id := ulid.Make().String()
	now := time.Now()
	stock.Id = &id
	stock.CreatedDate = &now
	return
}

func (stock *Stock) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	stock.UpdatedDate = &now
	return
}
