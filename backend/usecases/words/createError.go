package words

import (
	"rusEGE/database/models"
	"rusEGE/repositories"
	"rusEGE/web/schemas"
)

func CreateError(
	wr *repositories.GormWordRepository,
	data schemas.CreateWordErrorRequest,
	user *models.User,
) (*models.UserError, *models.Error, error) {
	userWord, err := wr.GetUserWord(data.Word)
	if err != nil {
		return nil, nil, err
	}

	userError, err := wr.CreateUserError(user.Id, userWord.Id)
	if err != nil {
		return nil, nil, err
	}

	word, err := wr.GetByWord(data.Word)

	if err != nil {
		return nil, nil, err
	}

	wordError, err := wr.CreateError(word.Id)
	if err != nil {
		return nil, nil, err
	}

	return userError, wordError, nil
}
