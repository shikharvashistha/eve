// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vm.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VmConfig struct {
	Kernel     string   `protobuf:"bytes,1,opt,name=kernel" json:"kernel,omitempty"`
	Ramdisk    string   `protobuf:"bytes,2,opt,name=ramdisk" json:"ramdisk,omitempty"`
	Memory     uint32   `protobuf:"varint,3,opt,name=memory" json:"memory,omitempty"`
	Maxmem     uint32   `protobuf:"varint,4,opt,name=maxmem" json:"maxmem,omitempty"`
	Vcpus      uint32   `protobuf:"varint,5,opt,name=vcpus" json:"vcpus,omitempty"`
	Maxcpus    uint32   `protobuf:"varint,6,opt,name=maxcpus" json:"maxcpus,omitempty"`
	Rootdev    string   `protobuf:"bytes,7,opt,name=rootdev" json:"rootdev,omitempty"`
	Extraargs  string   `protobuf:"bytes,8,opt,name=extraargs" json:"extraargs,omitempty"`
	Bootloader string   `protobuf:"bytes,9,opt,name=bootloader" json:"bootloader,omitempty"`
	Cpus       string   `protobuf:"bytes,10,opt,name=cpus" json:"cpus,omitempty"`
	Devicetree string   `protobuf:"bytes,11,opt,name=devicetree" json:"devicetree,omitempty"`
	Dtdev      []string `protobuf:"bytes,12,rep,name=dtdev" json:"dtdev,omitempty"`
	Irqs       []uint32 `protobuf:"varint,13,rep,packed,name=irqs" json:"irqs,omitempty"`
	Iomem      []string `protobuf:"bytes,14,rep,name=iomem" json:"iomem,omitempty"`
}

func (m *VmConfig) Reset()                    { *m = VmConfig{} }
func (m *VmConfig) String() string            { return proto.CompactTextString(m) }
func (*VmConfig) ProtoMessage()               {}
func (*VmConfig) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *VmConfig) GetKernel() string {
	if m != nil {
		return m.Kernel
	}
	return ""
}

func (m *VmConfig) GetRamdisk() string {
	if m != nil {
		return m.Ramdisk
	}
	return ""
}

func (m *VmConfig) GetMemory() uint32 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *VmConfig) GetMaxmem() uint32 {
	if m != nil {
		return m.Maxmem
	}
	return 0
}

func (m *VmConfig) GetVcpus() uint32 {
	if m != nil {
		return m.Vcpus
	}
	return 0
}

func (m *VmConfig) GetMaxcpus() uint32 {
	if m != nil {
		return m.Maxcpus
	}
	return 0
}

func (m *VmConfig) GetRootdev() string {
	if m != nil {
		return m.Rootdev
	}
	return ""
}

func (m *VmConfig) GetExtraargs() string {
	if m != nil {
		return m.Extraargs
	}
	return ""
}

func (m *VmConfig) GetBootloader() string {
	if m != nil {
		return m.Bootloader
	}
	return ""
}

func (m *VmConfig) GetCpus() string {
	if m != nil {
		return m.Cpus
	}
	return ""
}

func (m *VmConfig) GetDevicetree() string {
	if m != nil {
		return m.Devicetree
	}
	return ""
}

func (m *VmConfig) GetDtdev() []string {
	if m != nil {
		return m.Dtdev
	}
	return nil
}

func (m *VmConfig) GetIrqs() []uint32 {
	if m != nil {
		return m.Irqs
	}
	return nil
}

func (m *VmConfig) GetIomem() []string {
	if m != nil {
		return m.Iomem
	}
	return nil
}

func init() {
	proto.RegisterType((*VmConfig)(nil), "VmConfig")
}

func init() { proto.RegisterFile("vm.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x91, 0x3d, 0x4e, 0xf4, 0x30,
	0x10, 0x86, 0xb5, 0xff, 0x1b, 0x7f, 0x5f, 0x28, 0x2c, 0x84, 0x5c, 0xf0, 0x13, 0x51, 0xa5, 0x4a,
	0x0a, 0x2e, 0x80, 0xe0, 0x06, 0x29, 0x28, 0xe8, 0x1c, 0x7b, 0x08, 0xd6, 0xc6, 0x3b, 0xc1, 0x71,
	0xa2, 0xb0, 0x27, 0xe6, 0x18, 0xc8, 0xe3, 0xac, 0x96, 0xce, 0xcf, 0x33, 0x33, 0x7e, 0x6d, 0x0d,
	0xdb, 0x8f, 0xb6, 0xe8, 0x1c, 0x7a, 0x7c, 0xfc, 0x59, 0xb2, 0xfd, 0x9b, 0x7d, 0xc5, 0xe3, 0x87,
	0x69, 0xf8, 0x0d, 0xdb, 0x1e, 0xc0, 0x1d, 0xa1, 0x15, 0x8b, 0x6c, 0x91, 0x27, 0xd5, 0x4c, 0x5c,
	0xb0, 0x9d, 0x93, 0x56, 0x9b, 0xfe, 0x20, 0x96, 0x54, 0x38, 0x63, 0x98, 0xb0, 0x60, 0xd1, 0x7d,
	0x8b, 0x55, 0xb6, 0xc8, 0xd3, 0x6a, 0x26, 0xf2, 0x72, 0xb2, 0x60, 0xc5, 0x7a, 0xf6, 0x44, 0xfc,
	0x9a, 0x6d, 0x46, 0xd5, 0x0d, 0xbd, 0xd8, 0x90, 0x8e, 0x10, 0xee, 0xb7, 0x72, 0x22, 0xbf, 0x25,
	0x7f, 0x46, 0x4a, 0x46, 0xf4, 0x1a, 0x46, 0xb1, 0x9b, 0x93, 0x23, 0xf2, 0x5b, 0x96, 0xc0, 0xe4,
	0x9d, 0x94, 0xae, 0xe9, 0xc5, 0x9e, 0x6a, 0x17, 0xc1, 0xef, 0x19, 0xab, 0x11, 0x7d, 0x8b, 0x52,
	0x83, 0x13, 0x09, 0x95, 0xff, 0x18, 0xce, 0xd9, 0x9a, 0xe2, 0x18, 0x55, 0xe8, 0x1c, 0x66, 0x34,
	0x8c, 0x46, 0x81, 0x77, 0x00, 0xe2, 0x5f, 0x9c, 0xb9, 0x98, 0xf0, 0x76, 0x4d, 0x2f, 0xf9, 0x9f,
	0xad, 0xf2, 0xa4, 0x8a, 0x10, 0x6e, 0x32, 0xee, 0xab, 0x17, 0x69, 0xb6, 0xca, 0xd3, 0x8a, 0xce,
	0xa1, 0xd3, 0x60, 0xf8, 0xfc, 0x55, 0xec, 0x24, 0x78, 0x79, 0x66, 0x0f, 0x0a, 0x6d, 0x71, 0x02,
	0x0d, 0x5a, 0x16, 0xaa, 0xc5, 0x41, 0x17, 0x43, 0x0f, 0x2e, 0x04, 0xc4, 0x6d, 0xbc, 0xdf, 0x35,
	0xc6, 0x7f, 0x0e, 0x75, 0xa1, 0xd0, 0x96, 0xb1, 0xaf, 0x94, 0x9d, 0x29, 0x4f, 0x8a, 0xf6, 0x53,
	0x6f, 0xa9, 0xeb, 0xe9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x79, 0x6e, 0xfe, 0xbf, 0x01, 0x00,
	0x00,
}
