package serverdata

import (
	"context"
	serv "protobuf/gen/go"
	"services/service/data"
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
	res, err := s.taskData.GetUserTask(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &serv.GetTaskResponse{Task: res.Task, Status: res.Status}, nil
}

func (s *ServerData) SaveTask(ctx context.Context, req *serv.SaveTaskRequest) (*serv.SaveTaskResponse, error) {
	res, err := s.taskData.SaveUserTask(ctx, req.IdUser, req.GetTask(), "WAITING")
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


