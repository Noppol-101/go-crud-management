package domain

import (
	"gorm.io/gorm"
)

type Provinces struct {
	gorm.Model
	ProvinceCode string      `gorm:"type:varchar(2);not null;index"`
	NameEn       string      `gorm:"type:varchar(100)"`
	NameTh       string      `gorm:"type:varchar(100)"`
	Districts    []Districts `gorm:"foreignKey:ProvinceCode"`
}

type Districts struct {
	gorm.Model
	DistrictCode string `gorm:"type:varchar(10);not null;index"` // foreign key
	NameEN       string `gorm:"type:varchar(100)"`
	NameTH       string `gorm:"type:varchar(100)"`
	Postcode     string `gorm:"type:varchar(6)"`
	ProvinceCode string
	Province     Provinces `gorm:"foreignKey:ProvinceCode"`
}

// type Subdistricts struct {
// 	gorm.Model
// 	ProvinceCode     string
// 	DistrictCode     string
// 	SubdistrictsCode string    `gorm:"type:varchar(10);not null;index"` // foreign key
// 	NameEN           string    `gorm:"type:varchar(100)"`
// 	NameTH           string    `gorm:"type:varchar(100)"`
// 	Postcode         string    `gorm:"type:varchar(6)"`
// 	District         Districts `gorm:"foreignKey:DistrictCode"`
// }
