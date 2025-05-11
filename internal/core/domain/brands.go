package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Brands struct {
	gorm.Model
	BrandName  string     `gorm:"type:varchar(100)"`
	Address    string     `gorm:"type:varchar(100)"`
	City       string     `gorm:"type:varchar(100)"`
	Region     string     `gorm:"type:varchar(100)"`
	PostalCode string     `gorm:"type:varchar(100)"`
	Phone      string     `gorm:"type:varchar(100)"`
	Fax        string     `gorm:"type:varchar(100)"`
	HomePage   string     `gorm:"type:varchar(100)"`
	Products   []Products `gorm:"foreignKey:BrandId"`
	UUID       uuid.UUID  `gorm:"type:uuid"`
}
