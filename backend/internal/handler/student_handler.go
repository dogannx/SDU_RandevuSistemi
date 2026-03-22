package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/service"
)

type StudentHandler struct {
	studentService service.StudentService
}

func NewStudentHandler(studentService service.StudentService) *StudentHandler {
	return &StudentHandler{studentService: studentService}
}

func (h *StudentHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	student, err := h.studentService.GetByID(c.Context(), id)
	if err != nil {
		return Error(c, fiber.StatusNotFound, "NOT_FOUND", "Öğrenci bulunamadı")
	}

	return Success(c, fiber.StatusOK, student.ToResponse(), "Öğrenci bilgileri")
}

func (h *StudentHandler) GetAll(c *fiber.Ctx) error {
	students, err := h.studentService.GetAll(c.Context())
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", "Öğrenci listesi alınamadı")
	}

	var responses []any
	for _, s := range students {
		responses = append(responses, s.ToResponse())
	}

	return Success(c, fiber.StatusOK, responses, "Öğrenci listesi")
}
