package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/model"
)

type TeacherRepository struct {
	db *bun.DB
}

func NewTeacherRepository(db *bun.DB) *TeacherRepository {
	return &TeacherRepository{db: db}
}

func (r *TeacherRepository) FindAll(ctx context.Context) ([]model.Teacher, error) {
	var teachers []model.Teacher
	err := r.db.NewSelect().Model(&teachers).OrderExpr("name ASC").Scan(ctx)
	return teachers, err
}

func (r *TeacherRepository) FindByID(ctx context.Context, id string) (*model.Teacher, error) {
	teacher := new(model.Teacher)
	err := r.db.NewSelect().Model(teacher).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}
