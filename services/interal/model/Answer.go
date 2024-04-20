package model

type Answer struct {
	Id int
	Task_id int
	Answer int
}

func NewAnswer(task_id int, answer int) *Answer {
	return &Answer{
		Task_id: task_id,
		Answer: answer,
	}
}