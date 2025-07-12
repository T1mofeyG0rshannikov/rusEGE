package auth

import (
	"errors"
	"rusEGE/auth"
	"rusEGE/exceptions"
	"rusEGE/repositories"
	"rusEGE/security"
	"rusEGE/web/schemas"
)

func CreateUser(
	ur *repositories.GormUserRepository,
	wr *repositories.GormWordRepository,
	uwr *repositories.GormUserWordRepository,
	jwtProcessor *auth.JWTProcessor,
	hasher *security.ScryptHasher,
	data schemas.CreateUserRequest,
) (*auth.AccessToken, *auth.AccessToken, error) {
	_, err := ur.Get(data.Username)

	if err != nil {
		if !errors.Is(err, exceptions.ErrUserNotFound) {
			return nil, nil, err
		}
	} else {
		return nil, nil, exceptions.ErrUsernameExist
	}

	hashedPassword, err := hasher.HashPassword(data.Password)
	if err != nil {
		return nil, nil, err
	}

	user, err := ur.Create(data.Username, hashedPassword)

	if err != nil {
		return nil, nil, err
	}

	words, err := wr.All()

	if err != nil {
		return nil, nil, err
	}

	for _, word := range words {
		userWord, err := uwr.Create(user.Id, word.Word, word.TaskId, word.RuleId, &word.Exception, word.Description)

		if err != nil {
			return nil, nil, err
		}

		for _, option := range word.Options {
			_, err := uwr.CreateOption(userWord.Id, option.Letter)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	accessToken, err := jwtProcessor.GenerateToken(user.Username, 30)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := jwtProcessor.GenerateToken(user.Username, 60*24)

	if err != nil {
		return nil, nil, err
	}

	return &accessToken, &refreshToken, nil
}
