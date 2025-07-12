package mappers

import (
	"rusEGE/database/models"
	"rusEGE/interfaces"
)

func DbRuleToTaskRule(rule models.Rule) interfaces.TaskRule {
	return interfaces.TaskRule{
		Id:   rule.Id,
		Rule: rule.Rule,
	}
}


func DbRulesToTaskRules(rules []*models.Rule) []interfaces.TaskRule {
	taskRules := make([]interfaces.TaskRule, len(rules))

	for i, rule := range rules {
		taskRules[i] = DbRuleToTaskRule(*rule)
	}

	return taskRules
}
