// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: protos/auth/auth.proto

package auth

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

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletAddress   string `protobuf:"bytes,1,opt,name=walletAddress,proto3" json:"walletAddress,omitempty"`
	Signature       string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	SignatureExpiry int64  `protobuf:"varint,3,opt,name=signatureExpiry,proto3" json:"signatureExpiry,omitempty"`
	FirstName       string `protobuf:"bytes,4,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName        string `protobuf:"bytes,5,opt,name=lastName,proto3" json:"lastName,omitempty"`
	MiddleName      string `protobuf:"bytes,6,opt,name=middleName,proto3" json:"middleName,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_protos_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthRequest) GetWalletAddress() string {
	if x != nil {
		return x.WalletAddress
	}
	return ""
}

func (x *AuthRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *AuthRequest) GetSignatureExpiry() int64 {
	if x != nil {
		return x.SignatureExpiry
	}
	return 0
}

func (x *AuthRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *AuthRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *AuthRequest) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_protos_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_protos_auth_auth_proto protoreflect.FileDescriptor

var file_protos_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6b, 0x65, 0x79, 0x66, 0x69, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x22, 0xd5, 0x01, 0x0a, 0x0b,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xae, 0x01,
	0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x1e, 0x2e, 0x6b, 0x65, 0x79, 0x66, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x6b, 0x65, 0x79, 0x66, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4b, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x2e,
	0x6b, 0x65, 0x79, 0x66, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x6b, 0x65, 0x79, 0x66, 0x69, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19,
	0x5a, 0x17, 0x6b, 0x65, 0x79, 0x66, 0x69, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_auth_auth_proto_rawDescOnce sync.Once
	file_protos_auth_auth_proto_rawDescData = file_protos_auth_auth_proto_rawDesc
)

func file_protos_auth_auth_proto_rawDescGZIP() []byte {
	file_protos_auth_auth_proto_rawDescOnce.Do(func() {
		file_protos_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_auth_auth_proto_rawDescData)
	})
	return file_protos_auth_auth_proto_rawDescData
}

var file_protos_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_auth_auth_proto_goTypes = []interface{}{
	(*AuthRequest)(nil),  // 0: keyfi_protos.auth.AuthRequest
	(*AuthResponse)(nil), // 1: keyfi_protos.auth.AuthResponse
}
var file_protos_auth_auth_proto_depIdxs = []int32{
	0, // 0: keyfi_protos.auth.AuthenticationService.Login:input_type -> keyfi_protos.auth.AuthRequest
	0, // 1: keyfi_protos.auth.AuthenticationService.Register:input_type -> keyfi_protos.auth.AuthRequest
	1, // 2: keyfi_protos.auth.AuthenticationService.Login:output_type -> keyfi_protos.auth.AuthResponse
	1, // 3: keyfi_protos.auth.AuthenticationService.Register:output_type -> keyfi_protos.auth.AuthResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_auth_auth_proto_init() }
func file_protos_auth_auth_proto_init() {
	if File_protos_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_protos_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
			RawDescriptor: file_protos_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_auth_auth_proto_goTypes,
		DependencyIndexes: file_protos_auth_auth_proto_depIdxs,
		MessageInfos:      file_protos_auth_auth_proto_msgTypes,
	}.Build()
	File_protos_auth_auth_proto = out.File
	file_protos_auth_auth_proto_rawDesc = nil
	file_protos_auth_auth_proto_goTypes = nil
	file_protos_auth_auth_proto_depIdxs = nil
}
