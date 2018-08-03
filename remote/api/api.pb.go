// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AgentRegistration struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Hostname             string   `protobuf:"bytes,2,opt,name=Hostname,proto3" json:"Hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentRegistration) Reset()         { *m = AgentRegistration{} }
func (m *AgentRegistration) String() string { return proto.CompactTextString(m) }
func (*AgentRegistration) ProtoMessage()    {}
func (*AgentRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_117924c65de44d70, []int{0}
}
func (m *AgentRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentRegistration.Unmarshal(m, b)
}
func (m *AgentRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentRegistration.Marshal(b, m, deterministic)
}
func (dst *AgentRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentRegistration.Merge(dst, src)
}
func (m *AgentRegistration) XXX_Size() int {
	return xxx_messageInfo_AgentRegistration.Size(m)
}
func (m *AgentRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_AgentRegistration proto.InternalMessageInfo

func (m *AgentRegistration) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AgentRegistration) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type AgentPrivateToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentPrivateToken) Reset()         { *m = AgentPrivateToken{} }
func (m *AgentPrivateToken) String() string { return proto.CompactTextString(m) }
func (*AgentPrivateToken) ProtoMessage()    {}
func (*AgentPrivateToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_117924c65de44d70, []int{1}
}
func (m *AgentPrivateToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentPrivateToken.Unmarshal(m, b)
}
func (m *AgentPrivateToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentPrivateToken.Marshal(b, m, deterministic)
}
func (dst *AgentPrivateToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentPrivateToken.Merge(dst, src)
}
func (m *AgentPrivateToken) XXX_Size() int {
	return xxx_messageInfo_AgentPrivateToken.Size(m)
}
func (m *AgentPrivateToken) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentPrivateToken.DiscardUnknown(m)
}

var xxx_messageInfo_AgentPrivateToken proto.InternalMessageInfo

func (m *AgentPrivateToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AgentConfiguration struct {
	Token                string                    `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Commands             map[string]*RemoteCommand `protobuf:"bytes,2,rep,name=commands,proto3" json:"commands,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels               map[string]string         `protobuf:"bytes,3,rep,name=Labels,proto3" json:"Labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AgentID              string                    `protobuf:"bytes,4,opt,name=agentID,proto3" json:"agentID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AgentConfiguration) Reset()         { *m = AgentConfiguration{} }
func (m *AgentConfiguration) String() string { return proto.CompactTextString(m) }
func (*AgentConfiguration) ProtoMessage()    {}
func (*AgentConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_117924c65de44d70, []int{2}
}
func (m *AgentConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentConfiguration.Unmarshal(m, b)
}
func (m *AgentConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentConfiguration.Marshal(b, m, deterministic)
}
func (dst *AgentConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentConfiguration.Merge(dst, src)
}
func (m *AgentConfiguration) XXX_Size() int {
	return xxx_messageInfo_AgentConfiguration.Size(m)
}
func (m *AgentConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_AgentConfiguration proto.InternalMessageInfo

func (m *AgentConfiguration) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AgentConfiguration) GetCommands() map[string]*RemoteCommand {
	if m != nil {
		return m.Commands
	}
	return nil
}

func (m *AgentConfiguration) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *AgentConfiguration) GetAgentID() string {
	if m != nil {
		return m.AgentID
	}
	return ""
}

type CommandFinish struct {
	JobID                uint64   `protobuf:"varint,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	AgentID              string   `protobuf:"bytes,4,opt,name=agentID,proto3" json:"agentID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandFinish) Reset()         { *m = CommandFinish{} }
func (m *CommandFinish) String() string { return proto.CompactTextString(m) }
func (*CommandFinish) ProtoMessage()    {}
func (*CommandFinish) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_117924c65de44d70, []int{3}
}
func (m *CommandFinish) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandFinish.Unmarshal(m, b)
}
func (m *CommandFinish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandFinish.Marshal(b, m, deterministic)
}
func (dst *CommandFinish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandFinish.Merge(dst, src)
}
func (m *CommandFinish) XXX_Size() int {
	return xxx_messageInfo_CommandFinish.Size(m)
}
func (m *CommandFinish) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandFinish.DiscardUnknown(m)
}

