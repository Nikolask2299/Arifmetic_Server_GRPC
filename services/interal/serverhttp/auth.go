package serverhttp

import (
	"encoding/json"
	"fmt"

	"io"
	"net/http"
	serv1 "protobuf/gen/go"
	"services/interal/serverauth"
	htmlpath "services/pkg/htmlPath"
)

type HTTPAuth struct {
	auth *serverauth.AuthClient
	html htmlpath.HTML
}

func NewHTTPAuth(auth *serverauth.AuthClient, html htmlpath.HTML) *HTTPAuth {
	return &HTTPAuth{
		auth: auth,
		html: html,
	}
}


func (a *HTTPAuth) Autorisation(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data map[string]string

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.Unmarshal(body, &data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctxt, connect := a.auth.NewAuthClientConnection()		

		name := data["name"]
		password := data["password"]

		res, err := connect.RegisterUser(ctxt, &serv1.RegisterRequest{
			UserName:     name,
			Password: password,
		})	
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if res.GetUserId() == 0 {
			http.Error(w, "User not found", http.StatusBadRequest)
			return
		}
		
		fmt.Fprintln(w, "/arifmetic/auth")
		return
	}

	if r.Method == http.MethodGet {
		name := r.Header.Get("name")
		password := r.Header.Get("password")

		ctxt, connect := a.auth.NewAuthClientConnection()
		res, err := connect.LoginUser(ctxt, &serv1.LoginRequest{
			UserName:     name,
			Password: password,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Add("tocken", res.GetToken())
		fmt.Fprint(w, "/arifmetic/home")
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}	
