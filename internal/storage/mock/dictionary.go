package mock

import (
	"math/rand"
	"strings"

	"github.com/DwarfWizzard/practice-dictionary/internal/domain"
)

type DictionaryMock struct{}

func NewDictionaryMock() *DictionaryMock {
	return &DictionaryMock{}
}

func (s *DictionaryMock) GetWords(source *string, offset, limit *int) ([]domain.Word, error) {
	var words []domain.Word

	var seed int64
	for _, r := range *source {
		seed += int64(r)
	}
	rand.Seed(seed + int64(*offset))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")

	var original strings.Builder
	var translation strings.Builder

	for i := 0; i < *limit; i++ {
		var word domain.Word

		for j := 0; j < 5; j++ {
			original.WriteRune(chars[rand.Intn(len(chars))])
			translation.WriteRune(chars[rand.Intn(len(chars))])
		}

		word.Id = i + *offset
		word.Original = original.String()
		word.Translation = translation.String()

		words = append(words, word)

		original.Reset()
		translation.Reset()
	}

	return words, nil
}

func (s *DictionaryMock) GetTranslation(source, original *string) ([]domain.Word, error) {
	var words []domain.Word

	var seed int64
	for _, r := range *source {
		seed += int64(r)
	}
	for _, r := range *original {
		seed += int64(r)
	}

	rand.Seed(seed)

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö")
		
	var translation strings.Builder
	for i:=0; i<rand.Intn(10); i++ {
		var word domain.Word

		for j:=0; j<3+rand.Intn(20); j++ {
			translation.WriteRune(chars[rand.Intn(len(chars))])
		}

		word.Id = rand.Intn(10)*10+i
		word.Original = *original
		word.Translation = translation.String()

		translation.Reset()
		words = append(words, word)
	}
	return words, nil
}