// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: msg.proto

package proto

import (
	fmt "fmt"
	"github.com/bbdLe/iGame/comm"
	"github.com/bbdLe/iGame/comm/codec"
	_ "github.com/bbdLe/iGame/comm/codec/gogopb"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	"reflect"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ProtoID int32

const (
	ProtoID_CS_CMD_START               ProtoID = 0
	ProtoID_CS_CMD_LOGIN_REQ           ProtoID = 1
	ProtoID_CS_CMD_LOGIN_RES           ProtoID = 2
	ProtoID_CS_CMD_CHAT_REQ            ProtoID = 3
	ProtoID_CS_CMD_CHAT_RES            ProtoID = 4
	ProtoID_CS_CMD_VERIFY_REQ          ProtoID = 5
	ProtoID_CS_CMD_VERIFY_RES          ProtoID = 6
	ProtoID_CS_CMD_TRANSMIT_REQ        ProtoID = 7
	ProtoID_CS_CMD_TRANSMIT_RES        ProtoID = 8
	ProtoID_CS_CMD_KICK_CONN_REQ       ProtoID = 9
	ProtoID_CS_CMD_KICK_CONN_RES       ProtoID = 10
	ProtoID_CS_CMD_CONN_DISCONNECT_REQ ProtoID = 11
	ProtoID_CS_CMD_CONN_DISCONNECT_RES ProtoID = 12
)

var ProtoID_name = map[int32]string{
	0:  "CS_CMD_START",
	1:  "CS_CMD_LOGIN_REQ",
	2:  "CS_CMD_LOGIN_RES",
	3:  "CS_CMD_CHAT_REQ",
	4:  "CS_CMD_CHAT_RES",
	5:  "CS_CMD_VERIFY_REQ",
	6:  "CS_CMD_VERIFY_RES",
	7:  "CS_CMD_TRANSMIT_REQ",
	8:  "CS_CMD_TRANSMIT_RES",
	9:  "CS_CMD_KICK_CONN_REQ",
	10: "CS_CMD_KICK_CONN_RES",
	11: "CS_CMD_CONN_DISCONNECT_REQ",
	12: "CS_CMD_CONN_DISCONNECT_RES",
}

var ProtoID_value = map[string]int32{
	"CS_CMD_START":               0,
	"CS_CMD_LOGIN_REQ":           1,
	"CS_CMD_LOGIN_RES":           2,
	"CS_CMD_CHAT_REQ":            3,
	"CS_CMD_CHAT_RES":            4,
	"CS_CMD_VERIFY_REQ":          5,
	"CS_CMD_VERIFY_RES":          6,
	"CS_CMD_TRANSMIT_REQ":        7,
	"CS_CMD_TRANSMIT_RES":        8,
	"CS_CMD_KICK_CONN_REQ":       9,
	"CS_CMD_KICK_CONN_RES":       10,
	"CS_CMD_CONN_DISCONNECT_REQ": 11,
	"CS_CMD_CONN_DISCONNECT_RES": 12,
}

func (x ProtoID) String() string {
	return proto.EnumName(ProtoID_name, int32(x))
}

func (ProtoID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

//[id:CS_CMD_LOGIN_REQ]
type LoginReq struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Platform             string   `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	Uid                  string   `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Channel              string   `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel,omitempty"`
	AuthTime             int64    `protobuf:"varint,5,opt,name=AuthTime,proto3" json:"AuthTime,omitempty"`
	Token                string   `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}
func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *LoginReq) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *LoginReq) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *LoginReq) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *LoginReq) GetAuthTime() int64 {
	if m != nil {
		return m.AuthTime
	}
	return 0
}

func (m *LoginReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// [id:CS_CMD_LOGIN_RES]
type LoginRes struct {
	RetCode              int32    `protobuf:"varint,1,opt,name=ret_code,json=retCode,proto3" json:"ret_code,omitempty"`
	RetMsg               string   `protobuf:"bytes,2,opt,name=ret_msg,json=retMsg,proto3" json:"ret_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRes) Reset()         { *m = LoginRes{} }
func (m *LoginRes) String() string { return proto.CompactTextString(m) }
func (*LoginRes) ProtoMessage()    {}
func (*LoginRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}
func (m *LoginRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRes.Unmarshal(m, b)
}
func (m *LoginRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRes.Marshal(b, m, deterministic)
}
func (m *LoginRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRes.Merge(m, src)
}
func (m *LoginRes) XXX_Size() int {
	return xxx_messageInfo_LoginRes.Size(m)
}
func (m *LoginRes) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRes.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRes proto.InternalMessageInfo

func (m *LoginRes) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *LoginRes) GetRetMsg() string {
	if m != nil {
		return m.RetMsg
	}
	return ""
}

// [id:CS_CMD_VERIFY_REQ]
type VerifyReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Server               string   `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyReq) Reset()         { *m = VerifyReq{} }
func (m *VerifyReq) String() string { return proto.CompactTextString(m) }
func (*VerifyReq) ProtoMessage()    {}
func (*VerifyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}
func (m *VerifyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyReq.Unmarshal(m, b)
}
func (m *VerifyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyReq.Marshal(b, m, deterministic)
}
func (m *VerifyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyReq.Merge(m, src)
}
func (m *VerifyReq) XXX_Size() int {
	return xxx_messageInfo_VerifyReq.Size(m)
}
func (m *VerifyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyReq.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyReq proto.InternalMessageInfo

func (m *VerifyReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *VerifyReq) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

// [id:CS_CMD_VERIFY_RES]
type VerifyRes struct {
	RetCode              int32    `protobuf:"varint,1,opt,name=ret_code,json=retCode,proto3" json:"ret_code,omitempty"`
	RetMsg               string   `protobuf:"bytes,2,opt,name=ret_msg,json=retMsg,proto3" json:"ret_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyRes) Reset()         { *m = VerifyRes{} }
func (m *VerifyRes) String() string { return proto.CompactTextString(m) }
func (*VerifyRes) ProtoMessage()    {}
func (*VerifyRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
}
func (m *VerifyRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyRes.Unmarshal(m, b)
}
func (m *VerifyRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyRes.Marshal(b, m, deterministic)
}
func (m *VerifyRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyRes.Merge(m, src)
}
func (m *VerifyRes) XXX_Size() int {
	return xxx_messageInfo_VerifyRes.Size(m)
}
func (m *VerifyRes) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyRes.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyRes proto.InternalMessageInfo

func (m *VerifyRes) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *VerifyRes) GetRetMsg() string {
	if m != nil {
		return m.RetMsg
	}
	return ""
}

