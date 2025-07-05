package mappers

import (
	"rusEGE/database"
	"rusEGE/interfaces"
	"strings"
	"unicode"
)

func replaceUpperCaseWithSpaces(s string) string {
	var result strings.Builder
	for _, r := range s {
		if unicode.IsUpper(r) {
			result.WriteRune('_')
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func WordOptions(word string) []interfaces.Option {
	for _, letter := range word {
		if letter >= 'А' && letter <= 'Я' {
			return interfaces.LETTEROPTIONS[letter]
		}
	}

	return []interfaces.Option{}
}

func DbUserWordToWord(dbWords []*database.UserWord) []interfaces.Word {
	words := make([]interfaces.Word, len(dbWords))

	for i, dbWord := range dbWords {
		words[i] = interfaces.Word{
			Word:    replaceUpperCaseWithSpaces(dbWord.Word),
			Rule:    dbWord.Rule,
			Options: WordOptions(dbWord.Word),
		}
	}

	return words
}

func DbWordToWord(dbWords []*database.Word) []interfaces.Word {
	words := make([]interfaces.Word, len(dbWords))

	for i, dbWord := range dbWords {
		words[i] = interfaces.Word{
			Word:    replaceUpperCaseWithSpaces(dbWord.Word),
			Rule:    dbWord.Rule,
			Options: WordOptions(dbWord.Word),
		}
	}

	return words
}
