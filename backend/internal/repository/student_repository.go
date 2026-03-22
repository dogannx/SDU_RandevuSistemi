package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/model"
)

type StudentRepository struct {
	db *bun.DB
}

func NewStudentRepository(db *bun.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) Create(ctx context.Context, student *model.Student) error {
	_, err := r.db.NewInsert().Model(student).Exec(ctx)
	return err
}

func (r *StudentRepository) FindByEmail(ctx context.Context, email string) (*model.Student, error) {
	student := new(model.Student)
	err := r.db.NewSelect().Model(student).Where("email = ?", email).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (r *StudentRepository) FindByID(ctx context.Context, id string) (*model.Student, error) {
	student := new(model.Student)
	err := r.db.NewSelect().Model(student).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (r *StudentRepository) FindAll(ctx context.Context) ([]model.Student, error) {
	var students []model.Student
	err := r.db.NewSelect().Model(&students).OrderExpr("created_at DESC").Scan(ctx)
	return students, err
}
