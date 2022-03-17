// Copyright(c) 2017-2018 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: config/netcmn.proto

package config

import (
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

type ProxyProto int32

const (
	ProxyProto_PROXY_HTTP  ProxyProto = 0
	ProxyProto_PROXY_HTTPS ProxyProto = 1
	ProxyProto_PROXY_SOCKS ProxyProto = 2
	ProxyProto_PROXY_FTP   ProxyProto = 3
	ProxyProto_PROXY_OTHER ProxyProto = 255
)

// Enum value maps for ProxyProto.
var (
	ProxyProto_name = map[int32]string{
		0:   "PROXY_HTTP",
		1:   "PROXY_HTTPS",
		2:   "PROXY_SOCKS",
		3:   "PROXY_FTP",
		255: "PROXY_OTHER",
	}
	ProxyProto_value = map[string]int32{
		"PROXY_HTTP":  0,
		"PROXY_HTTPS": 1,
		"PROXY_SOCKS": 2,
		"PROXY_FTP":   3,
		"PROXY_OTHER": 255,
	}
)

func (x ProxyProto) Enum() *ProxyProto {
	p := new(ProxyProto)
	*p = x
	return p
}

func (x ProxyProto) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProxyProto) Descriptor() protoreflect.EnumDescriptor {
	return file_config_netcmn_proto_enumTypes[0].Descriptor()
}

func (ProxyProto) Type() protoreflect.EnumType {
	return &file_config_netcmn_proto_enumTypes[0]
}

func (x ProxyProto) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProxyProto.Descriptor instead.
func (ProxyProto) EnumDescriptor() ([]byte, []int) {
	return file_config_netcmn_proto_rawDescGZIP(), []int{0}
}

type DHCPType int32

const (
	DHCPType_DHCPNoop DHCPType = 0
	// Statically configure the DHCP for port
	DHCPType_Static DHCPType = 1
	// Don't run any DHCP, we are in passthrough mode for app
	DHCPType_DHCPNone DHCPType = 2
	// Run the DHCP client on this port
	DHCPType_Client DHCPType = 4
)

// Enum value maps for DHCPType.
var (
	DHCPType_name = map[int32]string{
		0: "DHCPNoop",
		1: "Static",
		2: "DHCPNone",
		4: "Client",
	}
	DHCPType_value = map[string]int32{
		"DHCPNoop": 0,
		"Static":   1,
		"DHCPNone": 2,
		"Client":   4,
	}
)

func (x DHCPType) Enum() *DHCPType {
	p := new(DHCPType)
	*p = x
	return p
}

func (x DHCPType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DHCPType) Descriptor() protoreflect.EnumDescriptor {
	return file_config_netcmn_proto_enumTypes[1].Descriptor()
}

func (DHCPType) Type() protoreflect.EnumType {
	return &file_config_netcmn_proto_enumTypes[1]
}

