// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package gmicro_srv_comment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddCommentRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	ArticleId            int32    `protobuf:"varint,2,opt,name=ArticleId,proto3" json:"ArticleId,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=Comment,proto3" json:"Comment,omitempty"`
	FromId               string   `protobuf:"bytes,4,opt,name=FromId,proto3" json:"FromId,omitempty"`
	ToId                 string   `protobuf:"bytes,5,opt,name=ToId,proto3" json:"ToId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddCommentRequest) Reset()         { *m = AddCommentRequest{} }
func (m *AddCommentRequest) String() string { return proto.CompactTextString(m) }
func (*AddCommentRequest) ProtoMessage()    {}
func (*AddCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *AddCommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddCommentRequest.Unmarshal(m, b)
}
func (m *AddCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddCommentRequest.Marshal(b, m, deterministic)
}
func (m *AddCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddCommentRequest.Merge(m, src)
}
func (m *AddCommentRequest) XXX_Size() int {
	return xxx_messageInfo_AddCommentRequest.Size(m)
}
func (m *AddCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddCommentRequest proto.InternalMessageInfo

func (m *AddCommentRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AddCommentRequest) GetArticleId() int32 {
	if m != nil {
		return m.ArticleId
	}
	return 0
}

func (m *AddCommentRequest) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *AddCommentRequest) GetFromId() string {
	if m != nil {
		return m.FromId
	}
	return ""
}

func (m *AddCommentRequest) GetToId() string {
	if m != nil {
		return m.ToId
	}
	return ""
}

type AddCommentRespond struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddCommentRespond) Reset()         { *m = AddCommentRespond{} }
func (m *AddCommentRespond) String() string { return proto.CompactTextString(m) }
func (*AddCommentRespond) ProtoMessage()    {}
func (*AddCommentRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *AddCommentRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddCommentRespond.Unmarshal(m, b)
}
func (m *AddCommentRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddCommentRespond.Marshal(b, m, deterministic)
}
func (m *AddCommentRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddCommentRespond.Merge(m, src)
}
func (m *AddCommentRespond) XXX_Size() int {
	return xxx_messageInfo_AddCommentRespond.Size(m)
}
func (m *AddCommentRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_AddCommentRespond.DiscardUnknown(m)
}

var xxx_messageInfo_AddCommentRespond proto.InternalMessageInfo

func (m *AddCommentRespond) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type DelCommentRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	ArticleId            int32    `protobuf:"varint,2,opt,name=ArticleId,proto3" json:"ArticleId,omitempty"`
	CommentId            int32    `protobuf:"varint,3,opt,name=CommentId,proto3" json:"CommentId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelCommentRequest) Reset()         { *m = DelCommentRequest{} }
func (m *DelCommentRequest) String() string { return proto.CompactTextString(m) }
func (*DelCommentRequest) ProtoMessage()    {}
func (*DelCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *DelCommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelCommentRequest.Unmarshal(m, b)
}
func (m *DelCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelCommentRequest.Marshal(b, m, deterministic)
}
func (m *DelCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelCommentRequest.Merge(m, src)
}
func (m *DelCommentRequest) XXX_Size() int {
	return xxx_messageInfo_DelCommentRequest.Size(m)
}
func (m *DelCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DelCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DelCommentRequest proto.InternalMessageInfo

func (m *DelCommentRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *DelCommentRequest) GetArticleId() int32 {
	if m != nil {
		return m.ArticleId
	}
	return 0
}

func (m *DelCommentRequest) GetCommentId() int32 {
	if m != nil {
		return m.CommentId
	}
	return 0
}

type DelCommentRespond struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelCommentRespond) Reset()         { *m = DelCommentRespond{} }
func (m *DelCommentRespond) String() string { return proto.CompactTextString(m) }
func (*DelCommentRespond) ProtoMessage()    {}
func (*DelCommentRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *DelCommentRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelCommentRespond.Unmarshal(m, b)
}
func (m *DelCommentRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelCommentRespond.Marshal(b, m, deterministic)
}
func (m *DelCommentRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelCommentRespond.Merge(m, src)
}
func (m *DelCommentRespond) XXX_Size() int {
	return xxx_messageInfo_DelCommentRespond.Size(m)
}
func (m *DelCommentRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_DelCommentRespond.DiscardUnknown(m)
}

var xxx_messageInfo_DelCommentRespond proto.InternalMessageInfo

