package serverhttp

import (
	"fmt"
	"net/http"

	"services/interal/model"
	"services/interal/serverdata"
	"services/interal/tocken"
	htmlpath "services/pkg/htmlPath"
	"services/pkg/serial"
)


type HTTPArifmetic struct {
	dataclient serverdata.DataClient
	html htmlpath.HTML
}


func NewHTTPArifmetic(dataclient serverdata.DataClient, html htmlpath.HTML) *HTTPArifmetic {
	return &HTTPArifmetic{
		dataclient: dataclient,
		html: html,
	}
}

func (h *HTTPArifmetic) ArifmeticServer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		userid, err := tocken.ParseTocken(r.Header.Get("tocken"), "secret-user")
		if err != nil {
			w.WriteHeader(304)
			h.html.HandleAuth(w, r)
			return			
		}

		task, err := serial.SerialTaskJSON(r)
		if err != nil {
			w.WriteHeader(304)
			h.html.HandleAuth(w, r)
			return
		}

		err = model.SerialTask(userid, task, h.dataclient)
		if err != nil {
			w.WriteHeader(304)
			h.html.HandleAuth(w, r)
			return
		}

		w.WriteHeader(200)
		fmt.Fprint(w)
		return
	}

	if r.Method == http.MethodGet {
		
		
		return
	}
}






