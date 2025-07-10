package words

import (
	"errors"
	"rusEGE/database/models"
	"rusEGE/exceptions"
	"rusEGE/repositories"
	"rusEGE/web/schemas"
	"strings"
)

func BulkCreateWord(
	tr *repositories.GormTaskRepository,
	wr *repositories.GormWordRepository,
	rr *repositories.GormRuleRepository,
	data schemas.BulkCreateWordRequest,
) error {
	task, err := tr.Get(data.TaskNumber)
	if err != nil {
		return err
	}

	lines := strings.Split(data.Content, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ", ")
		if len(parts) > 1 {
			wordString := strings.TrimSpace(parts[0])
			ruleString := strings.TrimSpace(parts[1])
			exception := strings.Contains(ruleString, "(искл)")

			ruleString = strings.Replace(ruleString, " (искл)", "", 1)

			var rule *models.Rule
			rule, err = rr.Get(ruleString)

			if err != nil && !errors.Is(err, exceptions.ErrRuleNotFound) {
				return err
			} else if errors.Is(err, exceptions.ErrRuleNotFound) {
				rule, err = rr.Create(&models.Rule{
					Rule: ruleString,
				})

				if err != nil {
					return err
				}
			}

			_, err := wr.Create(&models.Word{
				TaskId:    task.Id,
				Word:      wordString,
				RuleId:    rule.Id,
				Exception: exception,
			})

			if err != nil && !errors.Is(err, exceptions.ErrWordAlreadyExists) {
				return err
			}
		}
	}

	return nil
}
