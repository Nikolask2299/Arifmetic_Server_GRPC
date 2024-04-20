package agent

import (
	"fmt"
	"services/interal/model"
	"services/pkg/arifm"
	"time"
)

func NewCountDemon(count int, main *MainOrchestratorService) {
	for i := 0; i < count; i++ {
		go Demon(main)
	}
}

func Demon(main *MainOrchestratorService) {
	for {
		task := main.GetTask()
		if task == nil {
			time.Sleep(time.Second * 5)
			continue
		}

		res, err := arifm.ArifmeticServer(task.Task)
		if err != nil {
			fmt.Println("Demon:", err , task.Id, task.Task)
		}
		answer := model.NewAnswer(task.Id, res)
		main.AgentOut.PushAnswer(anwer)
	}
}

func (mainOrcServ *MainOrchestratorService) Output() {
	for {
		answ := mainOrcServ.AgentOut.GetAnswer()
		if answ == nil {
			time.Sleep(time.Second)
			continue
		}

		
	}	
}
