package sqlite

import (
	"context"
	"database/sql"
	"services/interal/model"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(storagePath string) (*Storage, error) {

	database, err := sql.Open("sqlite3",storagePath)
	if err != nil {
		return nil, err
	}

	return &Storage{db: database}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) SaveUser(ctx context.Context,  name string, password []byte) (id int64, err error) {
	var quest = "INSERT INTO users (user_name, pass_hash) VALUES ($1, $2)"

	result, err := s.db.ExecContext(ctx, quest, name, password)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Storage) GetUser(ctx context.Context, id int64, name string) (model.User, error) {
	var result model.User
	
	var quest1 = "SELECT id, user_name, pass_hash FROM users WHERE id = $1"
	var quest2 = "SELECT id, user_name, pass_hash FROM users WHERE user_name = $1"
	
	var res *sql.Row
	
	if name != "" {
		res = s.db.QueryRowContext(ctx, quest2, name)
	} else {
		res = s.db.QueryRowContext(ctx, quest1, id)
	}
	
	err := res.Scan(&result.ID, &result.Name, &result.PassHash)
	if err != nil {
		return model.User{}, err
	}

	return result, nil
}

func (s *Storage) SaveUserTask(ctx context.Context, user_id int64, task string, status string) (id int64, err error) {
	var quest = "INSERT INTO task (user_id, task, stat) VALUES ($1, $2, $3)"

	result, err := s.db.ExecContext(ctx, quest, user_id, task, status)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Storage) SaveUserAnswer(ctx context.Context, task_id int64, answer int) (id int64, err error) {
	var quest = "INSERT INTO answer (task_id, answer) VALUES ($1, $2)"

	result, err := s.db.ExecContext(ctx, quest, task_id, answer)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Storage) GetTask(ctx context.Context, id int64) (model.Task, error) {
	var result model.Task
	var quest = "SELECT id, user_id, task, stat FROM task WHERE id = $1"
	
	quer := s.db.QueryRowContext(ctx, quest, id)
	err := quer.Scan(&result.Id, &result.User_id, &result.Task, &result.Status)
	if err != nil {
		return model.Task{}, err
	}

	return result, nil
}

func (s *Storage) GetUserAnswer(ctx context.Context, id int64, task_id int64) (model.Answer, error) {	
	var result model.Answer
	
	var quest1 = "SELECT id, task_id, answer FROM answer WHERE id = $1"
	var quest2 = "SELECT id, task_id, answer FROM answer WHERE task_id = $1"
	
	var res *sql.Row
	
	if task_id != 0 {
		res = s.db.QueryRowContext(ctx, quest2, task_id)
	} else {
		res = s.db.QueryRowContext(ctx, quest1, id)
	}
	
	err := res.Scan(&result.Id, &result.Task_id, &result.Answer)
	if err != nil {
		return model.Answer{}, err
	}

	return result, nil
}

func (s *Storage) UpdateUserTask(ctx context.Context, id int64, status string) (int64, error) {
	var quest = "UPDATE task SET stat = $1 WHERE id = $2"
	
	res, err := s.db.ExecContext(ctx, quest, status, id)
	if err != nil {
		return 0, err
	}

	id, err = res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Storage) WorkTask(ctx context.Context) (model.Task, error) {
	var quest1 = "SELECT id, user_id, task, stat FROM task WHERE stat = $1"
	var quest2 = "UPDATE task SET stat = $1 WHERE id = $2"
	var retur model.Task

	res := s.db.QueryRowContext(ctx, quest1, "WAITING")
	err := res.Scan(&retur.Id, &retur.User_id, &retur.Task, &retur.Status)
	if err != nil {
		return model.Task{}, err
	}

	_, err = s.db.ExecContext(ctx, quest2, "WORKING", retur.Id)
	if err != nil {
		return model.Task{}, err
	}

	return retur, nil
}

