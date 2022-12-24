package auth

import (
	"time"

	"github.com/dev-hyunsang/clone-twitter-backend/ent"
	"github.com/golang-jwt/jwt"
)

func NewAuthJWT(user *ent.User) (string, error) {
	claims := jwt.MapClaims{
		"user_uuid":  user.UserUUID,
		"user_email": user.UserEmail,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("clone_todo_mate"))
	if err != nil {
		return "", err
	}

	return t, nil
}
