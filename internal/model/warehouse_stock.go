package model

import (
	"time"
)

type WarehouseStock struct {
	WarehouseId *string    `json:"warehouseId"`
	StockId     *string    `json:"stockId"`
	Quantity    *time.Time `json:"quantity"`
}
