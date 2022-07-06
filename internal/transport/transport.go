package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Transport struct {
	dt *DictionaryTransport
}

type TransportConfig struct {
	ds DictionaryService
}

func CreateHandlers(cfg TransportConfig) *Transport {
	return &Transport{
		dt: NewDictionaryTransport(&cfg.ds),
	}
}

func (t *Transport) InitRoutes() *fiber.App {
	router := fiber.New()

	api := router.Group("/api", logger.New())
	{
		api.Get("/:dict/:num", t.dt.GetTranlsations)
	}

	return router
}
