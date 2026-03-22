package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/service"
)

type TeacherHandler struct {
	teacherService service.TeacherService
}

func NewTeacherHandler(teacherService service.TeacherService) *TeacherHandler {
	return &TeacherHandler{teacherService: teacherService}
}

func (h *TeacherHandler) GetAll(c *fiber.Ctx) error {
	teachers, err := h.teacherService.GetAll(c.Context())
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", "Öğretmen listesi alınamadı")
	}

	return Success(c, fiber.StatusOK, teachers, "Öğretmen listesi")
}
