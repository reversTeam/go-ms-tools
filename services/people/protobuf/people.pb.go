// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.17.3
// source: services/people/protobuf/people.proto

package protobuf

import (
	protobuf "github.com/reversTeam/go-ms-tools/services/abs/protobuf"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PeopleEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" db:"id" xml:"id" validate:"required,uuid4_rfc4122"` // @gotags: json:"id" db:"id" xml:"id" validate:"required,uuid4_rfc4122"`
}

func (x *PeopleEntity) Reset() {
	*x = PeopleEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_people_protobuf_people_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeopleEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeopleEntity) ProtoMessage() {}

func (x *PeopleEntity) ProtoReflect() protoreflect.Message {
	mi := &file_services_people_protobuf_people_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeopleEntity.ProtoReflect.Descriptor instead.
func (*PeopleEntity) Descriptor() ([]byte, []int) {
	return file_services_people_protobuf_people_proto_rawDescGZIP(), []int{0}
}

func (x *PeopleEntity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PeopleCreateParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Firstname string `protobuf:"bytes,1,opt,name=firstname,proto3" json:"firstname" xml:"firstname" db:"firstname" validate:"required,alphanumunicode"` // @gotags: json:"firstname" xml:"firstname" db:"firstname" validate:"required,alphanumunicode"`
	Lastname  string `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname" xml:"lastname" db:"lastname" validate:"required,alphanumunicode"`     // @gotags: json:"lastname" xml:"lastname" db:"lastname" validate:"required,alphanumunicode"`
	Birthday  string `protobuf:"bytes,3,opt,name=birthday,proto3" json:"birthday" xml:"birthday" db:"birthday" validate:"required"`                     // @gotags: json:"birthday" xml:"birthday" db:"birthday" validate:"required"`
}

func (x *PeopleCreateParams) Reset() {
	*x = PeopleCreateParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_people_protobuf_people_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeopleCreateParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeopleCreateParams) ProtoMessage() {}

func (x *PeopleCreateParams) ProtoReflect() protoreflect.Message {
	mi := &file_services_people_protobuf_people_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeopleCreateParams.ProtoReflect.Descriptor instead.
func (*PeopleCreateParams) Descriptor() ([]byte, []int) {
	return file_services_people_protobuf_people_proto_rawDescGZIP(), []int{1}
}

func (x *PeopleCreateParams) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *PeopleCreateParams) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *PeopleCreateParams) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

type PeopleUpdateParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id" db:"id" xml:"id" validate:"required,uuid4_rfc4122"`                               // @gotags: json:"id" db:"id" xml:"id" validate:"required,uuid4_rfc4122"`
	Status     string                 `protobuf:"bytes,2,opt,name=Status,proto3" json:"status" xml:"status" db:"status" validate:"required,alpha"`                       // @gotags: json:"status" xml:"status" db:"status" validate:"required,alpha"`
	ExpireddAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=ExpireddAt,proto3" json:"expired_at" xml:"iexpired_atd" db:"expired_at"`                               // @gotags: json:"expired_at" xml:"iexpired_atd" db:"expired_at"
	Firstname  string                 `protobuf:"bytes,4,opt,name=firstname,proto3" json:"firstname" xml:"firstname" db:"firstname" validate:"required,alphanumunicode"` // @gotags: json:"firstname" xml:"firstname" db:"firstname" validate:"required,alphanumunicode"`
	Lastname   string                 `protobuf:"bytes,5,opt,name=lastname,proto3" json:"lastname" xml:"lastname" db:"lastname" validate:"required,alphanumunicode"`     // @gotags: json:"lastname" xml:"lastname" db:"lastname" validate:"required,alphanumunicode"`
	Birthday   string                 `protobuf:"bytes,6,opt,name=birthday,proto3" json:"birthday" xml:"birthday" db:"birthday" validate:"required"`                     // @gotags: json:"birthday" xml:"birthday" db:"birthday" validate:"required"`
}

func (x *PeopleUpdateParams) Reset() {
	*x = PeopleUpdateParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_people_protobuf_people_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeopleUpdateParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeopleUpdateParams) ProtoMessage() {}

func (x *PeopleUpdateParams) ProtoReflect() protoreflect.Message {
	mi := &file_services_people_protobuf_people_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeopleUpdateParams.ProtoReflect.Descriptor instead.
func (*PeopleUpdateParams) Descriptor() ([]byte, []int) {
	return file_services_people_protobuf_people_proto_rawDescGZIP(), []int{2}
}

func (x *PeopleUpdateParams) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PeopleUpdateParams) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PeopleUpdateParams) GetExpireddAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireddAt
	}
	return nil
}

func (x *PeopleUpdateParams) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *PeopleUpdateParams) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *PeopleUpdateParams) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

type PeopleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=Id,proto3" json:"id" xml:"id" db:"id"`                                // @gotags: json:"id" xml:"id" db:"id"
	Status    string                 `protobuf:"bytes,2,opt,name=Status,proto3" json:"status" xml:"status" db:"status"`                // @gotags: json:"status" xml:"status" db:"status"
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=CreatedAt,proto3" json:"created_at" xml:"created_at" db:"created_at"` // @gotags: json:"created_at" xml:"created_at" db:"created_at"
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=UpdatedAt,proto3" json:"updated_at" xml:"updated_at" db:"updated_at"` // @gotags: json:"updated_at" xml:"updated_at" db:"updated_at"
	ExpiredAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=ExpiredAt,proto3" json:"expired_at" xml:"expired_at" db:"expired_at"` // @gotags: json:"expired_at" xml:"expired_at" db:"expired_at"
	Firstname string                 `protobuf:"bytes,6,opt,name=Firstname,proto3" json:"firstname" xml:"firstname" db:"firstname"`    // @gotags: json:"firstname" xml:"firstname" db:"firstname"
	Lastname  string                 `protobuf:"bytes,7,opt,name=Lastname,proto3" json:"lastname" xml:"lastname" db:"lastname"`        // @gotags: json:"lastname" xml:"lastname" db:"lastname"
	Birthday  string                 `protobuf:"bytes,8,opt,name=Birthday,proto3" json:"birthday" xml:"birthday" db:"birthday"`        // @gotags: json:"birthday" xml:"birthday" db:"birthday"
}

