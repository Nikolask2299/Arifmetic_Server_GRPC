package agent

import (
	"sync"
	"time"
)

type AgentServiceInput struct {
	ChanInputTask chan *UserTask
	mux sync.Mutex
	timeout time.Duration
}

type UserTask struct {
	id int
	task string
}

func NewAgentServiceInput(timeout time.Duration) *AgentServiceInput {
	return &AgentServiceInput{ChanInputTask: make(chan *UserTask, 10), 
		mux: sync.Mutex{}, timeout: timeout}
}

func (a *MainOrchestratorService) GetTask() *UserTask {
	a.AgentInp.mux.Lock()
	defer a.AgentInp.mux.Unlock()
	select {
		case tsk := <- a.AgentInp.ChanInputTask:
			return tsk
		case <-time.After(2 * time.Second):
			for i, ts := range a.dataCash {
				if ts != nil {
					gf := ts
					a.dataCash[i] = nil
					return gf
				}
			}
			return nil
	}
}

func (a *MainOrchestratorService) Push(task UserTask) error {
	a.AgentInp.mux.Lock()
	defer a.AgentInp.mux.Unlock()
	select {
		case a.AgentInp.ChanInputTask <- &task:
			return nil
		case <-time.After(2 * time.Second):
			for i, val := range a.dataCash {
				if val == nil {
					a.dataCash[i] = &task
					return nil
				}
			}
			a.dataCash = append(a.dataCash, &task)
			return nil		
	}
}
