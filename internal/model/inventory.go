package model

import (
	"github.com/thanh-vt/splash-inventory-service/internal/constant"
	"time"
)

type Inventory struct {
	InventoryId *string             `json:"inventoryId"`
	Details     *[]*InventoryDetail `json:"details"`
	CreatedDate *time.Time          `json:"createdDate"`
	UpdatedDate *time.Time          `json:"updatedDate"`
	Status      *constant.Status    `json:"status"`
}
