// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/m3o/services/projects/invite/proto/invite.proto

package go_micro_service_projects_invite

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

type GenerateRequest struct {
	ProjectId            string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateRequest) Reset()         { *m = GenerateRequest{} }
func (m *GenerateRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateRequest) ProtoMessage()    {}
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f8db6bfc98891f, []int{0}
}

func (m *GenerateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateRequest.Unmarshal(m, b)
}
func (m *GenerateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateRequest.Marshal(b, m, deterministic)
}
func (m *GenerateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateRequest.Merge(m, src)
}
func (m *GenerateRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateRequest.Size(m)
}
func (m *GenerateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateRequest proto.InternalMessageInfo

func (m *GenerateRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *GenerateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *GenerateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GenerateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateResponse) Reset()         { *m = GenerateResponse{} }
func (m *GenerateResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateResponse) ProtoMessage()    {}
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f8db6bfc98891f, []int{1}
}

func (m *GenerateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateResponse.Unmarshal(m, b)
}
func (m *GenerateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateResponse.Marshal(b, m, deterministic)
}
func (m *GenerateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateResponse.Merge(m, src)
}
func (m *GenerateResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateResponse.Size(m)
}
func (m *GenerateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateResponse proto.InternalMessageInfo

type VerifyRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f8db6bfc98891f, []int{2}
}

