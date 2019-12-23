package repository

import (
	"github.com/betsegawlemma/restaurant-rest/entity"
	"github.com/betsegawlemma/restaurant-rest/menu"
	"github.com/jinzhu/gorm"
)

// IngredientGormRepo implements menu.IngredientRepository interface
type IngredientGormRepo struct {
	conn *gorm.DB
}

// NewIngredientGormRepo returns new object of IngredientGormRepo
func NewIngredientGormRepo(db *gorm.DB) menu.IngredientRepository {
	return &IngredientGormRepo{conn: db}
}

// Ingredients return all food ingredients stored in the databasee
func (ingredRepo *IngredientGormRepo) Ingredients() ([]entity.Ingredient, []error) {
	ingts := []entity.Ingredient{}
	errs := ingredRepo.conn.Find(&ingts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ingts, errs
}

// Ingredient retrieves a food ingredient from the database by its id
func (ingredRepo *IngredientGormRepo) Ingredient(id uint) (*entity.Ingredient, []error) {
	ingt := entity.Ingredient{}
	errs := ingredRepo.conn.First(&ingt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &ingt, errs
}

// UpdateIngredient updats a given food ingredinet in the database
func (ingredRepo *IngredientGormRepo) UpdateIngredient(ingredient *entity.Ingredient) (*entity.Ingredient, []error) {
	ingt := ingredient
	errs := ingredRepo.conn.Save(ingt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ingt, errs
}

// DeleteIngredient deletes a given food ingredint from the database
func (ingredRepo *IngredientGormRepo) DeleteIngredient(id uint) (*entity.Ingredient, []error) {
	ingt, errs := ingredRepo.Ingredient(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = ingredRepo.conn.Delete(ingt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ingt, errs
}

// StoreIngredient stores a given food ingredient in the database
func (ingredRepo *IngredientGormRepo) StoreIngredient(ingredient *entity.Ingredient) (*entity.Ingredient, []error) {
	ingt := ingredient
	errs := ingredRepo.conn.Create(ingt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ingt, errs
}
