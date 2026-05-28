package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/uptrace/bun"

	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/cache"
	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/model"
)

const (
	teachersAllCacheKey = "teachers:all"
	teachersCacheTTL    = 5 * time.Minute
)

type TeacherRepository struct {
	db    *bun.DB
	cache *cache.Cache
}

func NewTeacherRepository(db *bun.DB, c *cache.Cache) *TeacherRepository {
	return &TeacherRepository{db: db, cache: c}
}

func (r *TeacherRepository) FindAll(ctx context.Context) ([]model.Teacher, error) {
	var teachers []model.Teacher

	if err := r.cache.GetJSON(ctx, teachersAllCacheKey, &teachers); err == nil {
		log.Printf("[CACHE HIT] %s (%d öğretmen)", teachersAllCacheKey, len(teachers))
		return teachers, nil
	} else if !errors.Is(err, cache.ErrCacheMiss) {
		log.Printf("[CACHE GET ERROR] %v", err)
	}

	log.Printf("[CACHE MISS] %s → DB sorgusu", teachersAllCacheKey)
	err := r.db.NewSelect().Model(&teachers).OrderExpr("name ASC").Scan(ctx)
	if err != nil {
		return nil, err
	}

	if setErr := r.cache.SetJSON(ctx, teachersAllCacheKey, teachers, teachersCacheTTL); setErr != nil {
		log.Printf("[CACHE SET ERROR] %v", setErr)
	}

	return teachers, nil
}

func (r *TeacherRepository) FindByID(ctx context.Context, id string) (*model.Teacher, error) {
	teacher := new(model.Teacher)
	err := r.db.NewSelect().Model(teacher).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}
