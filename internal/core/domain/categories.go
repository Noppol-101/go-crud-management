package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	CategoryName string     `gorm:"type:varchar(100)" json:"category_name"`
	Description  string     `gorm:"type:varchar(100)" json:"description"`
	Products     []Products `gorm:"foreignKey:CategoryId"`
	UUID         uuid.UUID  `gorm:"type:uuid"`
}
