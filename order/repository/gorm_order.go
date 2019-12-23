package repository

import (
	"github.com/betsegawlemma/restaurant-rest/entity"
	"github.com/betsegawlemma/restaurant-rest/order"
	"github.com/jinzhu/gorm"
)

// OrderGormRepo implements the menu.OrderRepository interface
type OrderGormRepo struct {
	conn *gorm.DB
}

// NewOrderGormRepo returns new object of OrderGormRepo
func NewOrderGormRepo(db *gorm.DB) order.OrderRepository {
	return &OrderGormRepo{conn: db}
}

// Orders returns all customer orders stored in the database
func (orderRepo *OrderGormRepo) Orders() ([]entity.Order, []error) {
	orders := []entity.Order{}
	errs := orderRepo.conn.Find(&orders).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return orders, errs
}

// Order retrieve customer order by order id
func (orderRepo *OrderGormRepo) Order(id uint) (*entity.Order, []error) {
	order := entity.Order{}
	errs := orderRepo.conn.First(&order, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &order, errs
}

// UpdateOrder updates a given customer order in the database
func (orderRepo *OrderGormRepo) UpdateOrder(order *entity.Order) (*entity.Order, []error) {
	ordr := order
	errs := orderRepo.conn.Save(ordr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ordr, errs
}

// DeleteOrder deletes a given order from the database
func (orderRepo *OrderGormRepo) DeleteOrder(id uint) (*entity.Order, []error) {
	ordr, errs := orderRepo.Order(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = orderRepo.conn.Delete(ordr, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ordr, errs
}

// StoreOrder stores a given order in the database
func (orderRepo *OrderGormRepo) StoreOrder(order *entity.Order) (*entity.Order, []error) {
	ordr := order
	errs := orderRepo.conn.Create(ordr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ordr, errs
}

// CustomerOrders returns list of orders from the database for a given customer
func (orderRepo *OrderGormRepo) CustomerOrders(customer *entity.User) ([]entity.Order, []error) {
	custOrders := []entity.Order{}
	errs := orderRepo.conn.Model(customer).Related(&custOrders, "Orders").GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return custOrders, errs
}
