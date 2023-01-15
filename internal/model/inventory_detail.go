package model

import (
	"github.com/thanh-vt/splash-inventory-service/internal/constant"
	"time"
)

type InventoryDetail struct {
	Id          *string          `json:"id"`
	InventoryId *string          `json:"inventoryId"`
	ProductId   *string          `json:"productId"`
	ProductName *string          `json:"productName"`
	Quantity    *uint8           `json:"quantity"`
	CreatedDate *time.Time       `json:"createdDate"`
	UpdatedDate *time.Time       `json:"updatedDate"`
	Status      *constant.Status `json:"status"`
}
