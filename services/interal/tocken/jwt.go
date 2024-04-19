package tocken

import (
	"services/interal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewTocken(user model.User, secret string, duration time.Duration) (string, error) {
	tocken := jwt.New(jwt.SigningMethodHS256)
	claims := tocken.Claims.(jwt.MapClaims)

	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["nbf"] = time.Now().Add(time.Second * 1).Unix()
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["iat"] = time.Now().Unix()

	tockenString, err := tocken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tockenString, nil
}

