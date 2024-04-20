package htmlpath

import (
	"fmt"
	"html/template"
	"net/http"
)


type HTML struct {
	pathAuth string
	pathMain string
}

func NewHTML(pathAuth, pathMain string) *HTML {
	return &HTML{
		pathAuth: pathAuth,
		pathMain: pathMain,
	}
}

func (p *HTML)HandleAuth(rw http.ResponseWriter, r *http.Request) {
	
	tmpl, err := template.ParseFiles(p.pathAuth)
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

func (p *HTML)HandleHome(rw http.ResponseWriter, r *http.Request) {
	
	tmpl, err := template.ParseFiles(p.pathMain)
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