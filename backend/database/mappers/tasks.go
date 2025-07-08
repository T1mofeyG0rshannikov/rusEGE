package mappers

import (
	"fmt"
	"rusEGE/database/models"
	"rusEGE/interfaces"
	"strings"
)

func replaceUpperCaseWithSpaces(s string) string {
	var result strings.Builder
	for _, r := range s {
		_, exists := interfaces.LETTEROPTIONS[r]
		if exists {
			result.WriteRune('_')
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func WordOptions(word *models.Word) []interfaces.Option {
	fmt.Println(word.Rule.Options)
	if len(*word.Rule.Options) > 0 {
		var options []interfaces.Option
		for _, letter := range word.Word {
			if letter >= 'А' && letter <= 'Я' {
				for _, option := range *word.Rule.Options {
					options = append(options, interfaces.Option{
						Correct: option.Letter == string(letter),
						Letter:  option.Letter,
					})
				}
			}
		}

		return options
	} else {
		for _, letter := range word.Word {
			_, exists := interfaces.LETTEROPTIONS[letter]

			if exists {
				return interfaces.LETTEROPTIONS[letter]
			}
		}

		return []interfaces.Option{}
	}
}

func userWordOptions(word *models.UserWord) []interfaces.Option {
	fmt.Println(word.Rule.Options)
	for _, letter := range word.Word {
		_, exists := interfaces.LETTEROPTIONS[letter]

		if exists {
			return interfaces.LETTEROPTIONS[letter]
		}
	}

	return []interfaces.Option{}
}

func DbUserWordToWord(dbWord *models.UserWord) *interfaces.Word {
	return &interfaces.Word{
		Id:        dbWord.Id,
		Word:      replaceUpperCaseWithSpaces(dbWord.Word),
		Rule:      &dbWord.Rule.Rule,
		Options:   userWordOptions(dbWord),
		Exception: dbWord.Exception,
	}
}

func DbWordToWord(dbWord *models.Word) *interfaces.Word {
	return &interfaces.Word{
		Id:        dbWord.Id,
		Word:      replaceUpperCaseWithSpaces(dbWord.Word),
		Rule:      &dbWord.Rule.Rule,
		Options:   WordOptions(dbWord),
		Exception: dbWord.Exception,
	}
}

func DbUserWordsToWords(dbWords []*models.UserWord) []*interfaces.Word {
	words := make([]*interfaces.Word, len(dbWords))

	for i, dbWord := range dbWords {
		words[i] = DbUserWordToWord(dbWord)
	}

	return words
}

func DbWordsToWords(dbWords []*models.Word) []*interfaces.Word {
	words := make([]*interfaces.Word, len(dbWords))

	for i, dbWord := range dbWords {
		words[i] = DbWordToWord(dbWord)
	}

	return words
}
