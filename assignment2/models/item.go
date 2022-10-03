package models

type Item struct {
	ItemId      int `gorm:"primaryKey;autoIncrement"`
	ItemCode    string
	Description string
	Quantity    int
	OrderId     int
}

type CreateItemRequest struct {
	ItemCode    string `json:"itemcode" example:"A001"`
	Description string `json:"description" example:"Baju"`
	Quantity    int    `json:"quantity" example:"1"`
}

type UpdateItemRequest struct {
	ItemId      int    `json:"itemId" example:"1"`
	ItemCode    string `json:"itemCode" example:"A001"`
	Description string `json:"description" example:"Baju"`
	Quantity    int    `json:"quantity" example:"1"`
}

func (req *CreateItemRequest) ToModel(orderId int) (*Item, error) {
	item := &Item{
		OrderId:     orderId,
		ItemCode:    req.ItemCode,
		Description: req.Description,
		Quantity:    req.Quantity,
	}

	return item, nil
}

func (req *UpdateItemRequest) ToModel(orderId int) (*Item, error) {
	item := &Item{
		ItemId:      req.ItemId,
		OrderId:     orderId,
		ItemCode:    req.ItemCode,
		Description: req.Description,
		Quantity:    req.Quantity,
	}

	return item, nil
}
