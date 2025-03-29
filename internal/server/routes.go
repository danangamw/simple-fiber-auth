package server

import (
	"test-coding/internal/handler"
	"test-coding/internal/middleware"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (s *FiberServer) RegisterFiberRoutes() {
	// Apply CORS middleware
	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false, // credentials require explicit origins
		MaxAge:           300,
	}))

	s.App.Get("/", handler.PublicHandler)
	s.App.Get("/public", handler.PublicHandler)
	s.App.Get("/protected", middleware.AuthMiddleware(), handler.ProtectedHandler)
}
