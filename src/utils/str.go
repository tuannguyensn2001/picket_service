package utils

import (
	"math/rand"
)

func RandomWithLength(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]rune, n)

	for i := range result {
		r := rand.Intn(len(letterRunes))
		result[i] = letterRunes[r]
	}

	return string(result)

}
