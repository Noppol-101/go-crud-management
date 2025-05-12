package repository

import (
	"my-crud-management/internal/core/domain"
	"my-crud-management/internal/core/port"

	"gorm.io/gorm"
)

type categoryGormRepo struct {
	db *gorm.DB
}

func NewCategoryGormRepo(db *gorm.DB) port.CategoryRepository {
	return &categoryGormRepo{
		db: db,
	}
}

func (r *categoryGormRepo) SaveData(Category *domain.Categories) error {

	if result := r.db.Create(Category); result.Error != nil {
		return result.Error
	}
	return nil
}
