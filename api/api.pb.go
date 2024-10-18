// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc1
// source: api/api.proto

package api

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

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DayRate   float64 `protobuf:"fixed64,3,opt,name=day_rate,json=dayRate,proto3" json:"day_rate,omitempty"`
	MonthRate float64 `protobuf:"fixed64,4,opt,name=month_rate,json=monthRate,proto3" json:"month_rate,omitempty"`
	Image     string  `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	mi := &file_api_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *Car) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Car) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Car) GetDayRate() float64 {
	if x != nil {
		return x.DayRate
	}
	return 0
}

func (x *Car) GetMonthRate() float64 {
	if x != nil {
		return x.MonthRate
	}
	return 0
}

func (x *Car) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type GetCarsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	Page   int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetCarsRequest) Reset() {
	*x = GetCarsRequest{}
	mi := &file_api_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCarsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarsRequest) ProtoMessage() {}

func (x *GetCarsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarsRequest.ProtoReflect.Descriptor instead.
func (*GetCarsRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetCarsRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetCarsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetCarsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetCarsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*Car `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Page  int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Total int32  `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetCarsResponse) Reset() {
	*x = GetCarsResponse{}
	mi := &file_api_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCarsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarsResponse) ProtoMessage() {}

func (x *GetCarsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarsResponse.ProtoReflect.Descriptor instead.
func (*GetCarsResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetCarsResponse) GetData() []*Car {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetCarsResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetCarsResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetCarsResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CarId           int64  `protobuf:"varint,2,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	OrderDate       int64  `protobuf:"varint,3,opt,name=order_date,json=orderDate,proto3" json:"order_date,omitempty"`
	PickupDate      int64  `protobuf:"varint,4,opt,name=pickup_date,json=pickupDate,proto3" json:"pickup_date,omitempty"`
	DropoffDate     int64  `protobuf:"varint,5,opt,name=dropoff_date,json=dropoffDate,proto3" json:"dropoff_date,omitempty"`
	PickupLocation  string `protobuf:"bytes,6,opt,name=pickup_location,json=pickupLocation,proto3" json:"pickup_location,omitempty"`
	DropoffLocation string `protobuf:"bytes,7,opt,name=dropoff_location,json=dropoffLocation,proto3" json:"dropoff_location,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_api_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *Order) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetCarId() int64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *Order) GetOrderDate() int64 {
	if x != nil {
		return x.OrderDate
	}
	return 0
}

func (x *Order) GetPickupDate() int64 {
	if x != nil {
		return x.PickupDate
	}
	return 0
}

func (x *Order) GetDropoffDate() int64 {
	if x != nil {
		return x.DropoffDate
	}
	return 0
}

func (x *Order) GetPickupLocation() string {
	if x != nil {
		return x.PickupLocation
	}
	return ""
}

func (x *Order) GetDropoffLocation() string {
	if x != nil {
		return x.DropoffLocation
	}
	return ""
}

type GetOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	Page   int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit  int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetOrdersRequest) Reset() {
	*x = GetOrdersRequest{}
	mi := &file_api_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersRequest) ProtoMessage() {}

func (x *GetOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersRequest.ProtoReflect.Descriptor instead.
func (*GetOrdersRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrdersRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetOrdersRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetOrdersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*Order `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Page  int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Total int32    `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetOrdersResponse) Reset() {
	*x = GetOrdersResponse{}
	mi := &file_api_api_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersResponse) ProtoMessage() {}

func (x *GetOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersResponse.ProtoReflect.Descriptor instead.
func (*GetOrdersResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrdersResponse) GetData() []*Order {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetOrdersResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetOrdersResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetOrdersResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_api_api_proto protoreflect.FileDescriptor

var file_api_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0x79, 0x0a, 0x03, 0x43, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x64, 0x61, 0x79, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x07, 0x64, 0x61, 0x79, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x52, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22,
	0x52, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x6f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x22, 0xe5, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x61, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x69, 0x63, 0x6b, 0x75,
	0x70, 0x44, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x72, 0x6f, 0x70, 0x6f, 0x66, 0x66,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x72, 0x6f,
	0x70, 0x6f, 0x66, 0x66, 0x44, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x69, 0x63, 0x6b,
	0x75, 0x70, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x72, 0x6f, 0x70, 0x6f, 0x66, 0x66, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x72, 0x6f,
	0x70, 0x6f, 0x66, 0x66, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x73, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xa7, 0x03, 0x0a, 0x11, 0x43, 0x61, 0x72, 0x52,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x34, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x73, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x12, 0x08, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x72, 0x1a, 0x08, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x08, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x72, 0x1a, 0x08, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x61, 0x72, 0x12, 0x1f, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x12, 0x08, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x72, 0x1a, 0x08, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x72, 0x12, 0x1f, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x12, 0x08, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x72, 0x1a, 0x08,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x72, 0x12, 0x3a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a,
	0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x0a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_api_proto_rawDescOnce sync.Once
	file_api_api_proto_rawDescData = file_api_api_proto_rawDesc
)

func file_api_api_proto_rawDescGZIP() []byte {
	file_api_api_proto_rawDescOnce.Do(func() {
		file_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_proto_rawDescData)
	})
	return file_api_api_proto_rawDescData
}

var file_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_api_proto_goTypes = []any{
	(*Car)(nil),               // 0: api.Car
	(*GetCarsRequest)(nil),    // 1: api.GetCarsRequest
	(*GetCarsResponse)(nil),   // 2: api.GetCarsResponse
	(*Order)(nil),             // 3: api.Order
	(*GetOrdersRequest)(nil),  // 4: api.GetOrdersRequest
	(*GetOrdersResponse)(nil), // 5: api.GetOrdersResponse
}
var file_api_api_proto_depIdxs = []int32{
	0,  // 0: api.GetCarsResponse.data:type_name -> api.Car
	3,  // 1: api.GetOrdersResponse.data:type_name -> api.Order
	1,  // 2: api.CarRentalServices.GetCars:input_type -> api.GetCarsRequest
	0,  // 3: api.CarRentalServices.CreateCar:input_type -> api.Car
	0,  // 4: api.CarRentalServices.GetCarById:input_type -> api.Car
	0,  // 5: api.CarRentalServices.UpdateCar:input_type -> api.Car
	0,  // 6: api.CarRentalServices.DeleteCar:input_type -> api.Car
	4,  // 7: api.CarRentalServices.GetOrders:input_type -> api.GetOrdersRequest
	3,  // 8: api.CarRentalServices.CreateOrder:input_type -> api.Order
	3,  // 9: api.CarRentalServices.GetOrderById:input_type -> api.Order
	3,  // 10: api.CarRentalServices.UpdateOrder:input_type -> api.Order
	3,  // 11: api.CarRentalServices.DeleteOrder:input_type -> api.Order
	2,  // 12: api.CarRentalServices.GetCars:output_type -> api.GetCarsResponse
	0,  // 13: api.CarRentalServices.CreateCar:output_type -> api.Car
	0,  // 14: api.CarRentalServices.GetCarById:output_type -> api.Car
	0,  // 15: api.CarRentalServices.UpdateCar:output_type -> api.Car
	0,  // 16: api.CarRentalServices.DeleteCar:output_type -> api.Car
	5,  // 17: api.CarRentalServices.GetOrders:output_type -> api.GetOrdersResponse
	3,  // 18: api.CarRentalServices.CreateOrder:output_type -> api.Order
	3,  // 19: api.CarRentalServices.GetOrderById:output_type -> api.Order
	3,  // 20: api.CarRentalServices.UpdateOrder:output_type -> api.Order
	3,  // 21: api.CarRentalServices.DeleteOrder:output_type -> api.Order
	12, // [12:22] is the sub-list for method output_type
	2,  // [2:12] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_api_proto_init() }
func file_api_api_proto_init() {
	if File_api_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_proto_goTypes,
		DependencyIndexes: file_api_api_proto_depIdxs,
		MessageInfos:      file_api_api_proto_msgTypes,
	}.Build()
	File_api_api_proto = out.File
	file_api_api_proto_rawDesc = nil
	file_api_api_proto_goTypes = nil
	file_api_api_proto_depIdxs = nil
}
