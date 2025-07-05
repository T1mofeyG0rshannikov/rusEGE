package usecases

import (
	"rusEGE/auth"
	"rusEGE/database"
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
	data *schemas.CreateUserRequest,
) (*auth.AccessToken, error) {
	user, err := ur.Get(data.Username)

	if user != nil{
		return nil, exceptions.ErrUsernameExist
	}

	hashedPassword, err := hasher.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	userDB, err := ur.Create(&database.User{
		Username:     data.Username,
		HashPassword: hashedPassword,
	})

	if err != nil {
		return nil, err
	}

	words, err := wr.GetAll()

	if err != nil {
		return nil, err
	}

	for _, word := range words {
		wr.CreateUserWord(&database.UserWord{
			UserId: userDB.Id,
			Word: word.Word,
			TaskId: word.TaskId,
			Rule: word.Rule,
		})
	}

	accessToken, err := jwtProcessor.GenerateToken(userDB.Username)
	if err != nil {
		return nil, err
	}

	return &accessToken, nil
}