// [id:CS_CMD_TRANSMIT_REQ]
type TransmitReq struct {
	MsgId                int32    `protobuf:"varint,1,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	MsgData              []byte   `protobuf:"bytes,2,opt,name=msg_data,json=msgData,proto3" json:"msg_data,omitempty"`
	ClientId             int64    `protobuf:"varint,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransmitReq) Reset()         { *m = TransmitReq{} }
func (m *TransmitReq) String() string { return proto.CompactTextString(m) }
func (*TransmitReq) ProtoMessage()    {}
func (*TransmitReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4}
}
func (m *TransmitReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransmitReq.Unmarshal(m, b)
}
func (m *TransmitReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransmitReq.Marshal(b, m, deterministic)
}
func (m *TransmitReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransmitReq.Merge(m, src)
}
func (m *TransmitReq) XXX_Size() int {
	return xxx_messageInfo_TransmitReq.Size(m)
}
func (m *TransmitReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TransmitReq.DiscardUnknown(m)
}

var xxx_messageInfo_TransmitReq proto.InternalMessageInfo

func (m *TransmitReq) GetMsgId() int32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *TransmitReq) GetMsgData() []byte {
	if m != nil {
		return m.MsgData
	}
	return nil
}

func (m *TransmitReq) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

