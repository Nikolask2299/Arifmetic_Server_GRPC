package main

import (
	"fmt"
	"net/http"
	"services/interal/app"
	"services/interal/serverdata"
	"services/interal/storage/sqlite"
	"services/pkg/config"
	"strconv"
	"text/template"
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
	}

	servData := serverdata.NewServerData(storage, storage, storage)

	port, _ := strconv.Atoi(cfg.Port)

	dataApp := app.NewDataApp(servData, port)

	go dataApp.MustRun()

}