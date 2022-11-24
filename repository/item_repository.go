package repository

import (
	"golang-gorm-item-order/model/domain"

	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(item domain.Items) (domain.Items, error)
	Delete(ID int) error
	FindAll() ([]domain.Items, error)
	FindById(ID int) (domain.Items, error)
	Update(item domain.Items) (domain.Items, error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{db}
}

func (r *itemRepository) Create(item domain.Items) (domain.Items, error) {
	err := r.db.Create(&item).Error
	return item, err
}

func (r *itemRepository) Delete(ID int) error {
	err := r.db.Delete(&domain.Items{}, ID).Error
	return err
}

func (r *itemRepository) FindAll() ([]domain.Items, error) {
	var items []domain.Items
	err := r.db.Find(&items).Error
	return items, err
}

func (r *itemRepository) FindById(ID int) (domain.Items, error) {
	var item domain.Items
	err := r.db.Find(&item, ID).Error
	return item, err
}

func (r *itemRepository) Update(item domain.Items) (domain.Items, error) {
	err := r.db.Save(&item).Error
	return item, err
}
