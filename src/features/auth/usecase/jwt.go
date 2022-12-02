package auth_usecase

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"picket/src/entities"
	"strconv"
	"time"
)

func (u *usecase) GenerateToken(ctx context.Context, user entities.User) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Subject:   user.Email,
		ID:        strconv.Itoa(user.Id),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(u.config.GetSecretKey()))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func (u *usecase) VerifyToken(ctx context.Context, token string) (int, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.config.GetSecretKey()), nil
	})
	if err != nil {
		return -1, err
	}
	claims, ok := t.Claims.(*jwt.RegisteredClaims)
	if !ok || !t.Valid {
		return -1, jwt.ErrTokenNotValidYet
	}
	id, err := strconv.Atoi(claims.ID)
	if err != nil {
		return -1, err
	}
	return id, nil
}
