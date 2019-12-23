package entity

import "time"

// Category represents Food Menu Category
type Category struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Description string
	Image       string `gorm:"type:varchar(255)"`
	Items       []Item `gorm:"many2many:item_categories"`
}

// Role repesents application user roles
type Role struct {
	ID   uint
	Name string `gorm:"type:varchar(255)"`
}

// Item represents food menu items
type Item struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Price       float32
	Description string
	Categories  []Category   `gorm:"many2many:item_categories"`
	Image       string       `gorm:"type:varchar(255)"`
	Ingredients []Ingredient `gorm:"many2many:item_ingredients"`
}

// Ingredient represents ingredients in a food item
type Ingredient struct {
	ID          uint
	Name        string `gorm:"type:varchar(255);not null"`
	Description string
}

// Order represents customer order
type Order struct {
	ID       uint
	PlacedAt time.Time
	UserID   uint
	ItemID   uint
	Quantity uint
}

// User represents application user
type User struct {
	ID       uint
	UserName string `gorm:"type:varchar(255);not null"`
	FullName string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null; unique"`
	Phone    string `gorm:"type:varchar(100);not null; unique"`
	Password string `gorm:"type:varchar(255)"`
	Roles    []Role `gorm:"many2many:user_roles"`
	Orders   []Order
}

// Comment represents comments forwarded by application users
type Comment struct {
	ID       uint      `json:"id"`
	FullName string    `json:"fullname" gorm:"type:varchar(255)"`
	Message  string    `json:"message"`
	Phone    string    `json:"phone" gorm:"type:varchar(100);not null; unique"`
	Email    string    `json:"email" gorm:"type:varchar(255);not null; unique"`
	PlacedAt time.Time `json:"placedat"`
}

// Error represents error message
type Error struct {
	Code    int
	Message string
}
