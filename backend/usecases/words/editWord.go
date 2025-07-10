package words

import (
	"rusEGE/database/mappers"
	"rusEGE/interfaces"
	"rusEGE/repositories"
	"rusEGE/web/schemas"
)

func EditWord(
	wr *repositories.GormWordRepository,
	rr *repositories.GormRuleRepository,
	data schemas.EditWordRequest,
) (*interfaces.Word, error) {
	word, err := wr.Get(data.Id)

	if err != nil {
		return nil, err
	}

	if data.Rule != nil {
		rule, err := GetOrCreateRule(*data.Rule)

		if err != nil{
			return nil, err
		}

		word.RuleId = rule.Id
	}

	if data.Word != nil {
		word.Word = *data.Word
	}

	if data.Exception != nil {
		word.Exception = *data.Exception
	}

	if data.Options != nil {
		err := wr.DeleteOptions(word.Id)
		if err != nil {
			return nil, err
		}
		for _, option := range *data.Options {
			_, err := wr.CreateOption(word.Id, option)

			if err != nil {
				return nil, err
			}
		}
	}

	
	word, err = wr.Edit(word)
	if err != nil {
		return nil, err
	}

	word, err = wr.GetWithOptions(word.Id)
	if err != nil{
		return nil, err
	}

	return mappers.DbWordToWord(*word), nil
}
