// Copyright © 2021-2022 Durudex

// This file is part of Durudex: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

// Durudex is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with Durudex. If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: internal/delivery/grpc/pb/email.proto

package pb

import (
	types "github.com/durudex/durudex-email-service/internal/delivery/grpc/pb/types"
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

type UserCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Code     uint64 `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *UserCodeRequest) Reset() {
	*x = UserCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCodeRequest) ProtoMessage() {}

func (x *UserCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCodeRequest.ProtoReflect.Descriptor instead.
func (*UserCodeRequest) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_pb_email_proto_rawDescGZIP(), []int{0}
}

func (x *UserCodeRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserCodeRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserCodeRequest) GetCode() uint64 {
	if x != nil {
		return x.Code
	}
	return 0
}

type UserLoggedInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Ip    string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *UserLoggedInRequest) Reset() {
	*x = UserLoggedInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoggedInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoggedInRequest) ProtoMessage() {}

func (x *UserLoggedInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoggedInRequest.ProtoReflect.Descriptor instead.
func (*UserLoggedInRequest) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_pb_email_proto_rawDescGZIP(), []int{1}
}

func (x *UserLoggedInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserLoggedInRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type UserRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UserRegisterRequest) Reset() {
	*x = UserRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterRequest) ProtoMessage() {}

func (x *UserRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_pb_email_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterRequest.ProtoReflect.Descriptor instead.
func (*UserRegisterRequest) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_pb_email_proto_rawDescGZIP(), []int{2}
}

func (x *UserRegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_internal_delivery_grpc_pb_email_proto protoreflect.FileDescriptor

var file_internal_delivery_grpc_pb_email_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78,
	0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3b, 0x0a,
	0x13, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x47, 0x0a, 0x13, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0xcf, 0x01, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1e, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x41, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x12,
	0x22, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x41, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x22, 0x2e, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2e, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2f, 0x64, 0x75, 0x72, 0x75,
	0x64, 0x65, 0x78, 0x2d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_delivery_grpc_pb_email_proto_rawDescOnce sync.Once
	file_internal_delivery_grpc_pb_email_proto_rawDescData = file_internal_delivery_grpc_pb_email_proto_rawDesc
)

func file_internal_delivery_grpc_pb_email_proto_rawDescGZIP() []byte {
	file_internal_delivery_grpc_pb_email_proto_rawDescOnce.Do(func() {
		file_internal_delivery_grpc_pb_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_delivery_grpc_pb_email_proto_rawDescData)
	})
	return file_internal_delivery_grpc_pb_email_proto_rawDescData
}

var file_internal_delivery_grpc_pb_email_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_delivery_grpc_pb_email_proto_goTypes = []interface{}{
	(*UserCodeRequest)(nil),     // 0: durudex.email.UserCodeRequest
	(*UserLoggedInRequest)(nil), // 1: durudex.email.UserLoggedInRequest
	(*UserRegisterRequest)(nil), // 2: durudex.email.UserRegisterRequest
	(*types.Status)(nil),        // 3: types.Status
}
var file_internal_delivery_grpc_pb_email_proto_depIdxs = []int32{
	0, // 0: durudex.email.EmailService.UserCode:input_type -> durudex.email.UserCodeRequest
	1, // 1: durudex.email.EmailService.UserLoggedIn:input_type -> durudex.email.UserLoggedInRequest
	2, // 2: durudex.email.EmailService.UserRegister:input_type -> durudex.email.UserRegisterRequest
	3, // 3: durudex.email.EmailService.UserCode:output_type -> types.Status
	3, // 4: durudex.email.EmailService.UserLoggedIn:output_type -> types.Status
	3, // 5: durudex.email.EmailService.UserRegister:output_type -> types.Status
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_delivery_grpc_pb_email_proto_init() }
func file_internal_delivery_grpc_pb_email_proto_init() {
	if File_internal_delivery_grpc_pb_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_delivery_grpc_pb_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCodeRequest); i {
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
		file_internal_delivery_grpc_pb_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoggedInRequest); i {
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
		file_internal_delivery_grpc_pb_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterRequest); i {
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
			RawDescriptor: file_internal_delivery_grpc_pb_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_delivery_grpc_pb_email_proto_goTypes,
		DependencyIndexes: file_internal_delivery_grpc_pb_email_proto_depIdxs,
		MessageInfos:      file_internal_delivery_grpc_pb_email_proto_msgTypes,
	}.Build()
	File_internal_delivery_grpc_pb_email_proto = out.File
	file_internal_delivery_grpc_pb_email_proto_rawDesc = nil
	file_internal_delivery_grpc_pb_email_proto_goTypes = nil
	file_internal_delivery_grpc_pb_email_proto_depIdxs = nil
}
