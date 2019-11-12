// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v3alpha/server_info.proto

package envoy_admin_v3alpha

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type ServerInfo_State int32

const (
	ServerInfo_LIVE             ServerInfo_State = 0
	ServerInfo_DRAINING         ServerInfo_State = 1
	ServerInfo_PRE_INITIALIZING ServerInfo_State = 2
	ServerInfo_INITIALIZING     ServerInfo_State = 3
)

var ServerInfo_State_name = map[int32]string{
	0: "LIVE",
	1: "DRAINING",
	2: "PRE_INITIALIZING",
	3: "INITIALIZING",
}

var ServerInfo_State_value = map[string]int32{
	"LIVE":             0,
	"DRAINING":         1,
	"PRE_INITIALIZING": 2,
	"INITIALIZING":     3,
}

func (x ServerInfo_State) String() string {
	return proto.EnumName(ServerInfo_State_name, int32(x))
}

func (ServerInfo_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2bf208be306ad97, []int{0, 0}
}

type CommandLineOptions_IpVersion int32

const (
	CommandLineOptions_v4 CommandLineOptions_IpVersion = 0
	CommandLineOptions_v6 CommandLineOptions_IpVersion = 1
)

var CommandLineOptions_IpVersion_name = map[int32]string{
	0: "v4",
	1: "v6",
}

var CommandLineOptions_IpVersion_value = map[string]int32{
	"v4": 0,
	"v6": 1,
}

func (x CommandLineOptions_IpVersion) String() string {
	return proto.EnumName(CommandLineOptions_IpVersion_name, int32(x))
}

func (CommandLineOptions_IpVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2bf208be306ad97, []int{1, 0}
}

type CommandLineOptions_Mode int32

const (
	CommandLineOptions_Serve    CommandLineOptions_Mode = 0
	CommandLineOptions_Validate CommandLineOptions_Mode = 1
	CommandLineOptions_InitOnly CommandLineOptions_Mode = 2
)

var CommandLineOptions_Mode_name = map[int32]string{
	0: "Serve",
	1: "Validate",
	2: "InitOnly",
}

var CommandLineOptions_Mode_value = map[string]int32{
	"Serve":    0,
	"Validate": 1,
	"InitOnly": 2,
}

func (x CommandLineOptions_Mode) String() string {
	return proto.EnumName(CommandLineOptions_Mode_name, int32(x))
}

func (CommandLineOptions_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2bf208be306ad97, []int{1, 1}
}

