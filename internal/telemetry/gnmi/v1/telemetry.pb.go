// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: telemetry.proto

package v1

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

type Telemetry_Encoding int32

const (
	Telemetry_Encoding_GPB Telemetry_Encoding = 0
)

// Enum value maps for Telemetry_Encoding.
var (
	Telemetry_Encoding_name = map[int32]string{
		0: "Encoding_GPB",
	}
	Telemetry_Encoding_value = map[string]int32{
		"Encoding_GPB": 0,
	}
)

func (x Telemetry_Encoding) Enum() *Telemetry_Encoding {
	p := new(Telemetry_Encoding)
	*p = x
	return p
}

func (x Telemetry_Encoding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Telemetry_Encoding) Descriptor() protoreflect.EnumDescriptor {
	return file_telemetry_proto_enumTypes[0].Descriptor()
}

func (Telemetry_Encoding) Type() protoreflect.EnumType {
	return &file_telemetry_proto_enumTypes[0]
}

func (x Telemetry_Encoding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Telemetry_Encoding.Descriptor instead.
func (Telemetry_Encoding) EnumDescriptor() ([]byte, []int) {
	return file_telemetry_proto_rawDescGZIP(), []int{0, 0}
}

type Telemetry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeIdStr           string             `protobuf:"bytes,1,opt,name=node_id_str,json=nodeIdStr,proto3" json:"node_id_str,omitempty"`                                //�豸��hostname��Ϊ�豸�������е�Ψһ��ʶ���û����������޸ģ�GPB����ʱ����Ϊ1��
	SubscriptionIdStr   string             `protobuf:"bytes,2,opt,name=subscription_id_str,json=subscriptionIdStr,proto3" json:"subscription_id_str,omitempty"`        //�������ƣ���̬���ö���ʱ�Ķ������ƣ�GPB����ʱ����Ϊ2��
	SensorPath          string             `protobuf:"bytes,3,opt,name=sensor_path,json=sensorPath,proto3" json:"sensor_path,omitempty"`                               //����·����GPB����ʱ����Ϊ3��
	CollectionId        uint64             `protobuf:"varint,4,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`                        //��ʶ�����ִΣ�GPB����ʱ����Ϊ4��
	CollectionStartTime uint64             `protobuf:"varint,5,opt,name=collection_start_time,json=collectionStartTime,proto3" json:"collection_start_time,omitempty"` //��ʶ�����ִο�ʼʱ�䣬GPB����ʱ����Ϊ5��
	MsgTimestamp        uint64             `protobuf:"varint,6,opt,name=msg_timestamp,json=msgTimestamp,proto3" json:"msg_timestamp,omitempty"`                        //���ɱ���Ϣ��ʱ�����GPB����ʱ����Ϊ6��
	DataGpb             *TelemetryGPBTable `protobuf:"bytes,7,opt,name=data_gpb,json=dataGpb,proto3" json:"data_gpb,omitempty"`                                        //���ص����ݣ���TelemetryGPBTable���壬GPB����ʱ����Ϊ7��
	CollectionEndTime   uint64             `protobuf:"varint,8,opt,name=collection_end_time,json=collectionEndTime,proto3" json:"collection_end_time,omitempty"`       //��ʶ�����ִν���ʱ�䣬GPB����ʱ����Ϊ8��
	CurrentPeriod       uint32             `protobuf:"varint,9,opt,name=current_period,json=currentPeriod,proto3" json:"current_period,omitempty"`                     //�������ȣ���λ�Ǻ��룬GPB����ʱ����Ϊ9��
	ExceptDesc          string             `protobuf:"bytes,10,opt,name=except_desc,json=exceptDesc,proto3" json:"except_desc,omitempty"`                              //�쳣������Ϣ�������쳣ʱ�����ϱ��쳣��Ϣ��GPB����ʱ����Ϊ10��
	ProductName         string             `protobuf:"bytes,11,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`                           //��Ʒ��̬��
	Encoding            Telemetry_Encoding `protobuf:"varint,12,opt,name=encoding,proto3,enum=telemetry.Telemetry_Encoding" json:"encoding,omitempty"`                 //���ݱ��롣ΪGPBʱ��data_gpb�ֶ���Ч������ʱdata_str�ֶ���Ч
}

func (x *Telemetry) Reset() {
	*x = Telemetry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Telemetry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Telemetry) ProtoMessage() {}

func (x *Telemetry) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Telemetry.ProtoReflect.Descriptor instead.
func (*Telemetry) Descriptor() ([]byte, []int) {
	return file_telemetry_proto_rawDescGZIP(), []int{0}
}

func (x *Telemetry) GetNodeIdStr() string {
	if x != nil {
		return x.NodeIdStr
	}
	return ""
}

func (x *Telemetry) GetSubscriptionIdStr() string {
	if x != nil {
		return x.SubscriptionIdStr
	}
	return ""
}

func (x *Telemetry) GetSensorPath() string {
	if x != nil {
		return x.SensorPath
	}
	return ""
}

func (x *Telemetry) GetCollectionId() uint64 {
	if x != nil {
		return x.CollectionId
	}
	return 0
}

func (x *Telemetry) GetCollectionStartTime() uint64 {
	if x != nil {
		return x.CollectionStartTime
	}
	return 0
}

func (x *Telemetry) GetMsgTimestamp() uint64 {
	if x != nil {
		return x.MsgTimestamp
	}
	return 0
}

func (x *Telemetry) GetDataGpb() *TelemetryGPBTable {
	if x != nil {
		return x.DataGpb
	}
	return nil
}

