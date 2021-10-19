package model

type CartProduct struct {
	CartId      *uint64 `json:"cartId"`
	ProductCode *string `json:"productCode"`
	Quantity    *uint8  `json:"quantity"`
}
