package agent

import (
	"errors"
	"services/interal/model"
	"sync"
	"time"
)

type AgentServiceOutput struct {
	ChanOutputAnswer chan *model.Answer
	mux              sync.Mutex
}

func NewAgentServiceOutput() *AgentServiceOutput {
	return &AgentServiceOutput{
		ChanOutputAnswer: make(chan *model.Answer, 10),
		mux: sync.Mutex{},
	}
}

func(a *AgentServiceOutput) GetAnswer() *model.Answer {
	a.mux.Lock()
	defer a.mux.Unlock()
	select {
		case ans := <- a.ChanOutputAnswer:
			return ans
		case <-time.After(time.Second):
			return nil
	}
}

func (a *AgentServiceOutput) PushAnswer(usr *model.Answer) error {
	a.mux.Lock()
	defer a.mux.Unlock()	
	select {
		case a.ChanOutputAnswer <- usr:
			return nil
		case <- time.After(time.Second):
			return errors.New("AgentService is unavailable for push")
	}	
}

