package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Handler struct {
	dictHandler *DictionaryHandler
}

type HandlerConfig struct {
	dictService Dictionary
}

func CreateHandlers(cfg HandlerConfig) *Handler {
	return &Handler{
		dictHandler: NewDictionaryHandler(cfg.dictService),
	}
}

func (h *Handler) InitRoutes() *fiber.App {
	router := fiber.New()

	api := router.Group("/api", logger.New())
	{
		api.Get("/:source/:num", h.dictHandler.GetWords)
	}

	return router
}
