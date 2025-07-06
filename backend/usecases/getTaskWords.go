package usecases

import (
	"rusEGE/database/mappers"
	"rusEGE/database/models"
	"rusEGE/interfaces"
	"rusEGE/repositories"
	"rusEGE/web/schemas"

	"math/rand"
	"time"
)

func shuffleSlice(slice []*interfaces.Word) {
	rand.Seed(time.Now().UnixNano())

	n := len(slice)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

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
		words, err := wr.GetTaskUserWords(task.Id, user.Id)
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

	shuffleSlice(mappedWords)

	return mappedWords, nil
}
