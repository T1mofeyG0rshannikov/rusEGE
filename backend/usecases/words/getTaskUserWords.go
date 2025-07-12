package words

import (
	"rusEGE/database/models"
	"rusEGE/repositories"
)

func GetTaskUserWords(
	uwr *repositories.GormUserWordRepository,
	tr *repositories.GormTaskRepository,
	taskNumber uint,
	user *models.User,
) ([]*models.UserWord, error) {
	task, err := tr.Get(taskNumber)
	if err != nil {
		return nil, err
	}

	words, err := uwr.GetTaskWords(task.Id, user.Id, nil)
	if err != nil {
		return nil, err
	}

	return words, nil
}
