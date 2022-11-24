package service

import (
	"golang-gorm-item-order/model/domain"
	"golang-gorm-item-order/repository"
)

type OrderService interface {
	Create(order domain.Order) (domain.Order, error)
	Delete(ID int) (domain.Order, error)
	FindAll() ([]domain.Order, error)
	FindById(ID int) (domain.Order, error)
	Update(ID int, newOrder domain.Order) (domain.Order, error)
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(r repository.OrderRepository) *orderService {
	return &orderService{r}
}

func (s *orderService) Create(order domain.Order) (domain.Order, error) {
	newOrder, err := s.orderRepository.Create(order)
	return newOrder, err
}

func (s *orderService) Delete(ID int) (domain.Order, error) {
	order, err := s.orderRepository.FindById(ID)
	err = s.orderRepository.Delete(ID)
	return order, err
}

func (s *orderService) FindAll() ([]domain.Order, error) {
	orders, err := s.orderRepository.FindAll()
	return orders, err
}

func (s *orderService) FindById(ID int) (domain.Order, error) {
	return s.orderRepository.FindById(ID)
}

func (s *orderService) Update(ID int, newOrder domain.Order) (domain.Order, error) {
	order, err := s.orderRepository.FindById(ID)

	order.CustomerName = newOrder.CustomerName

	newNewOrder, err := s.orderRepository.Update(order)

	return newNewOrder, err
}
