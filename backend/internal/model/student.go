package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Student struct {
	bun.BaseModel `bun:"table:students"`

	ID        string    `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	Name      string    `bun:"name,notnull" json:"name"`
	Email     string    `bun:"email,notnull,unique" json:"email"`
	Password  string    `bun:"password,notnull" json:"-"`
	CreatedAt time.Time `bun:"created_at,default:now()" json:"createdAt"`
}

// Request DTO'lar

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Response DTO'lar

type AuthResponse struct {
	Token     string  `json:"token"`
	ExpiresIn int     `json:"expiresIn"`
	Student   Student `json:"student"`
}

type StudentResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

func (s *Student) ToResponse() StudentResponse {
	return StudentResponse{
		ID:        s.ID,
		Name:      s.Name,
		Email:     s.Email,
		CreatedAt: s.CreatedAt,
	}
}
