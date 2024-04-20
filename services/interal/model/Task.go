package model

type Task struct {
	Id int64
	User_id int
	Task string
	Status string
}

func NewTask(id int64, user_id int, task string, status string) *Task {
	return &Task{
		Id: id,
		User_id: user_id,
		Task: task,
		Status: status,
	}
}