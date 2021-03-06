// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/service/proto/auth.proto

package go_micro_auth

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

type Token struct {
	Token                string            `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Type                 string            `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Created              int64             `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Expiry               int64             `protobuf:"varint,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Subject              string            `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	Roles                []string          `protobuf:"bytes,6,rep,name=roles,proto3" json:"roles,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{0}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Token) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Token) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *Token) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Token) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Token) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Account struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Secret               *Token            `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	Roles                []string          `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{1}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetSecret() *Token {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *Account) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Account) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Resource struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Endpoint             string   `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{2}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Resource) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

type GenerateRequest struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Roles                []string          `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SecretExpiry         int64             `protobuf:"varint,4,opt,name=secret_expiry,json=secretExpiry,proto3" json:"secret_expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GenerateRequest) Reset()         { *m = GenerateRequest{} }
func (m *GenerateRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateRequest) ProtoMessage()    {}
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{3}
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

func (m *GenerateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GenerateRequest) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *GenerateRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *GenerateRequest) GetSecretExpiry() int64 {
	if m != nil {
		return m.SecretExpiry
	}
	return 0
}

type GenerateResponse struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateResponse) Reset()         { *m = GenerateResponse{} }
func (m *GenerateResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateResponse) ProtoMessage()    {}
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{4}
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

func (m *GenerateResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type GrantRequest struct {
	Role                 string    `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Resource             *Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GrantRequest) Reset()         { *m = GrantRequest{} }
func (m *GrantRequest) String() string { return proto.CompactTextString(m) }
func (*GrantRequest) ProtoMessage()    {}
func (*GrantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{5}
}

func (m *GrantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrantRequest.Unmarshal(m, b)
}
func (m *GrantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrantRequest.Marshal(b, m, deterministic)
}
func (m *GrantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrantRequest.Merge(m, src)
}
func (m *GrantRequest) XXX_Size() int {
	return xxx_messageInfo_GrantRequest.Size(m)
}
func (m *GrantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GrantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GrantRequest proto.InternalMessageInfo

func (m *GrantRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *GrantRequest) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type GrantResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrantResponse) Reset()         { *m = GrantResponse{} }
func (m *GrantResponse) String() string { return proto.CompactTextString(m) }
func (*GrantResponse) ProtoMessage()    {}
func (*GrantResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{6}
}

