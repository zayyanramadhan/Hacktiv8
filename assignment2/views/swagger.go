package views

type GetAllOrderWithItemsSuccess struct {
	Status  int         `json:"status" example:"200"`
	Message string      `json:"message" example:"OK"`
	Payload OrderGetAll `json:"payload"`
}

type SuccessDeleteOrder struct {
	Status  int    `json:"status" example:"200"`
	Message string `json:"message" example:"OK"`
}

type Failed struct {
	Message string `json:"message" example:"BAD_REQUEST"`
	Error   string `json:"error"`
}

type SuccessUpdateOrder struct {
	Status  int                       `json:"status" example:"200"`
	Message string                    `json:"message" example:"OK"`
	Payload structSuccessUpdateCreate `json:"payload"`
}

type structSuccessUpdateCreate struct {
	OrderId      int    `json:"OrderId" example:"1"`
	CustomerName string `json:"CustomerName" example:"Zayyan"`
	OrderedAt    string `json:"OrderedAt" example:"2022-10-01T21:21:21Z"`
}

type SuccesCreateOrder struct {
	Status  int                       `json:"status" example:"201"`
	Message string                    `json:"message" example:"CREATED"`
	Payload structSuccessUpdateCreate `json:"payload"`
}
