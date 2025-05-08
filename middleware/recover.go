package middleware

import (
	"fmt"
	"my-crud-management/internal/adapter/dto"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Recover() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				stack := debug.Stack()

				logrus.WithFields(logrus.Fields{
					"method": c.Method(),
					"route":  c.OriginalURL(),
					"panic":  fmt.Sprintf("%v", r),
					"stack":  string(stack),
				}).Error("recovered from panic : ")

				_ = c.Status(fiber.StatusInternalServerError).JSON(dto.Response{
					Status:    500,
					Message:   "internal server error",
					MessageTh: "ระบบมีปัญหา กรุณาติดต่อเจ้าหน้าที่",
					Error:     fmt.Sprintf("%v : %v %v", r, c.Method(), c.OriginalURL()),
				})
			}
		}()
		return c.Next()
	}
}
