package model

type Answer struct {
	Id int64
	Task_id int64
	Answer string
}

func NewAnswer(task_id int64, answer string) *Answer {
	return &Answer{
		Task_id: task_id,
		Answer: answer,
	}
}