func (m *DelCommentRespond) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type ListCommentRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	ArticleId            int32    `protobuf:"varint,2,opt,name=ArticleId,proto3" json:"ArticleId,omitempty"`
	Page                 int32    `protobuf:"varint,3,opt,name=Page,proto3" json:"Page,omitempty"`
	Size                 int32    `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCommentRequest) Reset()         { *m = ListCommentRequest{} }
func (m *ListCommentRequest) String() string { return proto.CompactTextString(m) }
func (*ListCommentRequest) ProtoMessage()    {}
func (*ListCommentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *ListCommentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCommentRequest.Unmarshal(m, b)
}
func (m *ListCommentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCommentRequest.Marshal(b, m, deterministic)
}
func (m *ListCommentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCommentRequest.Merge(m, src)
}
func (m *ListCommentRequest) XXX_Size() int {
	return xxx_messageInfo_ListCommentRequest.Size(m)
}
func (m *ListCommentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCommentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCommentRequest proto.InternalMessageInfo

func (m *ListCommentRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ListCommentRequest) GetArticleId() int32 {
	if m != nil {
		return m.ArticleId
	}
	return 0
}

func (m *ListCommentRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListCommentRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type Comment struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=Comment,proto3" json:"Comment,omitempty"`
	Time                 string   `protobuf:"bytes,3,opt,name=Time,proto3" json:"Time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *Comment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Comment.Unmarshal(m, b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return xxx_messageInfo_Comment.Size(m)
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Comment) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Comment) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

type ListCommentRespond struct {
	Count                int32              `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	Lists                map[int32]*Comment `protobuf:"bytes,2,rep,name=Lists,proto3" json:"Lists,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ListCommentRespond) Reset()         { *m = ListCommentRespond{} }
func (m *ListCommentRespond) String() string { return proto.CompactTextString(m) }
func (*ListCommentRespond) ProtoMessage()    {}
func (*ListCommentRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *ListCommentRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCommentRespond.Unmarshal(m, b)
}
func (m *ListCommentRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCommentRespond.Marshal(b, m, deterministic)
}
func (m *ListCommentRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCommentRespond.Merge(m, src)
}
func (m *ListCommentRespond) XXX_Size() int {
	return xxx_messageInfo_ListCommentRespond.Size(m)
}
func (m *ListCommentRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCommentRespond.DiscardUnknown(m)
}

var xxx_messageInfo_ListCommentRespond proto.InternalMessageInfo

func (m *ListCommentRespond) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListCommentRespond) GetLists() map[int32]*Comment {
	if m != nil {
		return m.Lists
	}
	return nil
}

func init() {
	proto.RegisterType((*AddCommentRequest)(nil), "gmicro.srv.comment.AddCommentRequest")
	proto.RegisterType((*AddCommentRespond)(nil), "gmicro.srv.comment.AddCommentRespond")
	proto.RegisterType((*DelCommentRequest)(nil), "gmicro.srv.comment.DelCommentRequest")
	proto.RegisterType((*DelCommentRespond)(nil), "gmicro.srv.comment.DelCommentRespond")
	proto.RegisterType((*ListCommentRequest)(nil), "gmicro.srv.comment.ListCommentRequest")
	proto.RegisterType((*Comment)(nil), "gmicro.srv.comment.Comment")
	proto.RegisterType((*ListCommentRespond)(nil), "gmicro.srv.comment.ListCommentRespond")
	proto.RegisterMapType((map[int32]*Comment)(nil), "gmicro.srv.comment.ListCommentRespond.ListsEntry")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xd1, 0x8a, 0xda, 0x50,
	0x10, 0x35, 0xd1, 0x6b, 0xc9, 0x08, 0x52, 0x87, 0x52, 0x82, 0xf5, 0x41, 0x02, 0xb5, 0x3e, 0x05,
	0xb4, 0x2f, 0xa5, 0x6f, 0x62, 0x5b, 0x09, 0xf4, 0x41, 0x62, 0x7d, 0x2b, 0x94, 0xd4, 0x7b, 0x91,
	0xd0, 0x24, 0xd7, 0x26, 0x51, 0x70, 0xff, 0x60, 0x3f, 0x66, 0x7f, 0x64, 0xbf, 0x6a, 0xb9, 0x93,
	0xab, 0xd1, 0x8d, 0x8b, 0x0b, 0xbe, 0x9d, 0x99, 0xcc, 0x3d, 0x73, 0xe6, 0xcc, 0x10, 0xb0, 0x82,
	0x4d, 0xe8, 0x6e, 0x52, 0x99, 0x4b, 0xc4, 0x75, 0x1c, 0xae, 0x52, 0xe9, 0x66, 0xe9, 0xce, 0x5d,
	0xc9, 0x38, 0x16, 0x49, 0xee, 0xdc, 0x1b, 0xd0, 0x99, 0x70, 0x3e, 0x2d, 0x42, 0x5f, 0xfc, 0xdf,
	0x8a, 0x2c, 0xc7, 0xb7, 0x50, 0x5f, 0x86, 0xdc, 0x36, 0xfa, 0xc6, 0xd0, 0xf2, 0x15, 0xc4, 0x1e,
	0x58, 0x93, 0x34, 0x0f, 0x57, 0x91, 0xf0, 0xb8, 0x6d, 0xf6, 0x8d, 0x21, 0xf3, 0xcb, 0x04, 0xda,
	0xf0, 0x46, 0x33, 0xd8, 0x75, 0x7a, 0x73, 0x08, 0xf1, 0x3d, 0x34, 0x7f, 0xa4, 0x32, 0xf6, 0xb8,
	0xdd, 0xa0, 0x0f, 0x3a, 0x42, 0x84, 0xc6, 0x2f, 0xe9, 0x71, 0x9b, 0x51, 0x96, 0xb0, 0xf3, 0xe9,
	0x5c, 0x4a, 0xb6, 0x91, 0x09, 0x15, 0x4e, 0x25, 0x17, 0xa4, 0x85, 0xf9, 0x84, 0x9d, 0x00, 0x3a,
	0xdf, 0x44, 0x74, 0xa3, 0xe6, 0x1e, 0x58, 0x9a, 0xc1, 0xe3, 0xa4, 0x9a, 0xf9, 0x65, 0x42, 0x69,
	0x39, 0x6d, 0xf1, 0xb2, 0x96, 0x08, 0xf0, 0x67, 0x98, 0xe5, 0x37, 0x8a, 0x41, 0x68, 0xcc, 0x83,
	0xb5, 0xd0, 0x3a, 0x08, 0xab, 0xdc, 0x22, 0xbc, 0x13, 0x64, 0x1c, 0xf3, 0x09, 0x3b, 0xb3, 0xa3,
	0xd1, 0xd8, 0x06, 0xd3, 0xe3, 0x5a, 0x8a, 0x79, 0xbe, 0x03, 0xf3, 0x7c, 0x07, 0xca, 0xeb, 0x30,
	0x16, 0x7a, 0x35, 0x84, 0x9d, 0x47, 0xe3, 0x99, 0xee, 0x62, 0xc2, 0x77, 0xc0, 0xa6, 0x72, 0x9b,
	0xe4, 0x9a, 0xb7, 0x08, 0x70, 0x06, 0x4c, 0xd5, 0x66, 0xb6, 0xd9, 0xaf, 0x0f, 0x5b, 0xe3, 0x91,
	0x5b, 0x3d, 0x24, 0xb7, 0x4a, 0x46, 0xa9, 0xec, 0x7b, 0x92, 0xa7, 0x7b, 0xbf, 0x78, 0xdf, 0x5d,
	0x02, 0x94, 0x49, 0x65, 0xd2, 0x3f, 0xb1, 0xd7, 0xad, 0x14, 0xc4, 0x11, 0xb0, 0x5d, 0x10, 0x6d,
	0x05, 0x4d, 0xd0, 0x1a, 0x7f, 0xb8, 0xd4, 0xe8, 0xd0, 0xa4, 0xa8, 0xfc, 0x6a, 0x7e, 0x31, 0xc6,
	0x0f, 0x26, 0xb4, 0xe7, 0xf2, 0xd8, 0x7f, 0x91, 0xee, 0xf0, 0x37, 0x40, 0x79, 0x4b, 0xf8, 0xf1,
	0x12, 0x51, 0xe5, 0xec, 0xbb, 0x57, 0xcb, 0x68, 0x2e, 0xa7, 0xa6, 0xd8, 0xcb, 0xeb, 0xb8, 0xcc,
	0x5e, 0x39, 0xd0, 0xee, 0xd5, 0xb2, 0x03, 0xfb, 0x1f, 0x68, 0x9d, 0xb8, 0x89, 0x83, 0xab, 0x76,
	0x17, 0xfc, 0x83, 0xd7, 0xad, 0xc5, 0xa9, 0xfd, 0x6d, 0xd2, 0xff, 0xe0, 0xf3, 0x53, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x11, 0xf5, 0x85, 0x9c, 0x1c, 0x04, 0x00, 0x00,
}
