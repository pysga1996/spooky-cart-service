package model

import (
	"encoding/json"
	"time"
)

type Cart struct {
	tableName       struct{}
	id              *uint64
	productQuantity *map[string]*uint8
	createTime      *time.Time
	updateTime      *time.Time
	status          *uint8
}

func (a *Cart) Id() *uint64 {
	return a.id
}

func (a *Cart) SetId(id *uint64) {
	a.id = id
}

func (a *Cart) ProductQuantity() *map[string]*uint8 {
	return a.productQuantity
}

func (a *Cart) SetProductQuantity(productQuantity *map[string]*uint8) {
	a.productQuantity = productQuantity
}

func (a *Cart) CreateTime() *time.Time {
	return a.createTime
}

func (a *Cart) SetCreateTime(createTime *time.Time) {
	a.createTime = createTime
}

func (a *Cart) UpdateTime() *time.Time {
	return a.updateTime
}

func (a *Cart) SetUpdateTime(updateTime *time.Time) {
	a.updateTime = updateTime
}

func (a *Cart) Status() *uint8 {
	return a.status
}

func (a *Cart) SetStatus(status *uint8) {
	a.status = status
}

func (a *Cart) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id              *uint64            `json:"id"`
		ProductQuantity *map[string]*uint8 `json:"productQuantity"`
		CreateTime      *time.Time         `json:"createTime"`
		UpdateTime      *time.Time         `json:"updateTime"`
		Status          *uint8             `json:"status"`
	}{
		Id:              a.id,
		ProductQuantity: a.productQuantity,
		CreateTime:      a.createTime,
		UpdateTime:      a.updateTime,
		Status:          a.status,
	})
}

func (a *Cart) UnmarshalJSON(bytes []byte) error {
	var tmp struct {
		Id              uint64            `json:"id"`
		ProductQuantity map[string]*uint8 `json:"productQuantity"`
		CreateTime      time.Time         `json:"createTime"`
		UpdateTime      time.Time         `json:"updateTime"`
		Status          uint8             `json:"status"`
	}
	err := json.Unmarshal(bytes, &tmp)
	if err != nil {
		return err
	}
	a.SetId(&tmp.Id)
	a.SetProductQuantity(&tmp.ProductQuantity)
	a.SetCreateTime(&tmp.CreateTime)
	a.SetUpdateTime(&tmp.UpdateTime)
	a.SetStatus(&tmp.Status)
	return nil
}
