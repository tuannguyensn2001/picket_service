package test_repository

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
	"testing"
)

func TestRepo_FindByTestId(t *testing.T) {
	rd := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	db, _ := gorm.Open(postgres.Open("postgres://postgres:secret@103.180.136.154/postgres"))
	repo := New(db, rd)

	ctx := context.TODO()

	//rd.Del(ctx, "test_1")

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, err := repo.FindByTestId(ctx, 1)
			log.Println(result, err)
		}()
	}

	wg.Wait()

	assert.True(t, true)
}
