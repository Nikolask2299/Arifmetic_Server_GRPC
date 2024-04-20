package agent

import (
	"errors"
	"services/interal/model"
	"sync"
	"time"
)

type AgentServiceInput struct {
	ChanInputTask chan *model.Task
	mux sync.Mutex
}

func NewAgentServiceInput(timeout time.Duration) *AgentServiceInput {
	return &AgentServiceInput{ChanInputTask: make(chan *model.Task, 10), 
		mux: sync.Mutex{}}
}

func (a *MainOrchestratorService) GetTask() *model.Task {
	a.AgentInp.mux.Lock()
	defer a.AgentInp.mux.Unlock()
	select {
		case tsk := <- a.AgentInp.ChanInputTask:
			return tsk
		case <-time.After(time.Second):			
			return errors.New("timeout")
	}
}

func (a *MainOrchestratorService) Push(task *model.Task) error {
	a.AgentInp.mux.Lock()
	defer a.AgentInp.mux.Unlock()
	select {
		case a.AgentInp.ChanInputTask <- task:
			return nil
		case <-time.After(time.Second):			
			return errors.New("timeout")
	}	
}
