package words

import (
	"fmt"
	"rusEGE/database/models"
	"rusEGE/repositories"
	"rusEGE/web/schemas"
)

func CreateError(
	wr *repositories.GormWordRepository,
	uwr *repositories.GormUserWordRepository,
	data schemas.CreateWordErrorRequest,
	user *models.User,
) (*models.UserError, *models.Error, error) {
	userWord, err := uwr.Get(data.Word)
	if err != nil {
		return nil, nil, err
	}

	userError, err := uwr.CreateError(user.Id, userWord.Id)
	fmt.Println(userError)
	if err != nil {
		return nil, nil, err
	}

	word, err := wr.GetByWord(data.Word)

	if err != nil {
		return nil, nil, err
	}

	wordError, err := wr.CreateError(word.Id)
	fmt.Println(wordError)
	if err != nil {
		return nil, nil, err
	}

	return userError, wordError, nil
}
