package model

type Task struct {
	Id int64
	User_id int64
	Task string
	Status string
}

func NewTask(id int64, user_id int64, task string, status string) *Task {
	return &Task{
		Id: id,
		User_id: user_id,
		Task: task,
		Status: status,
	}
}


