package port

import (
	"my-crud-management/internal/adapter/dto"
	"my-crud-management/internal/core/domain"
)

type CategoryRepository interface {
	SaveData(Category *domain.Categories) error
	// FindData()
	// FindDataById()
}

type CategoryService interface {
	CreateCategory(Category *domain.Categories) (dto.DtoCreateCategory, error)
	// FindCategory()
	// FindCategoryById()
}
