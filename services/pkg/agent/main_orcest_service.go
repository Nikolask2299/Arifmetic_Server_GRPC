package agent


type MainOrchestratorService struct {
	AgentInp *AgentServiceInput
	AgentOut *AgentServiceOutput
	dataCash []*UserTask
	database  map[int]*UserAnswer
	dataout *DataIndex
}

type DataIndex struct {
	dataindex map[int][]int
	dataBool map[int]bool
	countWork map[int]int
}

func NewMainOrchestratorService(agentInp *AgentServiceInput, agentOut *AgentServiceOutput) *MainOrchestratorService {
	return &MainOrchestratorService{
		AgentInp: agentInp,
		AgentOut: agentOut,
		dataCash: make([]*UserTask, 1, 100),
		database: make(map[int]*UserAnswer),
		dataout: &DataIndex{dataindex: make(map[int][]int), 
			dataBool: make(map[int]bool),
			countWork: make(map[int]int),
		},
	}
}

func (s *MainOrchestratorService) GetAnswerData(id int) *UserAnswer { 
	if vl, ok := s.database[id]; ok {
		delete(s.database, id)
		return vl
	} else {
		return nil
	}
}

