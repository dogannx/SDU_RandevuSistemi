package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/model"
	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/service"
)

type AppointmentHandler struct {
	apptService service.AppointmentService
}

func NewAppointmentHandler(apptService service.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{apptService: apptService}
}

func (h *AppointmentHandler) Create(c *fiber.Ctx) error {
	studentID := c.Locals("userID").(string)

	var req model.AppointmentCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return Error(c, fiber.StatusBadRequest, "BAD_REQUEST", "Geçersiz istek formatı")
	}

	if req.TeacherID == "" || req.Date == "" || req.Time == "" {
		return Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Öğretmen, tarih ve saat zorunludur")
	}

	appt, err := h.apptService.Create(c.Context(), studentID, req)
	if err != nil {
		if errors.Is(err, service.ErrTimeConflict) {
			return Error(c, fiber.StatusConflict, "CONFLICT", err.Error())
		}
		return Error(c, fiber.StatusBadRequest, "BAD_REQUEST", err.Error())
	}

	return Success(c, fiber.StatusCreated, appt, "Randevu oluşturuldu")
}

func (h *AppointmentHandler) GetMy(c *fiber.Ctx) error {
	studentID := c.Locals("userID").(string)

	appointments, err := h.apptService.GetByStudentID(c.Context(), studentID)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", "Randevular alınamadı")
	}

	return Success(c, fiber.StatusOK, appointments, "Randevu listesi")
}

func (h *AppointmentHandler) Update(c *fiber.Ctx) error {
	studentID := c.Locals("userID").(string)
	apptID := c.Params("id")

	var req model.AppointmentUpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return Error(c, fiber.StatusBadRequest, "BAD_REQUEST", "Geçersiz istek formatı")
	}

	appt, err := h.apptService.Update(c.Context(), apptID, studentID, req)
	if err != nil {
		if errors.Is(err, service.ErrAppointmentNotFound) {
			return Error(c, fiber.StatusNotFound, "NOT_FOUND", err.Error())
		}
		if errors.Is(err, service.ErrNotOwner) {
			return Error(c, fiber.StatusForbidden, "FORBIDDEN", err.Error())
		}
		if errors.Is(err, service.ErrTimeConflict) {
			return Error(c, fiber.StatusConflict, "CONFLICT", err.Error())
		}
		if errors.Is(err, service.ErrAlreadyCancelled) {
			return Error(c, fiber.StatusBadRequest, "BAD_REQUEST", err.Error())
		}
		return Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", "Randevu güncellenemedi")
	}

	return Success(c, fiber.StatusOK, appt, "Randevu güncellendi")
}

func (h *AppointmentHandler) Cancel(c *fiber.Ctx) error {
	studentID := c.Locals("userID").(string)
	apptID := c.Params("id")

	err := h.apptService.Cancel(c.Context(), apptID, studentID)
	if err != nil {
		if errors.Is(err, service.ErrAppointmentNotFound) {
			return Error(c, fiber.StatusNotFound, "NOT_FOUND", err.Error())
		}
		if errors.Is(err, service.ErrNotOwner) {
			return Error(c, fiber.StatusForbidden, "FORBIDDEN", err.Error())
		}
		if errors.Is(err, service.ErrAlreadyCancelled) {
			return Error(c, fiber.StatusBadRequest, "BAD_REQUEST", err.Error())
		}
		return Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", "Randevu iptal edilemedi")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *AppointmentHandler) Suggest(c *fiber.Ctx) error {
	var req model.SuggestRequest
	if err := c.BodyParser(&req); err != nil {
		return Error(c, fiber.StatusBadRequest, "BAD_REQUEST", "Geçersiz istek formatı")
	}

	if req.LessonName == "" || len(req.AvailableSlots) == 0 {
		return Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Ders adı ve müsait zaman dilimleri zorunludur")
	}

	result, err := h.apptService.Suggest(c.Context(), req)
	if err != nil {
		return Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", "Öneri oluşturulamadı")
	}

	return Success(c, fiber.StatusOK, result, "Randevu önerileri")
}
