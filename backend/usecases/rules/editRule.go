package rules

import (
	"rusEGE/database/models"
	"rusEGE/repositories"
	"rusEGE/web/schemas"
)

func EditRule(
	rr *repositories.GormRuleRepository,
	data schemas.EditRuleRequest,
) (*models.Rule, error) {
	rule, err := rr.GetById(data.Id)
	if err != nil {
		return nil, err
	}

	if data.NewRule != nil {
		rule.Rule = *data.NewRule
	}

	if data.Options != nil {
		err := rr.DeleteOptions(rule.Id)
		if err != nil {
			return nil, err
		}
		for _, option := range *data.Options {
			_, err := rr.CreateOption(&models.RuleOption{
				RuleId: rule.Id,
				Letter: option,
			})

			if err != nil {
				return nil, err
			}
		}
	}

	rule, err = rr.GetWithOptions(rule.Id)

	if err != nil {
		return nil, err
	}

	return rule, err
}
