package dictionary

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	httpServer *fiber.App
}

func (s *Server) Run(port string, createRouter func(app *fiber.App)) error {
	s.httpServer = fiber.New()
	createRouter(s.httpServer)
	return s.httpServer.Listen(port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown()
}
