// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DsType int32

const (
	DsType_DsUnknown DsType = 0
	DsType_DsHttp    DsType = 1
	DsType_DsHttps   DsType = 2
	DsType_DsS3      DsType = 3
	DsType_DsSFTP    DsType = 4
)

var DsType_name = map[int32]string{
	0: "DsUnknown",
	1: "DsHttp",
	2: "DsHttps",
	3: "DsS3",
	4: "DsSFTP",
}
var DsType_value = map[string]int32{
	"DsUnknown": 0,
	"DsHttp":    1,
	"DsHttps":   2,
	"DsS3":      3,
	"DsSFTP":    4,
}

func (x DsType) String() string {
	return proto.EnumName(DsType_name, int32(x))
}
func (DsType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type Format int32

const (
	Format_FmtUnknown Format = 0
	Format_RAW        Format = 1
	Format_QCOW       Format = 2
	Format_QCOW2      Format = 3
	Format_VHD        Format = 4
	Format_VMDK       Format = 5
	Format_OVA        Format = 6
	Format_VHDX       Format = 7
)

var Format_name = map[int32]string{
	0: "FmtUnknown",
	1: "RAW",
	2: "QCOW",
	3: "QCOW2",
	4: "VHD",
	5: "VMDK",
	6: "OVA",
	7: "VHDX",
}
var Format_value = map[string]int32{
	"FmtUnknown": 0,
	"RAW":        1,
	"QCOW":       2,
	"QCOW2":      3,
	"VHD":        4,
	"VMDK":       5,
	"OVA":        6,
	"VHDX":       7,
}

func (x Format) String() string {
	return proto.EnumName(Format_name, int32(x))
}
func (Format) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

type Target int32

const (
	Target_TgtUnknown Target = 0
	Target_Disk       Target = 1
	Target_Kernel     Target = 2
	Target_Initrd     Target = 3
	Target_RamDisk    Target = 4
)

var Target_name = map[int32]string{
	0: "TgtUnknown",
	1: "Disk",
	2: "Kernel",
	3: "Initrd",
	4: "RamDisk",
}
var Target_value = map[string]int32{
	"TgtUnknown": 0,
	"Disk":       1,
	"Kernel":     2,
	"Initrd":     3,
	"RamDisk":    4,
}

func (x Target) String() string {
	return proto.EnumName(Target_name, int32(x))
}
func (Target) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

type DriveType int32

const (
	DriveType_Unclassified DriveType = 0
	DriveType_CDROM        DriveType = 1
	DriveType_HDD          DriveType = 2
	DriveType_NET          DriveType = 3
	// this type is allocate the empty disk of maxsizebytes specified
	DriveType_HDD_EMPTY DriveType = 4
)

var DriveType_name = map[int32]string{
	0: "Unclassified",
	1: "CDROM",
	2: "HDD",
	3: "NET",
	4: "HDD_EMPTY",
}
var DriveType_value = map[string]int32{
	"Unclassified": 0,
	"CDROM":        1,
	"HDD":          2,
	"NET":          3,
	"HDD_EMPTY":    4,
}

func (x DriveType) String() string {
	return proto.EnumName(DriveType_name, int32(x))
}
func (DriveType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

type SignatureInfo struct {
	Intercertsurl string `protobuf:"bytes,1,opt,name=intercertsurl" json:"intercertsurl,omitempty"`
	Signercerturl string `protobuf:"bytes,2,opt,name=signercerturl" json:"signercerturl,omitempty"`
	Signature     []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignatureInfo) Reset()                    { *m = SignatureInfo{} }
func (m *SignatureInfo) String() string            { return proto.CompactTextString(m) }
func (*SignatureInfo) ProtoMessage()               {}
func (*SignatureInfo) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *SignatureInfo) GetIntercertsurl() string {
	if m != nil {
		return m.Intercertsurl
	}
	return ""
}

func (m *SignatureInfo) GetSignercerturl() string {
	if m != nil {
		return m.Signercerturl
	}
	return ""
}

func (m *SignatureInfo) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type DatastoreConfig struct {
	Id       string `protobuf:"bytes,100,opt,name=id" json:"id,omitempty"`
	DType    DsType `protobuf:"varint,1,opt,name=dType,enum=DsType" json:"dType,omitempty"`
	Fqdn     string `protobuf:"bytes,2,opt,name=fqdn" json:"fqdn,omitempty"`
	ApiKey   string `protobuf:"bytes,3,opt,name=apiKey" json:"apiKey,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	// depending on datastore types, it could be bucket or path
	Dpath string `protobuf:"bytes,5,opt,name=dpath" json:"dpath,omitempty"`
}

func (m *DatastoreConfig) Reset()                    { *m = DatastoreConfig{} }
func (m *DatastoreConfig) String() string            { return proto.CompactTextString(m) }
func (*DatastoreConfig) ProtoMessage()               {}
func (*DatastoreConfig) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *DatastoreConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DatastoreConfig) GetDType() DsType {
	if m != nil {
		return m.DType
	}
	return DsType_DsUnknown
}

func (m *DatastoreConfig) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *DatastoreConfig) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *DatastoreConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *DatastoreConfig) GetDpath() string {
	if m != nil {
		return m.Dpath
	}
	return ""
}

type Image struct {
	Uuidandversion *UUIDandVersion `protobuf:"bytes,1,opt,name=uuidandversion" json:"uuidandversion,omitempty"`
	// it could be relative path/name as well
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Sha256  string `protobuf:"bytes,3,opt,name=sha256" json:"sha256,omitempty"`
	Iformat Format `protobuf:"varint,4,opt,name=iformat,enum=Format" json:"iformat,omitempty"`
	// if its signed image
	Siginfo   *SignatureInfo `protobuf:"bytes,5,opt,name=siginfo" json:"siginfo,omitempty"`
	DsId      string         `protobuf:"bytes,6,opt,name=dsId" json:"dsId,omitempty"`
	SizeBytes int64          `protobuf:"varint,8,opt,name=sizeBytes" json:"sizeBytes,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *Image) GetUuidandversion() *UUIDandVersion {
	if m != nil {
		return m.Uuidandversion
	}
	return nil
}

func (m *Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Image) GetSha256() string {
	if m != nil {
		return m.Sha256
	}
	return ""
}

func (m *Image) GetIformat() Format {
	if m != nil {
		return m.Iformat
	}
	return Format_FmtUnknown
}

func (m *Image) GetSiginfo() *SignatureInfo {
	if m != nil {
		return m.Siginfo
	}
	return nil
}

func (m *Image) GetDsId() string {
	if m != nil {
		return m.DsId
	}
	return ""
}

func (m *Image) GetSizeBytes() int64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

type Drive struct {
	Image        *Image    `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Readonly     bool      `protobuf:"varint,5,opt,name=readonly" json:"readonly,omitempty"`
	Preserve     bool      `protobuf:"varint,6,opt,name=preserve" json:"preserve,omitempty"`
	Drvtype      DriveType `protobuf:"varint,8,opt,name=drvtype,enum=DriveType" json:"drvtype,omitempty"`
	Target       Target    `protobuf:"varint,9,opt,name=target,enum=Target" json:"target,omitempty"`
	Maxsizebytes int64     `protobuf:"varint,10,opt,name=maxsizebytes" json:"maxsizebytes,omitempty"`
}

func (m *Drive) Reset()                    { *m = Drive{} }
func (m *Drive) String() string            { return proto.CompactTextString(m) }
func (*Drive) ProtoMessage()               {}
func (*Drive) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *Drive) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Drive) GetReadonly() bool {
	if m != nil {
		return m.Readonly
	}
	return false
}

func (m *Drive) GetPreserve() bool {
	if m != nil {
		return m.Preserve
	}
	return false
}

func (m *Drive) GetDrvtype() DriveType {
	if m != nil {
		return m.Drvtype
	}
	return DriveType_Unclassified
}

func (m *Drive) GetTarget() Target {
	if m != nil {
		return m.Target
	}
	return Target_TgtUnknown
}

func (m *Drive) GetMaxsizebytes() int64 {
	if m != nil {
		return m.Maxsizebytes
	}
	return 0
}

func init() {
	proto.RegisterType((*SignatureInfo)(nil), "SignatureInfo")
	proto.RegisterType((*DatastoreConfig)(nil), "DatastoreConfig")
	proto.RegisterType((*Image)(nil), "Image")
	proto.RegisterType((*Drive)(nil), "Drive")
	proto.RegisterEnum("DsType", DsType_name, DsType_value)
	proto.RegisterEnum("Format", Format_name, Format_value)
	proto.RegisterEnum("Target", Target_name, Target_value)
	proto.RegisterEnum("DriveType", DriveType_name, DriveType_value)
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 698 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x54, 0xd1, 0x6e, 0xea, 0x46,
	0x10, 0xbd, 0x06, 0x6c, 0x60, 0x12, 0xc8, 0x6a, 0x55, 0x55, 0xd6, 0xd5, 0xbd, 0x4a, 0x8a, 0xf2,
	0x80, 0x78, 0x70, 0x24, 0xa2, 0xb6, 0xaf, 0x4d, 0xe2, 0x50, 0x50, 0x94, 0x26, 0x75, 0x80, 0xb4,
	0x55, 0xa5, 0x6a, 0xe3, 0x5d, 0x9c, 0x55, 0xf0, 0x2e, 0xdd, 0x5d, 0x48, 0xc9, 0xc7, 0xf4, 0x7b,
	0xfa, 0x2f, 0xfd, 0x89, 0x6a, 0x77, 0x81, 0x94, 0xfb, 0x36, 0xe7, 0xcc, 0xf1, 0xcc, 0x99, 0x19,
	0xdb, 0xd0, 0xd2, 0x46, 0x2a, 0x52, 0xb0, 0x64, 0xa1, 0xa4, 0x91, 0x1f, 0x8f, 0x28, 0x5b, 0xe5,
	0xb2, 0x2c, 0xa5, 0xf0, 0x44, 0x67, 0x0d, 0xad, 0x07, 0x5e, 0x08, 0x62, 0x96, 0x8a, 0x8d, 0xc4,
	0x4c, 0xe2, 0x53, 0x68, 0x71, 0x61, 0x98, 0xca, 0x99, 0x32, 0x7a, 0xa9, 0xe6, 0x71, 0x70, 0x12,
	0x74, 0x9b, 0xd9, 0x3e, 0x69, 0x55, 0x9a, 0x17, 0xc2, 0x33, 0x56, 0x55, 0xf1, 0xaa, 0x3d, 0x12,
	0x7f, 0x82, 0xa6, 0xde, 0x16, 0x8f, 0xab, 0x27, 0x41, 0xf7, 0x30, 0x7b, 0x27, 0x3a, 0x7f, 0x07,
	0x70, 0x94, 0x12, 0x43, 0xac, 0x43, 0x76, 0x25, 0xc5, 0x8c, 0x17, 0xb8, 0x0d, 0x15, 0x4e, 0x63,
	0xea, 0x8a, 0x55, 0x38, 0xc5, 0x9f, 0x21, 0xa4, 0xe3, 0xf5, 0x82, 0x39, 0x17, 0xed, 0x7e, 0x3d,
	0x49, 0xb5, 0x85, 0x99, 0x67, 0x31, 0x86, 0xda, 0xec, 0x4f, 0x2a, 0x36, 0xdd, 0x5d, 0x8c, 0xbf,
	0x86, 0x88, 0x2c, 0xf8, 0x0d, 0x5b, 0xbb, 0x8e, 0xcd, 0x6c, 0x83, 0xf0, 0x47, 0x68, 0x2c, 0x88,
	0xd6, 0xaf, 0x52, 0xd1, 0xb8, 0xe6, 0x32, 0x3b, 0x8c, 0xbf, 0x82, 0x90, 0x2e, 0x88, 0x79, 0x8e,
	0x43, 0x97, 0xf0, 0xa0, 0xf3, 0x6f, 0x00, 0xe1, 0xa8, 0x24, 0x05, 0xc3, 0xdf, 0x43, 0x7b, 0xb9,
	0xe4, 0x94, 0x08, 0xba, 0x62, 0x4a, 0x73, 0x29, 0x9c, 0x9f, 0x83, 0xfe, 0x51, 0x32, 0x99, 0x8c,
	0x52, 0x22, 0xe8, 0xd4, 0xd3, 0xd9, 0x17, 0x32, 0x6b, 0x50, 0x90, 0x92, 0x6d, 0x0d, 0xda, 0xd8,
	0x1a, 0xd4, 0xcf, 0xa4, 0xff, 0xed, 0x77, 0x5b, 0x83, 0x1e, 0xe1, 0x6f, 0xa0, 0xce, 0x67, 0x52,
	0x95, 0xc4, 0x38, 0x7f, 0x76, 0xda, 0x81, 0x83, 0xd9, 0x96, 0xc7, 0x5d, 0xa8, 0x6b, 0x5e, 0x70,
	0x31, 0x93, 0xce, 0xe9, 0x41, 0xbf, 0x9d, 0xec, 0x5d, 0x2f, 0xdb, 0xa6, 0x6d, 0x63, 0xaa, 0x47,
	0x34, 0x8e, 0x7c, 0x63, 0x1b, 0xfb, 0x73, 0xbc, 0xb1, 0xcb, 0xb5, 0x61, 0x3a, 0x6e, 0x9c, 0x04,
	0xdd, 0x6a, 0xf6, 0x4e, 0x74, 0xfe, 0x09, 0x20, 0x4c, 0x15, 0x5f, 0x31, 0xfc, 0x09, 0x42, 0x6e,
	0xc7, 0xde, 0x0c, 0x19, 0x25, 0x6e, 0x09, 0x99, 0x27, 0xed, 0x1e, 0x15, 0x23, 0x54, 0x8a, 0xf9,
	0xda, 0x99, 0x68, 0x64, 0x3b, 0xec, 0x76, 0xac, 0x98, 0x66, 0x6a, 0xc5, 0x5c, 0xe7, 0x46, 0xb6,
	0xc3, 0xf8, 0x14, 0xea, 0x54, 0xad, 0x8c, 0x3d, 0x66, 0xc3, 0x8d, 0x07, 0x89, 0x6b, 0xe7, 0xee,
	0xb9, 0x4d, 0xe1, 0x63, 0x88, 0x0c, 0x51, 0x05, 0x33, 0x71, 0x73, 0xb3, 0x83, 0xb1, 0x83, 0xd9,
	0x86, 0xc6, 0x1d, 0x38, 0x2c, 0xc9, 0x5f, 0xd6, 0xf6, 0x93, 0x9b, 0x03, 0xdc, 0x1c, 0x7b, 0x5c,
	0x6f, 0x00, 0x91, 0x7f, 0x4f, 0x70, 0x0b, 0x9a, 0xa9, 0x9e, 0x88, 0x17, 0x21, 0x5f, 0x05, 0xfa,
	0x80, 0xc1, 0x26, 0x86, 0xc6, 0x2c, 0x50, 0x80, 0x0f, 0xa0, 0xee, 0x63, 0x8d, 0x2a, 0xb8, 0x01,
	0xb5, 0x54, 0x3f, 0x9c, 0xa3, 0xaa, 0x97, 0x3c, 0x0c, 0xc6, 0xf7, 0xa8, 0xd6, 0xfb, 0x1d, 0x22,
	0x7f, 0x01, 0xdc, 0x06, 0x18, 0x94, 0xe6, 0xbd, 0x50, 0x1d, 0xaa, 0xd9, 0xc5, 0x23, 0x0a, 0xec,
	0x83, 0x3f, 0x5f, 0xdd, 0x3d, 0xa2, 0x0a, 0x6e, 0x42, 0x68, 0xa3, 0x3e, 0xaa, 0xda, 0xec, 0x74,
	0x98, 0xa2, 0x9a, 0xcd, 0x4e, 0x6f, 0xd3, 0x1b, 0x14, 0x5a, 0xea, 0x6e, 0x7a, 0x81, 0x22, 0x47,
	0x0d, 0xd3, 0x5f, 0x50, 0xbd, 0xf7, 0x23, 0x44, 0x7e, 0x36, 0x5b, 0x7d, 0x5c, 0xfc, 0xaf, 0xba,
	0x75, 0xc3, 0xf5, 0x0b, 0x0a, 0xac, 0x9b, 0x1b, 0xa6, 0x04, 0x9b, 0xa3, 0x8a, 0x8d, 0x47, 0x82,
	0x1b, 0x45, 0x51, 0xd5, 0x9a, 0xcf, 0x48, 0xe9, 0x44, 0xb5, 0xde, 0x08, 0x9a, 0xbb, 0x4d, 0x62,
	0x04, 0x87, 0x13, 0x91, 0xcf, 0x89, 0xd6, 0x7c, 0xc6, 0x19, 0x45, 0x1f, 0xac, 0xb1, 0xab, 0x34,
	0xbb, 0xbb, 0x45, 0x81, 0x75, 0x31, 0x4c, 0x53, 0x54, 0xb1, 0xc1, 0x4f, 0xd7, 0x63, 0x54, 0xb5,
	0x0b, 0x1a, 0xa6, 0xe9, 0x1f, 0xd7, 0xb7, 0xf7, 0xe3, 0x5f, 0x51, 0xed, 0xf2, 0x07, 0x38, 0xce,
	0x65, 0x99, 0xbc, 0x31, 0xca, 0x28, 0x49, 0xf2, 0xb9, 0x5c, 0xd2, 0x64, 0x69, 0xef, 0xc7, 0xf3,
	0xcd, 0x2f, 0xe4, 0xb7, 0xcf, 0x05, 0x37, 0xcf, 0xcb, 0xa7, 0x24, 0x97, 0xe5, 0x99, 0xd7, 0x9d,
	0x91, 0x05, 0x3f, 0x7b, 0xcb, 0xdd, 0x17, 0xfc, 0x14, 0x39, 0xd5, 0xf9, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x89, 0x48, 0x73, 0xce, 0x79, 0x04, 0x00, 0x00,
}
