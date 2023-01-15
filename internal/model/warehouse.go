package model

import (
	"github.com/thanh-vt/splash-inventory-service/internal/constant"
	"time"
)

type Warehouse struct {
	Id          *string          `json:"id"`
	Name        *string          `json:"name"`
	LocationId  *string          `json:"locationId"`
	CreatedDate *time.Time       `json:"createdDate"`
	UpdatedDate *time.Time       `json:"updatedDate"`
	Status      *constant.Status `json:"status"`
}
