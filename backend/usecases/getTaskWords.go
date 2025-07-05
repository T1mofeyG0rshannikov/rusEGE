package usecases

import (
	"rusEGE/database/models"
	"rusEGE/database/mappers"
	"rusEGE/interfaces"
	"rusEGE/repositories"
)

func GetTaskWords(
	tr *repositories.GormTaskRepository,
	wr *repositories.GormWordRepository,
	taskNumber uint,
	user *models.User,
) ([]interfaces.Word, error) {
	task, err := tr.Get(taskNumber)

	if err != nil {
		return nil, err
	}

	if user != nil {
		words, err := wr.GetTaskUserWords(task.Id, user.Id)
		if err != nil {
			return nil, err
		}

		return mappers.DbUserWordToWord(words), err
	} else {
		words, err := wr.GetTaskWords(task.Id)
		if err != nil {
			return nil, err
		}

		return mappers.DbWordToWord(words), err
	}
}
