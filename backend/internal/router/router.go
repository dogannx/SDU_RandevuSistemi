package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/handler"
	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/middleware"
)

type Handlers struct {
	Auth        *handler.AuthHandler
	Student     *handler.StudentHandler
	Teacher     *handler.TeacherHandler
	Appointment *handler.AppointmentHandler
}

func Setup(app *fiber.App, h Handlers, jwtSecret string) {
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	api := app.Group("/api/v1")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Auth routes (public)
	auth := api.Group("/auth")
	auth.Post("/register", h.Auth.Register)
	auth.Post("/login", h.Auth.Login)

	// Teacher routes (public)
	api.Get("/teachers", h.Teacher.GetAll)

	// Protected routes
	protected := api.Group("", middleware.AuthRequired(jwtSecret))

	// Student routes
	protected.Get("/students", h.Student.GetAll)
	protected.Get("/students/:id", h.Student.GetByID)

	// Appointment routes
	protected.Post("/appointments", h.Appointment.Create)
	protected.Get("/appointments", h.Appointment.GetMy)
	protected.Put("/appointments/:id", h.Appointment.Update)
	protected.Delete("/appointments/:id", h.Appointment.Cancel)
	protected.Post("/appointments/suggest", h.Appointment.Suggest)
}
