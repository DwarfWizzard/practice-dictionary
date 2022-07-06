package service

import "github.com/DwarfWizzard/practice-dictionary/internal/domain"

type DictionaryStorage interface {
	GetWords(source *string, offset, limit *int) ([]domain.Word, error)
}

type DictionaryService struct {
	dictStorage DictionaryStorage
}

func NewDictionaryService(dictStorage DictionaryStorage) *DictionaryService {
	return &DictionaryService{
		dictStorage: dictStorage,
	}
}

func (s *DictionaryService) GetWords(source *string, offset, limit *int) ([]domain.Word, error) {
	words, err := s.dictStorage.GetWords(source, offset, limit)
	if err != nil {
		return nil, err
	}
	return words, nil
}