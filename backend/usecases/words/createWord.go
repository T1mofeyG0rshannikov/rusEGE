package words

import (
	"rusEGE/database/mappers"
	"rusEGE/interfaces"
	"rusEGE/repositories"
	"rusEGE/web/schemas"
)

func CreateWord(
	tr *repositories.GormTaskRepository,
	wr *repositories.GormWordRepository,
	rr *repositories.GormRuleRepository,
	data schemas.CreateWordRequest,
) (*interfaces.Word, error) {
	task, err := tr.Get(data.TaskNumber)
	if err != nil {
		return nil, err
	}

	rule, err := GetOrCreateRule(data.Rule)
	if err != nil {
		return nil, err
	}

	word, err := wr.Create(task.Id, data.Word, rule.Id, &data.Exception, data.Description)

	if err != nil {
		return nil, err
	}

	return mappers.DbWordToWord(*word), nil
}