var xxx_messageInfo_CommandFinish proto.InternalMessageInfo

func (m *CommandFinish) GetJobID() uint64 {
	if m != nil {
		return m.JobID
	}
	return 0
}

func (m *CommandFinish) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CommandFinish) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CommandFinish) GetAgentID() string {
	if m != nil {
		return m.AgentID
	}
	return ""
}

type Help struct {
	Summary              string   `protobuf:"bytes,1,opt,name=Summary,proto3" json:"Summary,omitempty"`
	Args                 []string `protobuf:"bytes,2,rep,name=Args,proto3" json:"Args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Help) Reset()         { *m = Help{} }
func (m *Help) String() string { return proto.CompactTextString(m) }
func (*Help) ProtoMessage()    {}
func (*Help) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_117924c65de44d70, []int{4}
}
func (m *Help) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Help.Unmarshal(m, b)
}
func (m *Help) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Help.Marshal(b, m, deterministic)
}
func (dst *Help) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Help.Merge(dst, src)
}
func (m *Help) XXX_Size() int {
	return xxx_messageInfo_Help.Size(m)
}
func (m *Help) XXX_DiscardUnknown() {
	xxx_messageInfo_Help.DiscardUnknown(m)
}

var xxx_messageInfo_Help proto.InternalMessageInfo

func (m *Help) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *Help) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type RemoteCommand struct {
	Timeout              int64    `protobuf:"varint,1,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	AuthStrategy         string   `protobuf:"bytes,2,opt,name=AuthStrategy,proto3" json:"AuthStrategy,omitempty"`
	AllowedGroups        []string `protobuf:"bytes,3,rep,name=AllowedGroups,proto3" json:"AllowedGroups,omitempty"`
	ChannelStrategy      string   `protobuf:"bytes,4,opt,name=ChannelStrategy,proto3" json:"ChannelStrategy,omitempty"`
	AllowedChannels      []string `protobuf:"bytes,5,rep,name=AllowedChannels,proto3" json:"AllowedChannels,omitempty"`
	Help                 *Help    `protobuf:"bytes,6,opt,name=help,proto3" json:"help,omitempty"`
	HasHandshake         bool     `protobuf:"varint,7,opt,name=hasHandshake,proto3" json:"hasHandshake,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteCommand) Reset()         { *m = RemoteCommand{} }
func (m *RemoteCommand) String() string { return proto.CompactTextString(m) }
func (*RemoteCommand) ProtoMessage()    {}
func (*RemoteCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_117924c65de44d70, []int{5}
}
func (m *RemoteCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteCommand.Unmarshal(m, b)
}
func (m *RemoteCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteCommand.Marshal(b, m, deterministic)
}
func (dst *RemoteCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteCommand.Merge(dst, src)
}
func (m *RemoteCommand) XXX_Size() int {
	return xxx_messageInfo_RemoteCommand.Size(m)
}
func (m *RemoteCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteCommand.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteCommand proto.InternalMessageInfo

func (m *RemoteCommand) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *RemoteCommand) GetAuthStrategy() string {
	if m != nil {
		return m.AuthStrategy
	}
	return ""
}

func (m *RemoteCommand) GetAllowedGroups() []string {
	if m != nil {
		return m.AllowedGroups
	}
	return nil
}

func (m *RemoteCommand) GetChannelStrategy() string {
	if m != nil {
		return m.ChannelStrategy
	}
	return ""
}

func (m *RemoteCommand) GetAllowedChannels() []string {
	if m != nil {
		return m.AllowedChannels
	}
	return nil
}

func (m *RemoteCommand) GetHelp() *Help {
	if m != nil {
		return m.Help
	}
	return nil
}

func (m *RemoteCommand) GetHasHandshake() bool {
	if m != nil {
		return m.HasHandshake
	}
	return false
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_117924c65de44d70, []int{6}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type CommandRequest struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Args                 []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	UserID               string   `protobuf:"bytes,4,opt,name=userID,proto3" json:"userID,omitempty"`
	UserLink             string   `protobuf:"bytes,5,opt,name=userLink,proto3" json:"userLink,omitempty"`
	Channel              string   `protobuf:"bytes,6,opt,name=channel,proto3" json:"channel,omitempty"`
	ChannelID            string   `protobuf:"bytes,7,opt,name=channelID,proto3" json:"channelID,omitempty"`
	ChannelLink          string   `protobuf:"bytes,8,opt,name=channelLink,proto3" json:"channelLink,omitempty"`
	IsIM                 bool     `protobuf:"varint,9,opt,name=isIM,proto3" json:"isIM,omitempty"`
	JobID                uint64   `protobuf:"varint,10,opt,name=jobID,proto3" json:"jobID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandRequest) Reset()         { *m = CommandRequest{} }
func (m *CommandRequest) String() string { return proto.CompactTextString(m) }
func (*CommandRequest) ProtoMessage()    {}
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_117924c65de44d70, []int{7}
}
func (m *CommandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandRequest.Unmarshal(m, b)
}
func (m *CommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandRequest.Marshal(b, m, deterministic)
}
func (dst *CommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandRequest.Merge(dst, src)
}
func (m *CommandRequest) XXX_Size() int {
	return xxx_messageInfo_CommandRequest.Size(m)
}
func (m *CommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommandRequest proto.InternalMessageInfo

func (m *CommandRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *CommandRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *CommandRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CommandRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *CommandRequest) GetUserLink() string {
	if m != nil {
		return m.UserLink
	}
	return ""
}

func (m *CommandRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *CommandRequest) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *CommandRequest) GetChannelLink() string {
	if m != nil {
		return m.ChannelLink
	}
	return ""
}

