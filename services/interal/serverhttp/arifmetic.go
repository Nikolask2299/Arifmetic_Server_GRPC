package serverhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	serv1 "protobuf/gen/go"

	"services/interal/serverdata"
	"services/interal/tocken"
	htmlpath "services/pkg/htmlPath"
	"services/pkg/serial"
)

type OutTasks struct {
	Id int64 `json:"id"`
	Task string `json:"task"`
	Answer string `json:"answer"`
	Status string `json:"status"`
}

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
		userid, err := tocken.ParseTocken(r.Header.Get("tocken"), "secret-user")
		if err != nil {
			w.WriteHeader(305)
			fmt.Fprintln(w, "/arifmetic/auth")
			return			
		}

		ret, err := GetUserTasks(userid, h.dataclient)
		if err != nil {
			w.WriteHeader(305)
			fmt.Fprintln(w, "/arifmetic/auth")
			return
		}
		
		body, _ := json.Marshal(ret)
		fmt.Fprintln(w, string(body))
		return
	}

	w.WriteHeader(405)
	fmt.Fprintln(w, "Method not allowed")
	return
}





func GetUserTasks(id int64, dataClient serverdata.DataClient) ([]*OutTasks, error) {
	ctxt, conn := dataClient.NewDataClientConnection()
	resp, err := conn.GetUserTasks(ctxt, &serv1.UserTaskRequest{
		IdUser: id,
	})

	if err != nil {
		return nil, err
	}

	var out []*OutTasks

	for _, task := range resp.GetTaskmas() {
		resAns, err := conn.GetAnswer(ctxt, &serv1.GetAnswerRequest{
			Id: 0,
			TaskId: task.GetId(),
		})

		if err != nil {
			if task.GetStat() == "OK" {
				
				conn.UpdateTask(ctxt, &serv1.UpdateTaskRequest{
					Id: task.GetId(),
					Stat: "WAITING",
				})
			}
				out = append(out, &OutTasks{
					Id: task.GetId(),
					Task: task.GetTask(),
					Answer: "NONE",
					Status: "WAITING",
				})
				continue
			}
		
		out = append(out, &OutTasks{
				Id: task.GetId(),
				Task: task.GetTask(),
				Answer: resAns.GetAnswer(),
				Status: task.GetStat(),
			})
		}

	return out, nil
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



