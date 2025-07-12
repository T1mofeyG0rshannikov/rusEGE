package words

import (
	"errors"
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
		parts := strings.Split(line, "; ")
		if len(parts) > 1 {
			wordString := strings.TrimSpace(parts[0])
			ruleString := strings.TrimSpace(parts[1])
			description := strings.TrimSpace(parts[2])
			exception := strings.Contains(ruleString, "(искл)")

			ruleString = strings.Replace(ruleString, " (искл)", "", 1)

			rule, err := GetOrCreateRule(ruleString)

			if err != nil {
				return err
			}

			_, err = wr.Create(wordString, task.Id, rule.Id, &exception, &description)

			if err != nil && !errors.Is(err, exceptions.ErrWordAlreadyExists) {
				return err
			}
		}
	}

	return nil
}
