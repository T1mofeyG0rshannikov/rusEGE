package auth

import (
	"rusEGE/auth"
	"rusEGE/exceptions"
	"rusEGE/repositories"
	"rusEGE/security"
	"rusEGE/web/schemas"
)

func LoginUser(
	ur *repositories.GormUserRepository,
	hasher *security.ScryptHasher,
	jwtProcessor *auth.JWTProcessor,
	data *schemas.LoginRequest,
) (*auth.AccessToken, *auth.AccessToken, error) {
	user, err := ur.Get(data.Username)

	if err != nil {
		return nil, nil, err
	}

	if hasher.CheckPassword(data.Password, user.HashPassword) {
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

	return nil, nil, exceptions.ErrIncorrectPassword
}
