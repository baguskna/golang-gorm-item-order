package service

import (
	"golang-gorm-item-order/model/domain"
	"golang-gorm-item-order/repository"
)

type ItemService interface {
	Create(item domain.Items) (domain.Items, error)
	Delete(ID int) (domain.Items, error)
	FindAll() ([]domain.Items, error)
	FindById(ID int) (domain.Items, error)
	Update(ID int, newItem domain.Items) (domain.Items, error)
}

type itemService struct {
	itemRepository repository.ItemRepository
}

func NewItemService(r repository.ItemRepository) *itemService {
	return &itemService{r}
}

func (s *itemService) Create(item domain.Items) (domain.Items, error) {
	newItem, err := s.itemRepository.Create(item)
	return newItem, err
}

func (s *itemService) Delete(ID int) (domain.Items, error) {
	item, err := s.itemRepository.FindById(ID)
	err = s.itemRepository.Delete(ID)
	return item, err
}

func (s *itemService) FindAll() ([]domain.Items, error) {
	items, err := s.itemRepository.FindAll()
	return items, err
}

func (s *itemService) FindById(ID int) (domain.Items, error) {
	return s.itemRepository.FindById(ID)
}

func (s *itemService) Update(ID int, newItem domain.Items) (domain.Items, error) {
	item, err := s.itemRepository.FindById(ID)

	item.ItemCode = newItem.ItemCode
	item.Description = newItem.Description
	item.Quantity = newItem.Quantity

	newNewItem, err := s.itemRepository.Update(item)

	return newNewItem, err
}
