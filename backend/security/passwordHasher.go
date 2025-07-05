package security

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"crypto/hmac"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type PasswordHasher interface {
	GenerateSalt() (string, error)
	HashPassword(password, salt string) (string, error)
	CheckPassword(password, hash, salt string) bool
}

// ScryptHasher реализация PasswordHasher с использованием scrypt
type ScryptHasher struct {
	secretKey []byte
}

// NewScryptHasher конструктор для ScryptHasher
func NewScryptHasher() *ScryptHasher {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке файла .env: %v", err)
	}

	hasherSecretKey := os.Getenv("HASHERSECRETKEY")

	return &ScryptHasher{
		secretKey: []byte(hasherSecretKey),
	}
}

// HashPassword хеширует пароль с солью
func (s *ScryptHasher) HashPassword(password string) (string, error) {

	h := hmac.New(sha256.New, []byte("secretKnjbjkbbbkey"))

	// Записываем входную строку в HMAC
	h.Write([]byte(password))

	// Получаем хеш (в виде байтового массива)
	hashed := h.Sum(nil)

	hashedPassword := hex.EncodeToString(hashed[:])

	return hashedPassword, nil
}

// CheckPassword проверяет, соответствует ли пароль хешу и соли
func (s *ScryptHasher) CheckPassword(password, hash string) bool {
	hashedPassword, err := s.HashPassword(password)

	passwordHash1 := strings.Split(hashedPassword, "_")[0]
	passwordHash2 := strings.Split(hash, "_")[0]

	if err != nil {
		return false
	}

	return passwordHash1 == passwordHash2
}
