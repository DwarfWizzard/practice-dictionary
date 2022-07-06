package transport

import (
	"github.com/DwarfWizzard/practice-dictionary/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type Dictionary interface {
	GetWords() ([]domain.Word, error)
}

type DictionaryHandler struct {
	dictService Dictionary
}

func NewDictionaryHandler(dictService Dictionary) *DictionaryHandler {
	return &DictionaryHandler{
		dictService: dictService,
	}
}

func (d *DictionaryHandler) GetWords(c *fiber.Ctx) error{
	return nil
}

