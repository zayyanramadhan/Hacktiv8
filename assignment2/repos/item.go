package repos

import (
	"assignment2/models"

	"gorm.io/gorm"
)

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepo {
	return &itemRepo{db}
}

func (r *itemRepo) Create(item *models.Item) error {
	return r.db.Create(item).Error
}

func (r *itemRepo) CreateManyItem(item []models.Item) error {
	return r.db.Create(&item).Error
}

func (r *itemRepo) Update(item *models.Item) error {
	return r.db.First(&models.Item{}, item.ItemId).Updates(item).Error
}

func (r *itemRepo) DeleteItemByOrderId(orderId int) error {
	return r.db.Where("order_id = ?", orderId).Delete(&models.Item{}).Error
}

func (r *itemRepo) FindByOrderId(orderId int) ([]models.Item, error) {
	var items []models.Item
	err := r.db.Where("order_id = ?", orderId).Find(&items).Error
	return items, err
}
