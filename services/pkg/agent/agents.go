package agent

import (
	"fmt"
	"services/interal/model"
	"services/pkg/arifm"
	"strconv"
	"time"
)

func NewCountDemon(count int, inp *AgentServiceInput, out *AgentServiceOutput) {
	for i := 0; i < count; i++ {
		go Demon(inp, out)
	}
}

func Demon(inp *AgentServiceInput, out *AgentServiceOutput) {
	for {
		task := inp.GetTask()
		if task == nil {
			time.Sleep(time.Second * 5)
			continue
		}


		res, err := arifm.ArifmeticServer(task.Task)
		st := strconv.Itoa(int(res))
		if err != nil {
			fmt.Println("Demon:", err , task.Id, task.Task)
			st = "NONE"
		}

		answer := model.NewAnswer(task.Id, st)
		out.PushAnswer(answer)
	}
}
