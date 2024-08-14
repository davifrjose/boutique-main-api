package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/davifrjose/boutique-main-api/internal/adpater/auth/paseto"
	"github.com/davifrjose/boutique-main-api/internal/adpater/config"
	"github.com/davifrjose/boutique-main-api/internal/adpater/handler/http"
	"github.com/davifrjose/boutique-main-api/internal/adpater/logger"
	"github.com/davifrjose/boutique-main-api/internal/adpater/storage/postgres"
	repository "github.com/davifrjose/boutique-main-api/internal/adpater/storage/postgres/repositories"
	"github.com/davifrjose/boutique-main-api/internal/core/service"
)

//	@title						Go POS (Point of Sale) API
//	@version					1.0
//	@description				This is a simple RESTful Point of Sale (POS) Service API written in Go using Gin web framework, PostgreSQL database, and Redis cache.
//
//	@contact.name				Bagas Hizbullah
//	@contact.url				https://github.com/bagashiz/go-pos
//	@contact.email				bagash.office@simplelogin.com
//
//	@license.name				MIT
//	@license.url				https://github.com/bagashiz/go-pos/blob/main/LICENSE
//
//	@host						gopos.bagashiz.me
//	@BasePath					/v1
//	@schemes					http https
//
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and the access token.
func main() {
	// Load environment variables
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	// Set logger
	logger.Set(config.App)

	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// Init database
	ctx := context.Background()
	db, err := postgres.New(ctx, config.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Successfully connected to the database", "db", config.DB.Connection)

	// Migrate database
	err = db.Migrate()
	if err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully migrated the database")


	slog.Info("Successfully connected to the cache server")

	// Init token service
	token, err := paseto.New(config.Token)
	if err != nil {
		slog.Error("Error initializing token service", "error", err)
		os.Exit(1)
	}

	// Dependency injection
	// User
	userRepo := repository.NewUserRepository(db)
	addressRepo := repository.NewAddressRepository(db)
	userService := service.NewUserService(userRepo,addressRepo)
	userHandler := http.NewUserHandler(userService)

	// Auth
	authService := service.NewAuthService(userRepo, token)
	authHandler := http.NewAuthHandler(authService)

	// Init router
	router, err := http.NewRouter(
		config.HTTP,
		token,
		*userHandler,
		*authHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
