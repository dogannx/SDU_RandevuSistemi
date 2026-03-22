package service

import (
	"context"

	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/model"
	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/repository"
)

type teacherService struct {
	teacherRepo *repository.TeacherRepository
}

func NewTeacherService(teacherRepo *repository.TeacherRepository) TeacherService {
	return &teacherService{teacherRepo: teacherRepo}
}

func (s *teacherService) GetAll(ctx context.Context) ([]model.Teacher, error) {
	return s.teacherRepo.FindAll(ctx)
}
