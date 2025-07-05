package usecases

import (
	"rusEGE/database/models"
	"rusEGE/database/mappers"
	"rusEGE/interfaces"
	"rusEGE/repositories"

	"math/rand"
	"time"
)


func shuffleSlice(slice []interfaces.Word) {
	// Инициализируем генератор случайных чисел.
	rand.Seed(time.Now().UnixNano())

	// Реализация перемешивания Фишера-Йетса (Fisher-Yates shuffle).
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}


func GetTaskWords(
	tr *repositories.GormTaskRepository,
	wr *repositories.GormWordRepository,
	taskNumber uint,
	user *models.User,
) ([]interfaces.Word, error) {
	task, err := tr.Get(taskNumber)

	var mappedWords []interfaces.Word;

	if err != nil {
		return nil, err
	}

	if user != nil {
		words, err := wr.GetTaskUserWords(task.Id, user.Id)
		if err != nil {
			return nil, err
		}

		mappedWords = mappers.DbUserWordToWord(words)
	} else {
		words, err := wr.GetTaskWords(task.Id)
		if err != nil {
			return nil, err
		}

		mappedWords = mappers.DbWordToWord(words)
	}

	shuffleSlice(mappedWords)

	return mappedWords, nil
}
