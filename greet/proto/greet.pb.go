// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: greet.proto

package proto

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

type GreetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
}

func (x *GreetRequest) Reset() {
	*x = GreetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRequest) ProtoMessage() {}

func (x *GreetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRequest.ProtoReflect.Descriptor instead.
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return file_greet_proto_rawDescGZIP(), []int{0}
}

func (x *GreetRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

type GreetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetResponse) Reset() {
	*x = GreetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetResponse) ProtoMessage() {}

func (x *GreetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetResponse.ProtoReflect.Descriptor instead.
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return file_greet_proto_rawDescGZIP(), []int{1}
}

func (x *GreetResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type CaclRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber  int32 `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber int32 `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
}

func (x *CaclRequest) Reset() {
	*x = CaclRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaclRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaclRequest) ProtoMessage() {}

func (x *CaclRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaclRequest.ProtoReflect.Descriptor instead.
func (*CaclRequest) Descriptor() ([]byte, []int) {
	return file_greet_proto_rawDescGZIP(), []int{2}
}

func (x *CaclRequest) GetFirstNumber() int32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *CaclRequest) GetSecondNumber() int32 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type CaclResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CaclResponse) Reset() {
	*x = CaclResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaclResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaclResponse) ProtoMessage() {}

func (x *CaclResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaclResponse.ProtoReflect.Descriptor instead.
func (*CaclResponse) Descriptor() ([]byte, []int) {
	return file_greet_proto_rawDescGZIP(), []int{3}
}

func (x *CaclResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_greet_proto protoreflect.FileDescriptor

var file_greet_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x22, 0x2d, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x0d, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x55, 0x0a, 0x0b,
	0x43, 0x61, 0x63, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x0c, 0x43, 0x61, 0x63, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x73, 0x0a, 0x0c, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x63, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e,
	0x43, 0x61, 0x63, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x2e, 0x43, 0x61, 0x63, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x19, 0x5a, 0x17, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_greet_proto_rawDescOnce sync.Once
	file_greet_proto_rawDescData = file_greet_proto_rawDesc
)

func file_greet_proto_rawDescGZIP() []byte {
	file_greet_proto_rawDescOnce.Do(func() {
		file_greet_proto_rawDescData = protoimpl.X.CompressGZIP(file_greet_proto_rawDescData)
	})
	return file_greet_proto_rawDescData
}

var file_greet_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_greet_proto_goTypes = []interface{}{
	(*GreetRequest)(nil),  // 0: greet.GreetRequest
	(*GreetResponse)(nil), // 1: greet.GreetResponse
	(*CaclRequest)(nil),   // 2: greet.CaclRequest
	(*CaclResponse)(nil),  // 3: greet.CaclResponse
}
var file_greet_proto_depIdxs = []int32{
	0, // 0: greet.GreetService.Greet:input_type -> greet.GreetRequest
	2, // 1: greet.GreetService.Calc:input_type -> greet.CaclRequest
	1, // 2: greet.GreetService.Greet:output_type -> greet.GreetResponse
	3, // 3: greet.GreetService.Calc:output_type -> greet.CaclResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_greet_proto_init() }
func file_greet_proto_init() {
	if File_greet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRequest); i {
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
		file_greet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetResponse); i {
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
		file_greet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaclRequest); i {
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
		file_greet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaclResponse); i {
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
			RawDescriptor: file_greet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greet_proto_goTypes,
		DependencyIndexes: file_greet_proto_depIdxs,
		MessageInfos:      file_greet_proto_msgTypes,
	}.Build()
	File_greet_proto = out.File
	file_greet_proto_rawDesc = nil
	file_greet_proto_goTypes = nil
	file_greet_proto_depIdxs = nil
}
