package service

import "github.com/DwarfWizzard/practice-dictionary/internal/domain"

type DictionaryStorage interface {
	GetWords(source *string, limit, offset *int) ([]domain.Word, error)
	GetTranslation(source, original *string) ([]domain.Word, error)
}

type DictionaryService struct {
	storage DictionaryStorage
}

func NewDictionaryService(storage DictionaryStorage) *DictionaryService {
	return &DictionaryService{
		storage: storage,
	}
}

func (s *DictionaryService) GetWords(source *string, limit, offset *int) ([]domain.Word, error) {
	return s.storage.GetWords(source, limit, offset)
}

func (s *DictionaryService) GetTranslation(source, original *string) ([]domain.Word, error) {
	return s.storage.GetTranslation(source, original)
}