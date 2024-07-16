package models

import (
	"github.com/jinzhu/gorm"
)

// Task represents the structure of our task model
type Task struct {
	gorm.Model
	Name        string `gorm:"size:255;not null" json:"name"`
	Description string `gorm:"size:255" json:"description"`
}
