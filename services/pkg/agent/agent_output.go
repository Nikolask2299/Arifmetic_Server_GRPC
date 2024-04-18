package agent

import (
	"errors"
	"sync"
	"time"
)

type AgentServiceOutput struct {
	ChanOutputAnswer chan *UserAnswer
	mux              sync.Mutex
	timeout time.Duration
}
type UserAnswer struct {
	Id     int
	task   string
	answer int
}

func NewUserAnswer(task UserTask, answer int) *UserAnswer {
	return &UserAnswer{
		Id: task.id,
		task: task.task,
		answer: answer,
	}
}

func NewAgentServiceOutput(timeout time.Duration) *AgentServiceOutput {
	return &AgentServiceOutput{
		ChanOutputAnswer: make(chan *UserAnswer, 10),
		mux: sync.Mutex{},
		timeout: timeout,
	}
}

func(a *AgentServiceOutput) GetAnswer() *UserAnswer {
	a.mux.Lock()
	defer a.mux.Unlock()
	select {
		case ans := <- a.ChanOutputAnswer:
			return ans
		case <-time.After(2 * time.Second):
			return nil
	}
}

func (a *AgentServiceOutput) PushAnswer(usr *UserAnswer) error {
	a.mux.Lock()
	defer a.mux.Unlock()
	
	select {
		case a.ChanOutputAnswer <- usr:
			return nil
		case <- time.After(2 * time.Second):
			return errors.New("AgentService is unavailable for push")
	}	
}

