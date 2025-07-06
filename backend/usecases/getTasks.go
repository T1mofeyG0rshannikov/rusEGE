package usecases

import (
	"rusEGE/repositories"
	"rusEGE/interfaces"
)

func GetTasks(
	tr *repositories.GormTaskRepository,
) ([]*interfaces.Task, error) {
	dbTasks, err := tr.GetAll()
	if err != nil{
		return nil, err
	}

	tasks := make([]*interfaces.Task, len(dbTasks))
	for i, dbTask := range dbTasks {
		rules, err := tr.GetTaskRules(dbTask.Id)
		if err != nil{
			return nil, err
		}

		tasks[i] = &interfaces.Task{
			Number: dbTask.Number,
			Description: dbTask.Description,
			Rules: rules,
		}
	}

	return tasks, nil
}