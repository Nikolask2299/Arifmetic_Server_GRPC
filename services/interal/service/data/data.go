package data

import (
	"context"
	"services/interal/model"
)

type DataUser interface {
	SaveUser(ctx context.Context, name string, password []byte) (int64, error)
	GetUser(ctx context.Context, id int64, name string) (model.User, error)
	GetUserTasks(ctx context.Context, user_id int64) ([]model.Task, error)
}

type DataTask interface {
	SaveUserTask(ctx context.Context, user_id int64, task string, status string) (int64, error)
	GetTask(ctx context.Context, id int64) (model.Task, error)	
	UpdateUserTask(ctx context.Context, id int64, status string) (int64, error)
	WorkTask(ctx context.Context) (model.Task, error)
}

type DataAnswer interface {
	SaveUserAnswer(ctx context.Context, task_id int64, answer string) (int64, error)
	GetUserAnswer(ctx context.Context, id int64, task_id int64) (model.Answer, error)
}


