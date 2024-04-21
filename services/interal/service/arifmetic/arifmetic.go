package arifmetic

import (
	"fmt"
	serv1 "protobuf/gen/go"
	"services/interal/model"
	"services/interal/serverdata"
	"services/pkg/agent"
	"time"
)


type Arifmetic struct {
	chanInput *agent.AgentServiceInput
	chanOutput *agent.AgentServiceOutput
}

func NewArifmetic(chanInput *agent.AgentServiceInput, chanOutput *agent.AgentServiceOutput) *Arifmetic {
	return &Arifmetic{
		chanInput: chanInput,
		chanOutput: chanOutput,
	}
}

func AgentOutService(a *Arifmetic, dtclient serverdata.DataClient) {
	for {
		input := a.chanOutput.GetAnswer()
		if input == nil {
			time.Sleep(time.Second)
			continue
		}

		ctxt, conn := dtclient.NewDataClientConnection()
		_, err := conn.SaveAnswer(ctxt, &serv1.SaveAnswerRequest{
			TaskId: input.Task_id,
			Answer: int64(input.Answer),
		})
		if err != nil {
			fmt.Println("AgentOut ",err)
		}

		_, err = conn.UpdateTask(ctxt, &serv1.UpdateTaskRequest{
			Id: input.Task_id,
			Stat: "OK",
		})

		if err != nil {
			fmt.Println("AgentOut ",err)
		}

		time.Sleep(time.Second)
	}	
}


func AgentIntService(a *Arifmetic, dtclient serverdata.DataClient) {
	for {
		ctxt, conn := dtclient.NewDataClientConnection()
		
		res, err := conn.WorkTask(ctxt, &serv1.WorkRequest{
			Req: true,
		})

		if err != nil {
			fmt.Println("AgentInt ",err)
			time.Sleep(time.Second)
			continue
		}

		tsk := model.NewTask(res.GetId(), res.GetUserId(), res.GetTask(), res.GetStat())
		fmt.Println("AgentInt ",tsk)
		for {
			err := a.chanInput.Push(tsk)
			if err != nil {
				time.Sleep(time.Second)
				continue
			}
			break
		}
		
		time.Sleep(time.Second)
	}
}




