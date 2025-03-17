package routes

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/helloDevAman/movie-base/config"
	"github.com/helloDevAman/movie-base/internal/interfaces/handler"
	"github.com/helloDevAman/movie-base/internal/interfaces/middleware"
	"github.com/helloDevAman/movie-base/internal/repository"
	"github.com/helloDevAman/movie-base/internal/usecase"
)

type Route interface {
	LoadRoutes() (*gin.Engine, error)
	StartListning() error
	Shutdown(server *http.Server) error
}

type GinRouteLoader struct {
	cfg *config.Config
	db  *sql.DB
}

func LoadNewGinRoute(config *config.Config, db *sql.DB) *GinRouteLoader {
	return &GinRouteLoader{cfg: config, db: db}
}

// Load routes without storing the router in struct
func (g *GinRouteLoader) LoadRoutes() (*gin.Engine, error) {
	router := gin.Default()
	apiGroup := fmt.Sprintf("/%s/%s", g.cfg.App.Prefix, g.cfg.App.Version)
	api := router.Group(apiGroup)

	// Load API Routes (keep route definitions separate)
	RegisterRoutes(g.cfg, g.db, api)

	return router, nil
}

func RegisterRoutes(cfg *config.Config, db *sql.DB, api *gin.RouterGroup) {
	repo := repository.NewPostgresOTPRepository(db)
	useCase := usecase.NewSendOTPUseCase(repo)
	otpHandler := handler.NewSendOTPHandler(useCase)
	api.POST(sendOTP, middleware.LoggingMiddleware(otpHandler.ServeHTTP))
}

func (g *GinRouteLoader) StartListening(router *gin.Engine) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", g.cfg.App.Port),
		Handler: router,
	}

	// Channel to listen for OS interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Run the server in a goroutine so that it doesn't block
	go func() {
		log.Printf("Server is running on port %s...", g.cfg.App.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Block until we receive a termination signal
	<-quit
	log.Println("Shutdown signal received, shutting down server...")

	// Call Shutdown to gracefully stop the server
	return g.Shutdown(server)
}

func (g *GinRouteLoader) Shutdown(server *http.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	log.Println("Server exited gracefully.")
	return nil
}
