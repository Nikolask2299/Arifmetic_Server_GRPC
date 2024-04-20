package main

import (
	"fmt"
	"net/http"

	"services/interal/app"
	"services/interal/serverauth"
	"services/interal/serverdata"
	"services/interal/serverhttp"
	"services/interal/service/auth"
	"services/interal/storage/sqlite"
	"services/pkg/config"
	htmlpath "services/pkg/htmlPath"

	"time"
)


func main() {
	cfg := config.Mustload()
	
	storage, err := sqlite.NewStorage(cfg.StoragePath)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DataBase ok")
	}

	servData := serverdata.NewServerData(storage, storage, storage)
	dataApp := app.NewDataApp(servData, cfg.PortData)


	timeout, err := time.ParseDuration(cfg.Timeout)
	if err != nil {
		fmt.Println(err)
		timeout = 10 * time.Second
	}

	tockenTTL, err := time.ParseDuration(cfg.TockenTTL)
	if err != nil {
		fmt.Println(err)
		tockenTTL = 5 * time.Minute
	}

	
	
	clientData := serverdata.NewDataClient(timeout, cfg.PortData)

	auth := auth.NewAuth(*clientData, tockenTTL)
	servAuth := serverauth.NewServerAuth(*auth)
	authApp := app.NewAuthApp(servAuth, cfg.PortAuth)
	
	go authApp.MustRun()		
	go dataApp.MustRun()



	
	html := htmlpath.NewHTML(cfg.PathAuth, cfg.PathMain)

	clientAuth := serverauth.NewAuthClient(timeout, cfg.PortAuth)
	
	httpAuth := serverhttp.NewHTTPAuth(clientAuth, *html)

	http.HandleFunc("/arifmetic/auth/v1", httpAuth.Autorisation)
	http.HandleFunc("/arifmetic/auth", html.HandleAuth)
	http.HandleFunc("/arifmetic", )

	http.ListenAndServe("localhost:" + cfg.HttpPort, nil)
}