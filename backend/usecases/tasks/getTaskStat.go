package tasks

import (
	"rusEGE/database/models"
	"rusEGE/repositories"
	"strings"
)

func GetTaskStat(
	taskNumber uint,
	tr *repositories.GormTaskRepository,
	wr *repositories.GormWordRepository,
	user *models.User,
) (*[]map[string]interface{}, error) {
	task, err := tr.Get(taskNumber)

	if err != nil {
		return nil, err
	}

	words, err := wr.GetTaskUserWords(task.Id, user.Id, nil)

	if err != nil {
		return nil, err
	}

	var stat []map[string]interface{}

	for _, word := range words {
		errors, err := wr.GetUserWordErrors(word.Id)
		if err != nil {
			return nil, err
		}

		if len(*errors) != 0 {
			stat = append(stat, map[string]interface{}{
				"word":   strings.ToLower(word.Word),
				"errors": len(*errors),
			})
		}
	}

	return &stat, nil
}
