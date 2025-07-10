package usecases

import (
	"math/rand"
	"time"
)

func ShuffleSlice[T any](slice []T) {
	rand.Seed(time.Now().UnixNano())

	n := len(slice)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}