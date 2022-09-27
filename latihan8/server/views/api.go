package views

type Response struct {
	Message string      `json:"message" example:"GET_SUCCESS"`
	Status  int         `json:"status" example:"200"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type ResponseCreate struct {
	Message string      `json:"message" example:"CREATE_SUCCESS"`
	Status  int         `json:"status" example:"200"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type ResponseUpdate struct {
	Message string      `json:"message" example:"UPDATE_SUCCESS"`
	Status  int         `json:"status" example:"200"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type ResponseDelete struct {
	Message string      `json:"message" example:"DELETE_SUCCESS"`
	Status  int         `json:"status" example:"200"`
	Error   interface{} `json:"error,omitempty"`
}