// [id:CS_CMD_TRANSMIT_RES]
type TransmitRes struct {
	MsgId                int32    `protobuf:"varint,1,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	MsgData              []byte   `protobuf:"bytes,2,opt,name=msg_data,json=msgData,proto3" json:"msg_data,omitempty"`
	ClientId             int64    `protobuf:"varint,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransmitRes) Reset()         { *m = TransmitRes{} }
func (m *TransmitRes) String() string { return proto.CompactTextString(m) }
func (*TransmitRes) ProtoMessage()    {}
func (*TransmitRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{5}
}
func (m *TransmitRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransmitRes.Unmarshal(m, b)
}
func (m *TransmitRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransmitRes.Marshal(b, m, deterministic)
}
func (m *TransmitRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransmitRes.Merge(m, src)
}
func (m *TransmitRes) XXX_Size() int {
	return xxx_messageInfo_TransmitRes.Size(m)
}
func (m *TransmitRes) XXX_DiscardUnknown() {
	xxx_messageInfo_TransmitRes.DiscardUnknown(m)
}

var xxx_messageInfo_TransmitRes proto.InternalMessageInfo

func (m *TransmitRes) GetMsgId() int32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *TransmitRes) GetMsgData() []byte {
	if m != nil {
		return m.MsgData
	}
	return nil
}

func (m *TransmitRes) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

