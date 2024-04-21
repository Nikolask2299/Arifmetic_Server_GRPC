package serverhttp

import (
	"fmt"
	"net/http"
	serv1 "protobuf/gen/go"

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
			w.WriteHeader(305)
			fmt.Fprintln(w, "/arifmetic/auth")
			return			
		}

		task, err := serial.SerialTaskJSON(r)
		if err != nil {
			w.WriteHeader(305)
			fmt.Fprintln(w, "/arifmetic/auth")
			return
		}

		err = SerialTask(userid, task, h.dataclient)
		if err != nil {
			w.WriteHeader(305)
			fmt.Fprintln(w, "/arifmetic/auth")
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


func SerialTask(userID int64, tasks []string, dataClient serverdata.DataClient) error {
	ctxt, conn := dataClient.NewDataClientConnection()
	for _, task := range tasks {
		_, err := conn.SaveTask(ctxt, &serv1.SaveTaskRequest{
			IdUser: userID,
			Task: task,
		})

		if err != nil {
			return err
		}

	}
	return nil
}



