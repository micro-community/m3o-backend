// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/sprints/sprints.proto

package go_micro_api_distributed

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

type Status int32

const (
	Status_PENDING     Status = 0
	Status_IN_PROGRESS Status = 1
	Status_BLOCKED     Status = 2
	Status_COMPLETED   Status = 3
	Status_CANCELLED   Status = 4
)

var Status_name = map[int32]string{
	0: "PENDING",
	1: "IN_PROGRESS",
	2: "BLOCKED",
	3: "COMPLETED",
	4: "CANCELLED",
}

var Status_value = map[string]int32{
	"PENDING":     0,
	"IN_PROGRESS": 1,
	"BLOCKED":     2,
	"COMPLETED":   3,
	"CANCELLED":   4,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{0}
}

type Sprint struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StartTime            int64        `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              int64        `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Objectives           []*Objective `protobuf:"bytes,5,rep,name=objectives,proto3" json:"objectives,omitempty"`
	Tasks                []*Task      `protobuf:"bytes,6,rep,name=tasks,proto3" json:"tasks,omitempty"`
	Created              int64        `protobuf:"varint,7,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Sprint) Reset()         { *m = Sprint{} }
func (m *Sprint) String() string { return proto.CompactTextString(m) }
func (*Sprint) ProtoMessage()    {}
func (*Sprint) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{0}
}

func (m *Sprint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sprint.Unmarshal(m, b)
}
func (m *Sprint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sprint.Marshal(b, m, deterministic)
}
func (m *Sprint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sprint.Merge(m, src)
}
func (m *Sprint) XXX_Size() int {
	return xxx_messageInfo_Sprint.Size(m)
}
func (m *Sprint) XXX_DiscardUnknown() {
	xxx_messageInfo_Sprint.DiscardUnknown(m)
}

var xxx_messageInfo_Sprint proto.InternalMessageInfo

func (m *Sprint) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Sprint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Sprint) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Sprint) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *Sprint) GetObjectives() []*Objective {
	if m != nil {
		return m.Objectives
	}
	return nil
}

func (m *Sprint) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *Sprint) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

type Objective struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               Status   `protobuf:"varint,3,opt,name=status,proto3,enum=go.micro.api.distributed.Status" json:"status,omitempty"`
	Created              int64    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Objective) Reset()         { *m = Objective{} }
func (m *Objective) String() string { return proto.CompactTextString(m) }
func (*Objective) ProtoMessage()    {}
func (*Objective) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{1}
}

func (m *Objective) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Objective.Unmarshal(m, b)
}
func (m *Objective) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Objective.Marshal(b, m, deterministic)
}
func (m *Objective) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Objective.Merge(m, src)
}
func (m *Objective) XXX_Size() int {
	return xxx_messageInfo_Objective.Size(m)
}
func (m *Objective) XXX_DiscardUnknown() {
	xxx_messageInfo_Objective.DiscardUnknown(m)
}

var xxx_messageInfo_Objective proto.InternalMessageInfo

func (m *Objective) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Objective) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Objective) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_PENDING
}

func (m *Objective) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

type Task struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               Status   `protobuf:"varint,3,opt,name=status,proto3,enum=go.micro.api.distributed.Status" json:"status,omitempty"`
	Created              int64    `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{2}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_PENDING
}

func (m *Task) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

