package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"

	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/config"
	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/database"
	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/handler"
	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/repository"
	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/router"
	"github.com/dogannx/SDU_RandevuSistemi/backend/internal/service"
)

func main() {
	migrate := flag.Bool("migrate", false, "Migration'ları çalıştır")
	seed := flag.Bool("seed", false, "Seed verisini ekle")
	flag.Parse()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config yüklenemedi: %v", err)
	}

	db := database.NewPostgres(cfg)
	defer db.Close()

	if *migrate {
		migrationsDir := findMigrationsDir()
		if err := database.RunMigrations(db, migrationsDir); err != nil {
			log.Fatalf("Migration hatası: %v", err)
		}
		log.Println("Migration'lar başarıyla uygulandı")
		if !*seed {
			return
		}
	}

	if *seed {
		if err := database.RunSeed(db); err != nil {
			log.Fatalf("Seed hatası: %v", err)
		}
		return
	}

	// Repositories
	studentRepo := repository.NewStudentRepository(db)
	teacherRepo := repository.NewTeacherRepository(db)
	apptRepo := repository.NewAppointmentRepository(db)

	// Services
	authService := service.NewAuthService(studentRepo, cfg.JWTSecret)
	studentService := service.NewStudentService(studentRepo)
	teacherService := service.NewTeacherService(teacherRepo)
	apptService := service.NewAppointmentService(apptRepo, teacherRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authService)
	studentHandler := handler.NewStudentHandler(studentService)
	teacherHandler := handler.NewTeacherHandler(teacherService)
	apptHandler := handler.NewAppointmentHandler(apptService)

	app := fiber.New(fiber.Config{
		AppName: "Randevu Sistemi API v1",
	})

	router.Setup(app, router.Handlers{
		Auth:        authHandler,
		Student:     studentHandler,
		Teacher:     teacherHandler,
		Appointment: apptHandler,
	}, cfg.JWTSecret)

	log.Printf("Sunucu başlatılıyor: http://localhost:%s", cfg.AppPort)
	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatalf("Sunucu başlatılamadı: %v", err)
	}
}

func findMigrationsDir() string {
	candidates := []string{
		"internal/database/migrations",
		"backend/internal/database/migrations",
	}

	wd, _ := os.Getwd()
	for _, c := range candidates {
		path := filepath.Join(wd, c)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	log.Fatal("Migration klasörü bulunamadı")
	return ""
}
