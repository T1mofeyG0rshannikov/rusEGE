package utils


import "strings"


func IsUniqueConstraintError(err error) bool {
    // В зависимости от базы данных и драйвера, текст ошибки может отличаться.
    // Можно сделать общую проверку по тексту.
    if err == nil {
        return false
    }

    errMsg := err.Error()

    // Для SQLite
    if strings.Contains(errMsg, "UNIQUE constraint failed") {
        return true
    }

    // Для PostgreSQL
    if strings.Contains(errMsg, "duplicate key value violates unique constraint") {
        return true
    }

    // Для MySQL
    if strings.Contains(errMsg, "Error 1062") {
        return true
    }

    return false
}