func (m *CommandRequest) GetIsIM() bool {
	if m != nil {
		return m.IsIM
	}
	return false
}

func (m *CommandRequest) GetJobID() uint64 {
	if m != nil {
		return m.JobID
	}
	return 0
}

type LogEntry struct {
	JobID                uint64   `protobuf:"varint,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	Line                 string   `protobuf:"bytes,2,opt,name=line,proto3" json:"line,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_117924c65de44d70, []int{8}
}
func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (dst *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(dst, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetJobID() uint64 {
	if m != nil {
		return m.JobID
	}
	return 0
}

func (m *LogEntry) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

type ErrorLogEntry struct {
	JobID                uint64   `protobuf:"varint,1,opt,name=jobID,proto3" json:"jobID,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorLogEntry) Reset()         { *m = ErrorLogEntry{} }
func (m *ErrorLogEntry) String() string { return proto.CompactTextString(m) }
func (*ErrorLogEntry) ProtoMessage()    {}
func (*ErrorLogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_117924c65de44d70, []int{9}
}
func (m *ErrorLogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorLogEntry.Unmarshal(m, b)
}
func (m *ErrorLogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorLogEntry.Marshal(b, m, deterministic)
}
func (dst *ErrorLogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorLogEntry.Merge(dst, src)
}
func (m *ErrorLogEntry) XXX_Size() int {
	return xxx_messageInfo_ErrorLogEntry.Size(m)
}
func (m *ErrorLogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorLogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorLogEntry proto.InternalMessageInfo

func (m *ErrorLogEntry) GetJobID() uint64 {
	if m != nil {
		return m.JobID
	}
	return 0
}

func (m *ErrorLogEntry) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*AgentRegistration)(nil), "api.AgentRegistration")
	proto.RegisterType((*AgentPrivateToken)(nil), "api.AgentPrivateToken")
	proto.RegisterType((*AgentConfiguration)(nil), "api.AgentConfiguration")
	proto.RegisterMapType((map[string]*RemoteCommand)(nil), "api.AgentConfiguration.CommandsEntry")
	proto.RegisterMapType((map[string]string)(nil), "api.AgentConfiguration.LabelsEntry")
	proto.RegisterType((*CommandFinish)(nil), "api.CommandFinish")
	proto.RegisterType((*Help)(nil), "api.Help")
	proto.RegisterType((*RemoteCommand)(nil), "api.RemoteCommand")
	proto.RegisterType((*Empty)(nil), "api.Empty")
	proto.RegisterType((*CommandRequest)(nil), "api.CommandRequest")
	proto.RegisterType((*LogEntry)(nil), "api.LogEntry")
	proto.RegisterType((*ErrorLogEntry)(nil), "api.ErrorLogEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegistrationClient is the client API for Registration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistrationClient interface {
	Register(ctx context.Context, in *AgentRegistration, opts ...grpc.CallOption) (*AgentPrivateToken, error)
}

type registrationClient struct {
	cc *grpc.ClientConn
}

func NewRegistrationClient(cc *grpc.ClientConn) RegistrationClient {
	return &registrationClient{cc}
}

func (c *registrationClient) Register(ctx context.Context, in *AgentRegistration, opts ...grpc.CallOption) (*AgentPrivateToken, error) {
	out := new(AgentPrivateToken)
	err := c.cc.Invoke(ctx, "/api.Registration/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationServer is the server API for Registration service.
type RegistrationServer interface {
	Register(context.Context, *AgentRegistration) (*AgentPrivateToken, error)
}

func RegisterRegistrationServer(s *grpc.Server, srv RegistrationServer) {
	s.RegisterService(&_Registration_serviceDesc, srv)
}

func _Registration_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentRegistration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Registration/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).Register(ctx, req.(*AgentRegistration))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Registration",
	HandlerType: (*RegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Registration_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// CommandPipelineClient is the client API for CommandPipeline service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommandPipelineClient interface {
	RegisterAgent(ctx context.Context, in *AgentConfiguration, opts ...grpc.CallOption) (CommandPipeline_RegisterAgentClient, error)
	Finish(ctx context.Context, in *CommandFinish, opts ...grpc.CallOption) (*Empty, error)
}

type commandPipelineClient struct {
	cc *grpc.ClientConn
}

func NewCommandPipelineClient(cc *grpc.ClientConn) CommandPipelineClient {
	return &commandPipelineClient{cc}
}

func (c *commandPipelineClient) RegisterAgent(ctx context.Context, in *AgentConfiguration, opts ...grpc.CallOption) (CommandPipeline_RegisterAgentClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CommandPipeline_serviceDesc.Streams[0], "/api.CommandPipeline/RegisterAgent", opts...)
	if err != nil {
		return nil, err
	}
	x := &commandPipelineRegisterAgentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CommandPipeline_RegisterAgentClient interface {
	Recv() (*CommandRequest, error)
	grpc.ClientStream
}

type commandPipelineRegisterAgentClient struct {
	grpc.ClientStream
}

func (x *commandPipelineRegisterAgentClient) Recv() (*CommandRequest, error) {
	m := new(CommandRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *commandPipelineClient) Finish(ctx context.Context, in *CommandFinish, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.CommandPipeline/Finish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandPipelineServer is the server API for CommandPipeline service.
type CommandPipelineServer interface {
	RegisterAgent(*AgentConfiguration, CommandPipeline_RegisterAgentServer) error
	Finish(context.Context, *CommandFinish) (*Empty, error)
}

func RegisterCommandPipelineServer(s *grpc.Server, srv CommandPipelineServer) {
	s.RegisterService(&_CommandPipeline_serviceDesc, srv)
}

func _CommandPipeline_RegisterAgent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AgentConfiguration)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommandPipelineServer).RegisterAgent(m, &commandPipelineRegisterAgentServer{stream})
}

type CommandPipeline_RegisterAgentServer interface {
	Send(*CommandRequest) error
	grpc.ServerStream
}

type commandPipelineRegisterAgentServer struct {
	grpc.ServerStream
}

func (x *commandPipelineRegisterAgentServer) Send(m *CommandRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _CommandPipeline_Finish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandFinish)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandPipelineServer).Finish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CommandPipeline/Finish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandPipelineServer).Finish(ctx, req.(*CommandFinish))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommandPipeline_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.CommandPipeline",
	HandlerType: (*CommandPipelineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Finish",
			Handler:    _CommandPipeline_Finish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterAgent",
			Handler:       _CommandPipeline_RegisterAgent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}

// LogWriterClient is the client API for LogWriter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogWriterClient interface {
	Append(ctx context.Context, opts ...grpc.CallOption) (LogWriter_AppendClient, error)
	SetError(ctx context.Context, in *ErrorLogEntry, opts ...grpc.CallOption) (*Empty, error)
}

type logWriterClient struct {
	cc *grpc.ClientConn
}

func NewLogWriterClient(cc *grpc.ClientConn) LogWriterClient {
	return &logWriterClient{cc}
}

func (c *logWriterClient) Append(ctx context.Context, opts ...grpc.CallOption) (LogWriter_AppendClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogWriter_serviceDesc.Streams[0], "/api.LogWriter/Append", opts...)
	if err != nil {
		return nil, err
	}
	x := &logWriterAppendClient{stream}
	return x, nil
}

type LogWriter_AppendClient interface {
	Send(*LogEntry) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type logWriterAppendClient struct {
	grpc.ClientStream
}

func (x *logWriterAppendClient) Send(m *LogEntry) error {
	return x.ClientStream.SendMsg(m)
}

func (x *logWriterAppendClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logWriterClient) SetError(ctx context.Context, in *ErrorLogEntry, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.LogWriter/SetError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogWriterServer is the server API for LogWriter service.
type LogWriterServer interface {
	Append(LogWriter_AppendServer) error
	SetError(context.Context, *ErrorLogEntry) (*Empty, error)
}

func RegisterLogWriterServer(s *grpc.Server, srv LogWriterServer) {
	s.RegisterService(&_LogWriter_serviceDesc, srv)
}

func _LogWriter_Append_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LogWriterServer).Append(&logWriterAppendServer{stream})
}

type LogWriter_AppendServer interface {
	SendAndClose(*Empty) error
	Recv() (*LogEntry, error)
	grpc.ServerStream
}

type logWriterAppendServer struct {
	grpc.ServerStream
}

func (x *logWriterAppendServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *logWriterAppendServer) Recv() (*LogEntry, error) {
	m := new(LogEntry)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _LogWriter_SetError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ErrorLogEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogWriterServer).SetError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.LogWriter/SetError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogWriterServer).SetError(ctx, req.(*ErrorLogEntry))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogWriter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.LogWriter",
	HandlerType: (*LogWriterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetError",
			Handler:    _LogWriter_SetError_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Append",
			Handler:       _LogWriter_Append_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_api_117924c65de44d70) }

var fileDescriptor_api_117924c65de44d70 = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0x6d, 0x3e, 0x1b, 0xdf, 0x34, 0xaf, 0xef, 0xcd, 0xab, 0xfa, 0xac, 0xe8, 0x21, 0x45, 0x06,
	0x44, 0x40, 0xa8, 0x42, 0xa1, 0x0b, 0xa0, 0x6c, 0xac, 0x36, 0xd0, 0x48, 0x41, 0x54, 0x6e, 0x25,
	0xd6, 0x93, 0x76, 0x70, 0x86, 0xd8, 0x33, 0xc6, 0x1e, 0x17, 0x65, 0xc7, 0x9a, 0x25, 0xbf, 0x8c,
	0x9f, 0x84, 0xe6, 0xce, 0xd8, 0xb1, 0x4b, 0x0b, 0xbb, 0x7b, 0xae, 0xef, 0x39, 0x73, 0xbf, 0x66,
	0x0c, 0x0e, 0x4d, 0xf8, 0x41, 0x92, 0x4a, 0x25, 0x49, 0x8b, 0x26, 0xdc, 0x9b, 0xc2, 0x3f, 0x7e,
	0xc8, 0x84, 0x0a, 0x58, 0xc8, 0x33, 0x95, 0x52, 0xc5, 0xa5, 0x20, 0x7b, 0xd0, 0xb9, 0x90, 0x2b,
	0x26, 0xdc, 0xc6, 0xa8, 0x31, 0x76, 0x02, 0x03, 0xc8, 0x10, 0x7a, 0xa7, 0x32, 0x53, 0x82, 0xc6,
	0xcc, 0x6d, 0xe2, 0x87, 0x12, 0x7b, 0x8f, 0xad, 0xcc, 0x59, 0xca, 0xaf, 0xa9, 0x62, 0x86, 0x70,
	0xab, 0x8c, 0xf7, 0xa3, 0x09, 0x04, 0x63, 0x8f, 0xa5, 0xf8, 0xc8, 0xc3, 0xfc, 0xb7, 0x67, 0xfa,
	0xd0, 0xbb, 0x94, 0x71, 0x4c, 0xc5, 0x55, 0xe6, 0x36, 0x47, 0xad, 0x71, 0x7f, 0xf2, 0xf0, 0x40,
	0x57, 0xf0, 0xab, 0xc0, 0xc1, 0xb1, 0x8d, 0x9b, 0x0a, 0x95, 0xae, 0x83, 0x92, 0x46, 0x8e, 0xa0,
	0x3b, 0xa7, 0x0b, 0x16, 0x65, 0x6e, 0x0b, 0x05, 0xee, 0xdf, 0x25, 0x60, 0xa2, 0x0c, 0xdd, 0x52,
	0x88, 0x0b, 0xdb, 0x54, 0x47, 0xce, 0x4e, 0xdc, 0x36, 0xe6, 0x55, 0xc0, 0xe1, 0x7b, 0x18, 0xd4,
	0x4e, 0x24, 0x7f, 0x43, 0x6b, 0xc5, 0xd6, 0x36, 0x7d, 0x6d, 0x92, 0x31, 0x74, 0xae, 0x69, 0x94,
	0x9b, 0x6e, 0xf5, 0x27, 0x04, 0x0f, 0x0e, 0x58, 0x2c, 0x15, 0xb3, 0xd4, 0xc0, 0x04, 0xbc, 0x6a,
	0xbe, 0x68, 0x0c, 0x5f, 0x42, 0xbf, 0x92, 0xc1, 0x2d, 0x72, 0x7b, 0x55, 0x39, 0xa7, 0x42, 0xf5,
	0x64, 0x99, 0xcb, 0x1b, 0x2e, 0x78, 0xb6, 0xd4, 0xa1, 0x9f, 0xe4, 0x62, 0x76, 0x82, 0xf4, 0x76,
	0x60, 0x80, 0x2e, 0xe6, 0x52, 0x0a, 0xc5, 0x84, 0xb2, 0x12, 0x05, 0xd4, 0xf1, 0x2c, 0x4d, 0x65,
	0xea, 0xb6, 0x8c, 0x34, 0x82, 0xbb, 0x8b, 0xf7, 0x0e, 0xa1, 0x7d, 0xca, 0xa2, 0x44, 0x47, 0x9c,
	0xe7, 0x71, 0x4c, 0xd3, 0x22, 0xd1, 0x02, 0x12, 0x02, 0x6d, 0x3f, 0x0d, 0xcd, 0xd0, 0x9c, 0x00,
	0x6d, 0xef, 0x5b, 0x13, 0x06, 0xb5, 0xf2, 0x35, 0xff, 0x82, 0xc7, 0x4c, 0xe6, 0x0a, 0xf9, 0xad,
	0xa0, 0x80, 0xc4, 0x83, 0x1d, 0x3f, 0x57, 0xcb, 0x73, 0xbd, 0x92, 0x2c, 0x5c, 0xdb, 0x84, 0x6b,
	0x3e, 0xf2, 0x00, 0x06, 0x7e, 0x14, 0xc9, 0x2f, 0xec, 0xea, 0x6d, 0x2a, 0xf3, 0xc4, 0x0c, 0xd8,
	0x09, 0xea, 0x4e, 0x32, 0x86, 0xdd, 0xe3, 0x25, 0x15, 0x82, 0x45, 0xa5, 0x98, 0xa9, 0xe6, 0xa6,
	0x5b, 0x47, 0x5a, 0xaa, 0xfd, 0x92, 0xb9, 0x1d, 0x54, 0xbc, 0xe9, 0x26, 0xf7, 0xa0, 0xbd, 0x64,
	0x51, 0xe2, 0x76, 0x71, 0xb0, 0x0e, 0x0e, 0x56, 0x37, 0x24, 0x40, 0xb7, 0x4e, 0x7e, 0x49, 0xb3,
	0x53, 0xbd, 0x1b, 0x4b, 0xba, 0x62, 0xee, 0xf6, 0xa8, 0x31, 0xee, 0x05, 0x35, 0x9f, 0xb7, 0x0d,
	0x9d, 0x69, 0x9c, 0xa8, 0xb5, 0xf7, 0xbd, 0x09, 0x7f, 0x15, 0xeb, 0xc0, 0x3e, 0xe7, 0x2c, 0x53,
	0x66, 0x50, 0xe8, 0x29, 0xda, 0x6a, 0xa1, 0x6e, 0x2b, 0xad, 0xb4, 0x55, 0xdb, 0xfa, 0x5e, 0xe6,
	0x19, 0x4b, 0xf1, 0x5e, 0x9a, 0xf9, 0x95, 0x98, 0xec, 0x43, 0x57, 0xdb, 0xe5, 0x04, 0x2d, 0x2a,
	0x38, 0x73, 0x2e, 0x56, 0x6e, 0x67, 0xc3, 0xd1, 0x18, 0x4f, 0x37, 0x85, 0x62, 0x7d, 0xfa, 0x74,
	0x03, 0xc9, 0xff, 0xe0, 0x58, 0x73, 0x76, 0x82, 0x45, 0x39, 0xc1, 0xc6, 0x41, 0x46, 0xd0, 0xb7,
	0x00, 0x65, 0x7b, 0xf8, 0xbd, 0xea, 0xd2, 0xd9, 0xf3, 0x6c, 0xf6, 0xce, 0x75, 0xb0, 0x1f, 0x68,
	0x6f, 0x56, 0x15, 0x2a, 0xab, 0xea, 0x1d, 0x42, 0x6f, 0x2e, 0x43, 0x73, 0x13, 0x6e, 0x5f, 0x66,
	0x02, 0xed, 0x88, 0x8b, 0xe2, 0x32, 0xa0, 0xed, 0x1d, 0xc1, 0x60, 0xaa, 0x37, 0xf7, 0x0f, 0xd4,
	0x72, 0xdb, 0x9b, 0x95, 0x6d, 0x9f, 0xcc, 0x61, 0xa7, 0xf6, 0x08, 0xbe, 0x86, 0x9e, 0xc1, 0x2c,
	0x25, 0xfb, 0x9b, 0x37, 0xa3, 0x1a, 0x33, 0xac, 0xf8, 0xab, 0x2f, 0x9f, 0xb7, 0x35, 0xf9, 0xda,
	0x80, 0x5d, 0x3b, 0xd5, 0x33, 0x9e, 0x30, 0x9d, 0x1e, 0xf1, 0xf5, 0xfa, 0x1b, 0x45, 0xa4, 0x90,
	0xff, 0xee, 0x78, 0x8a, 0x86, 0xff, 0xe2, 0x87, 0xfa, 0x56, 0x78, 0x5b, 0xcf, 0x1a, 0xe4, 0x09,
	0x74, 0xed, 0x15, 0x27, 0xd5, 0x10, 0xe3, 0x1b, 0x02, 0xfa, 0xcc, 0x5a, 0x6d, 0x4d, 0x16, 0xe0,
	0xcc, 0x65, 0xf8, 0x21, 0xe5, 0xba, 0x82, 0x47, 0xd0, 0xf5, 0x93, 0x84, 0x89, 0x2b, 0x32, 0xc0,
	0xa0, 0xa2, 0x45, 0x75, 0xce, 0xb8, 0x41, 0x9e, 0x42, 0xef, 0x9c, 0x29, 0x6c, 0xa3, 0x3d, 0xa3,
	0xd6, 0xd2, 0x7a, 0xfc, 0xa2, 0x8b, 0xbf, 0x92, 0xe7, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xde,
	0xce, 0xc9, 0x96, 0x57, 0x06, 0x00, 0x00,
}