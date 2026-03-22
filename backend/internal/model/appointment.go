package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Appointment struct {
	bun.BaseModel `bun:"table:appointments"`

	ID        string    `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	StudentID string    `bun:"student_id,notnull,type:uuid" json:"studentId"`
	TeacherID string    `bun:"teacher_id,notnull,type:uuid" json:"teacherId"`
	Date      string    `bun:"date,notnull" json:"date"`
	Time      string    `bun:"time,notnull" json:"time"`
	Status    string    `bun:"status,default:'upcoming'" json:"status"`
	CreatedAt time.Time `bun:"created_at,default:now()" json:"createdAt"`

	// İlişkiler (join için)
	Teacher *Teacher `bun:"rel:belongs-to,join:teacher_id=id" json:"teacher,omitempty"`
}

// Request DTO'lar

type AppointmentCreateRequest struct {
	TeacherID string `json:"teacherId"`
	Date      string `json:"date"`
	Time      string `json:"time"`
}

type AppointmentUpdateRequest struct {
	Date *string `json:"date,omitempty"`
	Time *string `json:"time,omitempty"`
}

// AI Suggest DTO'lar

type SuggestRequest struct {
	LessonName     string   `json:"lessonName"`
	AvailableSlots []string `json:"availableSlots"`
}

type Suggestion struct {
	TeacherID   string  `json:"teacherId"`
	TeacherName string  `json:"teacherName"`
	Date        string  `json:"date"`
	Time        string  `json:"time"`
	Score       float64 `json:"score"`
}

type SuggestResponse struct {
	Suggestions []Suggestion `json:"suggestions"`
}
