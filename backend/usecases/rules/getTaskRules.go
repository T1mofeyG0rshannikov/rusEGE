package rules

import (
	"rusEGE/database/models"
	"rusEGE/repositories"
)

func GetTaskRules(
	tr *repositories.GormTaskRepository,
	rr *repositories.GormRuleRepository,
	taskNumber uint,
) ([]*models.Rule, error) {
	task, err := tr.Get(taskNumber)
	if err != nil{
		return nil, err
	}

	rules, err := rr.GetTaskRules(task.Id)

	if err != nil {
		return nil, err
	}

	return rules, nil
}
