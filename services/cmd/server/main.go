package main

import (
	
	"fmt"
	"net/http"

	"services/interal/app"
	"services/interal/serverauth"
	"services/interal/serverdata"
	"services/interal/service/auth"
	"services/interal/storage/sqlite"
	"services/pkg/config"

	"text/template"
	"time"
)

type HTML struct {
	path string
}

func (p *HTML)HandleHome(rw http.ResponseWriter, r *http.Request) {
	
	tmpl, err := template.ParseFiles(p.path)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), 400)
		return
	}

	err = tmpl.Execute(rw, nil)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), 400)
		return
	}
}

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
	 dataApp.MustRun()
}