func (m *VerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyRequest.Unmarshal(m, b)
}
func (m *VerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyRequest.Marshal(b, m, deterministic)
}
func (m *VerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRequest.Merge(m, src)
}
func (m *VerifyRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyRequest.Size(m)
}
func (m *VerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRequest proto.InternalMessageInfo

func (m *VerifyRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type VerifyResponse struct {
	ProjectName          string   `protobuf:"bytes,1,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyResponse) Reset()         { *m = VerifyResponse{} }
func (m *VerifyResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyResponse) ProtoMessage()    {}
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f8db6bfc98891f, []int{3}
}

func (m *VerifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyResponse.Unmarshal(m, b)
}
func (m *VerifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyResponse.Marshal(b, m, deterministic)
}
func (m *VerifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyResponse.Merge(m, src)
}
func (m *VerifyResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyResponse.Size(m)
}
func (m *VerifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyResponse proto.InternalMessageInfo

func (m *VerifyResponse) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

func (m *VerifyResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type RedeemRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedeemRequest) Reset()         { *m = RedeemRequest{} }
func (m *RedeemRequest) String() string { return proto.CompactTextString(m) }
func (*RedeemRequest) ProtoMessage()    {}
func (*RedeemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f8db6bfc98891f, []int{4}
}

func (m *RedeemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedeemRequest.Unmarshal(m, b)
}
func (m *RedeemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedeemRequest.Marshal(b, m, deterministic)
}
func (m *RedeemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedeemRequest.Merge(m, src)
}
func (m *RedeemRequest) XXX_Size() int {
	return xxx_messageInfo_RedeemRequest.Size(m)
}
func (m *RedeemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RedeemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RedeemRequest proto.InternalMessageInfo

func (m *RedeemRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *RedeemRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type RedeemResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedeemResponse) Reset()         { *m = RedeemResponse{} }
func (m *RedeemResponse) String() string { return proto.CompactTextString(m) }
func (*RedeemResponse) ProtoMessage()    {}
func (*RedeemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_96f8db6bfc98891f, []int{5}
}

func (m *RedeemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedeemResponse.Unmarshal(m, b)
}
func (m *RedeemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedeemResponse.Marshal(b, m, deterministic)
}
func (m *RedeemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedeemResponse.Merge(m, src)
}
func (m *RedeemResponse) XXX_Size() int {
	return xxx_messageInfo_RedeemResponse.Size(m)
}
func (m *RedeemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RedeemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RedeemResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenerateRequest)(nil), "go.micro.service.projects.invite.GenerateRequest")
	proto.RegisterType((*GenerateResponse)(nil), "go.micro.service.projects.invite.GenerateResponse")
	proto.RegisterType((*VerifyRequest)(nil), "go.micro.service.projects.invite.VerifyRequest")
	proto.RegisterType((*VerifyResponse)(nil), "go.micro.service.projects.invite.VerifyResponse")
	proto.RegisterType((*RedeemRequest)(nil), "go.micro.service.projects.invite.RedeemRequest")
	proto.RegisterType((*RedeemResponse)(nil), "go.micro.service.projects.invite.RedeemResponse")
}

func init() {
	proto.RegisterFile("github.com/m3o/services/projects/invite/proto/invite.proto", fileDescriptor_96f8db6bfc98891f)
}

var fileDescriptor_96f8db6bfc98891f = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4f, 0x02, 0x31,
	0x10, 0xc5, 0x01, 0x11, 0x65, 0x14, 0x24, 0x13, 0x13, 0x09, 0x89, 0x09, 0xd6, 0x8b, 0xa7, 0xae,
	0xe2, 0x95, 0xbb, 0xd9, 0x8b, 0x07, 0x4c, 0x3c, 0x78, 0x31, 0xcb, 0xee, 0x88, 0x35, 0x76, 0x8b,
	0x6d, 0x21, 0xf1, 0xeb, 0xfa, 0x49, 0x0c, 0xfd, 0x23, 0x60, 0x4c, 0x58, 0x6f, 0x7d, 0xb3, 0xf3,
	0xf6, 0xbd, 0xfe, 0x52, 0x18, 0xcf, 0x84, 0x7d, 0x5d, 0x4c, 0x79, 0xae, 0x64, 0x22, 0x45, 0xae,
	0x55, 0x62, 0x48, 0x2f, 0x45, 0x4e, 0x26, 0x99, 0x6b, 0xf5, 0x46, 0xb9, 0x35, 0x89, 0x28, 0x97,
	0xc2, 0xd2, 0x4a, 0x5b, 0x15, 0x04, 0x77, 0x02, 0x87, 0x33, 0xc5, 0x9d, 0x8b, 0x07, 0x17, 0x8f,
	0x26, 0xee, 0xf7, 0xd8, 0x13, 0x9c, 0xdc, 0x51, 0x49, 0x3a, 0xb3, 0x34, 0xa1, 0x8f, 0x05, 0x19,
	0x8b, 0xe7, 0x00, 0x61, 0xeb, 0x59, 0x14, 0xfd, 0xfa, 0xb0, 0x7e, 0xd5, 0x9e, 0xb4, 0xc3, 0x24,
	0x2d, 0xf0, 0x14, 0xf6, 0x49, 0x66, 0xe2, 0xbd, 0xdf, 0x70, 0x5f, 0xbc, 0x40, 0x84, 0x66, 0x99,
	0x49, 0xea, 0xef, 0xb9, 0xa1, 0x3b, 0x33, 0x84, 0xde, 0xfa, 0xdf, 0x66, 0xae, 0x4a, 0x43, 0xec,
	0x12, 0x3a, 0x8f, 0xa4, 0xc5, 0xcb, 0x67, 0x4c, 0x43, 0x68, 0xe6, 0xaa, 0xa0, 0x90, 0xe3, 0xce,
	0x2c, 0x85, 0x6e, 0x5c, 0xf2, 0x36, 0xbc, 0x80, 0xe3, 0xd8, 0xc9, 0xc5, 0xf8, 0xed, 0xa3, 0x30,
	0xbb, 0xcf, 0x24, 0xfd, 0xdd, 0x8b, 0x8d, 0xa1, 0x33, 0xa1, 0x82, 0x48, 0xc6, 0xbc, 0x33, 0x38,
	0x58, 0x18, 0xd2, 0xeb, 0xab, 0xb5, 0x56, 0x32, 0x2d, 0x7e, 0x8a, 0x34, 0x36, 0x8a, 0xf4, 0xa0,
	0x1b, 0xdd, 0xbe, 0xc8, 0xe8, 0xab, 0x01, 0x9d, 0xd4, 0xa1, 0x7b, 0xf0, 0x44, 0xd1, 0xc0, 0x61,
	0xbc, 0x25, 0xde, 0xf0, 0x5d, 0xc0, 0xf9, 0x2f, 0xda, 0x83, 0xd1, 0x7f, 0x2c, 0x01, 0x62, 0x0d,
	0x25, 0xb4, 0x3c, 0x21, 0x4c, 0x76, 0xfb, 0xb7, 0x80, 0x0f, 0xae, 0xab, 0x1b, 0x36, 0xe3, 0x3c,
	0x87, 0x2a, 0x71, 0x5b, 0xbc, 0xab, 0xc4, 0x6d, 0x23, 0x66, 0xb5, 0x69, 0xcb, 0xbd, 0xde, 0xdb,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x6b, 0x45, 0x13, 0xfd, 0x02, 0x00, 0x00,
}
