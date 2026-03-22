package service

import (
	"context"
	"errors"

	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/model"
	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/repository"
	"github.com/dogannx/SDU_RandevuSistemi/backend/pkg/hash"
	"github.com/dogannx/SDU_RandevuSistemi/backend/pkg/jwt"
)

var (
	ErrEmailExists      = errors.New("bu e-posta adresi zaten kullanımda")
	ErrInvalidCredentials = errors.New("e-posta veya şifre hatalı")
)

type authService struct {
	studentRepo *repository.StudentRepository
	jwtSecret   string
}

func NewAuthService(studentRepo *repository.StudentRepository, jwtSecret string) AuthService {
	return &authService{
		studentRepo: studentRepo,
		jwtSecret:   jwtSecret,
	}
}

func (s *authService) Register(ctx context.Context, req model.RegisterRequest) (*model.AuthResponse, error) {
	existing, _ := s.studentRepo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, ErrEmailExists
	}

	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	student := &model.Student{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := s.studentRepo.Create(ctx, student); err != nil {
		return nil, err
	}

	token, err := jwt.GenerateAccessToken(student.ID, student.Email, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		Token:     token,
		ExpiresIn: 900,
		Student:   *student,
	}, nil
}

func (s *authService) Login(ctx context.Context, req model.LoginRequest) (*model.AuthResponse, error) {
	student, err := s.studentRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !hash.CheckPassword(req.Password, student.Password) {
		return nil, ErrInvalidCredentials
	}

	token, err := jwt.GenerateAccessToken(student.ID, student.Email, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		Token:     token,
		ExpiresIn: 900,
		Student:   *student,
	}, nil
}