// [id:CS_CMD_CHAT_REQ ]
type ChatReq struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatReq) Reset()         { *m = ChatReq{} }
func (m *ChatReq) String() string { return proto.CompactTextString(m) }
func (*ChatReq) ProtoMessage()    {}
func (*ChatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{6}
}
func (m *ChatReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatReq.Unmarshal(m, b)
}
func (m *ChatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatReq.Marshal(b, m, deterministic)
}
func (m *ChatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatReq.Merge(m, src)
}
func (m *ChatReq) XXX_Size() int {
	return xxx_messageInfo_ChatReq.Size(m)
}
func (m *ChatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChatReq proto.InternalMessageInfo

func (m *ChatReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// [id:CS_CMD_CHAT_RES]
type ChatRes struct {
	RetCode              int32    `protobuf:"varint,1,opt,name=ret_code,json=retCode,proto3" json:"ret_code,omitempty"`
	RetMsg               string   `protobuf:"bytes,2,opt,name=ret_msg,json=retMsg,proto3" json:"ret_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatRes) Reset()         { *m = ChatRes{} }
func (m *ChatRes) String() string { return proto.CompactTextString(m) }
func (*ChatRes) ProtoMessage()    {}
func (*ChatRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{7}
}
func (m *ChatRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatRes.Unmarshal(m, b)
}
func (m *ChatRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatRes.Marshal(b, m, deterministic)
}
func (m *ChatRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatRes.Merge(m, src)
}
func (m *ChatRes) XXX_Size() int {
	return xxx_messageInfo_ChatRes.Size(m)
}
func (m *ChatRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatRes.DiscardUnknown(m)
}

var xxx_messageInfo_ChatRes proto.InternalMessageInfo

func (m *ChatRes) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *ChatRes) GetRetMsg() string {
	if m != nil {
		return m.RetMsg
	}
	return ""
}

// [id:CS_CMD_KICK_CONN_REQ]
type KickConnReq struct {
	ClientId             int64    `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KickConnReq) Reset()         { *m = KickConnReq{} }
func (m *KickConnReq) String() string { return proto.CompactTextString(m) }
func (*KickConnReq) ProtoMessage()    {}
func (*KickConnReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{8}
}
func (m *KickConnReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KickConnReq.Unmarshal(m, b)
}
func (m *KickConnReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KickConnReq.Marshal(b, m, deterministic)
}
func (m *KickConnReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KickConnReq.Merge(m, src)
}
func (m *KickConnReq) XXX_Size() int {
	return xxx_messageInfo_KickConnReq.Size(m)
}
func (m *KickConnReq) XXX_DiscardUnknown() {
	xxx_messageInfo_KickConnReq.DiscardUnknown(m)
}

var xxx_messageInfo_KickConnReq proto.InternalMessageInfo

func (m *KickConnReq) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

// [id:CS_CMD_KICK_CONN_RES]
type KickConnRes struct {
	ClientId             int64    `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	RetCode              int32    `protobuf:"varint,2,opt,name=ret_code,json=retCode,proto3" json:"ret_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KickConnRes) Reset()         { *m = KickConnRes{} }
func (m *KickConnRes) String() string { return proto.CompactTextString(m) }
func (*KickConnRes) ProtoMessage()    {}
func (*KickConnRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{9}
}
func (m *KickConnRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KickConnRes.Unmarshal(m, b)
}
func (m *KickConnRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KickConnRes.Marshal(b, m, deterministic)
}
func (m *KickConnRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KickConnRes.Merge(m, src)
}
func (m *KickConnRes) XXX_Size() int {
	return xxx_messageInfo_KickConnRes.Size(m)
}
func (m *KickConnRes) XXX_DiscardUnknown() {
	xxx_messageInfo_KickConnRes.DiscardUnknown(m)
}

var xxx_messageInfo_KickConnRes proto.InternalMessageInfo

func (m *KickConnRes) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *KickConnRes) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

// [id:CS_CMD_CONN_DISCONNECT_REQ]
type ConnDisconnectReq struct {
	ClientId             int64    `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnDisconnectReq) Reset()         { *m = ConnDisconnectReq{} }
func (m *ConnDisconnectReq) String() string { return proto.CompactTextString(m) }
func (*ConnDisconnectReq) ProtoMessage()    {}
func (*ConnDisconnectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{10}
}
func (m *ConnDisconnectReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnDisconnectReq.Unmarshal(m, b)
}
func (m *ConnDisconnectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnDisconnectReq.Marshal(b, m, deterministic)
}
func (m *ConnDisconnectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnDisconnectReq.Merge(m, src)
}
func (m *ConnDisconnectReq) XXX_Size() int {
	return xxx_messageInfo_ConnDisconnectReq.Size(m)
}
func (m *ConnDisconnectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnDisconnectReq.DiscardUnknown(m)
}

var xxx_messageInfo_ConnDisconnectReq proto.InternalMessageInfo

func (m *ConnDisconnectReq) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

// [id:CS_CMD_CONN_DISCONNECT_RES]
type ConnDisconnectRes struct {
	ClientId             int64    `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	RetCode              int32    `protobuf:"varint,2,opt,name=ret_code,json=retCode,proto3" json:"ret_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnDisconnectRes) Reset()         { *m = ConnDisconnectRes{} }
func (m *ConnDisconnectRes) String() string { return proto.CompactTextString(m) }
func (*ConnDisconnectRes) ProtoMessage()    {}
func (*ConnDisconnectRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{11}
}
func (m *ConnDisconnectRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnDisconnectRes.Unmarshal(m, b)
}
func (m *ConnDisconnectRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnDisconnectRes.Marshal(b, m, deterministic)
}
func (m *ConnDisconnectRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnDisconnectRes.Merge(m, src)
}
func (m *ConnDisconnectRes) XXX_Size() int {
	return xxx_messageInfo_ConnDisconnectRes.Size(m)
}
func (m *ConnDisconnectRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnDisconnectRes.DiscardUnknown(m)
}

var xxx_messageInfo_ConnDisconnectRes proto.InternalMessageInfo

func (m *ConnDisconnectRes) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *ConnDisconnectRes) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func init() {
	proto.RegisterEnum("proto.ProtoID", ProtoID_name, ProtoID_value)
	proto.RegisterType((*LoginReq)(nil), "proto.LoginReq")
	proto.RegisterType((*LoginRes)(nil), "proto.LoginRes")
	proto.RegisterType((*VerifyReq)(nil), "proto.VerifyReq")
	proto.RegisterType((*VerifyRes)(nil), "proto.VerifyRes")
	proto.RegisterType((*TransmitReq)(nil), "proto.TransmitReq")
	proto.RegisterType((*TransmitRes)(nil), "proto.TransmitRes")
	proto.RegisterType((*ChatReq)(nil), "proto.ChatReq")
	proto.RegisterType((*ChatRes)(nil), "proto.ChatRes")
	proto.RegisterType((*KickConnReq)(nil), "proto.KickConnReq")
	proto.RegisterType((*KickConnRes)(nil), "proto.KickConnRes")
	proto.RegisterType((*ConnDisconnectReq)(nil), "proto.ConnDisconnectReq")
	proto.RegisterType((*ConnDisconnectRes)(nil), "proto.ConnDisconnectRes")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x25, 0xed, 0x92, 0xb4, 0xb7, 0x95, 0xf0, 0xbc, 0x8e, 0x85, 0x21, 0xa1, 0x29, 0xbc, 0x4c,
	0x7b, 0x40, 0x48, 0x3c, 0xf1, 0x00, 0xa8, 0x4a, 0x0b, 0x44, 0x5d, 0x3b, 0x48, 0xa2, 0x49, 0x48,
	0x48, 0x51, 0x48, 0xbc, 0x34, 0x5a, 0x63, 0x0f, 0xdb, 0x9b, 0xc4, 0xcf, 0xf0, 0x37, 0xfc, 0x17,
	0xb2, 0xe3, 0x96, 0x95, 0x95, 0x21, 0x55, 0x3c, 0xd9, 0xe7, 0xdc, 0x7b, 0xcf, 0x39, 0xb5, 0x6e,
	0x03, 0xdd, 0x5a, 0x94, 0xcf, 0xaf, 0x38, 0x93, 0x0c, 0xdb, 0xfa, 0xf0, 0x7f, 0x58, 0xd0, 0x39,
	0x65, 0x65, 0x45, 0x23, 0xf2, 0x0d, 0x7b, 0xe0, 0xde, 0x10, 0x2e, 0x2a, 0x46, 0x3d, 0xeb, 0xc8,
	0x3a, 0xee, 0x46, 0x4b, 0x88, 0x0f, 0xa1, 0x73, 0xb5, 0xc8, 0xe4, 0x05, 0xe3, 0xb5, 0xd7, 0xd2,
	0xa5, 0x15, 0xc6, 0x08, 0xda, 0xd7, 0x55, 0xe1, 0xb5, 0x35, 0xad, 0xae, 0x4a, 0x27, 0x9f, 0x67,
	0x94, 0x92, 0x85, 0xb7, 0xd3, 0xe8, 0x18, 0xa8, 0x74, 0x86, 0xd7, 0x72, 0x9e, 0x54, 0x35, 0xf1,
	0xec, 0x23, 0xeb, 0xb8, 0x1d, 0xad, 0x30, 0x1e, 0x80, 0x2d, 0xd9, 0x25, 0xa1, 0x9e, 0xa3, 0x67,
	0x1a, 0xe0, 0xbf, 0x59, 0xe5, 0x13, 0xf8, 0x31, 0x74, 0x38, 0x91, 0x69, 0xce, 0x0a, 0xa2, 0x03,
	0xda, 0x91, 0xcb, 0x89, 0x0c, 0x58, 0x41, 0xf0, 0x01, 0xa8, 0x6b, 0x5a, 0x8b, 0xd2, 0xe4, 0x73,
	0x38, 0x91, 0x53, 0x51, 0xfa, 0xaf, 0xa0, 0x7b, 0x4e, 0x78, 0x75, 0xf1, 0x5d, 0xfd, 0xc0, 0x95,
	0x85, 0x75, 0xcb, 0x02, 0x3f, 0x02, 0x47, 0x10, 0x7e, 0x43, 0xf8, 0x72, 0xb4, 0x41, 0xfe, 0xdb,
	0xdf, 0xa3, 0xdb, 0x79, 0x7f, 0x81, 0x5e, 0xc2, 0x33, 0x2a, 0xea, 0x4a, 0x2a, 0xf7, 0x7d, 0x70,
	0x6a, 0x51, 0xa6, 0x55, 0x61, 0x04, 0xec, 0x5a, 0x94, 0x61, 0xa1, 0x94, 0x15, 0x5d, 0x64, 0x32,
	0xd3, 0xf3, 0xfd, 0xc8, 0xad, 0x45, 0x39, 0xca, 0x64, 0x86, 0x9f, 0x40, 0x37, 0x5f, 0x54, 0x84,
	0xca, 0xd4, 0x3c, 0x70, 0x3b, 0xea, 0x34, 0x44, 0x58, 0xac, 0xab, 0x8b, 0xff, 0xad, 0xfe, 0x0c,
	0xdc, 0x60, 0x9e, 0x49, 0xb3, 0x16, 0x39, 0xa3, 0x92, 0x50, 0xb9, 0x5c, 0x0b, 0x03, 0xfd, 0xd7,
	0xcb, 0xa6, 0xed, 0xde, 0xe7, 0x04, 0x7a, 0x93, 0x2a, 0xbf, 0x0c, 0x18, 0xd5, 0xeb, 0xb7, 0x96,
	0xc7, 0xfa, 0x23, 0xcf, 0xf8, 0x76, 0xaf, 0xb8, 0xb7, 0x77, 0x2d, 0x4b, 0x6b, 0x2d, 0x8b, 0xff,
	0x02, 0x76, 0x95, 0xc4, 0xa8, 0x12, 0x39, 0xa3, 0x94, 0xe4, 0xf2, 0x9f, 0xc6, 0x93, 0xbb, 0x13,
	0x5b, 0xdb, 0x9f, 0xfc, 0x6c, 0x81, 0xfb, 0x51, 0xfd, 0xf1, 0xc2, 0x11, 0x46, 0xd0, 0x0f, 0xe2,
	0x34, 0x98, 0x8e, 0xd2, 0x38, 0x19, 0x46, 0x09, 0x7a, 0x80, 0x07, 0x80, 0x0c, 0x73, 0x7a, 0xf6,
	0x3e, 0x9c, 0xa5, 0xd1, 0xf8, 0x13, 0xb2, 0x36, 0xb0, 0x31, 0x6a, 0xe1, 0x3d, 0x78, 0x68, 0xd8,
	0xe0, 0xc3, 0x30, 0xd1, 0xad, 0xed, 0xbb, 0x64, 0x8c, 0x76, 0xf0, 0x3e, 0xec, 0x1a, 0xf2, 0x7c,
	0x1c, 0x85, 0xef, 0x3e, 0xeb, 0x5e, 0x7b, 0x13, 0x1d, 0x23, 0x07, 0x1f, 0xc0, 0x9e, 0xa1, 0x93,
	0x68, 0x38, 0x8b, 0xa7, 0x61, 0xa3, 0xed, 0x6e, 0x2e, 0xc4, 0xa8, 0x83, 0x3d, 0x18, 0x98, 0xc2,
	0x24, 0x0c, 0x26, 0x69, 0x70, 0x36, 0x6b, 0x92, 0x77, 0xff, 0x52, 0x89, 0x11, 0xe0, 0xa7, 0x70,
	0xb8, 0x0c, 0xaa, 0xc8, 0x51, 0x18, 0xab, 0x73, 0x1c, 0x34, 0x66, 0xbd, 0x7b, 0xeb, 0x31, 0xea,
	0x7f, 0x75, 0xf4, 0xd7, 0xeb, 0xe5, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0xb1, 0x26, 0x1b,
	0xd1, 0x04, 0x00, 0x00,
}

func init() {

	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: int(ProtoID_CS_CMD_LOGIN_REQ),
		Type:  reflect.TypeOf((*LoginReq)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})

	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: int(ProtoID_CS_CMD_LOGIN_RES),
		Type:  reflect.TypeOf((*LoginRes)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})

	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: int(ProtoID_CS_CMD_VERIFY_REQ),
		Type:  reflect.TypeOf((*VerifyReq)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})

	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: int(ProtoID_CS_CMD_VERIFY_RES),
		Type:  reflect.TypeOf((*VerifyRes)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})

	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: int(ProtoID_CS_CMD_TRANSMIT_REQ),
		Type:  reflect.TypeOf((*TransmitReq)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})

	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: int(ProtoID_CS_CMD_TRANSMIT_RES),
		Type:  reflect.TypeOf((*TransmitRes)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})

	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: int(ProtoID_CS_CMD_CHAT_REQ),
		Type:  reflect.TypeOf((*ChatReq)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})

	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: int(ProtoID_CS_CMD_CHAT_RES),
		Type:  reflect.TypeOf((*ChatRes)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})

	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: int(ProtoID_CS_CMD_KICK_CONN_REQ),
		Type:  reflect.TypeOf((*KickConnReq)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})

	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: int(ProtoID_CS_CMD_KICK_CONN_RES),
		Type:  reflect.TypeOf((*KickConnRes)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})

	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: int(ProtoID_CS_CMD_CONN_DISCONNECT_REQ),
		Type:  reflect.TypeOf((*ConnDisconnectReq)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})

	comm.RegMessageMeta(&comm.MessageMeta{
		MsgId: int(ProtoID_CS_CMD_CONN_DISCONNECT_RES),
		Type:  reflect.TypeOf((*ConnDisconnectRes)(nil)).Elem(),
		Codec: codec.MustGetCodec("gogopb"),
	})

}