func (x DHCPType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DHCPType.Descriptor instead.
func (DHCPType) EnumDescriptor() ([]byte, []int) {
	return file_config_netcmn_proto_rawDescGZIP(), []int{1}
}

type NetworkType int32

const (
	NetworkType_NETWORKTYPENOOP NetworkType = 0
	NetworkType_V4              NetworkType = 4 // Legacy - interpreted same as Dual stack
	NetworkType_V6              NetworkType = 6 // Legacy - interpreted same as Dual stack
	NetworkType_CryptoV4        NetworkType = 24
	NetworkType_CryptoV6        NetworkType = 26
	NetworkType_CryptoEID       NetworkType = 14
	NetworkType_V4Only          NetworkType = 7
	NetworkType_V6Only          NetworkType = 8
	NetworkType_DualV4V6        NetworkType = 9
)

// Enum value maps for NetworkType.
var (
	NetworkType_name = map[int32]string{
		0:  "NETWORKTYPENOOP",
		4:  "V4",
		6:  "V6",
		24: "CryptoV4",
		26: "CryptoV6",
		14: "CryptoEID",
		7:  "V4Only",
		8:  "V6Only",
		9:  "DualV4V6",
	}
	NetworkType_value = map[string]int32{
		"NETWORKTYPENOOP": 0,
		"V4":              4,
		"V6":              6,
		"CryptoV4":        24,
		"CryptoV6":        26,
		"CryptoEID":       14,
		"V4Only":          7,
		"V6Only":          8,
		"DualV4V6":        9,
	}
)

func (x NetworkType) Enum() *NetworkType {
	p := new(NetworkType)
	*p = x
	return p
}

func (x NetworkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NetworkType) Descriptor() protoreflect.EnumDescriptor {
	return file_config_netcmn_proto_enumTypes[2].Descriptor()
}

func (NetworkType) Type() protoreflect.EnumType {
	return &file_config_netcmn_proto_enumTypes[2]
}

func (x NetworkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NetworkType.Descriptor instead.
func (NetworkType) EnumDescriptor() ([]byte, []int) {
	return file_config_netcmn_proto_rawDescGZIP(), []int{2}
}

type WirelessType int32

const (
	WirelessType_TypeNOOP WirelessType = 0
	WirelessType_WiFi     WirelessType = 1
	WirelessType_Cellular WirelessType = 2
)

// Enum value maps for WirelessType.
var (
	WirelessType_name = map[int32]string{
		0: "TypeNOOP",
		1: "WiFi",
		2: "Cellular",
	}
	WirelessType_value = map[string]int32{
		"TypeNOOP": 0,
		"WiFi":     1,
		"Cellular": 2,
	}
)

func (x WirelessType) Enum() *WirelessType {
	p := new(WirelessType)
	*p = x
	return p
}

func (x WirelessType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WirelessType) Descriptor() protoreflect.EnumDescriptor {
	return file_config_netcmn_proto_enumTypes[3].Descriptor()
}

func (WirelessType) Type() protoreflect.EnumType {
	return &file_config_netcmn_proto_enumTypes[3]
}

func (x WirelessType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WirelessType.Descriptor instead.
func (WirelessType) EnumDescriptor() ([]byte, []int) {
	return file_config_netcmn_proto_rawDescGZIP(), []int{3}
}

type WiFiKeyScheme int32

const (
	WiFiKeyScheme_SchemeNOOP WiFiKeyScheme = 0
	WiFiKeyScheme_WPAPSK     WiFiKeyScheme = 1 // WPA-PSK
	WiFiKeyScheme_WPAEAP     WiFiKeyScheme = 2 // WPA-EAP or WPA2 Enterprise
)

// Enum value maps for WiFiKeyScheme.
var (
	WiFiKeyScheme_name = map[int32]string{
		0: "SchemeNOOP",
		1: "WPAPSK",
		2: "WPAEAP",
	}
	WiFiKeyScheme_value = map[string]int32{
		"SchemeNOOP": 0,
		"WPAPSK":     1,
		"WPAEAP":     2,
	}
)

func (x WiFiKeyScheme) Enum() *WiFiKeyScheme {
	p := new(WiFiKeyScheme)
	*p = x
	return p
}

func (x WiFiKeyScheme) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WiFiKeyScheme) Descriptor() protoreflect.EnumDescriptor {
	return file_config_netcmn_proto_enumTypes[4].Descriptor()
}

func (WiFiKeyScheme) Type() protoreflect.EnumType {
	return &file_config_netcmn_proto_enumTypes[4]
}

func (x WiFiKeyScheme) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WiFiKeyScheme.Descriptor instead.
func (WiFiKeyScheme) EnumDescriptor() ([]byte, []int) {
	return file_config_netcmn_proto_rawDescGZIP(), []int{4}
}

type IpRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start string `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   string `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *IpRange) Reset() {
	*x = IpRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netcmn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpRange) ProtoMessage() {}

func (x *IpRange) ProtoReflect() protoreflect.Message {
	mi := &file_config_netcmn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpRange.ProtoReflect.Descriptor instead.
func (*IpRange) Descriptor() ([]byte, []int) {
	return file_config_netcmn_proto_rawDescGZIP(), []int{0}
}

func (x *IpRange) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *IpRange) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type ProxyServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proto  ProxyProto `protobuf:"varint,1,opt,name=proto,proto3,enum=org.lfedge.eve.config.ProxyProto" json:"proto,omitempty"`
	Server string     `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	Port   uint32     `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ProxyServer) Reset() {
	*x = ProxyServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netcmn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyServer) ProtoMessage() {}

