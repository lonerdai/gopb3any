// Code generated by protoc-gen-go.
// source: user/user.proto
// DO NOT EDIT!

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user/user.proto

It has these top-level messages:
	Pro
*/
package user

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Pro struct {
	Name      string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age       int32    `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Languages []string `protobuf:"bytes,3,rep,name=languages" json:"languages,omitempty"`
}

func (m *Pro) Reset()         { *m = Pro{} }
func (m *Pro) String() string { return proto.CompactTextString(m) }
func (*Pro) ProtoMessage()    {}

func init() {
}