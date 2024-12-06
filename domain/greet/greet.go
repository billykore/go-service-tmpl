package greet

import "gorm.io/gorm"

// Greet is greet entity.
type Greet struct {
	gorm.Model
	Name string
}
