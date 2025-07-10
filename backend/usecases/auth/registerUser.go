package auth

import (
	"errors"
	"rusEGE/auth"
	"rusEGE/database/models"
	"rusEGE/exceptions"
	"rusEGE/repositories"
	"rusEGE/security"
	"rusEGE/web/schemas"
)

func CreateUser(
	ur *repositories.GormUserRepository,
	wr *repositories.GormWordRepository,
	jwtProcessor *auth.JWTProcessor,
	hasher *security.ScryptHasher,
	data schemas.CreateUserRequest,
) (*auth.AccessToken, *auth.AccessToken, error) {
	_, err := ur.Get(data.Username)

	if err != nil{
		if !errors.Is(err, exceptions.ErrUserNotFound){
			return nil, nil, err
		}
	} else{
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

	words, err := wr.GetAll()

	if err != nil {
		return nil, nil, err
	}

	for _, word := range words {
		wr.CreateUserWord(&models.UserWord{
			UserId: user.Id,
			Word:   word.Word,
			TaskId: word.TaskId,
			RuleId: word.RuleId,
		})
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
