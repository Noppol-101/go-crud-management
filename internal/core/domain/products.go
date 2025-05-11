package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	ProductName string `gorm:"type:varchar(100)"`
	BrandId     uint
	Brand       Brands     `gorm:"foreignKey:BrandId"` // กำหนดความสัมพันธ์
	CategoryId  uint       // กำหนดฟิลด์นี้เพื่อเก็บ Category ID
	Category    Categories `gorm:"foreignKey:CategoryId"` // กำหนดความสัมพันธ์
	UnitPrice   string     `gorm:"type:varchar(100)"`
	UnitInStock string     `gorm:"type:varchar(100)"`
	UUID        uuid.UUID  `gorm:"type:uuid"`
}
