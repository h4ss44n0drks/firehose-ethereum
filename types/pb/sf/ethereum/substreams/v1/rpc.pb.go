// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.5
// source: sf/ethereum/substreams/v1/rpc.proto

package pbethss

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

type RpcCalls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calls []*RpcCall `protobuf:"bytes,1,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *RpcCalls) Reset() {
	*x = RpcCalls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_substreams_v1_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcCalls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcCalls) ProtoMessage() {}

func (x *RpcCalls) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_substreams_v1_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcCalls.ProtoReflect.Descriptor instead.
func (*RpcCalls) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_substreams_v1_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *RpcCalls) GetCalls() []*RpcCall {
	if x != nil {
		return x.Calls
	}
	return nil
}

type RpcCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToAddr []byte `protobuf:"bytes,1,opt,name=to_addr,json=toAddr,proto3" json:"to_addr,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RpcCall) Reset() {
	*x = RpcCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_substreams_v1_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcCall) ProtoMessage() {}

func (x *RpcCall) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_substreams_v1_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcCall.ProtoReflect.Descriptor instead.
func (*RpcCall) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_substreams_v1_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *RpcCall) GetToAddr() []byte {
	if x != nil {
		return x.ToAddr
	}
	return nil
}

func (x *RpcCall) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type RpcResponses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Responses []*RpcResponse `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *RpcResponses) Reset() {
	*x = RpcResponses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_substreams_v1_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcResponses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcResponses) ProtoMessage() {}

func (x *RpcResponses) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_substreams_v1_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcResponses.ProtoReflect.Descriptor instead.
func (*RpcResponses) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_substreams_v1_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *RpcResponses) GetResponses() []*RpcResponse {
	if x != nil {
		return x.Responses
	}
	return nil
}

type RpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Raw    []byte `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	Failed bool   `protobuf:"varint,2,opt,name=failed,proto3" json:"failed,omitempty"`
}

func (x *RpcResponse) Reset() {
	*x = RpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_substreams_v1_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcResponse) ProtoMessage() {}

func (x *RpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_substreams_v1_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcResponse.ProtoReflect.Descriptor instead.
func (*RpcResponse) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_substreams_v1_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *RpcResponse) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

func (x *RpcResponse) GetFailed() bool {
	if x != nil {
		return x.Failed
	}
	return false
}

var File_sf_ethereum_substreams_v1_rpc_proto protoreflect.FileDescriptor

var file_sf_ethereum_substreams_v1_rpc_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x66, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x22, 0x44, 0x0a, 0x08, 0x52, 0x70, 0x63, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x38, 0x0a, 0x05,
	0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x66,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x70, 0x63, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x22, 0x36, 0x0a, 0x07, 0x52, 0x70, 0x63, 0x43, 0x61, 0x6c,
	0x6c, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x54,
	0x0a, 0x0c, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x44,
	0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x70,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x73, 0x22, 0x37, 0x0a, 0x0b, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x42, 0x57, 0x5a,
	0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x68, 0x6f,
	0x73, 0x65, 0x2d, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x66, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x62, 0x65, 0x74, 0x68, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_ethereum_substreams_v1_rpc_proto_rawDescOnce sync.Once
	file_sf_ethereum_substreams_v1_rpc_proto_rawDescData = file_sf_ethereum_substreams_v1_rpc_proto_rawDesc
)

func file_sf_ethereum_substreams_v1_rpc_proto_rawDescGZIP() []byte {
	file_sf_ethereum_substreams_v1_rpc_proto_rawDescOnce.Do(func() {
		file_sf_ethereum_substreams_v1_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_ethereum_substreams_v1_rpc_proto_rawDescData)
	})
	return file_sf_ethereum_substreams_v1_rpc_proto_rawDescData
}

var file_sf_ethereum_substreams_v1_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sf_ethereum_substreams_v1_rpc_proto_goTypes = []interface{}{
	(*RpcCalls)(nil),     // 0: sf.ethereum.substreams.v1.RpcCalls
	(*RpcCall)(nil),      // 1: sf.ethereum.substreams.v1.RpcCall
	(*RpcResponses)(nil), // 2: sf.ethereum.substreams.v1.RpcResponses
	(*RpcResponse)(nil),  // 3: sf.ethereum.substreams.v1.RpcResponse
}
var file_sf_ethereum_substreams_v1_rpc_proto_depIdxs = []int32{
	1, // 0: sf.ethereum.substreams.v1.RpcCalls.calls:type_name -> sf.ethereum.substreams.v1.RpcCall
	3, // 1: sf.ethereum.substreams.v1.RpcResponses.responses:type_name -> sf.ethereum.substreams.v1.RpcResponse
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sf_ethereum_substreams_v1_rpc_proto_init() }
func file_sf_ethereum_substreams_v1_rpc_proto_init() {
	if File_sf_ethereum_substreams_v1_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_ethereum_substreams_v1_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcCalls); i {
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
		file_sf_ethereum_substreams_v1_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcCall); i {
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
		file_sf_ethereum_substreams_v1_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcResponses); i {
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
		file_sf_ethereum_substreams_v1_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcResponse); i {
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
			RawDescriptor: file_sf_ethereum_substreams_v1_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sf_ethereum_substreams_v1_rpc_proto_goTypes,
		DependencyIndexes: file_sf_ethereum_substreams_v1_rpc_proto_depIdxs,
		MessageInfos:      file_sf_ethereum_substreams_v1_rpc_proto_msgTypes,
	}.Build()
	File_sf_ethereum_substreams_v1_rpc_proto = out.File
	file_sf_ethereum_substreams_v1_rpc_proto_rawDesc = nil
	file_sf_ethereum_substreams_v1_rpc_proto_goTypes = nil
	file_sf_ethereum_substreams_v1_rpc_proto_depIdxs = nil
}
