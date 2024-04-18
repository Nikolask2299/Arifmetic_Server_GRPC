package agent

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

type OutAnswer struct {
	Id int `json:"id"`
	Task   string `json:"task"`
	Answer int `json:"answer"`
	Status string `json:"status"`
}


func NewUserTask(request *http.Request) ([]*UserTask, error) {
	res := make([]*UserTask, 0, 10)
	if request.Header.Get("Content-Type") == "application/json" {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}
		
		var buff []string

		err = json.Unmarshal(body, &buff)
		if err != nil {
			return nil, err
		}

		for _, tas := range buff {
			id := NewId()
			res = append(res, &UserTask{id: id, task: tas})
		}
	} else if request.Header.Get("Content-Type") == "text/plain" {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}
		
		id := NewId()
		res = append(res, &UserTask{id: id, task: string(body)})
	}
	return res, nil
}

func (a *MainOrchestratorService)MainOrchestrator(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		task, err := NewUserTask(r)
		if err != nil {
			fmt.Fprint(w, "Error creating incorrect user task")
		} else {
			index := make([]int, 0, len(task))
			for _, ts := range task {
				gf := ts
				a.dataout.dataBool[gf.id] = true
				index = append(index, gf.id)
				go func(gf *UserTask) {
					a.Push(*gf)
				}(gf) 
			}	
			mainIndex := NewId()
			a.dataout.dataindex[mainIndex] = index
			a.dataout.countWork[mainIndex] = len(task)
			fmt.Fprintln(w, mainIndex)
		}

	} else if r.Method == "GET" {

		idst := r.Header.Get("id")
		id, _ := strconv.Atoi(idst)
		masout := make([]OutAnswer, 0, 10)

		if vl, ok := a.dataout.dataindex[id]; ok {
			for _, ts := range vl {
				answ := a.GetAnswerData(ts)
				if answ != nil {
					outwansw1 := OutAnswer{
						Id: answ.Id,
						Task: answ.task,
						Answer: answ.answer,
						Status: "OK",
					}
					a.dataout.dataBool[answ.Id] = false
					a.dataout.countWork[id]--
					masout = append(masout, outwansw1)
				} else {
					outwansw2 := OutAnswer{}
					if a.dataout.dataBool[ts] {
						outwansw2.Id = ts
						outwansw2.Task = "NULL"
						outwansw2.Answer = 0
						outwansw2.Status = "WORKING"
						masout = append(masout, outwansw2)
					}
					
				}
			}
			
			body, _ := json.Marshal(masout)
			fmt.Fprintln(w, string(body))
		
			if a.dataout.countWork[id] == 0 {
				for _, ts := range vl {
					delete(a.dataout.dataBool, ts)
				}
				delete(a.dataout.dataindex, id)
				delete(a.dataout.countWork, id)		
			}
			
		} else {
			outwansw := OutAnswer{
				Id: id,
				Task: "NULL",
				Answer: 0,
				Status: "NOT FOUND",
			}
			masout = append(masout, outwansw)
			body, _ := json.Marshal(masout)
			fmt.Fprintln(w, string(body))
		}
	
	} else {
		fmt.Fprintln(w, "Invalid method")
	}
}



func NewId() int {
	var res int
	for i := 1; i < 5; i++ {
		num := i * rand.Intn(100)
		res += num
	}	
	return res % 1000
}