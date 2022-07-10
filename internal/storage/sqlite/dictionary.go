package sqlite

import (
	"fmt"
	"log"
	"net/url"

	"github.com/DwarfWizzard/practice-dictionary/internal/domain"
	"github.com/jmoiron/sqlx"
)

type DictionaryStorage struct {
	db *sqlx.DB
}

func NewDictionaryStorage(db *sqlx.DB) *DictionaryStorage {
	return &DictionaryStorage{
		db: db,
	}
}

type word struct {
	Id int `db:"id"`
	Word string `db:"word"`
}

func (s *DictionaryStorage) GetWords(source *string, offset, limit *int) ([]domain.Dictionary, error) {
	var dictionary []domain.Dictionary

	var words []word
	query := fmt.Sprintf("SELECT id, word FROM %s_words LIMIT %d OFFSET %d", *source, *limit, *offset)
	log.Print(query)
	err := s.db.Select(&words, query)
	if err != nil {
		return dictionary, err
	}

	for _, v := range words {
		var word domain.Dictionary
		var translation []string

		word.Id = v.Id
		word.Original = v.Word

		var query string
		switch *source {
		case "osetian":
			query = fmt.Sprintf("SELECT rw.word FROM osetian_russian AS os_ru INNER JOIN russian_words AS rw ON rw.id=os_ru.translation_id INNER JOIN osetian_words AS ow ON ow.id=os_ru.original_id WHERE ow.word = \"%s\"", v.Word)
		case "russian":
			query = fmt.Sprintf("SELECT ow.word FROM russian_osetian AS ru_os INNER JOIN russian_words AS rw ON rw.id=ru_os.original_id INNER JOIN osetian_words AS ow ON ow.id=ru_os.translation_id WHERE rw.word = \"%s\"", v.Word)
		}

		err := s.db.Select(&translation, query)
		if err != nil {
			return dictionary, err
		}

		word.Translation = translation
		dictionary = append(dictionary, word)
	}

	return dictionary, nil
}

func (s *DictionaryStorage) GetTranslation(source, original *string) (domain.Dictionary, error) {
	var dictionary domain.Dictionary

	decoded, err := url.QueryUnescape(*original)
	if err != nil {
		return dictionary, err
	}

	var query string
	switch *source {
	case "osetian":
		query = fmt.Sprintf("SELECT ow.id AS id, rw.word AS translation FROM osetian_russian AS os_ru INNER JOIN russian_words AS rw ON rw.id=os_ru.translation_id INNER JOIN osetian_words AS ow ON ow.id=os_ru.original_id WHERE ow.word = \"%s\"", decoded)
	case "russian":
		query = fmt.Sprintf("SELECT rw.id AS id, ow.word AS translation  FROM russian_osetian AS ru_os INNER JOIN russian_words AS rw ON rw.id=ru_os.original_id INNER JOIN osetian_words AS ow ON ow.id=ru_os.translation_id WHERE rw.word = \"%s\"", decoded)
	}
	
	dictionary.Original = decoded
	rows, err := s.db.Queryx(query)
	if err != nil {
		return dictionary, err
	}
	defer rows.Close()
	for rows.Next() {
		var word string
		var id int
		err := rows.Scan(&id, &word)
		if err != nil {
			return dictionary, err
		}

		dictionary.Id = id
		dictionary.Translation = append(dictionary.Translation, word)
	}

	return dictionary, nil
}
