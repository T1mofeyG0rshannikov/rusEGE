package mappers

import (
	"rusEGE/database/models"
	"rusEGE/interfaces"
)

func DbTaskToTask(task models.Task, rules []*models.Rule) interfaces.Task {
	return interfaces.Task{
		Number:      task.Number,
		Description: task.Description,
		Rules:       DbRulesToTaskRules(rules),
	}
}
