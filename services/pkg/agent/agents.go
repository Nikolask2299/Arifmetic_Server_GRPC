package agent

import (
	"fmt"
	"time"
	"services/pkg/arifm"
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
			time.Sleep(time.Second * 10)
			continue
		}

		res, err := arifm.ArifmeticServer(task.task)
		if err != nil {
			fmt.Println("Demon:", err , task.id, task.task)
		}
		anwer := NewUserAnswer(*task, res)
		main.AgentOut.PushAnswer(anwer)
	}
}

func (mainOrcServ *MainOrchestratorService) Output() {
	for {
		answ := mainOrcServ.AgentOut.GetAnswer()

		if answ == nil {
			time.Sleep(3 * time.Second)
			continue
		}

		mainOrcServ.database[answ.Id] = answ
	}	
}
