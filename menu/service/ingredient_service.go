package service

import (
	"github.com/betsegawlemma/restaurant-rest/entity"
	"github.com/betsegawlemma/restaurant-rest/menu"
)

// IngredientService implements menu.IngredientService interface
type IngredientService struct {
	ingredientRepo menu.IngredientRepository
}

// NewIngredientService returns new IngredientService object
func NewIngredientService(ingRepo menu.IngredientRepository) menu.IngredientService {
	return &IngredientService{ingredientRepo: ingRepo}
}

// Ingredients return all stored food item ingredients
func (is *IngredientService) Ingredients() ([]entity.Ingredient, []error) {
	ings, errs := is.ingredientRepo.Ingredients()
	if len(errs) > 0 {
		return nil, errs
	}
	return ings, errs
}

// Ingredient retrieves a food item ingredient by its id
func (is *IngredientService) Ingredient(id uint) (*entity.Ingredient, []error) {
	ing, errs := is.ingredientRepo.Ingredient(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return ing, errs
}

// UpdateIngredient updates a given food item ingredient
func (is *IngredientService) UpdateIngredient(ingredient *entity.Ingredient) (*entity.Ingredient, []error) {
	ing, errs := is.ingredientRepo.UpdateIngredient(ingredient)
	if len(errs) > 0 {
		return nil, errs
	}
	return ing, errs
}

// DeleteIngredient deletes a given food item ingredient
func (is *IngredientService) DeleteIngredient(id uint) (*entity.Ingredient, []error) {
	ing, errs := is.ingredientRepo.DeleteIngredient(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return ing, errs
}

// StoreIngredient stores a given food item ingredient
func (is *IngredientService) StoreIngredient(ingredient *entity.Ingredient) (*entity.Ingredient, []error) {
	ing, errs := is.ingredientRepo.StoreIngredient(ingredient)
	if len(errs) > 0 {
		return nil, errs
	}
	return ing, errs
}
