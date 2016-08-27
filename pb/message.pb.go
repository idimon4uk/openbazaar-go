// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Message
	Envelope
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Message_MessageType int32

const (
	Message_PING               Message_MessageType = 0
	Message_MESSAGE            Message_MessageType = 1
	Message_FOLLOW             Message_MessageType = 2
	Message_UNFOLLOW           Message_MessageType = 3
	Message_ORDER              Message_MessageType = 4
	Message_ORDER_CONFIRMATION Message_MessageType = 5
	Message_ORDER_FULFILLMENT  Message_MessageType = 6
	Message_RATING             Message_MessageType = 7
	Message_DISPUTE_OPEN       Message_MessageType = 8
	Message_DISPUTE_CLOSE      Message_MessageType = 9
	Message_REFUND             Message_MessageType = 10
	Message_OFFLINE_ACK        Message_MessageType = 11
)

var Message_MessageType_name = map[int32]string{
	0:  "PING",
	1:  "MESSAGE",
	2:  "FOLLOW",
	3:  "UNFOLLOW",
	4:  "ORDER",
	5:  "ORDER_CONFIRMATION",
	6:  "ORDER_FULFILLMENT",
	7:  "RATING",
	8:  "DISPUTE_OPEN",
	9:  "DISPUTE_CLOSE",
	10: "REFUND",
	11: "OFFLINE_ACK",
}
var Message_MessageType_value = map[string]int32{
	"PING":               0,
	"MESSAGE":            1,
	"FOLLOW":             2,
	"UNFOLLOW":           3,
	"ORDER":              4,
	"ORDER_CONFIRMATION": 5,
	"ORDER_FULFILLMENT":  6,
	"RATING":             7,
	"DISPUTE_OPEN":       8,
	"DISPUTE_CLOSE":      9,
	"REFUND":             10,
	"OFFLINE_ACK":        11,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (Message_MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

type Message struct {
	MessageType Message_MessageType  `protobuf:"varint,1,opt,name=messageType,enum=Message_MessageType" json:"messageType,omitempty"`
	Payload     *google_protobuf.Any `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Message) GetPayload() *google_protobuf.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Envelope struct {
	Message   *Message `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	PeerID    string   `protobuf:"bytes,2,opt,name=peerID" json:"peerID,omitempty"`
	Signature []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Envelope) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*Envelope)(nil), "Envelope")
	proto.RegisterEnum("Message_MessageType", Message_MessageType_name, Message_MessageType_value)
}

var fileDescriptor2 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x50, 0xcb, 0x6a, 0xe3, 0x30,
	0x14, 0x1d, 0xe7, 0xe1, 0xc7, 0x75, 0x32, 0xa3, 0x88, 0x99, 0x21, 0x2d, 0x5d, 0x14, 0xaf, 0xba,
	0x72, 0x20, 0x85, 0xee, 0x4d, 0x2c, 0x07, 0x53, 0x45, 0x0a, 0xb2, 0x4d, 0x97, 0xc1, 0x21, 0xaa,
	0x29, 0xa4, 0xb6, 0xc9, 0xa3, 0xe0, 0x5f, 0xeb, 0x07, 0xf4, 0xbb, 0xaa, 0xf8, 0x41, 0xb3, 0xb2,
	0xef, 0x79, 0xdc, 0x73, 0x75, 0x60, 0xfc, 0x2e, 0x8f, 0xc7, 0x34, 0x93, 0x6e, 0x79, 0x28, 0x4e,
	0xc5, 0xed, 0x4d, 0x56, 0x14, 0xd9, 0x5e, 0xce, 0xea, 0x69, 0x7b, 0x7e, 0x9d, 0xa5, 0x79, 0xd5,
	0x50, 0xce, 0x67, 0x0f, 0x8c, 0x55, 0x23, 0xc6, 0x4f, 0x60, 0xb7, 0xbe, 0xb8, 0x2a, 0xe5, 0x54,
	0xbb, 0xd7, 0x1e, 0x7e, 0xcf, 0xff, 0xba, 0x2d, 0xdd, 0x7d, 0x2f, 0x9c, 0xb8, 0x16, 0x62, 0x17,
	0x8c, 0x32, 0xad, 0xf6, 0x45, 0xba, 0x9b, 0xf6, 0x94, 0xc7, 0x56, 0x9e, 0x26, 0xd0, 0xed, 0x02,
	0x5d, 0x2f, 0xaf, 0x44, 0x27, 0x72, 0xbe, 0x34, 0xb0, 0xaf, 0x96, 0x61, 0x13, 0x06, 0xeb, 0x90,
	0x2d, 0xd1, 0x2f, 0x6c, 0xab, 0x63, 0x48, 0x14, 0x79, 0x4b, 0x82, 0x34, 0x0c, 0xa0, 0x07, 0x9c,
	0x52, 0xfe, 0x82, 0x7a, 0x78, 0x04, 0x66, 0xc2, 0xda, 0xa9, 0x8f, 0x2d, 0x18, 0x72, 0xe1, 0x13,
	0x81, 0x06, 0xf8, 0x3f, 0xe0, 0xfa, 0x77, 0xb3, 0xe0, 0x2c, 0x08, 0xc5, 0xca, 0x8b, 0x43, 0xce,
	0xd0, 0x10, 0xff, 0x83, 0x49, 0x83, 0x07, 0x09, 0x0d, 0x42, 0x4a, 0x57, 0x84, 0xc5, 0x48, 0xbf,
	0xec, 0x14, 0x4a, 0xa2, 0xc2, 0x0c, 0x8c, 0x60, 0xe4, 0x87, 0xd1, 0x3a, 0x89, 0xc9, 0x86, 0xaf,
	0x09, 0x43, 0x26, 0x9e, 0xc0, 0xb8, 0x43, 0x16, 0x94, 0x47, 0x04, 0x59, 0xb5, 0x81, 0x04, 0x09,
	0xf3, 0x11, 0xe0, 0x3f, 0x60, 0xf3, 0x20, 0xa0, 0x21, 0x23, 0x1b, 0x6f, 0xf1, 0x8c, 0x6c, 0x67,
	0x07, 0x26, 0xc9, 0x3f, 0xe4, 0xbe, 0x50, 0x8f, 0x70, 0xc0, 0x68, 0x3b, 0xa9, 0x8b, 0xb3, 0xe7,
	0x66, 0x57, 0x98, 0xe8, 0x08, 0x75, 0xac, 0x5e, 0x4a, 0x79, 0x08, 0xfd, 0xba, 0x27, 0x4b, 0xb4,
	0x13, 0xbe, 0x03, 0xeb, 0xf8, 0x96, 0xe5, 0xe9, 0xe9, 0x7c, 0x90, 0xd3, 0xbe, 0xa2, 0x46, 0xe2,
	0x07, 0xd8, 0xea, 0x75, 0x8b, 0x8f, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x43, 0x43, 0x7d,
	0xd5, 0x01, 0x00, 0x00,
}
