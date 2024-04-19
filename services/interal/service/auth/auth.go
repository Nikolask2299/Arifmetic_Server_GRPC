package auth

import (
	"context"
	serv1 "protobuf/gen/go"
	"services/interal/model"
	"services/interal/serverdata"

	"services/interal/tocken"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	dataclient serverdata.DataClient
	tokenTTL time.Duration
}


func NewAuth(client serverdata.DataClient, tokenTTL time.Duration) *Auth {
	return &Auth{
		dataclient: client,
		tokenTTL: tokenTTL,
	}
}

func (a *Auth) Login(ctx context.Context, username, password string) (string, error) {
	
	ctxt, connect := a.dataclient.NewDataClientConnection()

	resp, err := connect.GetUser(ctxt, &serv1.GetUserRequest{
		User: username,
		UserId: 0,
	})
	
	if err != nil {
		return "", err
	}

	user := model.User{
		ID: resp.GetIdUser(),
		Name: resp.GetUserName(),
		PassHash: []byte(resp.GetPassword()),
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		return "", err
	}

	token, err := tocken.NewTocken(user, "secret-user", a.tokenTTL)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *Auth) Register(ctx context.Context, username, password string) (int64, error) {

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	ctxt, connect := a.dataclient.NewDataClientConnection()

	resp, err := connect.SaveUser(ctxt, &serv1.SaveUserRequest{
		UserName: username,
		Password: string(passHash),
	})
	if err != nil {
		return 0, err
	}

	return resp.GetId(), nil
}
