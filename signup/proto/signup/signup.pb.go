// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: github.com/m3o/services/signup/proto/signup/signup.proto

package go_micro_service_signup

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SendVerificationEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SendVerificationEmailRequest) Reset() {
	*x = SendVerificationEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVerificationEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVerificationEmailRequest) ProtoMessage() {}

func (x *SendVerificationEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVerificationEmailRequest.ProtoReflect.Descriptor instead.
func (*SendVerificationEmailRequest) Descriptor() ([]byte, []int) {
	return file_github_com_micro_services_signup_proto_signup_signup_proto_rawDescGZIP(), []int{0}
}

func (x *SendVerificationEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SendVerificationEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendVerificationEmailResponse) Reset() {
	*x = SendVerificationEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVerificationEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVerificationEmailResponse) ProtoMessage() {}

func (x *SendVerificationEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVerificationEmailResponse.ProtoReflect.Descriptor instead.
func (*SendVerificationEmailResponse) Descriptor() ([]byte, []int) {
	return file_github_com_micro_services_signup_proto_signup_signup_proto_rawDescGZIP(), []int{1}
}

type VerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// Email token that was received in an email.
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VerifyRequest) Reset() {
	*x = VerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRequest) ProtoMessage() {}

func (x *VerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRequest.ProtoReflect.Descriptor instead.
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return file_github_com_micro_services_signup_proto_signup_signup_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *VerifyRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Auth token to be saved into `~/.micro`
	// For users who are already registered and paid,
	// the flow stops here.
	// For users who are yet to be registered
	// the token will be acquired in the `FinishSignup` step.
	AuthToken *AuthToken `protobuf:"bytes,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	// Payment provider custommer id that can be used to
	// acquire a payment method, see `micro login` flow for more.
	// @todo this is likely not needed
	CustomerID string `protobuf:"bytes,2,opt,name=customerID,proto3" json:"customerID,omitempty"`
	// Namesp
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *VerifyResponse) Reset() {
	*x = VerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyResponse) ProtoMessage() {}

func (x *VerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyResponse.ProtoReflect.Descriptor instead.
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return file_github_com_micro_services_signup_proto_signup_signup_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyResponse) GetAuthToken() *AuthToken {
	if x != nil {
		return x.AuthToken
	}
	return nil
}

func (x *VerifyResponse) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *VerifyResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type CompleteSignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// The token has to be passed here too for identification purposes.
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// This payment method ID is the one we got back from Stripe on the frontend (ie. `m3o.com/subscribe.html`)
	PaymentMethodID string `protobuf:"bytes,3,opt,name=paymentMethodID,proto3" json:"paymentMethodID,omitempty"`
	// The secret/password to use for the account
	Secret string `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *CompleteSignupRequest) Reset() {
	*x = CompleteSignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteSignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteSignupRequest) ProtoMessage() {}

func (x *CompleteSignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteSignupRequest.ProtoReflect.Descriptor instead.
func (*CompleteSignupRequest) Descriptor() ([]byte, []int) {
	return file_github_com_micro_services_signup_proto_signup_signup_proto_rawDescGZIP(), []int{4}
}

func (x *CompleteSignupRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CompleteSignupRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CompleteSignupRequest) GetPaymentMethodID() string {
	if x != nil {
		return x.PaymentMethodID
	}
	return ""
}

func (x *CompleteSignupRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type CompleteSignupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken *AuthToken `protobuf:"bytes,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	Namespace string     `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *CompleteSignupResponse) Reset() {
	*x = CompleteSignupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteSignupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteSignupResponse) ProtoMessage() {}

func (x *CompleteSignupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteSignupResponse.ProtoReflect.Descriptor instead.
func (*CompleteSignupResponse) Descriptor() ([]byte, []int) {
	return file_github_com_micro_services_signup_proto_signup_signup_proto_rawDescGZIP(), []int{5}
}

func (x *CompleteSignupResponse) GetAuthToken() *AuthToken {
	if x != nil {
		return x.AuthToken
	}
	return nil
}

func (x *CompleteSignupResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// lifted from https://github.com/micro/go-micro/blob/master/auth/service/proto/auth.proto
type AuthToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Created      int64  `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Expiry       int64  `protobuf:"varint,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *AuthToken) Reset() {
	*x = AuthToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthToken) ProtoMessage() {}

func (x *AuthToken) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthToken.ProtoReflect.Descriptor instead.
func (*AuthToken) Descriptor() ([]byte, []int) {
	return file_github_com_micro_services_signup_proto_signup_signup_proto_rawDescGZIP(), []int{6}
}

func (x *AuthToken) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AuthToken) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *AuthToken) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *AuthToken) GetExpiry() int64 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

