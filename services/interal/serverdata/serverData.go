package serverdata

import (
	"context"
	serv "protobuf/gen/go"
	"services/interal/service/data"

	"strconv"

	"google.golang.org/grpc"
)

type ServerData struct {
	serv.DataBaseServiceServer
	usrData data.DataUser
	taskData data.DataTask
	ansData data.DataAnswer
}

func NewServerData(usrData data.DataUser, taskData data.DataTask, ansData data.DataAnswer) *ServerData {
	return &ServerData{
		usrData: usrData,
		taskData: taskData,
		ansData: ansData,
	}
}

func Register(grpc *grpc.Server, s *ServerData) {
	serv.RegisterDataBaseServiceServer(grpc, s)
}

type DataBaseServiceServer interface {
	GetUser(context.Context, *serv.GetUserRequest) (*serv.GetUserResponse, error)
	SaveUser(context.Context, *serv.SaveUserRequest) (*serv.SaveUserResponse, error)
	GetTask(context.Context, *serv.GetTaskRequest) (*serv.GetTaskResponse, error)
	SaveTask(context.Context, *serv.SaveTaskRequest) (*serv.SaveTaskResponse, error)
	GetAnswer(context.Context, *serv.GetAnswerRequest) (*serv.GetAnswerResponse, error)
	SaveAnswer(context.Context, *serv.SaveAnswerRequest) (*serv.SaveAnswerResponse, error)
	UpdateTask(context.Context, *serv.UpdateTaskRequest) (*serv.UpdateTaskResponse, error)
	WorkTask(context.Context, *serv.WorkRequest) (*serv.WorkResponse, error)
	mustEmbedUnimplementedDataBaseServiceServer()
}

func (s *ServerData) GetUser(ctx context.Context, req *serv.GetUserRequest) (*serv.GetUserResponse, error) {
	res, err := s.usrData.GetUser(ctx, req.GetUserId(), req.GetUser())
	if err != nil {
		return nil, err
	}
	return &serv.GetUserResponse{IdUser: res.ID, UserName: res.Name, Password:  string(res.PassHash)}, nil
}

func (s *ServerData) SaveUser(ctx context.Context, req *serv.SaveUserRequest) (*serv.SaveUserResponse, error) {
	res, err := s.usrData.SaveUser(ctx, req.GetUserName(), []byte(req.GetPassword()))
	if err != nil {
		return nil, err
	}
	return &serv.SaveUserResponse{Id: res}, nil
}

func (s *ServerData) GetTask(ctx context.Context, req *serv.GetTaskRequest) (*serv.GetTaskResponse, error) {
	res, err := s.taskData.GetTask(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &serv.GetTaskResponse{Task: res.Task, Status: res.Status}, nil
}

func (s *ServerData) SaveTask(ctx context.Context, req *serv.SaveTaskRequest) (*serv.SaveTaskResponse, error) {
	res, err := s.taskData.SaveUserTask(ctx, req.GetIdUser(), req.GetTask(), "WAITING")
	if err != nil {
		return nil, err
	}
	return &serv.SaveTaskResponse{Id: res}, nil
}


func (s *ServerData) GetAnswer(ctx context.Context, req *serv.GetAnswerRequest) (*serv.GetAnswerResponse, error) {
	res, err := s.ansData.GetUserAnswer(ctx, req.GetId(), req.GetTaskId())
	if err != nil {
		return nil, err
	}
	return &serv.GetAnswerResponse{Id: int64(res.Id), Answer: strconv.Itoa(res.Answer)}, nil
}

func (s *ServerData) SaveAnswer(ctx context.Context, req *serv.SaveAnswerRequest) (*serv.SaveAnswerResponse, error) {
	res, err := s.ansData.SaveUserAnswer(ctx, req.GetTaskId(), int(req.GetAnswer()))
	if err != nil {
		return nil, err
	}
	return &serv.SaveAnswerResponse{Id: res}, nil
}

func (s *ServerData) UpdateTask(ctx context.Context, req *serv.UpdateTaskRequest) (*serv.UpdateTaskResponse, error) {
	res, err := s.taskData.UpdateUserTask(ctx, req.GetId(), req.GetStat())
	if err != nil {
		return nil, err
	}
	return &serv.UpdateTaskResponse{Id: res}, nil
}

func (s *ServerData) WorkTask(ctx context.Context, req *serv.WorkRequest) (*serv.WorkResponse, error) {
	res, err := s.taskData.WorkTask(ctx)
	if err != nil {
		return nil, err
	}
	return &serv.WorkResponse{
		Id: res.Id,
		UserId: int64(res.User_id),
		Task: res.Task,
		Stat: res.Status,
	}, nil
}
