// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serviceDataBase.proto

package serv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type WorkRequest struct {
	Req                  bool     `protobuf:"varint,1,opt,name=req,proto3" json:"req,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkRequest) Reset()         { *m = WorkRequest{} }
func (m *WorkRequest) String() string { return proto.CompactTextString(m) }
func (*WorkRequest) ProtoMessage()    {}
func (*WorkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{0}
}

func (m *WorkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkRequest.Unmarshal(m, b)
}
func (m *WorkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkRequest.Marshal(b, m, deterministic)
}
func (m *WorkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkRequest.Merge(m, src)
}
func (m *WorkRequest) XXX_Size() int {
	return xxx_messageInfo_WorkRequest.Size(m)
}
func (m *WorkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkRequest proto.InternalMessageInfo

func (m *WorkRequest) GetReq() bool {
	if m != nil {
		return m.Req
	}
	return false
}

type WorkResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Task                 string   `protobuf:"bytes,3,opt,name=task,proto3" json:"task,omitempty"`
	Stat                 string   `protobuf:"bytes,4,opt,name=stat,proto3" json:"stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkResponse) Reset()         { *m = WorkResponse{} }
func (m *WorkResponse) String() string { return proto.CompactTextString(m) }
func (*WorkResponse) ProtoMessage()    {}
func (*WorkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{1}
}

func (m *WorkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkResponse.Unmarshal(m, b)
}
func (m *WorkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkResponse.Marshal(b, m, deterministic)
}
func (m *WorkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkResponse.Merge(m, src)
}
func (m *WorkResponse) XXX_Size() int {
	return xxx_messageInfo_WorkResponse.Size(m)
}
func (m *WorkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkResponse proto.InternalMessageInfo

func (m *WorkResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WorkResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *WorkResponse) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *WorkResponse) GetStat() string {
	if m != nil {
		return m.Stat
	}
	return ""
}

type UpdateTaskRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Stat                 string   `protobuf:"bytes,2,opt,name=stat,proto3" json:"stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTaskRequest) Reset()         { *m = UpdateTaskRequest{} }
func (m *UpdateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskRequest) ProtoMessage()    {}
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{2}
}

func (m *UpdateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskRequest.Unmarshal(m, b)
}
func (m *UpdateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskRequest.Merge(m, src)
}
func (m *UpdateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskRequest.Size(m)
}
func (m *UpdateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskRequest proto.InternalMessageInfo

func (m *UpdateTaskRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateTaskRequest) GetStat() string {
	if m != nil {
		return m.Stat
	}
	return ""
}

type UpdateTaskResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTaskResponse) Reset()         { *m = UpdateTaskResponse{} }
func (m *UpdateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskResponse) ProtoMessage()    {}
func (*UpdateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{3}
}

func (m *UpdateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskResponse.Unmarshal(m, b)
}
func (m *UpdateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskResponse.Marshal(b, m, deterministic)
}
func (m *UpdateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskResponse.Merge(m, src)
}
func (m *UpdateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskResponse.Size(m)
}
func (m *UpdateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskResponse proto.InternalMessageInfo

func (m *UpdateTaskResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SaveAnswerRequest struct {
	TaskId               int64    `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Answer               int64    `protobuf:"varint,2,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveAnswerRequest) Reset()         { *m = SaveAnswerRequest{} }
func (m *SaveAnswerRequest) String() string { return proto.CompactTextString(m) }
func (*SaveAnswerRequest) ProtoMessage()    {}
func (*SaveAnswerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{4}
}

func (m *SaveAnswerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveAnswerRequest.Unmarshal(m, b)
}
func (m *SaveAnswerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveAnswerRequest.Marshal(b, m, deterministic)
}
func (m *SaveAnswerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveAnswerRequest.Merge(m, src)
}
func (m *SaveAnswerRequest) XXX_Size() int {
	return xxx_messageInfo_SaveAnswerRequest.Size(m)
}
func (m *SaveAnswerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveAnswerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveAnswerRequest proto.InternalMessageInfo

func (m *SaveAnswerRequest) GetTaskId() int64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *SaveAnswerRequest) GetAnswer() int64 {
	if m != nil {
		return m.Answer
	}
	return 0
}

type SaveAnswerResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveAnswerResponse) Reset()         { *m = SaveAnswerResponse{} }
func (m *SaveAnswerResponse) String() string { return proto.CompactTextString(m) }
func (*SaveAnswerResponse) ProtoMessage()    {}
func (*SaveAnswerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{5}
}

func (m *SaveAnswerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveAnswerResponse.Unmarshal(m, b)
}
func (m *SaveAnswerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveAnswerResponse.Marshal(b, m, deterministic)
}
func (m *SaveAnswerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveAnswerResponse.Merge(m, src)
}
func (m *SaveAnswerResponse) XXX_Size() int {
	return xxx_messageInfo_SaveAnswerResponse.Size(m)
}
func (m *SaveAnswerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveAnswerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveAnswerResponse proto.InternalMessageInfo

func (m *SaveAnswerResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetAnswerRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TaskId               int64    `protobuf:"varint,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAnswerRequest) Reset()         { *m = GetAnswerRequest{} }
func (m *GetAnswerRequest) String() string { return proto.CompactTextString(m) }
func (*GetAnswerRequest) ProtoMessage()    {}
func (*GetAnswerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{6}
}

func (m *GetAnswerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAnswerRequest.Unmarshal(m, b)
}
func (m *GetAnswerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAnswerRequest.Marshal(b, m, deterministic)
}
func (m *GetAnswerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAnswerRequest.Merge(m, src)
}
func (m *GetAnswerRequest) XXX_Size() int {
	return xxx_messageInfo_GetAnswerRequest.Size(m)
}
func (m *GetAnswerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAnswerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAnswerRequest proto.InternalMessageInfo

func (m *GetAnswerRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetAnswerRequest) GetTaskId() int64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

type GetAnswerResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Answer               string   `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAnswerResponse) Reset()         { *m = GetAnswerResponse{} }
func (m *GetAnswerResponse) String() string { return proto.CompactTextString(m) }
func (*GetAnswerResponse) ProtoMessage()    {}
func (*GetAnswerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{7}
}

func (m *GetAnswerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAnswerResponse.Unmarshal(m, b)
}
func (m *GetAnswerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAnswerResponse.Marshal(b, m, deterministic)
}
func (m *GetAnswerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAnswerResponse.Merge(m, src)
}
func (m *GetAnswerResponse) XXX_Size() int {
	return xxx_messageInfo_GetAnswerResponse.Size(m)
}
func (m *GetAnswerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAnswerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAnswerResponse proto.InternalMessageInfo

func (m *GetAnswerResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetAnswerResponse) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

type SaveTaskRequest struct {
	IdUser               int64    `protobuf:"varint,1,opt,name=id_user,json=idUser,proto3" json:"id_user,omitempty"`
	Task                 string   `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveTaskRequest) Reset()         { *m = SaveTaskRequest{} }
func (m *SaveTaskRequest) String() string { return proto.CompactTextString(m) }
func (*SaveTaskRequest) ProtoMessage()    {}
func (*SaveTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{8}
}

func (m *SaveTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveTaskRequest.Unmarshal(m, b)
}
func (m *SaveTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveTaskRequest.Marshal(b, m, deterministic)
}
func (m *SaveTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveTaskRequest.Merge(m, src)
}
func (m *SaveTaskRequest) XXX_Size() int {
	return xxx_messageInfo_SaveTaskRequest.Size(m)
}
func (m *SaveTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveTaskRequest proto.InternalMessageInfo

func (m *SaveTaskRequest) GetIdUser() int64 {
	if m != nil {
		return m.IdUser
	}
	return 0
}

func (m *SaveTaskRequest) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

type SaveTaskResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveTaskResponse) Reset()         { *m = SaveTaskResponse{} }
func (m *SaveTaskResponse) String() string { return proto.CompactTextString(m) }
func (*SaveTaskResponse) ProtoMessage()    {}
func (*SaveTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{9}
}

func (m *SaveTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveTaskResponse.Unmarshal(m, b)
}
func (m *SaveTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveTaskResponse.Marshal(b, m, deterministic)
}
func (m *SaveTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveTaskResponse.Merge(m, src)
}
func (m *SaveTaskResponse) XXX_Size() int {
	return xxx_messageInfo_SaveTaskResponse.Size(m)
}
func (m *SaveTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveTaskResponse proto.InternalMessageInfo

func (m *SaveTaskResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetTaskRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IdUser               int64    `protobuf:"varint,2,opt,name=id_user,json=idUser,proto3" json:"id_user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTaskRequest) Reset()         { *m = GetTaskRequest{} }
func (m *GetTaskRequest) String() string { return proto.CompactTextString(m) }
func (*GetTaskRequest) ProtoMessage()    {}
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{10}
}

func (m *GetTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskRequest.Unmarshal(m, b)
}
func (m *GetTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskRequest.Marshal(b, m, deterministic)
}
func (m *GetTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskRequest.Merge(m, src)
}
func (m *GetTaskRequest) XXX_Size() int {
	return xxx_messageInfo_GetTaskRequest.Size(m)
}
func (m *GetTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskRequest proto.InternalMessageInfo

func (m *GetTaskRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetTaskRequest) GetIdUser() int64 {
	if m != nil {
		return m.IdUser
	}
	return 0
}

type GetTaskResponse struct {
	Task                 string   `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTaskResponse) Reset()         { *m = GetTaskResponse{} }
func (m *GetTaskResponse) String() string { return proto.CompactTextString(m) }
func (*GetTaskResponse) ProtoMessage()    {}
func (*GetTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{11}
}

func (m *GetTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskResponse.Unmarshal(m, b)
}
func (m *GetTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskResponse.Marshal(b, m, deterministic)
}
func (m *GetTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskResponse.Merge(m, src)
}
func (m *GetTaskResponse) XXX_Size() int {
	return xxx_messageInfo_GetTaskResponse.Size(m)
}
func (m *GetTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskResponse proto.InternalMessageInfo

func (m *GetTaskResponse) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *GetTaskResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type SaveUserRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveUserRequest) Reset()         { *m = SaveUserRequest{} }
func (m *SaveUserRequest) String() string { return proto.CompactTextString(m) }
func (*SaveUserRequest) ProtoMessage()    {}
func (*SaveUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{12}
}

func (m *SaveUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveUserRequest.Unmarshal(m, b)
}
func (m *SaveUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveUserRequest.Marshal(b, m, deterministic)
}
func (m *SaveUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveUserRequest.Merge(m, src)
}
func (m *SaveUserRequest) XXX_Size() int {
	return xxx_messageInfo_SaveUserRequest.Size(m)
}
func (m *SaveUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveUserRequest proto.InternalMessageInfo

func (m *SaveUserRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *SaveUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SaveUserResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveUserResponse) Reset()         { *m = SaveUserResponse{} }
func (m *SaveUserResponse) String() string { return proto.CompactTextString(m) }
func (*SaveUserResponse) ProtoMessage()    {}
func (*SaveUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{13}
}

func (m *SaveUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveUserResponse.Unmarshal(m, b)
}
func (m *SaveUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveUserResponse.Marshal(b, m, deterministic)
}
func (m *SaveUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveUserResponse.Merge(m, src)
}
func (m *SaveUserResponse) XXX_Size() int {
	return xxx_messageInfo_SaveUserResponse.Size(m)
}
func (m *SaveUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveUserResponse proto.InternalMessageInfo

func (m *SaveUserResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetUserRequest struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{14}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *GetUserRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type GetUserResponse struct {
	IdUser               int64    `protobuf:"varint,1,opt,name=id_user,json=idUser,proto3" json:"id_user,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa90353ee2d235cd, []int{15}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetIdUser() int64 {
	if m != nil {
		return m.IdUser
	}
	return 0
}

func (m *GetUserResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *GetUserResponse) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*WorkRequest)(nil), "services.WorkRequest")
	proto.RegisterType((*WorkResponse)(nil), "services.WorkResponse")
	proto.RegisterType((*UpdateTaskRequest)(nil), "services.UpdateTaskRequest")
	proto.RegisterType((*UpdateTaskResponse)(nil), "services.UpdateTaskResponse")
	proto.RegisterType((*SaveAnswerRequest)(nil), "services.SaveAnswerRequest")
	proto.RegisterType((*SaveAnswerResponse)(nil), "services.SaveAnswerResponse")
	proto.RegisterType((*GetAnswerRequest)(nil), "services.GetAnswerRequest")
	proto.RegisterType((*GetAnswerResponse)(nil), "services.GetAnswerResponse")
	proto.RegisterType((*SaveTaskRequest)(nil), "services.SaveTaskRequest")
	proto.RegisterType((*SaveTaskResponse)(nil), "services.SaveTaskResponse")
	proto.RegisterType((*GetTaskRequest)(nil), "services.GetTaskRequest")
	proto.RegisterType((*GetTaskResponse)(nil), "services.GetTaskResponse")
	proto.RegisterType((*SaveUserRequest)(nil), "services.SaveUserRequest")
	proto.RegisterType((*SaveUserResponse)(nil), "services.SaveUserResponse")
	proto.RegisterType((*GetUserRequest)(nil), "services.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "services.GetUserResponse")
}

func init() {
	proto.RegisterFile("serviceDataBase.proto", fileDescriptor_fa90353ee2d235cd)
}

var fileDescriptor_fa90353ee2d235cd = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0x4f, 0x6b, 0xdb, 0x30,
	0x14, 0x27, 0x4e, 0x48, 0x9c, 0xb7, 0xd1, 0x24, 0x1a, 0x49, 0x3c, 0x77, 0xb0, 0x22, 0x76, 0xe8,
	0x29, 0xd0, 0xed, 0x30, 0x4a, 0x68, 0xa1, 0x25, 0x10, 0x72, 0xd9, 0xc1, 0x5d, 0x19, 0xec, 0x12,
	0xd4, 0x59, 0x03, 0x53, 0x5a, 0xa7, 0x96, 0x9c, 0x7e, 0xa7, 0x7d, 0xca, 0xa1, 0xbf, 0x96, 0x3c,
	0x79, 0xa7, 0xf8, 0xe9, 0xe9, 0xfd, 0xfe, 0xe8, 0xe9, 0x29, 0x30, 0x67, 0xb4, 0x3a, 0x16, 0xbf,
	0xe8, 0x86, 0x70, 0x72, 0x4b, 0x18, 0x5d, 0x1d, 0xaa, 0x92, 0x97, 0x28, 0xd6, 0xcb, 0x0c, 0x7f,
	0x84, 0x37, 0x3f, 0xca, 0xea, 0x31, 0xa3, 0x2f, 0x35, 0x65, 0x1c, 0x4d, 0xa1, 0x5f, 0xd1, 0x97,
	0xa4, 0x77, 0xd6, 0x3b, 0x8f, 0x33, 0xf1, 0x89, 0xf7, 0xf0, 0x56, 0x6d, 0x60, 0x87, 0xf2, 0x99,
	0x51, 0x74, 0x02, 0x51, 0x91, 0xcb, 0x0d, 0xfd, 0x2c, 0x2a, 0x72, 0xb4, 0x84, 0x51, 0xcd, 0x68,
	0xb5, 0x2f, 0xf2, 0x24, 0x92, 0x8b, 0x43, 0x11, 0xee, 0x72, 0x84, 0x60, 0xc0, 0x09, 0x7b, 0x4c,
	0xfa, 0x67, 0xbd, 0xf3, 0x71, 0x26, 0xbf, 0xc5, 0x1a, 0xe3, 0x84, 0x27, 0x03, 0xb5, 0x26, 0xbe,
	0xf1, 0x57, 0x98, 0xdd, 0x1f, 0x72, 0xc2, 0xe9, 0x77, 0xc2, 0xac, 0x8e, 0x36, 0x8b, 0x29, 0x8c,
	0x9c, 0xc2, 0x4f, 0x80, 0xdc, 0xc2, 0xb0, 0x3e, 0xbc, 0x81, 0xd9, 0x1d, 0x39, 0xd2, 0x9b, 0x67,
	0xf6, 0x4a, 0x2b, 0x03, 0xbf, 0x84, 0x91, 0xd0, 0xb3, 0xb7, 0x3b, 0x87, 0x22, 0xdc, 0xe5, 0x68,
	0x01, 0x43, 0x22, 0x77, 0x1a, 0x33, 0x2a, 0x12, 0x5c, 0x2e, 0x4a, 0x07, 0xd7, 0x1a, 0xa6, 0x5b,
	0xca, 0x7d, 0xaa, 0xc0, 0x79, 0x19, 0xea, 0xc8, 0xa5, 0xc6, 0x6b, 0x98, 0x39, 0xc5, 0x1d, 0xa7,
	0xed, 0xeb, 0x1b, 0x5b, 0x7d, 0xd7, 0x30, 0x11, 0xfa, 0xdc, 0x23, 0x5c, 0xc2, 0xa8, 0xc8, 0xf7,
	0xa2, 0x19, 0xc6, 0x63, 0x91, 0xdf, 0x33, 0x5a, 0xd9, 0xc6, 0x44, 0x4d, 0x63, 0x30, 0x86, 0x69,
	0x53, 0xdf, 0xe1, 0xee, 0x12, 0x4e, 0xb6, 0x94, 0xff, 0xaf, 0x4b, 0x0e, 0x65, 0xe4, 0x52, 0xe2,
	0x2b, 0x98, 0xd8, 0x52, 0x8d, 0x6e, 0x54, 0xf4, 0x9c, 0xeb, 0xb1, 0x80, 0xa1, 0xe8, 0x6c, 0xcd,
	0x8c, 0x3b, 0x15, 0xe1, 0x9d, 0x72, 0x27, 0xa0, 0x0c, 0x75, 0x0a, 0xb1, 0x08, 0xbf, 0x91, 0x27,
	0xaa, 0x21, 0x6c, 0x2c, 0x72, 0x07, 0xc2, 0xd8, 0x6b, 0x59, 0xe5, 0x1a, 0xc8, 0xc6, 0xc6, 0xa8,
	0x82, 0xea, 0x30, 0x7a, 0x25, 0x8d, 0xba, 0x6c, 0x08, 0x06, 0xf6, 0x20, 0xc7, 0x99, 0xfc, 0xee,
	0xbc, 0xf8, 0xf8, 0x41, 0x9a, 0xf5, 0x18, 0x3a, 0x7b, 0x91, 0x42, 0x5c, 0x1b, 0x1b, 0x5a, 0x6a,
	0x1d, 0xb2, 0xd1, 0xf7, 0x6d, 0x7c, 0xfe, 0x33, 0x80, 0x89, 0x99, 0xe9, 0x3b, 0x35, 0xcb, 0xe8,
	0x1a, 0x46, 0x9a, 0x17, 0x25, 0x2b, 0x33, 0xe0, 0x2b, 0xdf, 0x49, 0xfa, 0x3e, 0x90, 0xd1, 0x22,
	0x6f, 0x20, 0x36, 0x47, 0x83, 0x9c, 0x6d, 0xad, 0x93, 0x4f, 0xd3, 0x50, 0x4a, 0x43, 0x28, 0x09,
	0xa2, 0xcf, 0x2d, 0x09, 0xce, 0xad, 0x69, 0x49, 0xf0, 0x2e, 0x85, 0x96, 0x20, 0x01, 0x5a, 0x12,
	0x5c, 0x84, 0x34, 0x94, 0xd2, 0x10, 0x1b, 0x18, 0xdb, 0x31, 0x42, 0xa9, 0x47, 0xe5, 0x0d, 0x66,
	0x7a, 0x1a, 0xcc, 0x69, 0x94, 0x2d, 0x40, 0x33, 0xef, 0xe8, 0xd4, 0xe7, 0xf3, 0x71, 0x3e, 0x84,
	0x93, 0x0d, 0x50, 0xf3, 0x48, 0xb9, 0x40, 0xff, 0xbc, 0x79, 0x2e, 0x50, 0xe0, 0x5d, 0xbb, 0x84,
	0x58, 0xbc, 0xc3, 0x12, 0x66, 0xde, 0xec, 0x74, 0x1e, 0xef, 0x74, 0xd1, 0x5e, 0x56, 0xa5, 0xb7,
	0xf3, 0x9f, 0xef, 0x48, 0x55, 0xfc, 0x7e, 0x92, 0xe9, 0xd5, 0xf1, 0x62, 0x2d, 0x7e, 0x2f, 0x1e,
	0x86, 0xf2, 0xbf, 0xe0, 0xcb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x49, 0x46, 0x4b, 0x24,
	0x06, 0x00, 0x00,
}
