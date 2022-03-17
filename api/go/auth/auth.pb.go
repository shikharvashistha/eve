// Copyright(c) 2020 Zededa, Inc.
// All rights reserved.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: auth/auth.proto

package auth

import (
	evecommon "github.com/lf-edge/eve/api/go/evecommon"
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

type AuthBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *AuthBody) Reset() {
	*x = AuthBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthBody) ProtoMessage() {}

func (x *AuthBody) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthBody.ProtoReflect.Descriptor instead.
func (*AuthBody) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthBody) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type AuthContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtectedPayload *AuthBody `protobuf:"bytes,1,opt,name=protectedPayload,proto3" json:"protectedPayload,omitempty"` // envolope body, a marshalled protobuf data or it can be null
	// if the length of senderCertHash received is not N bytes, as described in hashAlgorithm, then the protobuf
	// message either is not AuthContainer type, or is corrupted. Otherwise, the
	// reciever may not have the sender's signing certificate
	Algo           evecommon.HashAlgorithm `protobuf:"varint,2,opt,name=algo,proto3,enum=org.lfedge.eve.common.HashAlgorithm" json:"algo,omitempty"` // hash algorithm used by sender Cert
	SenderCertHash []byte                  `protobuf:"bytes,3,opt,name=senderCertHash,proto3" json:"senderCertHash,omitempty"`                       // N bytes in length, 1st N bytes of sender siging cert sha256 hash
	SignatureHash  []byte                  `protobuf:"bytes,4,opt,name=signatureHash,proto3" json:"signatureHash,omitempty"`                         // signature of the sha256 hash of the payload
	SenderCert     []byte                  `protobuf:"bytes,5,opt,name=senderCert,proto3" json:"senderCert,omitempty"`                               // full senderCert needed for some payloads
}

func (x *AuthContainer) Reset() {
	*x = AuthContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthContainer) ProtoMessage() {}

func (x *AuthContainer) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthContainer.ProtoReflect.Descriptor instead.
func (*AuthContainer) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthContainer) GetProtectedPayload() *AuthBody {
	if x != nil {
		return x.ProtectedPayload
	}
	return nil
}

func (x *AuthContainer) GetAlgo() evecommon.HashAlgorithm {
	if x != nil {
		return x.Algo
	}
	return evecommon.HashAlgorithm_HASH_ALGORITHM_INVALID
}

func (x *AuthContainer) GetSenderCertHash() []byte {
	if x != nil {
		return x.SenderCertHash
	}
	return nil
}

func (x *AuthContainer) GetSignatureHash() []byte {
	if x != nil {
		return x.SignatureHash
	}
	return nil
}

func (x *AuthContainer) GetSenderCert() []byte {
	if x != nil {
		return x.SenderCert
	}
	return nil
}

var File_auth_auth_proto protoreflect.FileDescriptor

var file_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x19, 0x65, 0x76, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x65, 0x76, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x24, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x82, 0x02, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x10, 0x70, 0x72, 0x6f,
	0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65,
	0x2e, 0x65, 0x76, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x42, 0x6f,
	0x64, 0x79, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x38, 0x0a, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e,
	0x65, 0x76, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x41,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x04, 0x61, 0x6c, 0x67, 0x6f, 0x12, 0x26,
	0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x65,
	0x72, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x42, 0x39, 0x0a, 0x13,
	0x6f, 0x72, 0x67, 0x2e, 0x6c, 0x66, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x66, 0x2d, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_auth_proto_rawDescOnce sync.Once
	file_auth_auth_proto_rawDescData = file_auth_auth_proto_rawDesc
)

func file_auth_auth_proto_rawDescGZIP() []byte {
	file_auth_auth_proto_rawDescOnce.Do(func() {
		file_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_auth_proto_rawDescData)
	})
	return file_auth_auth_proto_rawDescData
}

var file_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_auth_proto_goTypes = []interface{}{
	(*AuthBody)(nil),             // 0: org.lfedge.eve.auth.AuthBody
	(*AuthContainer)(nil),        // 1: org.lfedge.eve.auth.AuthContainer
	(evecommon.HashAlgorithm)(0), // 2: org.lfedge.eve.common.HashAlgorithm
}
var file_auth_auth_proto_depIdxs = []int32{
	0, // 0: org.lfedge.eve.auth.AuthContainer.protectedPayload:type_name -> org.lfedge.eve.auth.AuthBody
	2, // 1: org.lfedge.eve.auth.AuthContainer.algo:type_name -> org.lfedge.eve.common.HashAlgorithm
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_auth_auth_proto_init() }
func file_auth_auth_proto_init() {
	if File_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthBody); i {
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
		file_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthContainer); i {
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
			RawDescriptor: file_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_auth_proto_goTypes,
		DependencyIndexes: file_auth_auth_proto_depIdxs,
		MessageInfos:      file_auth_auth_proto_msgTypes,
	}.Build()
	File_auth_auth_proto = out.File
	file_auth_auth_proto_rawDesc = nil
	file_auth_auth_proto_goTypes = nil
	file_auth_auth_proto_depIdxs = nil
}
