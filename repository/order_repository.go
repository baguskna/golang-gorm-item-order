package repository

import (
	"golang-gorm-item-order/model/domain"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order domain.Order) (domain.Order, error)
	Delete(ID int) error
	FindAll() ([]domain.Order, error)
	FindById(ID int) (domain.Order, error)
	Update(order domain.Order) (domain.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Create(order domain.Order) (domain.Order, error) {
	err := r.db.Create(&order).Error
	return order, err
}

func (r *orderRepository) Delete(ID int) error {
	err := r.db.Where("order_id = ?", ID).Delete(&domain.Items{}).Error
	err = r.db.Delete(&domain.Order{}, ID).Error
	return err
}

func (r *orderRepository) FindAll() ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Preload("Items").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) FindById(ID int) (domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("Items").Find(&order, ID).Error
	return order, err
}

func (r *orderRepository) Update(order domain.Order) (domain.Order, error) {
	err := r.db.Save(&order).Error
	return order, err
}