func (m *GrantResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrantResponse.Unmarshal(m, b)
}
func (m *GrantResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrantResponse.Marshal(b, m, deterministic)
}
func (m *GrantResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrantResponse.Merge(m, src)
}
func (m *GrantResponse) XXX_Size() int {
	return xxx_messageInfo_GrantResponse.Size(m)
}
func (m *GrantResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GrantResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GrantResponse proto.InternalMessageInfo

type VerifyRequest struct {
	Account              *Account  `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Resource             *Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *VerifyRequest) Reset()         { *m = VerifyRequest{} }
func (m *VerifyRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyRequest) ProtoMessage()    {}
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{7}
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

func (m *VerifyRequest) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *VerifyRequest) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type VerifyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyResponse) Reset()         { *m = VerifyResponse{} }
func (m *VerifyResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyResponse) ProtoMessage()    {}
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{8}
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

type RevokeRequest struct {
	Role                 string    `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Resource             *Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RevokeRequest) Reset()         { *m = RevokeRequest{} }
func (m *RevokeRequest) String() string { return proto.CompactTextString(m) }
func (*RevokeRequest) ProtoMessage()    {}
func (*RevokeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{9}
}

func (m *RevokeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeRequest.Unmarshal(m, b)
}
func (m *RevokeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeRequest.Marshal(b, m, deterministic)
}
func (m *RevokeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeRequest.Merge(m, src)
}
func (m *RevokeRequest) XXX_Size() int {
	return xxx_messageInfo_RevokeRequest.Size(m)
}
func (m *RevokeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeRequest proto.InternalMessageInfo

func (m *RevokeRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *RevokeRequest) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type RevokeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeResponse) Reset()         { *m = RevokeResponse{} }
func (m *RevokeResponse) String() string { return proto.CompactTextString(m) }
func (*RevokeResponse) ProtoMessage()    {}
func (*RevokeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{10}
}

func (m *RevokeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeResponse.Unmarshal(m, b)
}
func (m *RevokeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeResponse.Marshal(b, m, deterministic)
}
func (m *RevokeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeResponse.Merge(m, src)
}
func (m *RevokeResponse) XXX_Size() int {
	return xxx_messageInfo_RevokeResponse.Size(m)
}
func (m *RevokeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeResponse proto.InternalMessageInfo

type InspectRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectRequest) Reset()         { *m = InspectRequest{} }
func (m *InspectRequest) String() string { return proto.CompactTextString(m) }
func (*InspectRequest) ProtoMessage()    {}
func (*InspectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{11}
}

func (m *InspectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectRequest.Unmarshal(m, b)
}
func (m *InspectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectRequest.Marshal(b, m, deterministic)
}
func (m *InspectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectRequest.Merge(m, src)
}
func (m *InspectRequest) XXX_Size() int {
	return xxx_messageInfo_InspectRequest.Size(m)
}
func (m *InspectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InspectRequest proto.InternalMessageInfo

func (m *InspectRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type InspectResponse struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InspectResponse) Reset()         { *m = InspectResponse{} }
func (m *InspectResponse) String() string { return proto.CompactTextString(m) }
func (*InspectResponse) ProtoMessage()    {}
func (*InspectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{12}
}

func (m *InspectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InspectResponse.Unmarshal(m, b)
}
func (m *InspectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InspectResponse.Marshal(b, m, deterministic)
}
func (m *InspectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InspectResponse.Merge(m, src)
}
func (m *InspectResponse) XXX_Size() int {
	return xxx_messageInfo_InspectResponse.Size(m)
}
func (m *InspectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InspectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InspectResponse proto.InternalMessageInfo

func (m *InspectResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type RefreshRequest struct {
	Secret               string   `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	TokenExpiry          int64    `protobuf:"varint,2,opt,name=token_expiry,json=tokenExpiry,proto3" json:"token_expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshRequest) Reset()         { *m = RefreshRequest{} }
func (m *RefreshRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshRequest) ProtoMessage()    {}
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{13}
}

func (m *RefreshRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshRequest.Unmarshal(m, b)
}
func (m *RefreshRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshRequest.Marshal(b, m, deterministic)
}
func (m *RefreshRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshRequest.Merge(m, src)
}
func (m *RefreshRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshRequest.Size(m)
}
func (m *RefreshRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshRequest proto.InternalMessageInfo

func (m *RefreshRequest) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *RefreshRequest) GetTokenExpiry() int64 {
	if m != nil {
		return m.TokenExpiry
	}
	return 0
}

type RefreshResponse struct {
	Token                *Token   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshResponse) Reset()         { *m = RefreshResponse{} }
func (m *RefreshResponse) String() string { return proto.CompactTextString(m) }
func (*RefreshResponse) ProtoMessage()    {}
func (*RefreshResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21300bfacc51fc2a, []int{14}
}

func (m *RefreshResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshResponse.Unmarshal(m, b)
}
func (m *RefreshResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshResponse.Marshal(b, m, deterministic)
}
func (m *RefreshResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshResponse.Merge(m, src)
}
func (m *RefreshResponse) XXX_Size() int {
	return xxx_messageInfo_RefreshResponse.Size(m)
}
func (m *RefreshResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshResponse proto.InternalMessageInfo

func (m *RefreshResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*Token)(nil), "go.micro.auth.Token")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.auth.Token.MetadataEntry")
	proto.RegisterType((*Account)(nil), "go.micro.auth.Account")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.auth.Account.MetadataEntry")
	proto.RegisterType((*Resource)(nil), "go.micro.auth.Resource")
	proto.RegisterType((*GenerateRequest)(nil), "go.micro.auth.GenerateRequest")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.auth.GenerateRequest.MetadataEntry")
	proto.RegisterType((*GenerateResponse)(nil), "go.micro.auth.GenerateResponse")
	proto.RegisterType((*GrantRequest)(nil), "go.micro.auth.GrantRequest")
	proto.RegisterType((*GrantResponse)(nil), "go.micro.auth.GrantResponse")
	proto.RegisterType((*VerifyRequest)(nil), "go.micro.auth.VerifyRequest")
	proto.RegisterType((*VerifyResponse)(nil), "go.micro.auth.VerifyResponse")
	proto.RegisterType((*RevokeRequest)(nil), "go.micro.auth.RevokeRequest")
	proto.RegisterType((*RevokeResponse)(nil), "go.micro.auth.RevokeResponse")
	proto.RegisterType((*InspectRequest)(nil), "go.micro.auth.InspectRequest")
	proto.RegisterType((*InspectResponse)(nil), "go.micro.auth.InspectResponse")
	proto.RegisterType((*RefreshRequest)(nil), "go.micro.auth.RefreshRequest")
	proto.RegisterType((*RefreshResponse)(nil), "go.micro.auth.RefreshResponse")
}

func init() { proto.RegisterFile("auth/service/proto/auth.proto", fileDescriptor_21300bfacc51fc2a) }

var fileDescriptor_21300bfacc51fc2a = []byte{
	// 663 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xad, 0xed, 0xfc, 0x75, 0x52, 0x27, 0xd1, 0xaa, 0xea, 0x67, 0xf9, 0xa3, 0x25, 0x18, 0x84,
	0x2a, 0x54, 0xb9, 0x28, 0xbd, 0x41, 0x20, 0x10, 0x15, 0xad, 0xca, 0x8f, 0xca, 0x85, 0x85, 0x80,
	0x3b, 0xe4, 0x3a, 0x53, 0x6a, 0xd2, 0xd8, 0x66, 0xbd, 0x8e, 0xc8, 0x5b, 0xf0, 0x52, 0xbc, 0x05,
	0x77, 0xbc, 0x08, 0xda, 0xf5, 0xae, 0x6b, 0x3b, 0x09, 0x12, 0x05, 0xee, 0x76, 0x76, 0x66, 0xcf,
	0x99, 0x39, 0x73, 0xe2, 0xc0, 0xb6, 0x9f, 0xb1, 0x8b, 0xfd, 0x14, 0xe9, 0x2c, 0x0c, 0x70, 0x3f,
	0xa1, 0x31, 0x8b, 0xf7, 0xf9, 0x95, 0x2b, 0x8e, 0xc4, 0xfc, 0x18, 0xbb, 0xd3, 0x30, 0xa0, 0xb1,
	0xcb, 0x2f, 0x9d, 0xaf, 0x3a, 0x34, 0xdf, 0xc4, 0x13, 0x8c, 0xc8, 0x26, 0x34, 0x19, 0x3f, 0x58,
	0xda, 0x50, 0xdb, 0x5d, 0xf7, 0xf2, 0x80, 0x10, 0x68, 0xb0, 0x79, 0x82, 0x96, 0x2e, 0x2e, 0xc5,
	0x99, 0x58, 0xd0, 0x0e, 0x28, 0xfa, 0x0c, 0xc7, 0x96, 0x31, 0xd4, 0x76, 0x0d, 0x4f, 0x85, 0x64,
	0x0b, 0x5a, 0xf8, 0x25, 0x09, 0xe9, 0xdc, 0x6a, 0x88, 0x84, 0x8c, 0xf8, 0x8b, 0x34, 0x3b, 0xfb,
	0x84, 0x01, 0xb3, 0x9a, 0x02, 0x48, 0x85, 0x9c, 0x95, 0xc6, 0x97, 0x98, 0x5a, 0xad, 0xa1, 0xc1,
	0x59, 0x45, 0x40, 0x9e, 0x40, 0x67, 0x8a, 0xcc, 0x1f, 0xfb, 0xcc, 0xb7, 0xda, 0x43, 0x63, 0xb7,
	0x3b, 0x72, 0xdc, 0x4a, 0xdf, 0xae, 0xe8, 0xd9, 0x3d, 0x95, 0x45, 0xc7, 0x11, 0xa3, 0x73, 0xaf,
	0x78, 0x63, 0x3f, 0x02, 0xb3, 0x92, 0x22, 0x03, 0x30, 0x26, 0x38, 0x97, 0xa3, 0xf1, 0x23, 0x27,
	0x9e, 0xf9, 0x97, 0x99, 0x9a, 0x2c, 0x0f, 0x1e, 0xea, 0x0f, 0x34, 0xe7, 0xbb, 0x06, 0xed, 0xc3,
	0x20, 0x88, 0xb3, 0x88, 0x91, 0x1e, 0xe8, 0xe1, 0x58, 0x3e, 0xd3, 0xc3, 0x31, 0xd9, 0x83, 0x56,
	0x8a, 0x01, 0x45, 0x26, 0x9e, 0x75, 0x47, 0x9b, 0xcb, 0xda, 0xf2, 0x64, 0xcd, 0xd5, 0x70, 0x46,
	0x79, 0xb8, 0xa7, 0xa5, 0xe1, 0x1a, 0x62, 0xb8, 0x3b, 0x35, 0x14, 0xc9, 0xfe, 0x6f, 0xc6, 0x7b,
	0x0d, 0x1d, 0x0f, 0xd3, 0x38, 0xa3, 0x01, 0xf2, 0xed, 0x46, 0xfe, 0x14, 0xe5, 0x43, 0x71, 0x5e,
	0xba, 0x71, 0x1b, 0x3a, 0x18, 0x8d, 0x93, 0x38, 0x8c, 0x98, 0x58, 0xf9, 0xba, 0x57, 0xc4, 0xce,
	0x0f, 0x0d, 0xfa, 0x27, 0x18, 0x21, 0xf5, 0x19, 0x7a, 0xf8, 0x39, 0xc3, 0x74, 0x51, 0xb6, 0x42,
	0x08, 0xbd, 0x2c, 0xc4, 0xf3, 0x92, 0x10, 0x86, 0x10, 0x62, 0xaf, 0x26, 0x44, 0x0d, 0x77, 0x95,
	0x20, 0xe4, 0x36, 0x98, 0xb9, 0xe4, 0x1f, 0x2a, 0xf6, 0xdb, 0xc8, 0x2f, 0x8f, 0xc5, 0xdd, 0x9f,
	0xa9, 0x76, 0x04, 0x83, 0xab, 0x66, 0xd2, 0x24, 0x8e, 0x52, 0x24, 0xf7, 0xa1, 0xed, 0xe7, 0x9b,
	0x12, 0x18, 0xdd, 0xd1, 0xd6, 0xf2, 0x3d, 0x7a, 0xaa, 0xcc, 0x79, 0x07, 0x1b, 0x27, 0xd4, 0x8f,
	0x98, 0xd2, 0x89, 0x40, 0x83, 0x4b, 0xa1, 0xf4, 0xe7, 0x67, 0x72, 0x00, 0x1d, 0x2a, 0xf7, 0x23,
	0x4d, 0xf6, 0x5f, 0x0d, 0x56, 0xad, 0xcf, 0x2b, 0x0a, 0x9d, 0x3e, 0x98, 0x12, 0x38, 0xef, 0xcd,
	0x99, 0x81, 0xf9, 0x16, 0x69, 0x78, 0x3e, 0x57, 0x54, 0xbf, 0xdd, 0xec, 0xf5, 0x1a, 0x19, 0x40,
	0x4f, 0xf1, 0xca, 0x4e, 0xde, 0x83, 0xe9, 0xe1, 0x2c, 0x9e, 0xe0, 0x5f, 0x1f, 0x7a, 0x00, 0x3d,
	0x85, 0x2c, 0xb9, 0xee, 0x42, 0xef, 0x45, 0x94, 0x26, 0x18, 0x14, 0x0a, 0x2f, 0xfd, 0xaa, 0x39,
	0xcf, 0xa0, 0x5f, 0xd4, 0x5d, 0x7b, 0x99, 0xaf, 0x38, 0xfd, 0x39, 0xc5, 0xf4, 0x42, 0x91, 0x6d,
	0x15, 0x5f, 0x87, 0x9c, 0x4d, 0x7d, 0x07, 0x6e, 0xc1, 0x86, 0xe0, 0x55, 0xee, 0xd4, 0x85, 0x3b,
	0xbb, 0xe2, 0x2e, 0x37, 0xa7, 0xf3, 0x18, 0xfa, 0x05, 0x98, 0xec, 0xe8, 0x5e, 0xb9, 0xf5, 0x55,
	0x9f, 0x9a, 0xbc, 0x64, 0xf4, 0xcd, 0x80, 0xc6, 0x61, 0xc6, 0x2e, 0xc8, 0x29, 0x74, 0x94, 0x4f,
	0xc9, 0xce, 0xaf, 0x7f, 0x4d, 0xf6, 0xcd, 0x95, 0x79, 0x29, 0xe7, 0x1a, 0x39, 0x82, 0xa6, 0xf0,
	0x15, 0xf9, 0xbf, 0x5e, 0x5b, 0xb2, 0xb1, 0x7d, 0x63, 0x79, 0xb2, 0x40, 0x39, 0x81, 0x56, 0x6e,
	0x0a, 0x52, 0xaf, 0xac, 0x78, 0xd4, 0xde, 0x5e, 0x91, 0x2d, 0x03, 0xe5, 0x1b, 0x5f, 0x00, 0xaa,
	0x58, 0x6c, 0x01, 0xa8, 0x66, 0x93, 0x35, 0xf2, 0x12, 0xda, 0xd2, 0x00, 0xa4, 0x5e, 0x5b, 0x35,
	0x90, 0xbd, 0xb3, 0x2a, 0x5d, 0xc6, 0x92, 0xab, 0x23, 0x8b, 0xbc, 0x65, 0x7f, 0x2c, 0x60, 0xd5,
	0x36, 0xee, 0xac, 0x9d, 0xb5, 0xc4, 0x9f, 0xf4, 0xc1, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb7,
	0xf8, 0x55, 0xb6, 0xc5, 0x07, 0x00, 0x00,
}
