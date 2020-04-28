// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/services/event/service/proto/event.proto

package go_micro_service_events

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

type EventType int32

const (
	EventType_Unknown       EventType = 0
	EventType_BuildStarted  EventType = 1
	EventType_BuildFinished EventType = 2
	EventType_BuildFailed   EventType = 3
)

var EventType_name = map[int32]string{
	0: "Unknown",
	1: "BuildStarted",
	2: "BuildFinished",
	3: "BuildFailed",
}

var EventType_value = map[string]int32{
	"Unknown":       0,
	"BuildStarted":  1,
	"BuildFinished": 2,
	"BuildFailed":   3,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_86878713dbc31a8b, []int{0}
}

type Event struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId            string            `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Type                 EventType         `protobuf:"varint,3,opt,name=type,proto3,enum=go.micro.service.events.EventType" json:"type,omitempty"`
	Created              int64             `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_86878713dbc31a8b, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_Unknown
}

func (m *Event) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Event) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type CreateRequest struct {
	ProjectId            string            `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Type                 EventType         `protobuf:"varint,2,opt,name=type,proto3,enum=go.micro.service.events.EventType" json:"type,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86878713dbc31a8b, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *CreateRequest) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_Unknown
}

func (m *CreateRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type CreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86878713dbc31a8b, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

type ReadRequest struct {
	EventId              string   `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	ProjectId            string   `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86878713dbc31a8b, []int{3}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *ReadRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

type ReadResponse struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_86878713dbc31a8b, []int{4}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterEnum("go.micro.service.events.EventType", EventType_name, EventType_value)
	proto.RegisterType((*Event)(nil), "go.micro.service.events.Event")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.service.events.Event.MetadataEntry")
	proto.RegisterType((*CreateRequest)(nil), "go.micro.service.events.CreateRequest")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.service.events.CreateRequest.MetadataEntry")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.service.events.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.service.events.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.service.events.ReadResponse")
}

func init() {
	proto.RegisterFile("github.com/micro/services/event/service/proto/event.proto", fileDescriptor_86878713dbc31a8b)
}

