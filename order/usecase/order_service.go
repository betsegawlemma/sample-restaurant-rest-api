package service

import (
	"github.com/betsegawlemma/restaurant-rest/entity"
	"github.com/betsegawlemma/restaurant-rest/order"
)

// OrderService implements menu.OrderService interface
type OrderService struct {
	orderRepo order.OrderRepository
}

// NewOrderService returns new OrderService object
func NewOrderService(orderRepository order.OrderRepository) order.OrderService {
	return &OrderService{orderRepo: orderRepository}
}

// Orders returns all stored food orders
func (os *OrderService) Orders() ([]entity.Order, []error) {
	ords, errs := os.orderRepo.Orders()
	if len(errs) > 0 {
		return nil, errs
	}
	return ords, errs
}

// Order retrieves an order by its id
func (os *OrderService) Order(id uint) (*entity.Order, []error) {
	ord, errs := os.orderRepo.Order(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, errs
}

// CustomerOrders returns all orders of a given customer
func (os *OrderService) CustomerOrders(customer *entity.User) ([]entity.Order, []error) {
	ords, errs := os.orderRepo.CustomerOrders(customer)
	if len(errs) > 0 {
		return nil, errs
	}
	return ords, errs
}

// UpdateOrder updates a given order
func (os *OrderService) UpdateOrder(order *entity.Order) (*entity.Order, []error) {
	ord, errs := os.orderRepo.UpdateOrder(order)
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, errs
}

// DeleteOrder deletes a given order
func (os *OrderService) DeleteOrder(id uint) (*entity.Order, []error) {
	ord, errs := os.DeleteOrder(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, errs
}

// StoreOrder stores a given order
func (os *OrderService) StoreOrder(order *entity.Order) (*entity.Order, []error) {
	ord, errs := os.StoreOrder(order)
	if len(errs) > 0 {
		return nil, errs
	}
	return ord, errs
}
