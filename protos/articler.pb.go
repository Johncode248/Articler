// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: articler.proto

package protos

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_articler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_articler_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type LoginForm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginForm) Reset() {
	*x = LoginForm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginForm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginForm) ProtoMessage() {}

func (x *LoginForm) ProtoReflect() protoreflect.Message {
	mi := &file_articler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginForm.ProtoReflect.Descriptor instead.
func (*LoginForm) Descriptor() ([]byte, []int) {
	return file_articler_proto_rawDescGZIP(), []int{1}
}

func (x *LoginForm) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginForm) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ArticleForm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content      string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	ShortContent string `protobuf:"bytes,3,opt,name=shortContent,proto3" json:"shortContent,omitempty"`
	CreateTime   string `protobuf:"bytes,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	AuthorId     uint64 `protobuf:"varint,5,opt,name=authorId,proto3" json:"authorId,omitempty"`
	Token        string `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ArticleForm) Reset() {
	*x = ArticleForm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleForm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleForm) ProtoMessage() {}

func (x *ArticleForm) ProtoReflect() protoreflect.Message {
	mi := &file_articler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleForm.ProtoReflect.Descriptor instead.
func (*ArticleForm) Descriptor() ([]byte, []int) {
	return file_articler_proto_rawDescGZIP(), []int{2}
}

func (x *ArticleForm) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ArticleForm) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ArticleForm) GetShortContent() string {
	if x != nil {
		return x.ShortContent
	}
	return ""
}

func (x *ArticleForm) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *ArticleForm) GetAuthorId() uint64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *ArticleForm) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UpdateUserForm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UpdateUserForm) Reset() {
	*x = UpdateUserForm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserForm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserForm) ProtoMessage() {}

func (x *UpdateUserForm) ProtoReflect() protoreflect.Message {
	mi := &file_articler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserForm.ProtoReflect.Descriptor instead.
func (*UpdateUserForm) Descriptor() ([]byte, []int) {
	return file_articler_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateUserForm) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateUserForm) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateUserForm) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DelateArticleForm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdArticle string `protobuf:"bytes,1,opt,name=idArticle,proto3" json:"idArticle,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DelateArticleForm) Reset() {
	*x = DelateArticleForm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelateArticleForm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelateArticleForm) ProtoMessage() {}

func (x *DelateArticleForm) ProtoReflect() protoreflect.Message {
	mi := &file_articler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelateArticleForm.ProtoReflect.Descriptor instead.
func (*DelateArticleForm) Descriptor() ([]byte, []int) {
	return file_articler_proto_rawDescGZIP(), []int{4}
}

func (x *DelateArticleForm) GetIdArticle() string {
	if x != nil {
		return x.IdArticle
	}
	return ""
}

func (x *DelateArticleForm) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_articler_proto protoreflect.FileDescriptor

var file_articler_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22,
	0x43, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0xb3, 0x01, 0x0a, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x46, 0x6f, 0x72, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5e, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x47, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x32, 0xf3, 0x02, 0x0a, 0x0f, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x08, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x0a, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x1a, 0x08,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x08, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x46, 0x6f,
	0x72, 0x6d, 0x1a, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x29,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x6d, 0x1a, 0x08, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0c,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x1a, 0x08, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0c, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x1a, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x1a, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x12, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x08, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_articler_proto_rawDescOnce sync.Once
	file_articler_proto_rawDescData = file_articler_proto_rawDesc
)

func file_articler_proto_rawDescGZIP() []byte {
	file_articler_proto_rawDescOnce.Do(func() {
		file_articler_proto_rawDescData = protoimpl.X.CompressGZIP(file_articler_proto_rawDescData)
	})
	return file_articler_proto_rawDescData
}

var file_articler_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_articler_proto_goTypes = []interface{}{
	(*Message)(nil),           // 0: Message
	(*LoginForm)(nil),         // 1: LoginForm
	(*ArticleForm)(nil),       // 2: ArticleForm
	(*UpdateUserForm)(nil),    // 3: UpdateUserForm
	(*DelateArticleForm)(nil), // 4: DelateArticleForm
}
var file_articler_proto_depIdxs = []int32{
	0, // 0: ArticlerService.SayHello:input_type -> Message
	1, // 1: ArticlerService.Login:input_type -> LoginForm
	1, // 2: ArticlerService.Register:input_type -> LoginForm
	3, // 3: ArticlerService.UpdateUser:input_type -> UpdateUserForm
	0, // 4: ArticlerService.DeleteUser:input_type -> Message
	2, // 5: ArticlerService.CreateArticle:input_type -> ArticleForm
	2, // 6: ArticlerService.UpdateArticle:input_type -> ArticleForm
	4, // 7: ArticlerService.DeleteArticle:input_type -> DelateArticleForm
	0, // 8: ArticlerService.GetArticles:input_type -> Message
	0, // 9: ArticlerService.SayHello:output_type -> Message
	0, // 10: ArticlerService.Login:output_type -> Message
	0, // 11: ArticlerService.Register:output_type -> Message
	0, // 12: ArticlerService.UpdateUser:output_type -> Message
	0, // 13: ArticlerService.DeleteUser:output_type -> Message
	0, // 14: ArticlerService.CreateArticle:output_type -> Message
	0, // 15: ArticlerService.UpdateArticle:output_type -> Message
	0, // 16: ArticlerService.DeleteArticle:output_type -> Message
	0, // 17: ArticlerService.GetArticles:output_type -> Message
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_articler_proto_init() }
func file_articler_proto_init() {
	if File_articler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_articler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_articler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginForm); i {
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
		file_articler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleForm); i {
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
		file_articler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserForm); i {
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
		file_articler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelateArticleForm); i {
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
			RawDescriptor: file_articler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_articler_proto_goTypes,
		DependencyIndexes: file_articler_proto_depIdxs,
		MessageInfos:      file_articler_proto_msgTypes,
	}.Build()
	File_articler_proto = out.File
	file_articler_proto_rawDesc = nil
	file_articler_proto_goTypes = nil
	file_articler_proto_depIdxs = nil
}
