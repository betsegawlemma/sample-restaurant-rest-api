package service

import (
	"github.com/betsegawlemma/restaurant-rest/entity"
	"github.com/betsegawlemma/restaurant-rest/menu"
)

// ItemService implements menu.ItemService interface
type ItemService struct {
	itemRepo menu.ItemRepository
}

// NewItemService returns new ItemService object
func NewItemService(itemRepository menu.ItemRepository) menu.ItemService {
	return &ItemService{itemRepo: itemRepository}
}

// Items returns all stored food menu items
func (is *ItemService) Items() ([]entity.Item, []error) {
	itms, errs := is.itemRepo.Items()
	if len(errs) > 0 {
		return nil, errs
	}
	return itms, errs
}

// Item retrieves a food menu item by its id
func (is *ItemService) Item(id uint) (*entity.Item, []error) {
	itm, errs := is.itemRepo.Item(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

// UpdateItem updates a given food menu item
func (is *ItemService) UpdateItem(item *entity.Item) (*entity.Item, []error) {
	itm, errs := is.itemRepo.UpdateItem(item)
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

// DeleteItem deletes a given food menu item
func (is *ItemService) DeleteItem(id uint) (*entity.Item, []error) {
	itm, errs := is.itemRepo.DeleteItem(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

// StoreItem stores a given food menu item
func (is *ItemService) StoreItem(item *entity.Item) (*entity.Item, []error) {
	itm, errs := is.itemRepo.StoreItem(item)
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}
