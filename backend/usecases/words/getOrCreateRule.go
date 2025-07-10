package words

import (
	"errors"
	"rusEGE/database/models"
	"rusEGE/exceptions"
	"rusEGE/repositories"
)

func GetOrCreateRule(ruleContent string) (*models.Rule, error) {
	rr := repositories.NewGormRuleRepository(nil)
	var rule *models.Rule
	
	rule, err := rr.Get(ruleContent)
	if err != nil && !errors.Is(err, exceptions.ErrRuleNotFound) {
		return nil, err
	} else if errors.Is(err, exceptions.ErrRuleNotFound) {
		rule, err = rr.Create(ruleContent)

		if err != nil {
			return nil, err
		}
	}

	return rule, err
}