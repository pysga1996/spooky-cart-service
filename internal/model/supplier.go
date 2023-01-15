package model

import (
	"github.com/thanh-vt/splash-inventory-service/internal/constant"
	"time"
)

type Supplier struct {
	Code        *string          `json:"code"`
	Name        *string          `json:"name"`
	CreatedDate *time.Time       `json:"createdDate"`
	UpdatedDate *time.Time       `json:"updatedDate"`
	Status      *constant.Status `json:"status"`
}
