package words

import (
	"rusEGE/database/models"
	"rusEGE/repositories"
	"rusEGE/web/schemas"
	"strings"
)

func CreateUserWord(
	uwr *repositories.GormUserWordRepository,
	tr *repositories.GormTaskRepository,
	data schemas.CreateUserWordRequest,
	user *models.User,
) (*models.UserWord, error) {
	task, err := tr.Get(data.Task)

	if err != nil {
		return nil, err
	}

	word, err := uwr.Create(user.Id, data.Word, task.Id, data.Rule, data.Exception, data.Description)

	letters := strings.Split(data.Letters, ",")

	for _, letter := range letters {
		uwr.CreateOption(word.Id, letter)
	}

	if err != nil {
		return nil, err
	}

	return word, nil
}
