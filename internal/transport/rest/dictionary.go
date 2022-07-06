package rest

import (
	"strconv"

	"github.com/DwarfWizzard/practice-dictionary/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type DictionaryService interface {
	GetWords(source string, offset, limit int) ([]domain.Word, error)
}

type DictionaryHandler struct {
	DictService DictionaryService
}

func NewDictionaryHandler(dictService DictionaryService) *DictionaryHandler {
	return &DictionaryHandler{
		DictService: dictService,
	}
}

func (d *DictionaryHandler) GetWords(c *fiber.Ctx) error {
	source := c.Params("source")
	offset, _ := strconv.Atoi(c.Params("offset"))
	limit, _ := strconv.Atoi(c.Params("limit"))

	words, err := d.DictService.GetWords(source, offset, limit)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(words)
}
