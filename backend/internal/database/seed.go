package database

import (
	"context"
	"log"

	"github.com/uptrace/bun"
)

type SeedTeacher struct {
	bun.BaseModel `bun:"table:teachers"`

	ID       string   `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Name     string   `bun:"name,notnull"`
	Subjects []string `bun:"subjects,array,notnull"`
	Bio      string   `bun:"bio"`
}

func RunSeed(db *bun.DB) error {
	teachers := []SeedTeacher{
		{
			Name:     "Ahmet Yıldız",
			Subjects: []string{"Matematik", "Geometri"},
			Bio:      "15 yıllık deneyimli Matematik öğretmeni",
		},
		{
			Name:     "Ayşe Kaya",
			Subjects: []string{"Fizik", "Astronomi"},
			Bio:      "SDÜ Fizik Bölümü mezunu, 10 yıllık deneyim",
		},
		{
			Name:     "Mehmet Demir",
			Subjects: []string{"Kimya", "Biyokimya"},
			Bio:      "Organik kimya uzmanı",
		},
		{
			Name:     "Fatma Çelik",
			Subjects: []string{"Biyoloji", "Genetik"},
			Bio:      "Moleküler biyoloji alanında doktora",
		},
		{
			Name:     "Ali Öztürk",
			Subjects: []string{"Türkçe", "Edebiyat"},
			Bio:      "Dil bilimi ve edebiyat uzmanı",
		},
		{
			Name:     "Zeynep Arslan",
			Subjects: []string{"İngilizce"},
			Bio:      "Cambridge sertifikalı İngilizce eğitmeni",
		},
	}

	ctx := context.Background()

	for _, t := range teachers {
		exists, err := db.NewSelect().
			Model((*SeedTeacher)(nil)).
			Where("name = ?", t.Name).
			Exists(ctx)
		if err != nil {
			return err
		}
		if exists {
			continue
		}

		_, err = db.NewInsert().Model(&t).Exec(ctx)
		if err != nil {
			return err
		}
		log.Printf("Öğretmen eklendi: %s", t.Name)
	}

	log.Println("Seed işlemi tamamlandı")
	return nil
}
