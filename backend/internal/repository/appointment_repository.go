package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/model"
)

type AppointmentRepository struct {
	db *bun.DB
}

func NewAppointmentRepository(db *bun.DB) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (r *AppointmentRepository) Create(ctx context.Context, appt *model.Appointment) error {
	_, err := r.db.NewInsert().Model(appt).Exec(ctx)
	return err
}

func (r *AppointmentRepository) FindByStudentID(ctx context.Context, studentID string) ([]model.Appointment, error) {
	var appointments []model.Appointment
	err := r.db.NewSelect().
		Model(&appointments).
		Relation("Teacher").
		Where("appointment.student_id = ?", studentID).
		OrderExpr("appointment.date DESC, appointment.time DESC").
		Scan(ctx)
	return appointments, err
}

func (r *AppointmentRepository) FindByID(ctx context.Context, id string) (*model.Appointment, error) {
	appt := new(model.Appointment)
	err := r.db.NewSelect().
		Model(appt).
		Relation("Teacher").
		Where("appointment.id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return appt, nil
}

func (r *AppointmentRepository) Update(ctx context.Context, appt *model.Appointment) error {
	_, err := r.db.NewUpdate().
		Model(appt).
		Column("date", "time").
		Where("id = ?", appt.ID).
		Exec(ctx)
	return err
}

func (r *AppointmentRepository) Cancel(ctx context.Context, id string) error {
	_, err := r.db.NewUpdate().
		Model((*model.Appointment)(nil)).
		Set("status = ?", "cancelled").
		Where("id = ?", id).
		Exec(ctx)
	return err
}

func (r *AppointmentRepository) HasConflict(ctx context.Context, teacherID, date, time, excludeID string) (bool, error) {
	q := r.db.NewSelect().
		Model((*model.Appointment)(nil)).
		Where("teacher_id = ?", teacherID).
		Where("date = ?", date).
		Where("time = ?", time).
		Where("status != ?", "cancelled")

	if excludeID != "" {
		q = q.Where("id != ?", excludeID)
	}

	return q.Exists(ctx)
}
