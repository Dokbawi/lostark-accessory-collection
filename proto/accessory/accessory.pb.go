// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: proto/accessory/accessory.proto

package accessory

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

type GetAverageTradePriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token           string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Tier            int32  `protobuf:"varint,2,opt,name=tier,proto3" json:"tier,omitempty"`
	PolishingEffect int32  `protobuf:"varint,3,opt,name=polishing_effect,json=polishingEffect,proto3" json:"polishing_effect,omitempty"`
	CategoryCode    int32  `protobuf:"varint,4,opt,name=category_code,json=categoryCode,proto3" json:"category_code,omitempty"`
	ItemGrade       string `protobuf:"bytes,5,opt,name=item_grade,json=itemGrade,proto3" json:"item_grade,omitempty"`
}

func (x *GetAverageTradePriceRequest) Reset() {
	*x = GetAverageTradePriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_accessory_accessory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAverageTradePriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAverageTradePriceRequest) ProtoMessage() {}

func (x *GetAverageTradePriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_accessory_accessory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAverageTradePriceRequest.ProtoReflect.Descriptor instead.
func (*GetAverageTradePriceRequest) Descriptor() ([]byte, []int) {
	return file_proto_accessory_accessory_proto_rawDescGZIP(), []int{0}
}

func (x *GetAverageTradePriceRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetAverageTradePriceRequest) GetTier() int32 {
	if x != nil {
		return x.Tier
	}
	return 0
}

func (x *GetAverageTradePriceRequest) GetPolishingEffect() int32 {
	if x != nil {
		return x.PolishingEffect
	}
	return 0
}

func (x *GetAverageTradePriceRequest) GetCategoryCode() int32 {
	if x != nil {
		return x.CategoryCode
	}
	return 0
}

func (x *GetAverageTradePriceRequest) GetItemGrade() string {
	if x != nil {
		return x.ItemGrade
	}
	return ""
}

type GetAverageTradePriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Grade        string  `protobuf:"bytes,3,opt,name=grade,proto3" json:"grade,omitempty"`
	AuctionPrice int32   `protobuf:"varint,4,opt,name=auction_price,json=auctionPrice,proto3" json:"auction_price,omitempty"`
	BuyPrice     int32   `protobuf:"varint,5,opt,name=buy_price,json=buyPrice,proto3" json:"buy_price,omitempty"`
	TradeLeft    int32   `protobuf:"varint,6,opt,name=trade_left,json=tradeLeft,proto3" json:"trade_left,omitempty"`
	Quality      float32 `protobuf:"fixed32,7,opt,name=quality,proto3" json:"quality,omitempty"`
}

func (x *GetAverageTradePriceResponse) Reset() {
	*x = GetAverageTradePriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_accessory_accessory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAverageTradePriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAverageTradePriceResponse) ProtoMessage() {}

