package service

import (
	"github.com/betsegawlemma/restaurant-rest/entity"
	"github.com/betsegawlemma/restaurant-rest/menu"
)

// CategoryService implements menu.CategoryService interface
type CategoryService struct {
	categoryRepo menu.CategoryRepository
}

// NewCategoryService will create new CategoryService object
func NewCategoryService(CatRepo menu.CategoryRepository) menu.CategoryService {
	return &CategoryService{categoryRepo: CatRepo}
}

// Categories returns list of categories
func (cs *CategoryService) Categories() ([]entity.Category, []error) {

	categories, errs := cs.categoryRepo.Categories()

	if len(errs) > 0 {
		return nil, errs
	}

	return categories, nil
}

// StoreCategory persists new category information
func (cs *CategoryService) StoreCategory(category *entity.Category) (*entity.Category, []error) {

	cat, errs := cs.categoryRepo.StoreCategory(category)

	if len(errs) > 0 {
		return nil, errs
	}

	return cat, nil
}

// Category returns a category object with a given id
func (cs *CategoryService) Category(id uint) (*entity.Category, []error) {

	c, errs := cs.categoryRepo.Category(id)

	if len(errs) > 0 {
		return c, errs
	}

	return c, nil
}

// UpdateCategory updates a cateogory with new data
func (cs *CategoryService) UpdateCategory(category *entity.Category) (*entity.Category, []error) {

	cat, errs := cs.categoryRepo.UpdateCategory(category)

	if len(errs) > 0 {
		return nil, errs
	}

	return cat, nil
}

// DeleteCategory delete a category by its id
func (cs *CategoryService) DeleteCategory(id uint) (*entity.Category, []error) {

	cat, errs := cs.categoryRepo.DeleteCategory(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return cat, nil
}

// ItemsInCategory returns list of menu items in a given category
func (cs *CategoryService) ItemsInCategory(category *entity.Category) ([]entity.Item, []error) {

	cts, errs := cs.categoryRepo.ItemsInCategory(category)

	if len(errs) > 0 {
		return nil, errs
	}

	return cts, nil

}
