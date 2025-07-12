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
	uwr *repositories.GormUserWordRepository,
	data schemas.GetWordsRequest,
	user *models.User,
) ([]*interfaces.Word, error) {
	task, err := tr.Get(data.Task)

	var mappedWords []*interfaces.Word

	if err != nil {
		return nil, err
	}

	if user != nil {
		words, err := uwr.GetTaskWords(task.Id, user.Id, data.RuleIds)
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
