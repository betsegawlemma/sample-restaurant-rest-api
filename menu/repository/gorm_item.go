package repository

import (
	"github.com/betsegawlemma/restaurant-rest/entity"
	"github.com/betsegawlemma/restaurant-rest/menu"
	"github.com/jinzhu/gorm"
)

// ItemGormRepo implements the menu.ItemRepository interface
type ItemGormRepo struct {
	conn *gorm.DB
}

// NewItemGormRepo will create a new object of ItemGormRepo
func NewItemGormRepo(db *gorm.DB) menu.ItemRepository {
	return &ItemGormRepo{conn: db}
}

// Items returns all food menus stored in the database
func (itemRepo *ItemGormRepo) Items() ([]entity.Item, []error) {
	items := []entity.Item{}
	errs := itemRepo.conn.Find(&items).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return items, errs
}

// Item retrieves a food menu by its id from the database
func (itemRepo *ItemGormRepo) Item(id uint) (*entity.Item, []error) {
	item := entity.Item{}
	errs := itemRepo.conn.First(&item, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &item, errs
}

// UpdateItem updates a given food menu item in the database
func (itemRepo *ItemGormRepo) UpdateItem(item *entity.Item) (*entity.Item, []error) {
	itm := item
	errs := itemRepo.conn.Save(itm).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

// DeleteItem deletes a given food menu item from the database
func (itemRepo *ItemGormRepo) DeleteItem(id uint) (*entity.Item, []error) {
	itm, errs := itemRepo.Item(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = itemRepo.conn.Delete(itm, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

// StoreItem stores a given food menu item in the database
func (itemRepo *ItemGormRepo) StoreItem(item *entity.Item) (*entity.Item, []error) {
	itm := item
	errs := itemRepo.conn.Create(itm).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}
