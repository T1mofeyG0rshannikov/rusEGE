package tasks

import (
	"rusEGE/database/mappers"
	"rusEGE/interfaces"
	"rusEGE/repositories"
)

func GetTasks(
	tr *repositories.GormTaskRepository,
	rr *repositories.GormRuleRepository,
) ([]interfaces.Task, error) {
	dbTasks, err := tr.All()
	if err != nil {
		return nil, err
	}

	tasks := make([]interfaces.Task, len(dbTasks))
	for i, dbTask := range dbTasks {
		rules, err := rr.GetTaskRules(dbTask.Id)
		
		if err != nil {
			return nil, err
		}

		tasks[i] = mappers.DbTaskToTask(*dbTask, rules)		
	}

	return tasks, nil
}
