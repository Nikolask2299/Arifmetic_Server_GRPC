package tocken

import (
	"errors"
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

func ParseTocken(tockenString string, secret string) (int64, error) {
	tockenParsed, err := jwt.Parse(tockenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := tockenParsed.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("could not parse claims")
	}

	currTime := time.Now().Unix()

	if claims["nbf"].(int64) > currTime {
		return 0, errors.New("token has expired")
	}	
	
	if claims["exp"].(int64) > currTime {
		return 0, errors.New("token has expired")
	}

	id := claims["id"].(int64)

	return id, nil
}
