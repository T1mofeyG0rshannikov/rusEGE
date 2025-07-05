package usecases

import (
	"rusEGE/database"
	"rusEGE/database/mappers"
	"rusEGE/interfaces"
	"rusEGE/repositories"
)

func GetTaskWords(
	taskNumber uint,
	user *database.User,
) ([]interfaces.Word, error) {
	db := database.GetDB()

	tr := repositories.NewGormTaskRepository(db)
	wr := repositories.NewGormWordRepository(db)

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
