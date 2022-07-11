package rest

import (
	"log"
	"strconv"

	"github.com/DwarfWizzard/practice-dictionary/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type DictionaryService interface {
	GetWords(source *string, limit, offset *int) ([]domain.Dictionary, error)
	GetTranslation(source, original *string) (domain.Dictionary, error)
}

type DictionaryHandler struct {
	service DictionaryService
}

func NewDictionaryHandler(service DictionaryService) *DictionaryHandler {
	return &DictionaryHandler{
		service: service,
	}
}

func (h *DictionaryHandler) GetWords(c *fiber.Ctx) error {
	source := c.Params("source")
	limit, err1 := strconv.Atoi(c.Params("limit"))
	offset, err2 := strconv.Atoi(c.Params("offset"))

	if err1 != nil && err2 != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(map[string]string{
			"error": "неверные параметры",
		})
	}

	if source != "osetian" && source != "russian" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(map[string]string{
			"error": "неверные параметры",
		})
	}

	words, err := h.service.GetWords(&source, &limit, &offset)
	if err != nil {
		log.Println(err.Error())
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(words)
}

func (h *DictionaryHandler) GetTranslation(c *fiber.Ctx) error {
	source := c.Params("source")
	word := c.Params("word")
	
	if source != "osetian" && source != "russian" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(map[string]string{
			"error": "неверные параметры",
		})
	}
	
	words, err := h.service.GetTranslation(&source, &word)
	if err != nil {
		log.Println(err.Error())
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(map[string]string{
			"error": err.Error(),
		})		
	}

	return c.JSON(words)
}