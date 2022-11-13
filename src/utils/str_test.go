package utils

import (
	"github.com/magiconair/properties/assert"
	"log"
	"math/rand"
	"testing"
)

func TestRandomWithLength(t *testing.T) {
	for i := 1; i <= 10; i++ {
		log.Print(rand.Intn(50))
	}
	length := 5
	result := RandomWithLength(5)
	assert.Equal(t, length, len(result))
}