func (x *GetAverageTradePriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_accessory_accessory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAverageTradePriceResponse.ProtoReflect.Descriptor instead.
func (*GetAverageTradePriceResponse) Descriptor() ([]byte, []int) {
	return file_proto_accessory_accessory_proto_rawDescGZIP(), []int{1}
}

func (x *GetAverageTradePriceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetAverageTradePriceResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAverageTradePriceResponse) GetGrade() string {
	if x != nil {
		return x.Grade
	}
	return ""
}

func (x *GetAverageTradePriceResponse) GetAuctionPrice() int32 {
	if x != nil {
		return x.AuctionPrice
	}
	return 0
}

func (x *GetAverageTradePriceResponse) GetBuyPrice() int32 {
	if x != nil {
		return x.BuyPrice
	}
	return 0
}

func (x *GetAverageTradePriceResponse) GetTradeLeft() int32 {
	if x != nil {
		return x.TradeLeft
	}
	return 0
}

func (x *GetAverageTradePriceResponse) GetQuality() float32 {
	if x != nil {
		return x.Quality
	}
	return 0
}

type GetAverageTradePriceResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items    []*GetAverageTradePriceResponse `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Progress float32                         `protobuf:"fixed32,2,opt,name=progress,proto3" json:"progress,omitempty"`
}

func (x *GetAverageTradePriceResponseList) Reset() {
	*x = GetAverageTradePriceResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_accessory_accessory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAverageTradePriceResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAverageTradePriceResponseList) ProtoMessage() {}

func (x *GetAverageTradePriceResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_accessory_accessory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAverageTradePriceResponseList.ProtoReflect.Descriptor instead.
func (*GetAverageTradePriceResponseList) Descriptor() ([]byte, []int) {
	return file_proto_accessory_accessory_proto_rawDescGZIP(), []int{2}
}

func (x *GetAverageTradePriceResponseList) GetItems() []*GetAverageTradePriceResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *GetAverageTradePriceResponseList) GetProgress() float32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

var File_proto_accessory_accessory_proto protoreflect.FileDescriptor

var file_proto_accessory_accessory_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x79, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x79, 0x22, 0xb6, 0x01, 0x0a,
	0x1b, 0x47, 0x65, 0x74, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x74, 0x69, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x6f, 0x6c, 0x69, 0x73, 0x68,
	0x69, 0x6e, 0x67, 0x5f, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x70, 0x6f, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x45, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d,
	0x47, 0x72, 0x61, 0x64, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x79, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x62, 0x75, 0x79, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x6c, 0x65, 0x66, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x72, 0x61, 0x64, 0x65, 0x4c, 0x65, 0x66,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x7d, 0x0a, 0x20, 0x47,
	0x65, 0x74, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x3d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x32, 0x83, 0x01, 0x0a, 0x10, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x6f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x19, 0x5a, 0x17, 0x6c, 0x6f, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_accessory_accessory_proto_rawDescOnce sync.Once
	file_proto_accessory_accessory_proto_rawDescData = file_proto_accessory_accessory_proto_rawDesc
)

func file_proto_accessory_accessory_proto_rawDescGZIP() []byte {
	file_proto_accessory_accessory_proto_rawDescOnce.Do(func() {
		file_proto_accessory_accessory_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_accessory_accessory_proto_rawDescData)
	})
	return file_proto_accessory_accessory_proto_rawDescData
}

var file_proto_accessory_accessory_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_accessory_accessory_proto_goTypes = []any{
	(*GetAverageTradePriceRequest)(nil),      // 0: accessory.GetAverageTradePriceRequest
	(*GetAverageTradePriceResponse)(nil),     // 1: accessory.GetAverageTradePriceResponse
	(*GetAverageTradePriceResponseList)(nil), // 2: accessory.GetAverageTradePriceResponseList
}
var file_proto_accessory_accessory_proto_depIdxs = []int32{
	1, // 0: accessory.GetAverageTradePriceResponseList.items:type_name -> accessory.GetAverageTradePriceResponse
	0, // 1: accessory.AccessoryService.GetAverageTradePrice:input_type -> accessory.GetAverageTradePriceRequest
	2, // 2: accessory.AccessoryService.GetAverageTradePrice:output_type -> accessory.GetAverageTradePriceResponseList
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_accessory_accessory_proto_init() }
func file_proto_accessory_accessory_proto_init() {
	if File_proto_accessory_accessory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_accessory_accessory_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetAverageTradePriceRequest); i {
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
		file_proto_accessory_accessory_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAverageTradePriceResponse); i {
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
		file_proto_accessory_accessory_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAverageTradePriceResponseList); i {
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
			RawDescriptor: file_proto_accessory_accessory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_accessory_accessory_proto_goTypes,
		DependencyIndexes: file_proto_accessory_accessory_proto_depIdxs,
		MessageInfos:      file_proto_accessory_accessory_proto_msgTypes,
	}.Build()
	File_proto_accessory_accessory_proto = out.File
	file_proto_accessory_accessory_proto_rawDesc = nil
	file_proto_accessory_accessory_proto_goTypes = nil
	file_proto_accessory_accessory_proto_depIdxs = nil
}