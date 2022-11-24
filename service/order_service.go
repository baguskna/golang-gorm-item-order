package service

import (
	"golang-gorm-item-order/model/domain"
	"golang-gorm-item-order/repository"
)

type OrderService interface {
	Create(item domain.Order) (domain.Order, error)
	Delete(ID int) (domain.Order, error)
	FindAll() ([]domain.Order, error)
	FindById(ID int) (domain.Order, error)
	Update(ID int, newItem domain.Order) (domain.Order, error)
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(r repository.OrderRepository) *orderService {
	return &orderService{r}
}

func (s *orderService) Create(item domain.Order) (domain.Order, error) {
	newItem, err := s.orderRepository.Create(item)
	return newItem, err
}

func (s *orderService) Delete(ID int) (domain.Order, error) {
	item, err := s.orderRepository.FindById(ID)
	err = s.orderRepository.Delete(ID)
	return item, err
}

func (s *orderService) FindAll() ([]domain.Order, error) {
	items, err := s.orderRepository.FindAll()
	return items, err
}

func (s *orderService) FindById(ID int) (domain.Order, error) {
	return s.orderRepository.FindById(ID)
}

func (s *orderService) Update(ID int, newItem domain.Order) (domain.Order, error) {
	item, err := s.orderRepository.FindById(ID)

	// item.ItemCode = newItem.ItemCode
	// item.Description = newItem.Description
	// item.Quantity = newItem.Quantity

	newNewItem, err := s.orderRepository.Update(item)

	return newNewItem, err
}