func (x *Telemetry) GetCollectionEndTime() uint64 {
	if x != nil {
		return x.CollectionEndTime
	}
	return 0
}

func (x *Telemetry) GetCurrentPeriod() uint32 {
	if x != nil {
		return x.CurrentPeriod
	}
	return 0
}

func (x *Telemetry) GetExceptDesc() string {
	if x != nil {
		return x.ExceptDesc
	}
	return ""
}

func (x *Telemetry) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *Telemetry) GetEncoding() Telemetry_Encoding {
	if x != nil {
		return x.Encoding
	}
	return Telemetry_Encoding_GPB
}

type TelemetryGPBTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Row []*TelemetryRowGPB `protobuf:"bytes,1,rep,name=row,proto3" json:"row,omitempty"` //���鶨�壬��ʶ������TelemetryRowGPB�ṹ���ظ���GPB����ʱ����Ϊ1��
}

func (x *TelemetryGPBTable) Reset() {
	*x = TelemetryGPBTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryGPBTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryGPBTable) ProtoMessage() {}

func (x *TelemetryGPBTable) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryGPBTable.ProtoReflect.Descriptor instead.
func (*TelemetryGPBTable) Descriptor() ([]byte, []int) {
	return file_telemetry_proto_rawDescGZIP(), []int{1}
}

func (x *TelemetryGPBTable) GetRow() []*TelemetryRowGPB {
	if x != nil {
		return x.Row
	}
	return nil
}

type TelemetryRowGPB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"` //������ǰʵ����ʱ�����GPB����ʱ����Ϊ1��
	Content   []byte `protobuf:"bytes,11,opt,name=content,proto3" json:"content,omitempty"`     //���صĲ���ʵ�����ݣ�GPB����ʱ����Ϊ11����Ҫ���sensor_path�ֶΣ��ſ����жϴ˴������ĸ�proto�ļ����롣
}

func (x *TelemetryRowGPB) Reset() {
	*x = TelemetryRowGPB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemetryRowGPB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryRowGPB) ProtoMessage() {}

func (x *TelemetryRowGPB) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryRowGPB.ProtoReflect.Descriptor instead.
func (*TelemetryRowGPB) Descriptor() ([]byte, []int) {
	return file_telemetry_proto_rawDescGZIP(), []int{2}
}

func (x *TelemetryRowGPB) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *TelemetryRowGPB) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_telemetry_proto protoreflect.FileDescriptor

var file_telemetry_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x22, 0xa7, 0x04, 0x0a,
	0x09, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0b, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x53, 0x74, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x73, 0x74,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x53, 0x74, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x32, 0x0a, 0x15, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x13, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6d, 0x73, 0x67,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x37, 0x0a, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x67, 0x70, 0x62, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x47, 0x50, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x47,
	0x70, 0x62, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x11, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x63,
	0x65, 0x70, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x44, 0x65, 0x73, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a,
	0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x54, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x08,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x1c, 0x0a, 0x08, 0x45, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x47, 0x50, 0x42, 0x10, 0x00, 0x22, 0x41, 0x0a, 0x11, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x47, 0x50, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x72,
	0x6f, 0x77, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x6f,
	0x77, 0x47, 0x50, 0x42, 0x52, 0x03, 0x72, 0x6f, 0x77, 0x22, 0x49, 0x0a, 0x0f, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x6f, 0x77, 0x47, 0x50, 0x42, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_telemetry_proto_rawDescOnce sync.Once
	file_telemetry_proto_rawDescData = file_telemetry_proto_rawDesc
)

func file_telemetry_proto_rawDescGZIP() []byte {
	file_telemetry_proto_rawDescOnce.Do(func() {
		file_telemetry_proto_rawDescData = protoimpl.X.CompressGZIP(file_telemetry_proto_rawDescData)
	})
	return file_telemetry_proto_rawDescData
}

var file_telemetry_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_telemetry_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_telemetry_proto_goTypes = []interface{}{
	(Telemetry_Encoding)(0),   // 0: telemetry.Telemetry.Encoding
	(*Telemetry)(nil),         // 1: telemetry.Telemetry
	(*TelemetryGPBTable)(nil), // 2: telemetry.TelemetryGPBTable
	(*TelemetryRowGPB)(nil),   // 3: telemetry.TelemetryRowGPB
}
var file_telemetry_proto_depIdxs = []int32{
	2, // 0: telemetry.Telemetry.data_gpb:type_name -> telemetry.TelemetryGPBTable
	0, // 1: telemetry.Telemetry.encoding:type_name -> telemetry.Telemetry.Encoding
	3, // 2: telemetry.TelemetryGPBTable.row:type_name -> telemetry.TelemetryRowGPB
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_telemetry_proto_init() }
func file_telemetry_proto_init() {
	if File_telemetry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_telemetry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Telemetry); i {
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
		file_telemetry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryGPBTable); i {
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
		file_telemetry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemetryRowGPB); i {
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
			RawDescriptor: file_telemetry_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_telemetry_proto_goTypes,
		DependencyIndexes: file_telemetry_proto_depIdxs,
		EnumInfos:         file_telemetry_proto_enumTypes,
		MessageInfos:      file_telemetry_proto_msgTypes,
	}.Build()
	File_telemetry_proto = out.File
	file_telemetry_proto_rawDesc = nil
	file_telemetry_proto_goTypes = nil
	file_telemetry_proto_depIdxs = nil
}