package menu

import "github.com/betsegawlemma/restaurant-rest/entity"

// CategoryService specifies food menu category services
type CategoryService interface {
	Categories() ([]entity.Category, []error)
	Category(id uint) (*entity.Category, []error)
	UpdateCategory(category *entity.Category) (*entity.Category, []error)
	DeleteCategory(id uint) (*entity.Category, []error)
	StoreCategory(category *entity.Category) (*entity.Category, []error)
	ItemsInCategory(category *entity.Category) ([]entity.Item, []error)
}

// ItemService specifies food menu item related services
type ItemService interface {
	Items() ([]entity.Item, []error)
	Item(id uint) (*entity.Item, []error)
	UpdateItem(menu *entity.Item) (*entity.Item, []error)
	DeleteItem(id uint) (*entity.Item, []error)
	StoreItem(item *entity.Item) (*entity.Item, []error)
}

// IngredientService speifies food item ingredints related services
type IngredientService interface {
	Ingredients() ([]entity.Ingredient, []error)
	Ingredient(id uint) (*entity.Ingredient, []error)
	UpdateIngredient(ingredient *entity.Ingredient) (*entity.Ingredient, []error)
	DeleteIngredient(id uint) (*entity.Ingredient, []error)
	StoreIngredient(ingredient *entity.Ingredient) (*entity.Ingredient, []error)
}
