package models

import "time"

type Order struct {
	OrderId      int `gorm:"primaryKey;autoIncrement"`
	CustomerName string
	OrderedAt    time.Time
}

type CreateOrderRequest struct {
	CustomerName string              `json:"customerName" example:"John Doe"`
	OrderAt      string              `json:"orderedAt" example:"2021-05-01T00:00:00Z"`
	Items        []CreateItemRequest `json:"items"`
}

type UpdateOrderRequest struct {
	// OrderID      int                 `json:"orderid"`
	CustomerName string              `json:"customerName" example:"Zayyan"`
	OrderAt      string              `json:"orderedAt" example:"2022-10-01T21:21:21Z"`
	Items        []UpdateItemRequest `json:"items"`
}

func (req *CreateOrderRequest) ToModel() (*Order, error) {
	convertTime, err := time.Parse(time.RFC3339, req.OrderAt)
	if err != nil {
		return nil, err
	}

	order := &Order{
		CustomerName: req.CustomerName,
		OrderedAt:    convertTime,
	}

	return order, nil
}

func (req *UpdateOrderRequest) ToModel(id int) (*Order, error) {
	convertTime, err := time.Parse(time.RFC3339, req.OrderAt)
	if err != nil {
		return nil, err
	}

	order := &Order{
		OrderId:      id,
		CustomerName: req.CustomerName,
		OrderedAt:    convertTime,
	}

	return order, nil
}
