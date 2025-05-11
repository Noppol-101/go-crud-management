package service

import (
	"my-crud-management/internal/adapter/dto"
	"my-crud-management/internal/core/domain"
	"my-crud-management/internal/core/port"

	"github.com/google/uuid"
)

type categoryServiceImpl struct {
	categoryRepo port.CategoryRepository
}

func NewCategoryService(categoryRepo port.CategoryRepository) port.CategoryService {
	return &categoryServiceImpl{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryServiceImpl) CreateCategory(Category domain.Categories) (dto.DtoCreateCategory, error) {

	if Category.UUID == uuid.Nil {
		Category.UUID = uuid.New()
	}

	if err := s.categoryRepo.SaveData(Category); err != nil {
		return dto.DtoCreateCategory{}, err
	}

	dataCreate := dto.DtoCreateCategory{
		CategoryName: Category.CategoryName,
		Description:  Category.Description,
		UUID:         Category.UUID.String(),
	}

	return dataCreate, nil
}
