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

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/yaninyzwitty/graphql-cocroach-go/graph"
	"github.com/yaninyzwitty/graphql-cocroach-go/pkg"
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

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

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
