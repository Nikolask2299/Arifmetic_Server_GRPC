package tocken

import (
	"errors"
	"fmt"
	"services/interal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewTocken(user model.User, secret string, duration time.Duration) (string, error) {
	tocken := jwt.New(jwt.SigningMethodHS256)
	claims := tocken.Claims.(jwt.MapClaims)

	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["nbf"] = time.Now().Add(time.Second).Unix()
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
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			panic(fmt.Errorf("unexpected signing method: %v", token.Header["alg"]))
		}
		
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	var id float64

	if claims, ok := tockenParsed.Claims.(jwt.MapClaims); ok {
		id = claims["id"].(float64)
	} else {
		return 0, errors.New("invalid tocken")
	} 

	return int64(id), nil
}
