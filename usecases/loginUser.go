package usecases


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
) (*auth.AccessToken, error) {
	user, err := ur.Get(data.Username)

	if err != nil {
		return nil, err
	}

	if hasher.CheckPassword(data.Password, user.HashPassword) {
		accessToken, err := jwtProcessor.GenerateToken(user.Username)

		if err != nil {
			return nil, err
		}

		return &accessToken, nil
	}

	return nil, exceptions.ErrIncorrectPassword
}