func (x *ProxyServer) ProtoReflect() protoreflect.Message {
	mi := &file_config_netcmn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyServer.ProtoReflect.Descriptor instead.
func (*ProxyServer) Descriptor() ([]byte, []int) {
	return file_config_netcmn_proto_rawDescGZIP(), []int{1}
}

func (x *ProxyServer) GetProto() ProxyProto {
	if x != nil {
		return x.Proto
	}
	return ProxyProto_PROXY_HTTP
}

func (x *ProxyServer) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *ProxyServer) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type ProxyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// enable network level proxy in the form of WPAD
	NetworkProxyEnable bool `protobuf:"varint,1,opt,name=networkProxyEnable,proto3" json:"networkProxyEnable,omitempty"`
	// dedicated per protocol information
	Proxies []*ProxyServer `protobuf:"bytes,2,rep,name=proxies,proto3" json:"proxies,omitempty"`
	// exceptions separated by commas
	Exceptions string `protobuf:"bytes,3,opt,name=exceptions,proto3" json:"exceptions,omitempty"`
	// or pacfile can be in place of others
	// base64 encoded
	Pacfile string `protobuf:"bytes,4,opt,name=pacfile,proto3" json:"pacfile,omitempty"`
	// Direct URL for wpad.dat download
	NetworkProxyURL string `protobuf:"bytes,5,opt,name=networkProxyURL,proto3" json:"networkProxyURL,omitempty"`
	// Uploaded proxy certificate or certificate chain for MITM
	// this may be needed either in explicit (has ProxyServer items), automatic
	// (networkProxyEnable) or transparent (network layer not aware of proxy)
	ProxyCertPEM [][]byte `protobuf:"bytes,6,rep,name=proxyCertPEM,proto3" json:"proxyCertPEM,omitempty"`
}

func (x *ProxyConfig) Reset() {
	*x = ProxyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netcmn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyConfig) ProtoMessage() {}

func (x *ProxyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_netcmn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyConfig.ProtoReflect.Descriptor instead.
func (*ProxyConfig) Descriptor() ([]byte, []int) {
	return file_config_netcmn_proto_rawDescGZIP(), []int{2}
}

func (x *ProxyConfig) GetNetworkProxyEnable() bool {
	if x != nil {
		return x.NetworkProxyEnable
	}
	return false
}

func (x *ProxyConfig) GetProxies() []*ProxyServer {
	if x != nil {
		return x.Proxies
	}
	return nil
}

func (x *ProxyConfig) GetExceptions() string {
	if x != nil {
		return x.Exceptions
	}
	return ""
}

func (x *ProxyConfig) GetPacfile() string {
	if x != nil {
		return x.Pacfile
	}
	return ""
}

func (x *ProxyConfig) GetNetworkProxyURL() string {
	if x != nil {
		return x.NetworkProxyURL
	}
	return ""
}

func (x *ProxyConfig) GetProxyCertPEM() [][]byte {
	if x != nil {
		return x.ProxyCertPEM
	}
	return nil
}

// deprecated use ZnetStaticDNSEntry
type ZedServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostName string   `protobuf:"bytes,1,opt,name=HostName,proto3" json:"HostName,omitempty"`
	EID      []string `protobuf:"bytes,2,rep,name=EID,proto3" json:"EID,omitempty"`
}

func (x *ZedServer) Reset() {
	*x = ZedServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netcmn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZedServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZedServer) ProtoMessage() {}

func (x *ZedServer) ProtoReflect() protoreflect.Message {
	mi := &file_config_netcmn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZedServer.ProtoReflect.Descriptor instead.
func (*ZedServer) Descriptor() ([]byte, []int) {
	return file_config_netcmn_proto_rawDescGZIP(), []int{3}
}

func (x *ZedServer) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *ZedServer) GetEID() []string {
	if x != nil {
		return x.EID
	}
	return nil
}

// These are list of static mapping that can be added to network
type ZnetStaticDNSEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostName string   `protobuf:"bytes,1,opt,name=HostName,proto3" json:"HostName,omitempty"`
	Address  []string `protobuf:"bytes,2,rep,name=Address,proto3" json:"Address,omitempty"`
}

func (x *ZnetStaticDNSEntry) Reset() {
	*x = ZnetStaticDNSEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netcmn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZnetStaticDNSEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZnetStaticDNSEntry) ProtoMessage() {}