func (x *PeopleResponse) Reset() {
	*x = PeopleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_people_protobuf_people_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeopleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeopleResponse) ProtoMessage() {}

func (x *PeopleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_people_protobuf_people_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeopleResponse.ProtoReflect.Descriptor instead.
func (*PeopleResponse) Descriptor() ([]byte, []int) {
	return file_services_people_protobuf_people_proto_rawDescGZIP(), []int{3}
}

func (x *PeopleResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PeopleResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PeopleResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PeopleResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *PeopleResponse) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

func (x *PeopleResponse) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *PeopleResponse) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *PeopleResponse) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

var File_services_people_protobuf_people_proto protoreflect.FileDescriptor

var file_services_people_protobuf_people_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x54,
	0x65, 0x61, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x73, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1e, 0x0a, 0x0c, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x6a, 0x0a, 0x12, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x22, 0xce, 0x01, 0x0a, 0x12,
	0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x22, 0xbc, 0x02, 0x0a,
	0x0e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x32, 0x9c, 0x04, 0x0a, 0x06,
	0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x5a, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x30, 0x01, 0x12, 0x72, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x65, 0x6f,
	0x70, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f,
	0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x6b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x25, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2e, 0x50,
	0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x6e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x62,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x3a, 0x01, 0x2a, 0x32, 0x0c, 0x2f, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x65, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x25, 0x2e,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2e, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x62, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x2a, 0x0c, 0x2f, 0x70,
	0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x54,
	0x65, 0x61, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x73, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_people_protobuf_people_proto_rawDescOnce sync.Once
	file_services_people_protobuf_people_proto_rawDescData = file_services_people_protobuf_people_proto_rawDesc
)

func file_services_people_protobuf_people_proto_rawDescGZIP() []byte {
	file_services_people_protobuf_people_proto_rawDescOnce.Do(func() {
		file_services_people_protobuf_people_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_people_protobuf_people_proto_rawDescData)
	})
	return file_services_people_protobuf_people_proto_rawDescData
}

var file_services_people_protobuf_people_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_services_people_protobuf_people_proto_goTypes = []interface{}{
	(*PeopleEntity)(nil),          // 0: go.micro.service.people.PeopleEntity
	(*PeopleCreateParams)(nil),    // 1: go.micro.service.people.PeopleCreateParams
	(*PeopleUpdateParams)(nil),    // 2: go.micro.service.people.PeopleUpdateParams
	(*PeopleResponse)(nil),        // 3: go.micro.service.people.PeopleResponse
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 5: google.protobuf.Empty
	(*protobuf.Response)(nil),     // 6: go.micro.service.abs.Response
}
var file_services_people_protobuf_people_proto_depIdxs = []int32{
	4, // 0: go.micro.service.people.PeopleUpdateParams.ExpireddAt:type_name -> google.protobuf.Timestamp
	4, // 1: go.micro.service.people.PeopleResponse.CreatedAt:type_name -> google.protobuf.Timestamp
	4, // 2: go.micro.service.people.PeopleResponse.UpdatedAt:type_name -> google.protobuf.Timestamp
	4, // 3: go.micro.service.people.PeopleResponse.ExpiredAt:type_name -> google.protobuf.Timestamp
	5, // 4: go.micro.service.people.People.List:input_type -> google.protobuf.Empty
	1, // 5: go.micro.service.people.People.Create:input_type -> go.micro.service.people.PeopleCreateParams
	0, // 6: go.micro.service.people.People.Get:input_type -> go.micro.service.people.PeopleEntity
	2, // 7: go.micro.service.people.People.Update:input_type -> go.micro.service.people.PeopleUpdateParams
	0, // 8: go.micro.service.people.People.Delete:input_type -> go.micro.service.people.PeopleEntity
	3, // 9: go.micro.service.people.People.List:output_type -> go.micro.service.people.PeopleResponse
	3, // 10: go.micro.service.people.People.Create:output_type -> go.micro.service.people.PeopleResponse
	3, // 11: go.micro.service.people.People.Get:output_type -> go.micro.service.people.PeopleResponse
	6, // 12: go.micro.service.people.People.Update:output_type -> go.micro.service.abs.Response
	6, // 13: go.micro.service.people.People.Delete:output_type -> go.micro.service.abs.Response
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_services_people_protobuf_people_proto_init() }
func file_services_people_protobuf_people_proto_init() {
	if File_services_people_protobuf_people_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_people_protobuf_people_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeopleEntity); i {
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
		file_services_people_protobuf_people_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeopleCreateParams); i {
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
		file_services_people_protobuf_people_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeopleUpdateParams); i {
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
		file_services_people_protobuf_people_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeopleResponse); i {
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
			RawDescriptor: file_services_people_protobuf_people_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_people_protobuf_people_proto_goTypes,
		DependencyIndexes: file_services_people_protobuf_people_proto_depIdxs,
		MessageInfos:      file_services_people_protobuf_people_proto_msgTypes,
	}.Build()
	File_services_people_protobuf_people_proto = out.File
	file_services_people_protobuf_people_proto_rawDesc = nil
	file_services_people_protobuf_people_proto_goTypes = nil
	file_services_people_protobuf_people_proto_depIdxs = nil
}
