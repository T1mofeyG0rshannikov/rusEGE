package words

import (
	"errors"
	"rusEGE/database/models"
	"rusEGE/exceptions"
	"rusEGE/interfaces"
	"rusEGE/repositories"
	"rusEGE/web/schemas"
	"rusEGE/database/mappers"
)

func CreateWord(
	tr *repositories.GormTaskRepository,
	wr *repositories.GormWordRepository,
	rr *repositories.GormRuleRepository,
	data schemas.CreateWordRequest,
) (*interfaces.Word, error){
	task, err := tr.Get(data.TaskNumber)
	if err != nil {
		return nil, err
	}

	var rule *models.Rule
	rule, err = rr.Get(data.Rule)

	if err != nil && !errors.Is(err, exceptions.ErrRuleNotFound){
		return nil, err
	} else if errors.Is(err, exceptions.ErrRuleNotFound){
		rule, err = rr.Create(&models.Rule{
			Rule: data.Rule,
		})

		if err != nil{
			return nil, err
		}
	}

	word, err := wr.Create(&models.Word{
		TaskId: task.Id,
		Word:   data.Word,
		RuleId:   rule.Id,
	})

	if err != nil{
		return nil, err
	}

	return mappers.DbWordToWord(word), nil
}