func (x *ZnetStaticDNSEntry) ProtoReflect() protoreflect.Message {
	mi := &file_config_netcmn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZnetStaticDNSEntry.ProtoReflect.Descriptor instead.
func (*ZnetStaticDNSEntry) Descriptor() ([]byte, []int) {
	return file_config_netcmn_proto_rawDescGZIP(), []int{4}
}

func (x *ZnetStaticDNSEntry) GetHostName() string {
	if x != nil {
		return x.HostName
	}
	return ""
}

func (x *ZnetStaticDNSEntry) GetAddress() []string {
	if x != nil {
		return x.Address
	}
	return nil
}

// Common for IPv4 and IPv6
type Ipspec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dhcp DHCPType `protobuf:"varint,2,opt,name=dhcp,proto3,enum=org.lfedge.eve.config.DHCPType" json:"dhcp,omitempty"`
	// subnet is CIDR format...x.y.z.l/nn
	Subnet  string   `protobuf:"bytes,3,opt,name=subnet,proto3" json:"subnet,omitempty"`
	Gateway string   `protobuf:"bytes,5,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Domain  string   `protobuf:"bytes,6,opt,name=domain,proto3" json:"domain,omitempty"`
	Ntp     string   `protobuf:"bytes,7,opt,name=ntp,proto3" json:"ntp,omitempty"`
	Dns     []string `protobuf:"bytes,8,rep,name=dns,proto3" json:"dns,omitempty"`
	// for IPAM management when dhcp is turned on.
	// If none provided, system will default pool.
	DhcpRange *IpRange `protobuf:"bytes,9,opt,name=dhcpRange,proto3" json:"dhcpRange,omitempty"`
}

func (x *Ipspec) Reset() {
	*x = Ipspec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_netcmn_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ipspec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ipspec) ProtoMessage() {}

func (x *Ipspec) ProtoReflect() protoreflect.Message {
	mi := &file_config_netcmn_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ipspec.ProtoReflect.Descriptor instead.
func (*Ipspec) Descriptor() ([]byte, []int) {
	return file_config_netcmn_proto_rawDescGZIP(), []int{5}
}

func (x *Ipspec) GetDhcp() DHCPType {
	if x != nil {
		return x.Dhcp
	}
	return DHCPType_DHCPNoop
}

func (x *Ipspec) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

func (x *Ipspec) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

func (x *Ipspec) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Ipspec) GetNtp() string {
	if x != nil {
		return x.Ntp
	}
	return ""
}

func (x *Ipspec) GetDns() []string {
	if x != nil {
		return x.Dns
	}
	return nil
}

func (x *Ipspec) GetDhcpRange() *IpRange {
	if x != nil {
		return x.DhcpRange
	}
	return nil
}

var File_config_netcmn_proto protoreflect.FileDescriptor

var file_config_netcmn_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6e, 0x65, 0x74, 0x63, 0x6d, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x31, 0x0a, 0x07,
	0x69, 0x70, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22,
	0x72, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x37,
	0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x22, 0x83, 0x02, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x2e, 0x0a, 0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67,
	0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x55, 0x52, 0x4c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x55, 0x52, 0x4c, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x65,
	0x72, 0x74, 0x50, 0x45, 0x4d, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x43, 0x65, 0x72, 0x74, 0x50, 0x45, 0x4d, 0x22, 0x39, 0x0a, 0x09, 0x5a, 0x65, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x49, 0x44, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x45, 0x49, 0x44, 0x22, 0x4a, 0x0a, 0x12, 0x5a, 0x6e, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x44, 0x4e, 0x53, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x6f,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0xe9, 0x01, 0x0a, 0x06, 0x69, 0x70, 0x73, 0x70, 0x65, 0x63, 0x12, 0x33, 0x0a, 0x04, 0x64,
	0x68, 0x63, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x44, 0x48, 0x43, 0x50, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x64, 0x68, 0x63, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x74,
	0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x74, 0x70, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6e, 0x73, 0x12, 0x3c,
	0x0a, 0x09, 0x64, 0x68, 0x63, 0x70, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65,
	0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x69, 0x70, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x09, 0x64, 0x68, 0x63, 0x70, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x2a, 0x5f, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52,
	0x4f, 0x58, 0x59, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x52,
	0x4f, 0x58, 0x59, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x53, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x50,
	0x52, 0x4f, 0x58, 0x59, 0x5f, 0x53, 0x4f, 0x43, 0x4b, 0x53, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09,
	0x50, 0x52, 0x4f, 0x58, 0x59, 0x5f, 0x46, 0x54, 0x50, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0b, 0x50,
	0x52, 0x4f, 0x58, 0x59, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0xff, 0x01, 0x2a, 0x3e, 0x0a,
	0x08, 0x44, 0x48, 0x43, 0x50, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x48, 0x43,
	0x50, 0x4e, 0x6f, 0x6f, 0x70, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x48, 0x43, 0x50, 0x4e, 0x6f, 0x6e, 0x65, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x10, 0x04, 0x2a, 0x83, 0x01,
	0x0a, 0x0b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a,
	0x0f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x54, 0x59, 0x50, 0x45, 0x4e, 0x4f, 0x4f, 0x50,
	0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x34, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x36,
	0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x56, 0x34, 0x10, 0x18,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x56, 0x36, 0x10, 0x1a, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x45, 0x49, 0x44, 0x10, 0x0e, 0x12, 0x0a, 0x0a,
	0x06, 0x56, 0x34, 0x4f, 0x6e, 0x6c, 0x79, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x36, 0x4f,
	0x6e, 0x6c, 0x79, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x75, 0x61, 0x6c, 0x56, 0x34, 0x56,
	0x36, 0x10, 0x09, 0x2a, 0x34, 0x0a, 0x0c, 0x57, 0x69, 0x72, 0x65, 0x6c, 0x65, 0x73, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x4f, 0x4f, 0x50, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x69, 0x46, 0x69, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43,
	0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x10, 0x02, 0x2a, 0x37, 0x0a, 0x0d, 0x57, 0x69, 0x46,
	0x69, 0x4b, 0x65, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x65, 0x4e, 0x4f, 0x4f, 0x50, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x50,
	0x41, 0x50, 0x53, 0x4b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x50, 0x41, 0x45, 0x41, 0x50,
	0x10, 0x02, 0x42, 0x3d, 0x0a, 0x15, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f,
	0x65, 0x76, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_netcmn_proto_rawDescOnce sync.Once
	file_config_netcmn_proto_rawDescData = file_config_netcmn_proto_rawDesc
)

func file_config_netcmn_proto_rawDescGZIP() []byte {
	file_config_netcmn_proto_rawDescOnce.Do(func() {
		file_config_netcmn_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_netcmn_proto_rawDescData)
	})
	return file_config_netcmn_proto_rawDescData
}

var file_config_netcmn_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_config_netcmn_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_config_netcmn_proto_goTypes = []interface{}{
	(ProxyProto)(0),            // 0: org.lfedge.eve.config.proxyProto
	(DHCPType)(0),              // 1: org.lfedge.eve.config.DHCPType
	(NetworkType)(0),           // 2: org.lfedge.eve.config.NetworkType
	(WirelessType)(0),          // 3: org.lfedge.eve.config.WirelessType
	(WiFiKeyScheme)(0),         // 4: org.lfedge.eve.config.WiFiKeyScheme
	(*IpRange)(nil),            // 5: org.lfedge.eve.config.ipRange
	(*ProxyServer)(nil),        // 6: org.lfedge.eve.config.ProxyServer
	(*ProxyConfig)(nil),        // 7: org.lfedge.eve.config.ProxyConfig
	(*ZedServer)(nil),          // 8: org.lfedge.eve.config.ZedServer
	(*ZnetStaticDNSEntry)(nil), // 9: org.lfedge.eve.config.ZnetStaticDNSEntry
	(*Ipspec)(nil),             // 10: org.lfedge.eve.config.ipspec
}
var file_config_netcmn_proto_depIdxs = []int32{
	0, // 0: org.lfedge.eve.config.ProxyServer.proto:type_name -> org.lfedge.eve.config.proxyProto
	6, // 1: org.lfedge.eve.config.ProxyConfig.proxies:type_name -> org.lfedge.eve.config.ProxyServer
	1, // 2: org.lfedge.eve.config.ipspec.dhcp:type_name -> org.lfedge.eve.config.DHCPType
	5, // 3: org.lfedge.eve.config.ipspec.dhcpRange:type_name -> org.lfedge.eve.config.ipRange
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_config_netcmn_proto_init() }
func file_config_netcmn_proto_init() {
	if File_config_netcmn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_netcmn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpRange); i {
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
		file_config_netcmn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyServer); i {
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
		file_config_netcmn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyConfig); i {
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
		file_config_netcmn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZedServer); i {
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
		file_config_netcmn_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZnetStaticDNSEntry); i {
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
		file_config_netcmn_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ipspec); i {
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
			RawDescriptor: file_config_netcmn_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_netcmn_proto_goTypes,
		DependencyIndexes: file_config_netcmn_proto_depIdxs,
		EnumInfos:         file_config_netcmn_proto_enumTypes,
		MessageInfos:      file_config_netcmn_proto_msgTypes,
	}.Build()
	File_config_netcmn_proto = out.File
	file_config_netcmn_proto_rawDesc = nil
	file_config_netcmn_proto_goTypes = nil
	file_config_netcmn_proto_depIdxs = nil
}
