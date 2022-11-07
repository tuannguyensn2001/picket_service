package utils

import (
	"context"
	"strconv"
)

func GetAuth(ctx context.Context) (int, error) {
	val, err := strconv.Atoi(ctx.Value("userId").(string))
	if err != nil {
		return -1, err
	}
	return val, nil
}
