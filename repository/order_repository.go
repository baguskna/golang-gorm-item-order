package repository

import (
	"golang-gorm-item-order/model/domain"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(item domain.Order) (domain.Order, error)
	Delete(ID int) error
	FindAll() ([]domain.Order, error)
	FindById(ID int) (domain.Order, error)
	Update(item domain.Order) (domain.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Create(item domain.Order) (domain.Order, error) {
	err := r.db.Create(&item).Error
	return item, err
}

func (r *orderRepository) Delete(ID int) error {
	err := r.db.Delete(&domain.Order{}, ID).Error
	return err
}

func (r *orderRepository) FindAll() ([]domain.Order, error) {
	var items []domain.Order
	err := r.db.Preload("Items").Find(&items).Error
	return items, err
}

func (r *orderRepository) FindById(ID int) (domain.Order, error) {
	var item domain.Order
	err := r.db.Preload("Items").Find(&item, ID).Error
	return item, err
}

func (r *orderRepository) Update(item domain.Order) (domain.Order, error) {
	err := r.db.Save(&item).Error
	return item, err
}
