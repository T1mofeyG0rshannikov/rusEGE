package words

import (
	"rusEGE/database/mappers"
	"rusEGE/database/models"
	"rusEGE/interfaces"
	"rusEGE/repositories"
	"rusEGE/web/schemas"
	utils "rusEGE/usecases"
)


func GetTaskWords(
	tr *repositories.GormTaskRepository,
	wr *repositories.GormWordRepository,
	data schemas.GetWordsRequest,
	user *models.User,
) ([]*interfaces.Word, error) {
	task, err := tr.Get(data.Task)

	var mappedWords []*interfaces.Word

	if err != nil {
		return nil, err
	}

	if user != nil {
		words, err := wr.GetTaskUserWords(task.Id, user.Id, data.RuleIds)
		if err != nil {
			return nil, err
		}

		mappedWords = mappers.DbUserWordsToWords(words)
	} else {
		words, err := wr.GetTaskWords(task.Id, data.RuleIds)
		if err != nil {
			return nil, err
		}

		mappedWords = mappers.DbWordsToWords(words)
	}

	utils.ShuffleSlice(mappedWords)

	return mappedWords, nil
}
