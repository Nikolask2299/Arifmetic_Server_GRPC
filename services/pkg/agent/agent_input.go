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

func NewAgentServiceInput() *AgentServiceInput {
	return &AgentServiceInput{ChanInputTask: make(chan *model.Task, 10), 
		mux: sync.Mutex{}}
}

func (a *AgentServiceInput) GetTask() *model.Task {
	a.mux.Lock()
	defer a.mux.Unlock()
	select {
		case tsk := <- a.ChanInputTask:
			return tsk
		case <-time.After(time.Second):			
			return nil
	}
}

func (a *AgentServiceInput) Push(task *model.Task) error {
	a.mux.Lock()
	defer a.mux.Unlock()
	select {
		case a.ChanInputTask <- task:
			return nil
		case <-time.After(time.Second):			
			return errors.New("timeout")
	}	
}
