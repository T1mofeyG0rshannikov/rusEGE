package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"rusEGE/exceptions"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type JWTProcessor struct {
	secretKey []byte
	issuer    string
}

func NewJWTProcessor() *JWTProcessor {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке файла .env: %v", err)
	}

	secretKey := os.Getenv("SECRETKEY")
	issuer := os.Getenv("ISSUER")

	return &JWTProcessor{
		secretKey: []byte(secretKey),
		issuer:    issuer,
	}
}

// Claims - структура для хранения данных, которые будут включены в JWT токен.
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken создает новый JWT токен.
func (jp *JWTProcessor) GenerateToken(username string, storageTimeInMins uint) (AccessToken, error) {
	expirationTime := time.Minute * time.Duration(int64(storageTimeInMins))
	expiration := time.Now().Add(expirationTime)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    jp.issuer,
			Subject:   "authorization",
		},
	}

	// Создаем новый токен с использованием алгоритма HS256.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен с использованием секретного ключа.
	signedToken, err := token.SignedString(jp.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return AccessToken(signedToken), nil
}

// ValidateToken проверяет JWT токен и возвращает Claims, если токен действителен.
// Если токен недействителен, возвращает ошибку.
func (jp *JWTProcessor) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Парсим токен и проверяем его подпись.
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что используется ожидаемый алгоритм подписи.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jp.secretKey, nil
	})

	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return nil, exceptions.ErrTokenExpired
		}
		return nil, exceptions.ErrInvalidJwtToken
	}

	// Проверяем, что токен действителен.
	if !token.Valid {
		return nil, exceptions.ErrInvalidJwtToken
	}

	// Проверяем издателя (опционально, но рекомендуется).
	if claims.Issuer != jp.issuer {
		return nil, exceptions.ErrInvalidJwtToken
	}

	return claims, nil
}

// ExtractClaims извлекает Claims из токена без проверки подписи.
// Используется только для случаев, когда необходимо прочитать Claims, не доверяя токену.
// В большинстве случаев следует использовать ValidateToken.
func (jp *JWTProcessor) ExtractClaims(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, claims) // ParseUnverified
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}
