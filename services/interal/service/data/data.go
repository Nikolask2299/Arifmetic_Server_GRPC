package data

import (
	"context"
	"services/interal/model"
)

type DataUser interface {
	SaveUser(ctx context.Context, name string, password []byte) (int64, error)
	GetUser(ctx context.Context, id int64, name string) (model.User, error)
}

type DataTask interface {
	SaveUserTask(ctx context.Context, user_id int64, task string, status string) (int64, error)
	GetUserTask(ctx context.Context, id int64) (model.Task, error)	
}

type DataAnswer interface {
	SaveUserAnswer(ctx context.Context, task_id int64, answer int) (int64, error)
	GetUserAnswer(ctx context.Context, id int64, task_id int64) (model.Answer, error)
}

