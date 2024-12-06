package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/yaninyzwitty/graphql-cocroach-go/graph"
	"github.com/yaninyzwitty/graphql-cocroach-go/internal/database"
	"github.com/yaninyzwitty/graphql-cocroach-go/internal/service"
	"github.com/yaninyzwitty/graphql-cocroach-go/pkg"
)

var (
	password string
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	file, err := os.Open("config.yaml")
	if err != nil {
		slog.Error("failed to open config.yaml ")
		os.Exit(1)
	}

	var cfg pkg.Config
	if err := cfg.LoadConfig(file); err != nil {
		slog.Error("failed to load config.yaml")
		os.Exit(1)
	}

	err = godotenv.Load()
	if err != nil {
		slog.Error("failed to load .env file")
		os.Exit(1)
	}

	if s := os.Getenv("DB_PASSWORD"); s != "" {
		password = s
	}

	databaseConfig := &database.DatabaseCfg{
		User:     cfg.Database.User,
		Password: password,
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		Database: cfg.Database.Database,
		SSLMode:  cfg.Database.SSLMode,
	}

	pool, err := database.NewPgxPool(ctx, databaseConfig, cfg.Database.MaxRetries)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)

	}

	defer pool.Close()

	err = database.PingDatabase(ctx, pool)
	if err != nil {
		slog.Error("failed to ping database", "error", err)
		os.Exit(1)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	socialService := service.NewSocialService()

	resolvers := &graph.Resolver{
		DB:            pool,
		SocialService: socialService,
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolvers}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}
	stopCH := make(chan os.Signal, 1)
	signal.Notify(stopCH, os.Interrupt, syscall.SIGTERM)

	// start a server
	go func() {
		slog.Info("server is listening on :" + fmt.Sprintf("%d", cfg.Server.Port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("failed to start server")
			os.Exit(1)
		}

	}()
	<-stopCH
	slog.Info("shuttting down the server...")
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server")
		os.Exit(1)
	} else {
		slog.Info("server stopped down gracefully")

	}

}
