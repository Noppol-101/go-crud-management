package handler

import (
	"my-crud-management/internal/adapter/dto"
	"my-crud-management/internal/core/domain"
	"my-crud-management/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	categorySrv port.CategoryService
}

func NewCategoryHandler(categorySrv port.CategoryService) CategoryHandler {
	return CategoryHandler{
		categorySrv: categorySrv,
	}
}

func (h CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	category := new(domain.Categories)
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	result, err := h.categorySrv.CreateCategory(category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.Response{
		Status:    fiber.StatusCreated,
		Message:   "create data successful",
		MessageTh: "สร้างข้อมูลสำเร็จ",
		Data:      result,
	})
}
