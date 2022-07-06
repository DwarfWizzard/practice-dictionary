package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Handler struct {
	DictHandler *DictionaryHandler
}

type HandlerConfig struct {
	DictService DictionaryService
}

func NewHandlers(cfg *HandlerConfig) *Handler {
	return &Handler{
		DictHandler: NewDictionaryHandler(cfg.DictService),
	}
}

func (h *Handler) InitRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/:source/:limit/:offset", h.DictHandler.GetWords)
}
