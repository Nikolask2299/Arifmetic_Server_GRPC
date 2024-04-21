package model

import (
	serv1 "protobuf/gen/go"
	"services/interal/serverdata"
)

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

func SerialTask(userID int64, tasks []string, dataClient serverdata.DataClient) error {
	ctxt, conn := dataClient.NewDataClientConnection()
	for _, task := range tasks {
		_, err := conn.SaveTask(ctxt, &serv1.SaveTaskRequest{
			IdUser: userID,
			Task: task,
		})

		if err != nil {
			return err
		}

	}
	return nil
}
