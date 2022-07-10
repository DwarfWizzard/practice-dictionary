package service

import "github.com/DwarfWizzard/practice-dictionary/internal/domain"

type DictionaryStorage interface {
	GetWords(source *string, limit, offset *int) ([]domain.Dictionary, error)
	GetTranslation(source, original *string) (domain.Dictionary, error)
}

type DictionaryService struct {
	storage DictionaryStorage
}

func NewDictionaryService(storage DictionaryStorage) *DictionaryService {
	return &DictionaryService{
		storage: storage,
	}
}

func (s *DictionaryService) GetWords(source *string, limit, offset *int) ([]domain.Dictionary, error) {
	return s.storage.GetWords(source, limit, offset)
}

func (s *DictionaryService) GetTranslation(source, original *string) (domain.Dictionary, error) {
	return s.storage.GetTranslation(source, original)
}