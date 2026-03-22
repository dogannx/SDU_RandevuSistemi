package model

import "github.com/uptrace/bun"

type Teacher struct {
	bun.BaseModel `bun:"table:teachers"`

	ID       string   `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	Name     string   `bun:"name,notnull" json:"name"`
	Subjects []string `bun:"subjects,array,notnull" json:"subjects"`
	Bio      *string  `bun:"bio" json:"bio,omitempty"`
}
