package agent


type MainOrchestratorService struct {
	AgentInp *AgentServiceInput
	AgentOut *AgentServiceOutput
}

func NewMainOrchestratorService(agentInp *AgentServiceInput, agentOut *AgentServiceOutput) *MainOrchestratorService {
	return &MainOrchestratorService{
		AgentInp: agentInp,
		AgentOut: agentOut,
	}
}


