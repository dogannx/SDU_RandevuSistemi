package service

import (
	"context"

	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/model"
	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/repository"
)

type studentService struct {
	studentRepo *repository.StudentRepository
}

func NewStudentService(studentRepo *repository.StudentRepository) StudentService {
	return &studentService{studentRepo: studentRepo}
}

func (s *studentService) GetByID(ctx context.Context, id string) (*model.Student, error) {
	return s.studentRepo.FindByID(ctx, id)
}

func (s *studentService) GetAll(ctx context.Context) ([]model.Student, error) {
	return s.studentRepo.FindAll(ctx)
}
