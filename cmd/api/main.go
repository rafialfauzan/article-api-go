package main

import (
	"golang-api/internal/config"
	"golang-api/internal/handler"
	"golang-api/internal/repository"
	"golang-api/internal/routes"
	"golang-api/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load konfigurasi
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Inisialisasi database
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Inisialisasi repository, service, dan handler
	articleRepo := repository.NewArticleRepository(db)
	articleService := service.NewArticleService(articleRepo)
	articleHandler := handler.NewArticleHandler(articleService)

	// Inisialisasi repository, service, dan handler untuk User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Setup router
	router := gin.Default()

	// Grup rute API
	api := router.Group("/api")
	{
		// Setup rute untuk Article
		routes.SetupArticleRoutes(api, articleHandler)
		// Setup rute untuk User
		routes.SetupUserRoutes(api, userHandler)
	}

	// Jalankan server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