var fileDescriptor_86878713dbc31a8b = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xed, 0xda, 0xf9, 0x68, 0x26, 0x4d, 0x30, 0x23, 0x24, 0x4c, 0x24, 0x50, 0x64, 0xf1, 0x11,
	0x21, 0xb4, 0x91, 0x02, 0xaa, 0xf8, 0xb8, 0x81, 0x5a, 0xe8, 0x01, 0x84, 0x0c, 0x3d, 0x71, 0x40,
	0x5b, 0xef, 0xa8, 0x5d, 0x9a, 0x78, 0x8d, 0x77, 0x13, 0xe4, 0x5f, 0xc1, 0x1f, 0xe2, 0x6f, 0x71,
	0x47, 0x59, 0xdb, 0xa5, 0x45, 0x72, 0x02, 0x07, 0x6e, 0x33, 0x6f, 0xdf, 0xbc, 0x9d, 0xf7, 0xec,
	0x85, 0x67, 0xa7, 0xca, 0x9e, 0x2d, 0x4f, 0x78, 0xa2, 0x17, 0xd3, 0x85, 0x4a, 0x72, 0x3d, 0x35,
	0x94, 0xaf, 0x54, 0x42, 0x66, 0x4a, 0x2b, 0x4a, 0x6d, 0xdd, 0x4e, 0xb3, 0x5c, 0x5b, 0x5d, 0x62,
	0xdc, 0xd5, 0x78, 0xf3, 0x54, 0x73, 0x37, 0xc2, 0x2b, 0x0e, 0x77, 0xa7, 0x26, 0xfa, 0xee, 0x41,
	0xfb, 0x60, 0x5d, 0xe2, 0x10, 0x3c, 0x25, 0x43, 0x36, 0x66, 0x93, 0x5e, 0xec, 0x29, 0x89, 0xb7,
	0x01, 0xb2, 0x5c, 0x7f, 0xa1, 0xc4, 0x7e, 0x56, 0x32, 0xf4, 0x1c, 0xde, 0xab, 0x90, 0x23, 0x89,
	0xfb, 0xd0, 0xb2, 0x45, 0x46, 0xa1, 0x3f, 0x66, 0x93, 0xe1, 0x2c, 0xe2, 0x0d, 0x17, 0x70, 0x27,
	0xfe, 0xb1, 0xc8, 0x28, 0x76, 0x7c, 0x0c, 0xa1, 0x9b, 0xe4, 0x24, 0x2c, 0xc9, 0xb0, 0x35, 0x66,
	0x13, 0x3f, 0xae, 0x5b, 0x7c, 0x03, 0xbb, 0x0b, 0xb2, 0x42, 0x0a, 0x2b, 0xc2, 0xf6, 0xd8, 0x9f,
	0xf4, 0x67, 0x8f, 0x36, 0xab, 0xf2, 0xb7, 0x15, 0xfd, 0x20, 0xb5, 0x79, 0x11, 0x5f, 0x4c, 0x8f,
	0x5e, 0xc0, 0xe0, 0xca, 0x11, 0x06, 0xe0, 0x9f, 0x53, 0x51, 0x99, 0x5b, 0x97, 0x78, 0x03, 0xda,
	0x2b, 0x31, 0x5f, 0x52, 0x65, 0xac, 0x6c, 0x9e, 0x7b, 0x4f, 0x59, 0xf4, 0x93, 0xc1, 0xe0, 0x95,
	0x5b, 0x29, 0xa6, 0xaf, 0x4b, 0x32, 0xf6, 0x8f, 0x24, 0x58, 0x53, 0x12, 0xde, 0x3f, 0x26, 0xf1,
	0xfe, 0x92, 0x5f, 0xdf, 0xf9, 0x7d, 0xd2, 0x38, 0x7b, 0x65, 0xa1, 0xff, 0xe3, 0x3b, 0x80, 0x61,
	0x7d, 0x8b, 0xc9, 0x74, 0x6a, 0x28, 0x7a, 0x0d, 0xfd, 0x98, 0x84, 0xac, 0x63, 0xb8, 0x05, 0xbb,
	0x6e, 0x9b, 0xdf, 0x21, 0x74, 0x5d, 0x7f, 0xb4, 0xed, 0x5f, 0x89, 0x0e, 0x61, 0xaf, 0x14, 0x2a,
	0x85, 0x71, 0x1f, 0x3a, 0xa5, 0xaf, 0x90, 0x39, 0xdf, 0x77, 0x36, 0x67, 0x16, 0x57, 0xec, 0x87,
	0xef, 0xa0, 0x77, 0x11, 0x22, 0xf6, 0xa1, 0x7b, 0x9c, 0x9e, 0xa7, 0xfa, 0x5b, 0x1a, 0xec, 0x60,
	0x00, 0x7b, 0x2f, 0x97, 0x6a, 0x2e, 0x3f, 0x58, 0x91, 0x5b, 0x92, 0x01, 0xc3, 0xeb, 0x30, 0x70,
	0xc8, 0xa1, 0x4a, 0x95, 0x39, 0x23, 0x19, 0x78, 0x78, 0x0d, 0xfa, 0x25, 0x24, 0xd4, 0x9c, 0x64,
	0xe0, 0xcf, 0x7e, 0x30, 0xe8, 0x38, 0x41, 0x83, 0x9f, 0xa0, 0x53, 0xba, 0xc7, 0xfb, 0x7f, 0xf7,
	0x11, 0x46, 0x0f, 0xb6, 0xf2, 0xaa, 0x18, 0x77, 0xf0, 0x18, 0x5a, 0x6b, 0xff, 0x78, 0xb7, 0x71,
	0xe4, 0x52, 0xce, 0xa3, 0x7b, 0x5b, 0x58, 0xb5, 0xec, 0x49, 0xc7, 0xbd, 0xed, 0xc7, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xc8, 0x23, 0xe5, 0xf4, 0x18, 0x04, 0x00, 0x00,
}