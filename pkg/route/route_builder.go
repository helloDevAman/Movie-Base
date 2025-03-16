package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/helloDevAman/movie-base/config"
)

type Route interface {
	LoadRoutes() (*gin.Engine, error)
	StartListning() error
	Shutdown(server *http.Server) error
}

type GinRouteLoader struct {
	config *config.Config
}

func LoadNewGinRoute(config *config.Config) *GinRouteLoader {
	return &GinRouteLoader{config: config}
}

// Load routes without storing the router in struct
func (g *GinRouteLoader) LoadRoutes() (*gin.Engine, error) {
	router := gin.Default()
	apiGroup := fmt.Sprintf("/%s/%s", g.config.App.Prefix, g.config.App.Version)
	api := router.Group(apiGroup)

	// Load API Routes (keep route definitions separate)
	RegisterRoutes(api)

	return router, nil
}

func RegisterRoutes(api *gin.RouterGroup) {
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func (g *GinRouteLoader) StartListening(router *gin.Engine) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", g.config.App.Port),
		Handler: router,
	}

	// Channel to listen for OS interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Run the server in a goroutine so that it doesn't block
	go func() {
		log.Printf("Server is running on port %s...", g.config.App.Port)
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
