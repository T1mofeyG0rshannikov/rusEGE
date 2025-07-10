package rules

import (
	"rusEGE/database/models"
	"rusEGE/interfaces"
	"rusEGE/repositories"
)

func GetRulesStat(
	tr *repositories.GormTaskRepository,
	rr *repositories.GormRuleRepository,
	taskNumber uint,
	user *models.User,
) (*[]interfaces.Rule, error) {
	task, err := tr.Get(taskNumber)
	if err != nil{
		return nil, err
	}

	rules, err := tr.GetTaskRules(task.Id)
	taskRules := make([]interfaces.Rule, len(rules))

	if err != nil {
		return nil, err
	}

	if len(rules) > 0 && user != nil {
		for i, dbRule := range rules {
			ruleErrors, err := rr.GetErrorsCount(dbRule.Id, user.Id)

			if err != nil {
				return nil, err
			}

			taskRules[i] = interfaces.Rule{
				Id:     dbRule.Id,
				Rule:   dbRule.Rule,
				Errors: *ruleErrors,
			}
		}
	} 

	return &taskRules, nil
}
