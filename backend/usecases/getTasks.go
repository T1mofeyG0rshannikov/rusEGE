package usecases

import (
	"rusEGE/database/models"
	"rusEGE/interfaces"
	"rusEGE/repositories"
)

func GetTasks(
	tr *repositories.GormTaskRepository,
	rr *repositories.GormRuleRepository,
	user *models.User,
) ([]*interfaces.Task, error) {
	dbTasks, err := tr.GetAll()
	if err != nil {
		return nil, err
	}

	tasks := make([]*interfaces.Task, len(dbTasks))
	for i, dbTask := range dbTasks {

		rules, err := tr.GetTaskRules(dbTask.Id)
		taskRules := make([]interfaces.TaskRule, len(rules))

		if err != nil {
			return nil, err
		}

		if len(rules) > 0 && user != nil {
			for i, dbRule := range rules {
				ruleErrors, err := rr.GetErrorsCount(dbRule.Id, user.Id)

				if err != nil {
					return nil, err
				}

				taskRules[i] = interfaces.TaskRule{
					Id:     dbRule.Id,
					Rule:   dbRule.Rule,
					Errors: ruleErrors,
				}
			}

			tasks[i] = &interfaces.Task{
				Number:      dbTask.Number,
				Description: dbTask.Description,
				Rules:       &taskRules,
			}
		} else {
			for i, dbRule := range rules {
				taskRules[i] = interfaces.TaskRule{
					Id:   dbRule.Id,
					Rule: dbRule.Rule,
				}
			}

			tasks[i] = &interfaces.Task{
				Number:      dbTask.Number,
				Description: dbTask.Description,
				Rules:       &taskRules,
			}
		}
	}

	return tasks, nil
}
