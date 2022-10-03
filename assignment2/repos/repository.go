package repos

import "assignment2/models"

type ItemRepo interface {
	Create(item *models.Item) error
	CreateManyItem(item []models.Item) error
	Update(item *models.Item) error
	DeleteItemByOrderId(orderId int) error
	FindByOrderId(orderId int) ([]models.Item, error)
}

type OrderRepo interface {
	Create(order *models.Order) error
	Update(order *models.Order) error
	Delete(id int) error
	FindAll() ([]models.Order, error)
}