type CreateSprintRequest struct {
	Sprint               *Sprint  `protobuf:"bytes,1,opt,name=sprint,proto3" json:"sprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSprintRequest) Reset()         { *m = CreateSprintRequest{} }
func (m *CreateSprintRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSprintRequest) ProtoMessage()    {}
func (*CreateSprintRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{3}
}

func (m *CreateSprintRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSprintRequest.Unmarshal(m, b)
}
func (m *CreateSprintRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSprintRequest.Marshal(b, m, deterministic)
}
func (m *CreateSprintRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSprintRequest.Merge(m, src)
}
func (m *CreateSprintRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSprintRequest.Size(m)
}
func (m *CreateSprintRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSprintRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSprintRequest proto.InternalMessageInfo

func (m *CreateSprintRequest) GetSprint() *Sprint {
	if m != nil {
		return m.Sprint
	}
	return nil
}

type ListSprintsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSprintsRequest) Reset()         { *m = ListSprintsRequest{} }
func (m *ListSprintsRequest) String() string { return proto.CompactTextString(m) }
func (*ListSprintsRequest) ProtoMessage()    {}
func (*ListSprintsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{4}
}

func (m *ListSprintsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSprintsRequest.Unmarshal(m, b)
}
func (m *ListSprintsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSprintsRequest.Marshal(b, m, deterministic)
}
func (m *ListSprintsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSprintsRequest.Merge(m, src)
}
func (m *ListSprintsRequest) XXX_Size() int {
	return xxx_messageInfo_ListSprintsRequest.Size(m)
}
func (m *ListSprintsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSprintsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSprintsRequest proto.InternalMessageInfo

type ReadSprintRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadSprintRequest) Reset()         { *m = ReadSprintRequest{} }
func (m *ReadSprintRequest) String() string { return proto.CompactTextString(m) }
func (*ReadSprintRequest) ProtoMessage()    {}
func (*ReadSprintRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{5}
}

func (m *ReadSprintRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadSprintRequest.Unmarshal(m, b)
}
func (m *ReadSprintRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadSprintRequest.Marshal(b, m, deterministic)
}
func (m *ReadSprintRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadSprintRequest.Merge(m, src)
}
func (m *ReadSprintRequest) XXX_Size() int {
	return xxx_messageInfo_ReadSprintRequest.Size(m)
}
func (m *ReadSprintRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadSprintRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadSprintRequest proto.InternalMessageInfo

func (m *ReadSprintRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateTaskRequest struct {
	SprintId             string   `protobuf:"bytes,1,opt,name=sprint_id,json=sprintId,proto3" json:"sprint_id,omitempty"`
	Task                 *Task    `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{6}
}

func (m *CreateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRequest.Unmarshal(m, b)
}
func (m *CreateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRequest.Marshal(b, m, deterministic)
}
func (m *CreateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRequest.Merge(m, src)
}
func (m *CreateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRequest.Size(m)
}
func (m *CreateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRequest proto.InternalMessageInfo

func (m *CreateTaskRequest) GetSprintId() string {
	if m != nil {
		return m.SprintId
	}
	return ""
}

func (m *CreateTaskRequest) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type UpdateTaskRequest struct {
	SprintId             string   `protobuf:"bytes,1,opt,name=sprint_id,json=sprintId,proto3" json:"sprint_id,omitempty"`
	Task                 *Task    `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTaskRequest) Reset()         { *m = UpdateTaskRequest{} }
func (m *UpdateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskRequest) ProtoMessage()    {}
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{7}
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

func (m *UpdateTaskRequest) GetSprintId() string {
	if m != nil {
		return m.SprintId
	}
	return ""
}

func (m *UpdateTaskRequest) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type DeleteTaskRequest struct {
	SprintId             string   `protobuf:"bytes,1,opt,name=sprint_id,json=sprintId,proto3" json:"sprint_id,omitempty"`
	TaskId               string   `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskRequest) Reset()         { *m = DeleteTaskRequest{} }
func (m *DeleteTaskRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskRequest) ProtoMessage()    {}
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{8}
}

