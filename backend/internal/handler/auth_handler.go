package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/model"
	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req model.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return Error(c, fiber.StatusBadRequest, "BAD_REQUEST", "Geçersiz istek formatı")
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		return Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Ad, e-posta ve şifre zorunludur")
	}

	if len(req.Password) < 8 {
		return Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "Şifre en az 8 karakter olmalıdır")
	}

	result, err := h.authService.Register(c.Context(), req)
	if err != nil {
		if errors.Is(err, service.ErrEmailExists) {
			return Error(c, fiber.StatusConflict, "CONFLICT", err.Error())
		}
		return Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", "Kayıt işlemi başarısız")
	}

	return Success(c, fiber.StatusCreated, result, "Kayıt başarılı")
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return Error(c, fiber.StatusBadRequest, "BAD_REQUEST", "Geçersiz istek formatı")
	}

	if req.Email == "" || req.Password == "" {
		return Error(c, fiber.StatusBadRequest, "VALIDATION_ERROR", "E-posta ve şifre zorunludur")
	}

	result, err := h.authService.Login(c.Context(), req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return Error(c, fiber.StatusUnauthorized, "UNAUTHORIZED", err.Error())
		}
		return Error(c, fiber.StatusInternalServerError, "INTERNAL_ERROR", "Giriş işlemi başarısız")
	}

	return Success(c, fiber.StatusOK, result, "Giriş başarılı")
}
