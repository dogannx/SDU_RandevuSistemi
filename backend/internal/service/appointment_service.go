package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/model"
	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/repository"
)

var (
	ErrAppointmentNotFound = errors.New("randevu bulunamadı")
	ErrNotOwner            = errors.New("bu randevu size ait değil")
	ErrTimeConflict        = errors.New("seçilen saat dolu veya başka bir randevuyla çakışıyor")
	ErrAlreadyCancelled    = errors.New("randevu zaten iptal edilmiş")
)

type appointmentService struct {
	apptRepo    *repository.AppointmentRepository
	teacherRepo *repository.TeacherRepository
}

func NewAppointmentService(apptRepo *repository.AppointmentRepository, teacherRepo *repository.TeacherRepository) AppointmentService {
	return &appointmentService{
		apptRepo:    apptRepo,
		teacherRepo: teacherRepo,
	}
}

func (s *appointmentService) Create(ctx context.Context, studentID string, req model.AppointmentCreateRequest) (*model.Appointment, error) {
	_, err := s.teacherRepo.FindByID(ctx, req.TeacherID)
	if err != nil {
		return nil, errors.New("öğretmen bulunamadı")
	}

	conflict, err := s.apptRepo.HasConflict(ctx, req.TeacherID, req.Date, req.Time, "")
	if err != nil {
		return nil, err
	}
	if conflict {
		return nil, ErrTimeConflict
	}

	appt := &model.Appointment{
		StudentID: studentID,
		TeacherID: req.TeacherID,
		Date:      req.Date,
		Time:      req.Time,
		Status:    "upcoming",
	}

	if err := s.apptRepo.Create(ctx, appt); err != nil {
		return nil, err
	}

	return s.apptRepo.FindByID(ctx, appt.ID)
}

func (s *appointmentService) GetByStudentID(ctx context.Context, studentID string) ([]model.Appointment, error) {
	return s.apptRepo.FindByStudentID(ctx, studentID)
}

func (s *appointmentService) Update(ctx context.Context, id, studentID string, req model.AppointmentUpdateRequest) (*model.Appointment, error) {
	appt, err := s.apptRepo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrAppointmentNotFound
	}

	if appt.StudentID != studentID {
		return nil, ErrNotOwner
	}

	if appt.Status == "cancelled" {
		return nil, ErrAlreadyCancelled
	}

	if req.Date != nil {
		appt.Date = *req.Date
	}
	if req.Time != nil {
		appt.Time = *req.Time
	}

	conflict, err := s.apptRepo.HasConflict(ctx, appt.TeacherID, appt.Date, appt.Time, appt.ID)
	if err != nil {
		return nil, err
	}
	if conflict {
		return nil, ErrTimeConflict
	}

	if err := s.apptRepo.Update(ctx, appt); err != nil {
		return nil, err
	}

	return s.apptRepo.FindByID(ctx, appt.ID)
}

func (s *appointmentService) Cancel(ctx context.Context, id, studentID string) error {
	appt, err := s.apptRepo.FindByID(ctx, id)
	if err != nil {
		return ErrAppointmentNotFound
	}

	if appt.StudentID != studentID {
		return ErrNotOwner
	}

	if appt.Status == "cancelled" {
		return ErrAlreadyCancelled
	}

	return s.apptRepo.Cancel(ctx, id)
}

func (s *appointmentService) Suggest(ctx context.Context, req model.SuggestRequest) (*model.SuggestResponse, error) {
	allTeachers, err := s.teacherRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var matchingTeachers []model.Teacher
	for _, t := range allTeachers {
		for _, subj := range t.Subjects {
			if subj == req.LessonName {
				matchingTeachers = append(matchingTeachers, t)
				break
			}
		}
	}

	if len(matchingTeachers) == 0 {
		return &model.SuggestResponse{Suggestions: []model.Suggestion{}}, nil
	}

	dayMap := map[string]time.Weekday{
		"Pazartesi": time.Monday,
		"Salı":      time.Tuesday,
		"Çarşamba":  time.Wednesday,
		"Perşembe":  time.Thursday,
		"Cuma":      time.Friday,
	}

	var suggestions []model.Suggestion
	score := 0.95

	for _, slot := range req.AvailableSlots {
		dayName, timeRange := splitSlot(slot)
		if dayName == "" {
			continue
		}

		targetDay, ok := dayMap[dayName]
		if !ok {
			continue
		}

		startHour, endHour := parseTimeRange(timeRange)
		if startHour < 0 {
			continue
		}

		date := nextWeekday(targetDay)

		for hour := startHour; hour < endHour; hour++ {
			timeStr := fmt.Sprintf("%02d:00", hour)

			for _, teacher := range matchingTeachers {
				conflict, _ := s.apptRepo.HasConflict(ctx, teacher.ID, date, timeStr, "")
				if conflict {
					continue
				}

				suggestions = append(suggestions, model.Suggestion{
					TeacherID:   teacher.ID,
					TeacherName: teacher.Name,
					Date:        date,
					Time:        timeStr,
					Score:       score,
				})
				score -= 0.05
				if score < 0.5 {
					score = 0.5
				}

				if len(suggestions) >= 5 {
					return &model.SuggestResponse{Suggestions: suggestions}, nil
				}
			}
		}
	}

	return &model.SuggestResponse{Suggestions: suggestions}, nil
}

func splitSlot(slot string) (string, string) {
	for i := len(slot) - 1; i >= 0; i-- {
		if slot[i] == ' ' {
			return slot[:i], slot[i+1:]
		}
	}
	return "", ""
}

func parseTimeRange(tr string) (int, int) {
	// "09:00-12:00" → 9, 12
	var startH, endH int
	_, err := fmt.Sscanf(tr, "%d:00-%d:00", &startH, &endH)
	if err != nil {
		return -1, -1
	}
	return startH, endH
}

func nextWeekday(target time.Weekday) string {
	now := time.Now()
	current := now.Weekday()
	daysUntil := int(target) - int(current)
	if daysUntil <= 0 {
		daysUntil += 7
	}
	return now.AddDate(0, 0, daysUntil).Format("2006-01-02")
}
