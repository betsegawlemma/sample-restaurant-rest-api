package menu

import "github.com/betsegawlemma/restaurant-rest/entity"

// CategoryRepository specifies food menu category database operations
type CategoryRepository interface {
	Categories() ([]entity.Category, []error)
	Category(id uint) (*entity.Category, []error)
	UpdateCategory(category *entity.Category) (*entity.Category, []error)
	DeleteCategory(id uint) (*entity.Category, []error)
	StoreCategory(category *entity.Category) (*entity.Category, []error)
	ItemsInCategory(category *entity.Category) ([]entity.Item, []error)
}

// ItemRepository specifies food menu item related database operations
type ItemRepository interface {
	Items() ([]entity.Item, []error)
	Item(id uint) (*entity.Item, []error)
	UpdateItem(menu *entity.Item) (*entity.Item, []error)
	DeleteItem(id uint) (*entity.Item, []error)
	StoreItem(item *entity.Item) (*entity.Item, []error)
}

// IngredientRepository speifies food item ingredints related database operations
type IngredientRepository interface {
	Ingredients() ([]entity.Ingredient, []error)
	Ingredient(id uint) (*entity.Ingredient, []error)
	UpdateIngredient(ingredient *entity.Ingredient) (*entity.Ingredient, []error)
	DeleteIngredient(id uint) (*entity.Ingredient, []error)
	StoreIngredient(ingredient *entity.Ingredient) (*entity.Ingredient, []error)
}
