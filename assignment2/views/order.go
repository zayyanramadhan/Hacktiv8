package views

import "time"

type OrderGetAll struct {
	OrderId      int          `json:"orderId"`
	OrderedAt    time.Time    `json:"orderedAt"`
	CustomerName string       `json:"customerName"`
	Items        []ItemGetAll `json:"items"`
}