type ServerInfo struct {
	Version              string              `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	State                ServerInfo_State    `protobuf:"varint,2,opt,name=state,proto3,enum=envoy.admin.v3alpha.ServerInfo_State" json:"state,omitempty"`
	UptimeCurrentEpoch   *duration.Duration  `protobuf:"bytes,3,opt,name=uptime_current_epoch,json=uptimeCurrentEpoch,proto3" json:"uptime_current_epoch,omitempty"`
	UptimeAllEpochs      *duration.Duration  `protobuf:"bytes,4,opt,name=uptime_all_epochs,json=uptimeAllEpochs,proto3" json:"uptime_all_epochs,omitempty"`
	HotRestartVersion    string              `protobuf:"bytes,5,opt,name=hot_restart_version,json=hotRestartVersion,proto3" json:"hot_restart_version,omitempty"`
	CommandLineOptions   *CommandLineOptions `protobuf:"bytes,6,opt,name=command_line_options,json=commandLineOptions,proto3" json:"command_line_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ServerInfo) Reset()         { *m = ServerInfo{} }
func (m *ServerInfo) String() string { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()    {}
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2bf208be306ad97, []int{0}
}

func (m *ServerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerInfo.Unmarshal(m, b)
}
func (m *ServerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerInfo.Marshal(b, m, deterministic)
}
func (m *ServerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerInfo.Merge(m, src)
}
func (m *ServerInfo) XXX_Size() int {
	return xxx_messageInfo_ServerInfo.Size(m)
}
func (m *ServerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServerInfo proto.InternalMessageInfo

func (m *ServerInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServerInfo) GetState() ServerInfo_State {
	if m != nil {
		return m.State
	}
	return ServerInfo_LIVE
}

func (m *ServerInfo) GetUptimeCurrentEpoch() *duration.Duration {
	if m != nil {
		return m.UptimeCurrentEpoch
	}
	return nil
}

func (m *ServerInfo) GetUptimeAllEpochs() *duration.Duration {
	if m != nil {
		return m.UptimeAllEpochs
	}
	return nil
}

func (m *ServerInfo) GetHotRestartVersion() string {
	if m != nil {
		return m.HotRestartVersion
	}
	return ""
}

func (m *ServerInfo) GetCommandLineOptions() *CommandLineOptions {
	if m != nil {
		return m.CommandLineOptions
	}
	return nil
}

type CommandLineOptions struct {
	BaseId                     uint64                       `protobuf:"varint,1,opt,name=base_id,json=baseId,proto3" json:"base_id,omitempty"`
	Concurrency                uint32                       `protobuf:"varint,2,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	ConfigPath                 string                       `protobuf:"bytes,3,opt,name=config_path,json=configPath,proto3" json:"config_path,omitempty"`
	ConfigYaml                 string                       `protobuf:"bytes,4,opt,name=config_yaml,json=configYaml,proto3" json:"config_yaml,omitempty"`
	AllowUnknownStaticFields   bool                         `protobuf:"varint,5,opt,name=allow_unknown_static_fields,json=allowUnknownStaticFields,proto3" json:"allow_unknown_static_fields,omitempty"`
	RejectUnknownDynamicFields bool                         `protobuf:"varint,26,opt,name=reject_unknown_dynamic_fields,json=rejectUnknownDynamicFields,proto3" json:"reject_unknown_dynamic_fields,omitempty"`
	AdminAddressPath           string                       `protobuf:"bytes,6,opt,name=admin_address_path,json=adminAddressPath,proto3" json:"admin_address_path,omitempty"`
	LocalAddressIpVersion      CommandLineOptions_IpVersion `protobuf:"varint,7,opt,name=local_address_ip_version,json=localAddressIpVersion,proto3,enum=envoy.admin.v3alpha.CommandLineOptions_IpVersion" json:"local_address_ip_version,omitempty"`
	LogLevel                   string                       `protobuf:"bytes,8,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	ComponentLogLevel          string                       `protobuf:"bytes,9,opt,name=component_log_level,json=componentLogLevel,proto3" json:"component_log_level,omitempty"`
	LogFormat                  string                       `protobuf:"bytes,10,opt,name=log_format,json=logFormat,proto3" json:"log_format,omitempty"`
	LogFormatEscaped           bool                         `protobuf:"varint,27,opt,name=log_format_escaped,json=logFormatEscaped,proto3" json:"log_format_escaped,omitempty"`
	LogPath                    string                       `protobuf:"bytes,11,opt,name=log_path,json=logPath,proto3" json:"log_path,omitempty"`
	ServiceCluster             string                       `protobuf:"bytes,13,opt,name=service_cluster,json=serviceCluster,proto3" json:"service_cluster,omitempty"`
	ServiceNode                string                       `protobuf:"bytes,14,opt,name=service_node,json=serviceNode,proto3" json:"service_node,omitempty"`
	ServiceZone                string                       `protobuf:"bytes,15,opt,name=service_zone,json=serviceZone,proto3" json:"service_zone,omitempty"`
	FileFlushInterval          *duration.Duration           `protobuf:"bytes,16,opt,name=file_flush_interval,json=fileFlushInterval,proto3" json:"file_flush_interval,omitempty"`
	DrainTime                  *duration.Duration           `protobuf:"bytes,17,opt,name=drain_time,json=drainTime,proto3" json:"drain_time,omitempty"`
	ParentShutdownTime         *duration.Duration           `protobuf:"bytes,18,opt,name=parent_shutdown_time,json=parentShutdownTime,proto3" json:"parent_shutdown_time,omitempty"`
	Mode                       CommandLineOptions_Mode      `protobuf:"varint,19,opt,name=mode,proto3,enum=envoy.admin.v3alpha.CommandLineOptions_Mode" json:"mode,omitempty"`
	DisableHotRestart          bool                         `protobuf:"varint,22,opt,name=disable_hot_restart,json=disableHotRestart,proto3" json:"disable_hot_restart,omitempty"`
	EnableMutexTracing         bool                         `protobuf:"varint,23,opt,name=enable_mutex_tracing,json=enableMutexTracing,proto3" json:"enable_mutex_tracing,omitempty"`
	RestartEpoch               uint32                       `protobuf:"varint,24,opt,name=restart_epoch,json=restartEpoch,proto3" json:"restart_epoch,omitempty"`
	CpusetThreads              bool                         `protobuf:"varint,25,opt,name=cpuset_threads,json=cpusetThreads,proto3" json:"cpuset_threads,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                     `json:"-"`
	XXX_unrecognized           []byte                       `json:"-"`
	XXX_sizecache              int32                        `json:"-"`
}

func (m *CommandLineOptions) Reset()         { *m = CommandLineOptions{} }
func (m *CommandLineOptions) String() string { return proto.CompactTextString(m) }
func (*CommandLineOptions) ProtoMessage()    {}
func (*CommandLineOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2bf208be306ad97, []int{1}
}

func (m *CommandLineOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandLineOptions.Unmarshal(m, b)
}
func (m *CommandLineOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandLineOptions.Marshal(b, m, deterministic)
}
func (m *CommandLineOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandLineOptions.Merge(m, src)
}
func (m *CommandLineOptions) XXX_Size() int {
	return xxx_messageInfo_CommandLineOptions.Size(m)
}
func (m *CommandLineOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandLineOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CommandLineOptions proto.InternalMessageInfo

func (m *CommandLineOptions) GetBaseId() uint64 {
	if m != nil {
		return m.BaseId
	}
	return 0
}

func (m *CommandLineOptions) GetConcurrency() uint32 {
	if m != nil {
		return m.Concurrency
	}
	return 0
}

func (m *CommandLineOptions) GetConfigPath() string {
	if m != nil {
		return m.ConfigPath
	}
	return ""
}

func (m *CommandLineOptions) GetConfigYaml() string {
	if m != nil {
		return m.ConfigYaml
	}
	return ""
}

func (m *CommandLineOptions) GetAllowUnknownStaticFields() bool {
	if m != nil {
		return m.AllowUnknownStaticFields
	}
	return false
}

func (m *CommandLineOptions) GetRejectUnknownDynamicFields() bool {
	if m != nil {
		return m.RejectUnknownDynamicFields
	}
	return false
}

func (m *CommandLineOptions) GetAdminAddressPath() string {
	if m != nil {
		return m.AdminAddressPath
	}
	return ""
}

func (m *CommandLineOptions) GetLocalAddressIpVersion() CommandLineOptions_IpVersion {
	if m != nil {
		return m.LocalAddressIpVersion
	}
	return CommandLineOptions_v4
}

func (m *CommandLineOptions) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *CommandLineOptions) GetComponentLogLevel() string {
	if m != nil {
		return m.ComponentLogLevel
	}
	return ""
}

func (m *CommandLineOptions) GetLogFormat() string {
	if m != nil {
		return m.LogFormat
	}
	return ""
}

func (m *CommandLineOptions) GetLogFormatEscaped() bool {
	if m != nil {
		return m.LogFormatEscaped
	}
	return false
}

func (m *CommandLineOptions) GetLogPath() string {
	if m != nil {
		return m.LogPath
	}
	return ""
}

func (m *CommandLineOptions) GetServiceCluster() string {
	if m != nil {
		return m.ServiceCluster
	}
	return ""
}

func (m *CommandLineOptions) GetServiceNode() string {
	if m != nil {
		return m.ServiceNode
	}
	return ""
}

func (m *CommandLineOptions) GetServiceZone() string {
	if m != nil {
		return m.ServiceZone
	}
	return ""
}

func (m *CommandLineOptions) GetFileFlushInterval() *duration.Duration {
	if m != nil {
		return m.FileFlushInterval
	}
	return nil
}

func (m *CommandLineOptions) GetDrainTime() *duration.Duration {
	if m != nil {
		return m.DrainTime
	}
	return nil
}

func (m *CommandLineOptions) GetParentShutdownTime() *duration.Duration {
	if m != nil {
		return m.ParentShutdownTime
	}
	return nil
}

func (m *CommandLineOptions) GetMode() CommandLineOptions_Mode {
	if m != nil {
		return m.Mode
	}
	return CommandLineOptions_Serve
}

func (m *CommandLineOptions) GetDisableHotRestart() bool {
	if m != nil {
		return m.DisableHotRestart
	}
	return false
}

func (m *CommandLineOptions) GetEnableMutexTracing() bool {
	if m != nil {
		return m.EnableMutexTracing
	}
	return false
}

func (m *CommandLineOptions) GetRestartEpoch() uint32 {
	if m != nil {
		return m.RestartEpoch
	}
	return 0
}

func (m *CommandLineOptions) GetCpusetThreads() bool {
	if m != nil {
		return m.CpusetThreads
	}
	return false
}

func init() {
	proto.RegisterEnum("envoy.admin.v3alpha.ServerInfo_State", ServerInfo_State_name, ServerInfo_State_value)
	proto.RegisterEnum("envoy.admin.v3alpha.CommandLineOptions_IpVersion", CommandLineOptions_IpVersion_name, CommandLineOptions_IpVersion_value)
	proto.RegisterEnum("envoy.admin.v3alpha.CommandLineOptions_Mode", CommandLineOptions_Mode_name, CommandLineOptions_Mode_value)
	proto.RegisterType((*ServerInfo)(nil), "envoy.admin.v3alpha.ServerInfo")
	proto.RegisterType((*CommandLineOptions)(nil), "envoy.admin.v3alpha.CommandLineOptions")
}

func init() {
	proto.RegisterFile("envoy/admin/v3alpha/server_info.proto", fileDescriptor_c2bf208be306ad97)
}

var fileDescriptor_c2bf208be306ad97 = []byte{
	// 955 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xed, 0x4e, 0xdc, 0x46,
	0x14, 0xcd, 0xc2, 0x02, 0xbb, 0x97, 0x2f, 0x33, 0x40, 0x32, 0x80, 0xd2, 0x12, 0x2a, 0x14, 0x7e,
	0xa4, 0xa6, 0x4d, 0xaa, 0xa8, 0x52, 0x55, 0xa9, 0x84, 0x8f, 0xd4, 0x94, 0x10, 0x64, 0x28, 0x52,
	0xf2, 0x67, 0x34, 0x6b, 0xcf, 0xee, 0x0e, 0x1d, 0xcf, 0x58, 0xf6, 0x78, 0xc3, 0xf6, 0x21, 0xfa,
	0x62, 0x7d, 0xa9, 0x68, 0xee, 0x78, 0x77, 0x89, 0x82, 0x44, 0x7e, 0x81, 0xcf, 0x3d, 0xe7, 0x78,
	0xe6, 0xde, 0xbb, 0xc7, 0xb0, 0x2b, 0xf4, 0xc0, 0x0c, 0xf7, 0x79, 0x9a, 0x49, 0xbd, 0x3f, 0x78,
	0xc5, 0x55, 0xde, 0xe7, 0xfb, 0xa5, 0x28, 0x06, 0xa2, 0x60, 0x52, 0x77, 0x4d, 0x98, 0x17, 0xc6,
	0x1a, 0xb2, 0x8a, 0xb4, 0x10, 0x69, 0x61, 0x4d, 0xdb, 0xfc, 0xae, 0x67, 0x4c, 0x4f, 0x89, 0x7d,
	0xa4, 0x74, 0xaa, 0xee, 0x7e, 0x5a, 0x15, 0xdc, 0x4a, 0xa3, 0xbd, 0x68, 0xe7, 0xff, 0x69, 0x80,
	0x4b, 0xb4, 0x8a, 0x74, 0xd7, 0x10, 0x0a, 0x73, 0x03, 0x51, 0x94, 0xd2, 0x68, 0xda, 0xd8, 0x6e,
	0xec, 0xb5, 0xe3, 0xd1, 0x23, 0xf9, 0x0d, 0x66, 0x4a, 0xcb, 0xad, 0xa0, 0x53, 0xdb, 0x8d, 0xbd,
	0xa5, 0x97, 0xbb, 0xe1, 0x3d, 0x6f, 0x0b, 0x27, 0x4e, 0xe1, 0xa5, 0x23, 0xc7, 0x5e, 0x43, 0xfe,
	0x82, 0xb5, 0x2a, 0xb7, 0x32, 0x13, 0x2c, 0xa9, 0x8a, 0x42, 0x68, 0xcb, 0x44, 0x6e, 0x92, 0x3e,
	0x9d, 0xde, 0x6e, 0xec, 0xcd, 0xbf, 0xdc, 0x08, 0xfd, 0x21, 0xc3, 0xd1, 0x21, 0xc3, 0xa3, 0xfa,
	0x90, 0x31, 0xf1, 0xb2, 0x43, 0xaf, 0x3a, 0x76, 0x22, 0x72, 0x0c, 0x2b, 0xb5, 0x19, 0x57, 0xca,
	0x1b, 0x95, 0xb4, 0xf9, 0x90, 0xd3, 0xb2, 0xd7, 0x1c, 0x28, 0x85, 0x2e, 0x25, 0x09, 0x61, 0xb5,
	0x6f, 0x2c, 0x2b, 0x44, 0x69, 0x79, 0x61, 0xd9, 0xe8, 0xda, 0x33, 0x78, 0xed, 0x95, 0xbe, 0xb1,
	0xb1, 0xaf, 0x5c, 0xd7, 0x0d, 0xf8, 0x00, 0x6b, 0x89, 0xc9, 0x32, 0xae, 0x53, 0xa6, 0xa4, 0x16,
	0xcc, 0xe4, 0xce, 0xb7, 0xa4, 0xb3, 0xf8, 0xe6, 0xe7, 0xf7, 0xf6, 0xe3, 0xd0, 0x0b, 0xce, 0xa4,
	0x16, 0xef, 0x3d, 0x3d, 0x26, 0xc9, 0x57, 0xd8, 0xce, 0x5b, 0x98, 0xc1, 0x76, 0x91, 0x16, 0x34,
	0xcf, 0xa2, 0xeb, 0xe3, 0xe0, 0x11, 0x59, 0x80, 0xd6, 0x51, 0x7c, 0x10, 0x9d, 0x47, 0xe7, 0x6f,
	0x83, 0x06, 0x59, 0x83, 0xe0, 0x22, 0x3e, 0x66, 0xd1, 0x79, 0x74, 0x15, 0x1d, 0x9c, 0x45, 0x1f,
	0x1d, 0x3a, 0x45, 0x02, 0x58, 0xf8, 0x02, 0x99, 0xde, 0xf9, 0x0f, 0x80, 0x7c, 0xfd, 0x4e, 0xf2,
	0x04, 0xe6, 0x3a, 0xbc, 0x14, 0x4c, 0xa6, 0x38, 0xd5, 0x66, 0x3c, 0xeb, 0x1e, 0xa3, 0x94, 0x6c,
	0xc3, 0x7c, 0x62, 0xb4, 0x9f, 0x49, 0x32, 0xc4, 0xd1, 0x2e, 0xc6, 0x77, 0x21, 0xf2, 0x3d, 0x32,
	0xba, 0xb2, 0xc7, 0x72, 0x6e, 0xfd, 0xc0, 0xda, 0x31, 0x78, 0xe8, 0x82, 0xdb, 0xfe, 0x1d, 0xc2,
	0x90, 0x67, 0x0a, 0xe7, 0x30, 0x26, 0x7c, 0xe0, 0x99, 0x22, 0xbf, 0xc3, 0x16, 0x57, 0xca, 0x7c,
	0x62, 0x95, 0xfe, 0x47, 0x9b, 0x4f, 0x9a, 0xb9, 0x95, 0x90, 0x09, 0xeb, 0x4a, 0xa1, 0xd2, 0x12,
	0xfb, 0xdd, 0x8a, 0x29, 0x52, 0xfe, 0xf6, 0x8c, 0x4b, 0x24, 0x9c, 0x60, 0x9d, 0x1c, 0xc0, 0xd3,
	0x42, 0xdc, 0x88, 0xc4, 0x8e, 0xf5, 0xe9, 0x50, 0xf3, 0x6c, 0x62, 0xb0, 0x89, 0x06, 0x9b, 0x9e,
	0x54, 0x3b, 0x1c, 0x79, 0x4a, 0x6d, 0xf1, 0x02, 0x08, 0x8e, 0x85, 0xf1, 0x34, 0x2d, 0x44, 0x59,
	0xfa, 0xab, 0xcc, 0xe2, 0x49, 0x03, 0xac, 0x1c, 0xf8, 0x02, 0x5e, 0xe8, 0x06, 0xa8, 0x32, 0x09,
	0x57, 0x63, 0xb6, 0xcc, 0xc7, 0xcb, 0x31, 0x87, 0xbb, 0xff, 0xf3, 0x37, 0xce, 0x3a, 0x8c, 0xf2,
	0x7a, 0x79, 0xe2, 0x75, 0xb4, 0xac, 0x5f, 0x33, 0x86, 0xc9, 0x16, 0xb4, 0x95, 0xe9, 0x31, 0x25,
	0x06, 0x42, 0xd1, 0x16, 0x1e, 0xa8, 0xa5, 0x4c, 0xef, 0xcc, 0x3d, 0xbb, 0x05, 0x4d, 0x4c, 0x96,
	0x1b, 0xed, 0x7e, 0x2f, 0x13, 0x5a, 0xdb, 0x2f, 0xe8, 0xb8, 0x74, 0x36, 0xe2, 0x3f, 0x05, 0x70,
	0xac, 0xae, 0x29, 0x32, 0x6e, 0x29, 0x20, 0xcd, 0xd9, 0x9f, 0x20, 0xe0, 0xba, 0x30, 0x29, 0x33,
	0x51, 0x26, 0x3c, 0x17, 0x29, 0xdd, 0xc2, 0xee, 0x05, 0x63, 0xda, 0xb1, 0xc7, 0xc9, 0x06, 0xb8,
	0x83, 0xf8, 0x4e, 0xcd, 0xfb, 0x24, 0x50, 0xc6, 0x4f, 0xfc, 0x39, 0x2c, 0xbb, 0xf0, 0x91, 0x89,
	0x60, 0x89, 0xaa, 0x4a, 0x2b, 0x0a, 0xba, 0x88, 0x8c, 0xa5, 0x1a, 0x3e, 0xf4, 0x28, 0x79, 0x06,
	0x0b, 0x23, 0xa2, 0x36, 0xa9, 0xa0, 0x4b, 0xc8, 0x9a, 0xaf, 0xb1, 0x73, 0x93, 0x8a, 0xbb, 0x94,
	0x7f, 0x8d, 0x16, 0x74, 0xf9, 0x0b, 0xca, 0x47, 0xa3, 0x05, 0x89, 0x60, 0xb5, 0x2b, 0x95, 0x60,
	0x5d, 0x55, 0x95, 0x7d, 0x26, 0xb5, 0x15, 0xc5, 0x80, 0x2b, 0x1a, 0x3c, 0xf4, 0x83, 0x5f, 0x71,
	0xaa, 0x13, 0x27, 0x8a, 0x6a, 0x0d, 0xf9, 0x15, 0x20, 0x2d, 0xb8, 0xd4, 0xcc, 0x25, 0x01, 0x5d,
	0x79, 0xc8, 0xa1, 0x8d, 0xe4, 0x2b, 0x99, 0x61, 0x80, 0xe5, 0x1c, 0x83, 0xab, 0xec, 0x57, 0x36,
	0x75, 0x6b, 0x88, 0x1e, 0xe4, 0xc1, 0x00, 0xf3, 0xb2, 0xcb, 0x5a, 0x85, 0x66, 0x7f, 0x40, 0x33,
	0x73, 0xfd, 0x58, 0xc5, 0x6d, 0x7a, 0xf1, 0xad, 0xdb, 0xf4, 0xce, 0xa4, 0x22, 0x46, 0xa5, 0x5b,
	0x8d, 0x54, 0x96, 0xbc, 0xa3, 0x04, 0xbb, 0x93, 0x61, 0xf4, 0x31, 0x0e, 0x73, 0xa5, 0x2e, 0xfd,
	0x39, 0x8e, 0x30, 0xf2, 0x13, 0xac, 0x09, 0x8d, 0xf4, 0xac, 0xb2, 0xe2, 0x96, 0xd9, 0x82, 0x27,
	0x52, 0xf7, 0xe8, 0x13, 0x14, 0x10, 0x5f, 0x7b, 0xe7, 0x4a, 0x57, 0xbe, 0x42, 0x7e, 0x80, 0xc5,
	0x51, 0x32, 0xfa, 0xa8, 0xa6, 0x98, 0x0d, 0x0b, 0x35, 0xe8, 0x93, 0x78, 0x17, 0x96, 0x92, 0xbc,
	0x2a, 0x85, 0x65, 0xb6, 0x5f, 0x08, 0x9e, 0x96, 0x74, 0x03, 0x0d, 0x17, 0x3d, 0x7a, 0xe5, 0xc1,
	0x9d, 0x2d, 0x68, 0x4f, 0x56, 0x7e, 0x16, 0xa6, 0x06, 0xbf, 0x04, 0x8f, 0xf0, 0xef, 0xeb, 0xa0,
	0xb1, 0xf3, 0x23, 0x34, 0xdd, 0xc5, 0x48, 0x1b, 0x66, 0xf0, 0xeb, 0xe1, 0xb3, 0xef, 0x9a, 0x2b,
	0x99, 0x72, 0x2b, 0x82, 0x86, 0x7b, 0x8a, 0xb4, 0xb4, 0xef, 0xb5, 0x1a, 0x06, 0x53, 0xa7, 0xcd,
	0xd6, 0x42, 0xb0, 0x78, 0xda, 0x6c, 0xad, 0x05, 0xeb, 0xa7, 0xcd, 0xd6, 0x7a, 0xf0, 0x38, 0x6e,
	0x67, 0xfc, 0x16, 0x33, 0xa5, 0x8c, 0x03, 0xf7, 0xaf, 0xe9, 0xdc, 0x30, 0xcd, 0x33, 0xc1, 0x94,
	0xd0, 0x6f, 0x5e, 0xc3, 0x33, 0x69, 0x7c, 0x83, 0xf3, 0xc2, 0xdc, 0x0e, 0xef, 0xeb, 0xf5, 0x9b,
	0xe5, 0xc9, 0x67, 0xeb, 0xc2, 0x0d, 0xf0, 0xa2, 0xd1, 0x99, 0xc5, 0x49, 0xbe, 0xfa, 0x1c, 0x00,
	0x00, 0xff, 0xff, 0xb5, 0x02, 0x7b, 0xe7, 0x7b, 0x07, 0x00, 0x00,
}