package model

import (
	"github.com/thanh-vt/splash-inventory-service/internal/constant"
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
