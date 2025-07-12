package mappers

import (
	"rusEGE/database/models"
	"rusEGE/interfaces"
	"strings"
)


func replaceUpperCaseWithSpaces(s string) string {
	var result strings.Builder
	for _, r := range s {
		if !(r >= 'а' && r <= 'я') && r != ' ' && r != ' ' && r != '(' && r != ')' && r != 'ё' {
			result.WriteRune('_')
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}


func WordOptions(word *models.Word) []interfaces.Option {
	var options []interfaces.Option
	for _, letter := range word.Word {
		if !(letter >= 'а' && letter <= 'я') {
			if word.Options != nil && len(word.Options) > 0 {
				for _, option := range word.Options {
					options = append(options, interfaces.Option{
						Correct: option.Letter == string(letter),
						Letter:  option.Letter,
					})
				}
			} else if word.Rule.Options != nil && len(word.Rule.Options) > 0 {
				for _, option := range word.Rule.Options {
					options = append(options, interfaces.Option{
						Correct: option.Letter == string(letter),
						Letter:  option.Letter,
					})
				}
			} else {
				_, exists := interfaces.LETTEROPTIONS[letter]

				if exists {
					return interfaces.LETTEROPTIONS[letter]
				}
			}
		}
	}
	return options
}

func userWordOptions(word *models.UserWord) []interfaces.Option {
	var options []interfaces.Option
	for _, letter := range word.Word {
		if !(letter >= 'а' && letter <= 'я') {
			if word.Options != nil && len(word.Options) > 0 {
				for _, option := range word.Options {
					options = append(options, interfaces.Option{
						Correct: option.Letter == string(letter),
						Letter:  option.Letter,
					})
				}
			} else if word.Rule.Options != nil && len(word.Rule.Options) > 0 {
				for _, option := range word.Rule.Options {
					options = append(options, interfaces.Option{
						Correct: option.Letter == string(letter),
						Letter:  option.Letter,
					})
				}
			} else {
				_, exists := interfaces.LETTEROPTIONS[letter]

				if exists {
					return interfaces.LETTEROPTIONS[letter]
				}
			}
		}
	}
	return options
}

func DbUserWordToWord(dbWord models.UserWord) *interfaces.Word {
	return &interfaces.Word{
		Id:          dbWord.Id,
		Word:        replaceUpperCaseWithSpaces(dbWord.Word),
		Original:    dbWord.Word,
		Rule:        &dbWord.Rule.Rule,
		Options:     userWordOptions(&dbWord),
		Exception:   dbWord.Exception,
		Description: dbWord.Description,
	}
}

func DbWordToWord(dbWord models.Word) *interfaces.Word {
	return &interfaces.Word{
		Id:          dbWord.Id,
		Word:        replaceUpperCaseWithSpaces(dbWord.Word),
		Original:    dbWord.Word,
		Rule:        &dbWord.Rule.Rule,
		Options:     WordOptions(&dbWord),
		Exception:   dbWord.Exception,
		Description: dbWord.Description,
	}
}

func DbUserWordsToWords(dbWords []*models.UserWord) []*interfaces.Word {
	words := make([]*interfaces.Word, len(dbWords))

	for i, dbWord := range dbWords {
		words[i] = DbUserWordToWord(*dbWord)
	}

	return words
}

func DbWordsToWords(dbWords []*models.Word) []*interfaces.Word {
	words := make([]*interfaces.Word, len(dbWords))

	for i, dbWord := range dbWords {
		words[i] = DbWordToWord(*dbWord)
	}

	return words
}
