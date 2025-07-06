package usecases

import (
	"errors"
	"rusEGE/database/models"
	"rusEGE/exceptions"
	"rusEGE/repositories"
	"rusEGE/web/schemas"
)

func EditWord(
	wr *repositories.GormWordRepository,
	data schemas.EditWordRequest,
) (*models.Word, error) {
	word, err := wr.Get(&data.Id)

	if err != nil {
		return nil, err
	}

	if data.Rule != nil {
		var rule *models.Rule
		rule, err := wr.GetRule(*data.Rule)
		if err != nil && !errors.Is(err, exceptions.ErrRuleNotFound) {
			return nil, err
		} else if errors.Is(err, exceptions.ErrRuleNotFound) {
			rule, err = wr.CreateRule(&models.Rule{
				Rule: *data.Rule,
			})

			if err != nil {
				return nil, err
			}
		}

		word.RuleId = rule.Id
	}

	if data.Word != nil {
		word.Word = *data.Word
	}

	if data.Exception != nil {
		word.Exception = *data.Exception
	}

	word, err = wr.Edit(word)
	if err != nil {
		return nil, err
	}

	return word, nil
}
