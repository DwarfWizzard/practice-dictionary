package mock

import (
	"math/rand"
	"strings"

	"github.com/DwarfWizzard/practice-dictionary/internal/domain"
)

type DictionaryMock struct {}

func NewDictionaryMock() *DictionaryMock {
	return &DictionaryMock{}
}

func (s *DictionaryMock) GetWords(source string, offset, limit int) ([]domain.Word, error) {
	var words []domain.Word

	rand.Seed(int64(offset))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
    "abcdefghijklmnopqrstuvwxyzåäö" +
    "0123456789")

	var original strings.Builder
	var translation strings.Builder

	for i := 0; i < limit; i++ {
		var word domain.Word

		for j := 0; j<5; j++ {
			original.WriteRune(chars[rand.Intn(len(chars))])
			translation.WriteRune(chars[rand.Intn(len(chars))])
		}

		word.Original = original.String()
		word.Translation = translation.String()

		words = append(words, word)

		original.Reset()
		translation.Reset()
	}

	return words, nil
}