func (m *DeleteTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskRequest.Unmarshal(m, b)
}
func (m *DeleteTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskRequest.Merge(m, src)
}
func (m *DeleteTaskRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskRequest.Size(m)
}
func (m *DeleteTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskRequest proto.InternalMessageInfo

func (m *DeleteTaskRequest) GetSprintId() string {
	if m != nil {
		return m.SprintId
	}
	return ""
}

func (m *DeleteTaskRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type CreateObjectiveRequest struct {
	SprintId             string     `protobuf:"bytes,1,opt,name=sprint_id,json=sprintId,proto3" json:"sprint_id,omitempty"`
	Objective            *Objective `protobuf:"bytes,2,opt,name=objective,proto3" json:"objective,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateObjectiveRequest) Reset()         { *m = CreateObjectiveRequest{} }
func (m *CreateObjectiveRequest) String() string { return proto.CompactTextString(m) }
func (*CreateObjectiveRequest) ProtoMessage()    {}
func (*CreateObjectiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{9}
}

func (m *CreateObjectiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateObjectiveRequest.Unmarshal(m, b)
}
func (m *CreateObjectiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateObjectiveRequest.Marshal(b, m, deterministic)
}
func (m *CreateObjectiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateObjectiveRequest.Merge(m, src)
}
func (m *CreateObjectiveRequest) XXX_Size() int {
	return xxx_messageInfo_CreateObjectiveRequest.Size(m)
}
func (m *CreateObjectiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateObjectiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateObjectiveRequest proto.InternalMessageInfo

func (m *CreateObjectiveRequest) GetSprintId() string {
	if m != nil {
		return m.SprintId
	}
	return ""
}

func (m *CreateObjectiveRequest) GetObjective() *Objective {
	if m != nil {
		return m.Objective
	}
	return nil
}

type UpdateObjectiveRequest struct {
	SprintId             string     `protobuf:"bytes,1,opt,name=sprint_id,json=sprintId,proto3" json:"sprint_id,omitempty"`
	Objective            *Objective `protobuf:"bytes,2,opt,name=objective,proto3" json:"objective,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateObjectiveRequest) Reset()         { *m = UpdateObjectiveRequest{} }
func (m *UpdateObjectiveRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateObjectiveRequest) ProtoMessage()    {}
func (*UpdateObjectiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{10}
}

func (m *UpdateObjectiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateObjectiveRequest.Unmarshal(m, b)
}
func (m *UpdateObjectiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateObjectiveRequest.Marshal(b, m, deterministic)
}
func (m *UpdateObjectiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateObjectiveRequest.Merge(m, src)
}
func (m *UpdateObjectiveRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateObjectiveRequest.Size(m)
}
func (m *UpdateObjectiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateObjectiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateObjectiveRequest proto.InternalMessageInfo

func (m *UpdateObjectiveRequest) GetSprintId() string {
	if m != nil {
		return m.SprintId
	}
	return ""
}

func (m *UpdateObjectiveRequest) GetObjective() *Objective {
	if m != nil {
		return m.Objective
	}
	return nil
}

type DeleteObjectiveRequest struct {
	SprintId             string   `protobuf:"bytes,1,opt,name=sprint_id,json=sprintId,proto3" json:"sprint_id,omitempty"`
	ObjectiveId          string   `protobuf:"bytes,2,opt,name=objective_id,json=objectiveId,proto3" json:"objective_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteObjectiveRequest) Reset()         { *m = DeleteObjectiveRequest{} }
func (m *DeleteObjectiveRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteObjectiveRequest) ProtoMessage()    {}
func (*DeleteObjectiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{11}
}

func (m *DeleteObjectiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteObjectiveRequest.Unmarshal(m, b)
}
func (m *DeleteObjectiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteObjectiveRequest.Marshal(b, m, deterministic)
}
func (m *DeleteObjectiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteObjectiveRequest.Merge(m, src)
}
func (m *DeleteObjectiveRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteObjectiveRequest.Size(m)
}
func (m *DeleteObjectiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteObjectiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteObjectiveRequest proto.InternalMessageInfo

func (m *DeleteObjectiveRequest) GetSprintId() string {
	if m != nil {
		return m.SprintId
	}
	return ""
}

func (m *DeleteObjectiveRequest) GetObjectiveId() string {
	if m != nil {
		return m.ObjectiveId
	}
	return ""
}

type CreateSprintResponse struct {
	Sprint               *Sprint  `protobuf:"bytes,1,opt,name=sprint,proto3" json:"sprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSprintResponse) Reset()         { *m = CreateSprintResponse{} }
func (m *CreateSprintResponse) String() string { return proto.CompactTextString(m) }
func (*CreateSprintResponse) ProtoMessage()    {}
func (*CreateSprintResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{12}
}

func (m *CreateSprintResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSprintResponse.Unmarshal(m, b)
}
func (m *CreateSprintResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSprintResponse.Marshal(b, m, deterministic)
}
func (m *CreateSprintResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSprintResponse.Merge(m, src)
}
func (m *CreateSprintResponse) XXX_Size() int {
	return xxx_messageInfo_CreateSprintResponse.Size(m)
}
func (m *CreateSprintResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSprintResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSprintResponse proto.InternalMessageInfo

func (m *CreateSprintResponse) GetSprint() *Sprint {
	if m != nil {
		return m.Sprint
	}
	return nil
}

type ListSprintsResponse struct {
	Sprints              []*Sprint `protobuf:"bytes,1,rep,name=sprints,proto3" json:"sprints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListSprintsResponse) Reset()         { *m = ListSprintsResponse{} }
func (m *ListSprintsResponse) String() string { return proto.CompactTextString(m) }
func (*ListSprintsResponse) ProtoMessage()    {}
func (*ListSprintsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{13}
}

func (m *ListSprintsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSprintsResponse.Unmarshal(m, b)
}
func (m *ListSprintsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSprintsResponse.Marshal(b, m, deterministic)
}
func (m *ListSprintsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSprintsResponse.Merge(m, src)
}
func (m *ListSprintsResponse) XXX_Size() int {
	return xxx_messageInfo_ListSprintsResponse.Size(m)
}
func (m *ListSprintsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSprintsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSprintsResponse proto.InternalMessageInfo

func (m *ListSprintsResponse) GetSprints() []*Sprint {
	if m != nil {
		return m.Sprints
	}
	return nil
}

type ReadSprintResponse struct {
	Sprint               *Sprint  `protobuf:"bytes,1,opt,name=sprint,proto3" json:"sprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadSprintResponse) Reset()         { *m = ReadSprintResponse{} }
func (m *ReadSprintResponse) String() string { return proto.CompactTextString(m) }
func (*ReadSprintResponse) ProtoMessage()    {}
func (*ReadSprintResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{14}
}

func (m *ReadSprintResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadSprintResponse.Unmarshal(m, b)
}
func (m *ReadSprintResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadSprintResponse.Marshal(b, m, deterministic)
}
func (m *ReadSprintResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadSprintResponse.Merge(m, src)
}
func (m *ReadSprintResponse) XXX_Size() int {
	return xxx_messageInfo_ReadSprintResponse.Size(m)
}
func (m *ReadSprintResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadSprintResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadSprintResponse proto.InternalMessageInfo

func (m *ReadSprintResponse) GetSprint() *Sprint {
	if m != nil {
		return m.Sprint
	}
	return nil
}

type CreateTaskResponse struct {
	Task                 *Task    `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskResponse) Reset()         { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()    {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{15}
}

func (m *CreateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse.Unmarshal(m, b)
}
func (m *CreateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse.Marshal(b, m, deterministic)
}
func (m *CreateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse.Merge(m, src)
}
func (m *CreateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponse.Size(m)
}
func (m *CreateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponse proto.InternalMessageInfo

func (m *CreateTaskResponse) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type UpdateTaskResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTaskResponse) Reset()         { *m = UpdateTaskResponse{} }
func (m *UpdateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskResponse) ProtoMessage()    {}
func (*UpdateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{16}
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

type DeleteTaskResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskResponse) Reset()         { *m = DeleteTaskResponse{} }
func (m *DeleteTaskResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskResponse) ProtoMessage()    {}
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{17}
}

func (m *DeleteTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskResponse.Unmarshal(m, b)
}
func (m *DeleteTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskResponse.Merge(m, src)
}
func (m *DeleteTaskResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskResponse.Size(m)
}
func (m *DeleteTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskResponse proto.InternalMessageInfo

type CreateObjectiveResponse struct {
	Objective            *Objective `protobuf:"bytes,1,opt,name=objective,proto3" json:"objective,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateObjectiveResponse) Reset()         { *m = CreateObjectiveResponse{} }
func (m *CreateObjectiveResponse) String() string { return proto.CompactTextString(m) }
func (*CreateObjectiveResponse) ProtoMessage()    {}
func (*CreateObjectiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{18}
}

func (m *CreateObjectiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateObjectiveResponse.Unmarshal(m, b)
}
func (m *CreateObjectiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateObjectiveResponse.Marshal(b, m, deterministic)
}
func (m *CreateObjectiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateObjectiveResponse.Merge(m, src)
}
func (m *CreateObjectiveResponse) XXX_Size() int {
	return xxx_messageInfo_CreateObjectiveResponse.Size(m)
}
func (m *CreateObjectiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateObjectiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateObjectiveResponse proto.InternalMessageInfo

func (m *CreateObjectiveResponse) GetObjective() *Objective {
	if m != nil {
		return m.Objective
	}
	return nil
}

type UpdateObjectiveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateObjectiveResponse) Reset()         { *m = UpdateObjectiveResponse{} }
func (m *UpdateObjectiveResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateObjectiveResponse) ProtoMessage()    {}
func (*UpdateObjectiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{19}
}

func (m *UpdateObjectiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateObjectiveResponse.Unmarshal(m, b)
}
func (m *UpdateObjectiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateObjectiveResponse.Marshal(b, m, deterministic)
}
func (m *UpdateObjectiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateObjectiveResponse.Merge(m, src)
}
func (m *UpdateObjectiveResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateObjectiveResponse.Size(m)
}
func (m *UpdateObjectiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateObjectiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateObjectiveResponse proto.InternalMessageInfo

type DeleteObjectiveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteObjectiveResponse) Reset()         { *m = DeleteObjectiveResponse{} }
func (m *DeleteObjectiveResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteObjectiveResponse) ProtoMessage()    {}
func (*DeleteObjectiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b280c0e093567fd, []int{20}
}

func (m *DeleteObjectiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteObjectiveResponse.Unmarshal(m, b)
}
func (m *DeleteObjectiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteObjectiveResponse.Marshal(b, m, deterministic)
}
func (m *DeleteObjectiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteObjectiveResponse.Merge(m, src)
}
func (m *DeleteObjectiveResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteObjectiveResponse.Size(m)
}
func (m *DeleteObjectiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteObjectiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteObjectiveResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("go.micro.api.distributed.Status", Status_name, Status_value)
	proto.RegisterType((*Sprint)(nil), "go.micro.api.distributed.Sprint")
	proto.RegisterType((*Objective)(nil), "go.micro.api.distributed.Objective")
	proto.RegisterType((*Task)(nil), "go.micro.api.distributed.Task")
	proto.RegisterType((*CreateSprintRequest)(nil), "go.micro.api.distributed.CreateSprintRequest")
	proto.RegisterType((*ListSprintsRequest)(nil), "go.micro.api.distributed.ListSprintsRequest")
	proto.RegisterType((*ReadSprintRequest)(nil), "go.micro.api.distributed.ReadSprintRequest")
	proto.RegisterType((*CreateTaskRequest)(nil), "go.micro.api.distributed.CreateTaskRequest")
	proto.RegisterType((*UpdateTaskRequest)(nil), "go.micro.api.distributed.UpdateTaskRequest")
	proto.RegisterType((*DeleteTaskRequest)(nil), "go.micro.api.distributed.DeleteTaskRequest")
	proto.RegisterType((*CreateObjectiveRequest)(nil), "go.micro.api.distributed.CreateObjectiveRequest")
	proto.RegisterType((*UpdateObjectiveRequest)(nil), "go.micro.api.distributed.UpdateObjectiveRequest")
	proto.RegisterType((*DeleteObjectiveRequest)(nil), "go.micro.api.distributed.DeleteObjectiveRequest")
	proto.RegisterType((*CreateSprintResponse)(nil), "go.micro.api.distributed.CreateSprintResponse")
	proto.RegisterType((*ListSprintsResponse)(nil), "go.micro.api.distributed.ListSprintsResponse")
	proto.RegisterType((*ReadSprintResponse)(nil), "go.micro.api.distributed.ReadSprintResponse")
	proto.RegisterType((*CreateTaskResponse)(nil), "go.micro.api.distributed.CreateTaskResponse")
	proto.RegisterType((*UpdateTaskResponse)(nil), "go.micro.api.distributed.UpdateTaskResponse")
	proto.RegisterType((*DeleteTaskResponse)(nil), "go.micro.api.distributed.DeleteTaskResponse")
	proto.RegisterType((*CreateObjectiveResponse)(nil), "go.micro.api.distributed.CreateObjectiveResponse")
	proto.RegisterType((*UpdateObjectiveResponse)(nil), "go.micro.api.distributed.UpdateObjectiveResponse")
	proto.RegisterType((*DeleteObjectiveResponse)(nil), "go.micro.api.distributed.DeleteObjectiveResponse")
}

func init() { proto.RegisterFile("proto/sprints/sprints.proto", fileDescriptor_2b280c0e093567fd) }

var fileDescriptor_2b280c0e093567fd = []byte{
	// 742 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x96, 0xc1, 0x73, 0xd2, 0x4e,
	0x14, 0xc7, 0x09, 0x50, 0x28, 0x8f, 0xfe, 0x5a, 0xfa, 0xda, 0x29, 0x29, 0x9d, 0x9f, 0x83, 0xe9,
	0xa5, 0xa3, 0x6d, 0x54, 0xf4, 0xe0, 0x78, 0xab, 0x90, 0xa9, 0x8c, 0x08, 0x34, 0xd4, 0x19, 0x0f,
	0xce, 0x74, 0xd2, 0xee, 0x8e, 0x13, 0x5b, 0x08, 0xb2, 0x5b, 0xa7, 0x27, 0x2f, 0x9e, 0xfd, 0x8b,
	0xbd, 0x38, 0xd9, 0x5d, 0x48, 0x20, 0x24, 0x05, 0x3b, 0xea, 0x89, 0xe4, 0xed, 0x7b, 0xef, 0xf3,
	0x76, 0xf3, 0x7d, 0x6f, 0x81, 0xbd, 0xe1, 0xc8, 0xe3, 0xde, 0x13, 0x36, 0x1c, 0xb9, 0x03, 0xce,
	0xc6, 0xbf, 0xa6, 0xb0, 0xa2, 0xfe, 0xc9, 0x33, 0xfb, 0xee, 0xe5, 0xc8, 0x33, 0x9d, 0xa1, 0x6b,
	0x12, 0x97, 0xf1, 0x91, 0x7b, 0x71, 0xc3, 0x29, 0x31, 0x7e, 0x6a, 0x90, 0xeb, 0x09, 0x5f, 0x5c,
	0x87, 0xb4, 0x4b, 0x74, 0xad, 0xaa, 0x1d, 0x14, 0xec, 0xb4, 0x4b, 0x10, 0x21, 0x3b, 0x70, 0xfa,
	0x54, 0x4f, 0x0b, 0x8b, 0x78, 0xc6, 0xff, 0x01, 0x18, 0x77, 0x46, 0xfc, 0x9c, 0xbb, 0x7d, 0xaa,
	0x67, 0xaa, 0xda, 0x41, 0xc6, 0x2e, 0x08, 0xcb, 0x99, 0xdb, 0xa7, 0xb8, 0x0b, 0xab, 0x74, 0x40,
	0xe4, 0x62, 0x56, 0x2c, 0xe6, 0xe9, 0x80, 0x88, 0xa5, 0x3a, 0x80, 0x77, 0xf1, 0x99, 0x5e, 0x72,
	0xf7, 0x2b, 0x65, 0xfa, 0x4a, 0x35, 0x73, 0x50, 0xac, 0xed, 0x9b, 0x71, 0x75, 0x99, 0x9d, 0xb1,
	0xaf, 0x1d, 0x0a, 0xc3, 0x17, 0xb0, 0xc2, 0x1d, 0x76, 0xc5, 0xf4, 0x9c, 0x88, 0x7f, 0x10, 0x1f,
	0x7f, 0xe6, 0xb0, 0x2b, 0x5b, 0x3a, 0xa3, 0x0e, 0xf9, 0xcb, 0x11, 0x75, 0x38, 0x25, 0x7a, 0x5e,
	0x16, 0xa5, 0x5e, 0x8d, 0xef, 0x1a, 0x14, 0x26, 0xa4, 0x85, 0x0e, 0xe0, 0x25, 0xe4, 0x18, 0x77,
	0xf8, 0x0d, 0x13, 0x9b, 0x5f, 0xaf, 0x55, 0xe3, 0x4b, 0xe8, 0x09, 0x3f, 0x5b, 0xf9, 0x87, 0xab,
	0xc8, 0x4e, 0x57, 0xf1, 0x0d, 0xb2, 0x7e, 0xb9, 0xff, 0x8c, 0xdf, 0x81, 0xad, 0xba, 0x78, 0x94,
	0x42, 0xb0, 0xe9, 0x97, 0x1b, 0xca, 0xb8, 0x40, 0x09, 0x83, 0x28, 0xa9, 0x98, 0x88, 0x92, 0x81,
	0xca, 0xdf, 0xd8, 0x06, 0x6c, 0xb9, 0x8c, 0x4b, 0x2b, 0x53, 0xf9, 0x8c, 0x7d, 0xd8, 0xb4, 0xa9,
	0x43, 0xa6, 0x21, 0x33, 0x7b, 0x36, 0x08, 0x6c, 0xca, 0x5a, 0xc4, 0x07, 0x54, 0x4e, 0x7b, 0x50,
	0x90, 0x99, 0xcf, 0x27, 0xbe, 0xab, 0xd2, 0xd0, 0x24, 0x58, 0x83, 0xac, 0xff, 0x99, 0xc5, 0x29,
	0xdd, 0x2d, 0x09, 0xe1, 0xeb, 0x53, 0xde, 0x0f, 0xc9, 0x9f, 0xa6, 0x34, 0x61, 0xb3, 0x41, 0xaf,
	0xe9, 0x12, 0x94, 0x32, 0xe4, 0xfd, 0x48, 0x7f, 0x49, 0x7e, 0xf4, 0x9c, 0xff, 0xda, 0x24, 0xc6,
	0x2d, 0xec, 0xc8, 0x63, 0x09, 0xfa, 0x62, 0x91, 0x7c, 0xc7, 0x50, 0x98, 0x74, 0x8f, 0x2a, 0x7d,
	0xa1, 0x9e, 0x0b, 0xa2, 0x7c, 0xb2, 0x3c, 0xaa, 0xbf, 0x4e, 0xfe, 0x00, 0x3b, 0xf2, 0xf8, 0x96,
	0x23, 0x3f, 0x84, 0xb5, 0x49, 0x8e, 0xe0, 0x20, 0x8b, 0x13, 0x5b, 0x93, 0x18, 0x5d, 0xd8, 0x9e,
	0x16, 0x3c, 0x1b, 0x7a, 0x03, 0x46, 0xef, 0xa1, 0xf8, 0x53, 0xd8, 0x9a, 0x52, 0xbc, 0x4a, 0xf8,
	0x0a, 0xf2, 0x6a, 0x10, 0xeb, 0x9a, 0x98, 0x58, 0x77, 0x67, 0x1c, 0x07, 0x18, 0x6d, 0xc0, 0x70,
	0xbb, 0xdc, 0xbb, 0xc4, 0x37, 0x80, 0xe1, 0xce, 0x52, 0xf9, 0xc6, 0xba, 0xd6, 0x96, 0xd0, 0xf5,
	0x36, 0x60, 0xb8, 0x7b, 0x64, 0x26, 0xdf, 0x1a, 0x56, 0xbb, 0xb2, 0x7e, 0x84, 0x72, 0x44, 0xb8,
	0x0a, 0x3d, 0x25, 0x11, 0xed, 0xb7, 0x24, 0xb2, 0x0b, 0xe5, 0x88, 0x38, 0x15, 0x78, 0x17, 0xca,
	0x11, 0xf5, 0xc8, 0xa5, 0x47, 0xa7, 0x90, 0x93, 0xb3, 0x11, 0x8b, 0x90, 0xef, 0x5a, 0xed, 0x46,
	0xb3, 0x7d, 0x52, 0x4a, 0xe1, 0x06, 0x14, 0x9b, 0xed, 0xf3, 0xae, 0xdd, 0x39, 0xb1, 0xad, 0x5e,
	0xaf, 0xa4, 0xf9, 0xab, 0xaf, 0x5b, 0x9d, 0xfa, 0x5b, 0xab, 0x51, 0x4a, 0xe3, 0x7f, 0x50, 0xa8,
	0x77, 0xde, 0x75, 0x5b, 0xd6, 0x99, 0xd5, 0x28, 0x65, 0xc4, 0xeb, 0x71, 0xbb, 0x6e, 0xb5, 0x5a,
	0x56, 0xa3, 0x94, 0xad, 0xfd, 0x58, 0x05, 0x6c, 0x04, 0xd5, 0x2a, 0x1d, 0xa0, 0x07, 0x6b, 0x61,
	0xa1, 0xe1, 0x51, 0xfc, 0xfe, 0xe6, 0x4c, 0xe0, 0x8a, 0xb9, 0xa8, 0xbb, 0xda, 0x73, 0x0a, 0xaf,
	0xa1, 0x18, 0xd2, 0x21, 0x1e, 0xc6, 0x27, 0x88, 0x0e, 0xe8, 0xca, 0xd1, 0x82, 0xde, 0x13, 0x9a,
	0x0b, 0x10, 0x48, 0x14, 0x1f, 0xc7, 0x87, 0x47, 0xe6, 0x7e, 0xe5, 0x70, 0x31, 0xe7, 0x30, 0x2a,
	0x50, 0x6f, 0x12, 0x2a, 0x72, 0x7b, 0x24, 0xa1, 0xa2, 0x0d, 0x21, 0x51, 0x81, 0xbc, 0x93, 0x50,
	0x91, 0x2b, 0x24, 0x09, 0x35, 0xa7, 0x63, 0x04, 0x2a, 0xe8, 0x99, 0x24, 0x54, 0xe4, 0x1e, 0x49,
	0x42, 0xcd, 0x69, 0xc3, 0x14, 0xde, 0xc2, 0xc6, 0x4c, 0x23, 0xe2, 0xd3, 0xbb, 0x0e, 0x66, 0x76,
	0xf0, 0x56, 0x9e, 0x2d, 0x11, 0x11, 0x26, 0xcf, 0x34, 0x69, 0x12, 0x79, 0xfe, 0x65, 0x93, 0x44,
	0x8e, 0x9b, 0x00, 0x82, 0x3c, 0x33, 0x03, 0x92, 0xc8, 0xf3, 0x2f, 0x9b, 0x24, 0x72, 0xcc, 0x80,
	0x31, 0x52, 0x17, 0x39, 0xf1, 0xbf, 0xfb, 0xf9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0x3d,
	0x55, 0x2f, 0x96, 0x0b, 0x00, 0x00,
}