var File_github_com_micro_services_signup_proto_signup_signup_proto protoreflect.FileDescriptor

var file_github_com_micro_services_signup_proto_signup_signup_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x2f,
	0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73,
	0x69, 0x67, 0x6e, 0x75, 0x70, 0x22, 0x34, 0x0a, 0x1c, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x1f, 0x0a, 0x1d, 0x53,
	0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x0a, 0x0d,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x85, 0x01, 0x0a,
	0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x22, 0x78, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x85,
	0x01, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x32, 0xdf, 0x02, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x12, 0x86, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x35, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73,
	0x69, 0x67, 0x6e, 0x75, 0x70, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x36, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x06, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x12, 0x26, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x2e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x2e, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_micro_services_signup_proto_signup_signup_proto_rawDescOnce sync.Once
	file_github_com_micro_services_signup_proto_signup_signup_proto_rawDescData = file_github_com_micro_services_signup_proto_signup_signup_proto_rawDesc
)

func file_github_com_micro_services_signup_proto_signup_signup_proto_rawDescGZIP() []byte {
	file_github_com_micro_services_signup_proto_signup_signup_proto_rawDescOnce.Do(func() {
		file_github_com_micro_services_signup_proto_signup_signup_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_micro_services_signup_proto_signup_signup_proto_rawDescData)
	})
	return file_github_com_micro_services_signup_proto_signup_signup_proto_rawDescData
}

var file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_micro_services_signup_proto_signup_signup_proto_goTypes = []interface{}{
	(*SendVerificationEmailRequest)(nil),  // 0: go.micro.service.signup.SendVerificationEmailRequest
	(*SendVerificationEmailResponse)(nil), // 1: go.micro.service.signup.SendVerificationEmailResponse
	(*VerifyRequest)(nil),                 // 2: go.micro.service.signup.VerifyRequest
	(*VerifyResponse)(nil),                // 3: go.micro.service.signup.VerifyResponse
	(*CompleteSignupRequest)(nil),         // 4: go.micro.service.signup.CompleteSignupRequest
	(*CompleteSignupResponse)(nil),        // 5: go.micro.service.signup.CompleteSignupResponse
	(*AuthToken)(nil),                     // 6: go.micro.service.signup.AuthToken
}
var file_github_com_micro_services_signup_proto_signup_signup_proto_depIdxs = []int32{
	6, // 0: go.micro.service.signup.VerifyResponse.authToken:type_name -> go.micro.service.signup.AuthToken
	6, // 1: go.micro.service.signup.CompleteSignupResponse.authToken:type_name -> go.micro.service.signup.AuthToken
	0, // 2: go.micro.service.signup.Signup.SendVerificationEmail:input_type -> go.micro.service.signup.SendVerificationEmailRequest
	2, // 3: go.micro.service.signup.Signup.Verify:input_type -> go.micro.service.signup.VerifyRequest
	4, // 4: go.micro.service.signup.Signup.CompleteSignup:input_type -> go.micro.service.signup.CompleteSignupRequest
	1, // 5: go.micro.service.signup.Signup.SendVerificationEmail:output_type -> go.micro.service.signup.SendVerificationEmailResponse
	3, // 6: go.micro.service.signup.Signup.Verify:output_type -> go.micro.service.signup.VerifyResponse
	5, // 7: go.micro.service.signup.Signup.CompleteSignup:output_type -> go.micro.service.signup.CompleteSignupResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_micro_services_signup_proto_signup_signup_proto_init() }
func file_github_com_micro_services_signup_proto_signup_signup_proto_init() {
	if File_github_com_micro_services_signup_proto_signup_signup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVerificationEmailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVerificationEmailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteSignupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteSignupResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_micro_services_signup_proto_signup_signup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_micro_services_signup_proto_signup_signup_proto_goTypes,
		DependencyIndexes: file_github_com_micro_services_signup_proto_signup_signup_proto_depIdxs,
		MessageInfos:      file_github_com_micro_services_signup_proto_signup_signup_proto_msgTypes,
	}.Build()
	File_github_com_micro_services_signup_proto_signup_signup_proto = out.File
	file_github_com_micro_services_signup_proto_signup_signup_proto_rawDesc = nil
	file_github_com_micro_services_signup_proto_signup_signup_proto_goTypes = nil
	file_github_com_micro_services_signup_proto_signup_signup_proto_depIdxs = nil
}
