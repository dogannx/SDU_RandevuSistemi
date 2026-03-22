package service

import (
	"context"

	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/model"
)

type AuthService interface {
	Register(ctx context.Context, req model.RegisterRequest) (*model.AuthResponse, error)
	Login(ctx context.Context, req model.LoginRequest) (*model.AuthResponse, error)
}

type StudentService interface {
	GetByID(ctx context.Context, id string) (*model.Student, error)
	GetAll(ctx context.Context) ([]model.Student, error)
}

type TeacherService interface {
	GetAll(ctx context.Context) ([]model.Teacher, error)
}

type AppointmentService interface {
	Create(ctx context.Context, studentID string, req model.AppointmentCreateRequest) (*model.Appointment, error)
	GetByStudentID(ctx context.Context, studentID string) ([]model.Appointment, error)
	Update(ctx context.Context, id, studentID string, req model.AppointmentUpdateRequest) (*model.Appointment, error)
	Cancel(ctx context.Context, id, studentID string) error
	Suggest(ctx context.Context, req model.SuggestRequest) (*model.SuggestResponse, error)
}
