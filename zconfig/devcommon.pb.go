// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devcommon.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// XXX duplicate of definition in zmet.proto with uniq names (ZCio vs Zio)
// Types of I/O adapters that can be assigned to applications
type ZCioType int32

const (
	ZCioType_ZCioNop   ZCioType = 0
	ZCioType_ZCioEth   ZCioType = 1
	ZCioType_ZCioUSB   ZCioType = 2
	ZCioType_ZCioCOM   ZCioType = 3
	ZCioType_ZCioHDMI  ZCioType = 4
	ZCioType_ZCioOther ZCioType = 255
)

var ZCioType_name = map[int32]string{
	0:   "ZCioNop",
	1:   "ZCioEth",
	2:   "ZCioUSB",
	3:   "ZCioCOM",
	4:   "ZCioHDMI",
	255: "ZCioOther",
}
var ZCioType_value = map[string]int32{
	"ZCioNop":   0,
	"ZCioEth":   1,
	"ZCioUSB":   2,
	"ZCioCOM":   3,
	"ZCioHDMI":  4,
	"ZCioOther": 255,
}

func (x ZCioType) String() string {
	return proto.EnumName(ZCioType_name, int32(x))
}
func (ZCioType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type ZcServiceType int32

const (
	ZcServiceType_zcloudInvalidSrv ZcServiceType = 0
	// mapping service for zededa overlay service
	ZcServiceType_mapServer ZcServiceType = 1
	// if device has support feature enabled, this is cloud service from
	// device can be reached
	ZcServiceType_supportServer ZcServiceType = 2
)

var ZcServiceType_name = map[int32]string{
	0: "zcloudInvalidSrv",
	1: "mapServer",
	2: "supportServer",
}
var ZcServiceType_value = map[string]int32{
	"zcloudInvalidSrv": 0,
	"mapServer":        1,
	"supportServer":    2,
}

func (x ZcServiceType) String() string {
	return proto.EnumName(ZcServiceType_name, int32(x))
}
func (ZcServiceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type UUIDandVersion struct {
	Uuid    string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *UUIDandVersion) Reset()                    { *m = UUIDandVersion{} }
func (m *UUIDandVersion) String() string            { return proto.CompactTextString(m) }
func (*UUIDandVersion) ProtoMessage()               {}
func (*UUIDandVersion) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *UUIDandVersion) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *UUIDandVersion) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Adapter bundles corresponding to a subset of what is in ZioBundle
type Adapter struct {
	Type ZCioType `protobuf:"varint,1,opt,name=type,enum=ZCioType" json:"type,omitempty"`
	Name string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Adapter) Reset()                    { *m = Adapter{} }
func (m *Adapter) String() string            { return proto.CompactTextString(m) }
func (*Adapter) ProtoMessage()               {}
func (*Adapter) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Adapter) GetType() ZCioType {
	if m != nil {
		return m.Type
	}
	return ZCioType_ZCioNop
}

func (m *Adapter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// This is way to tell the device if there is service in cloud somewhere,
// what type it is how to access it
type ZcServicePoint struct {
	ZsType     ZcServiceType `protobuf:"varint,3,opt,name=zsType,enum=ZcServiceType" json:"zsType,omitempty"`
	NameOrIp   string        `protobuf:"bytes,1,opt,name=NameOrIp" json:"NameOrIp,omitempty"`
	Credential string        `protobuf:"bytes,2,opt,name=Credential" json:"Credential,omitempty"`
}

func (m *ZcServicePoint) Reset()                    { *m = ZcServicePoint{} }
func (m *ZcServicePoint) String() string            { return proto.CompactTextString(m) }
func (*ZcServicePoint) ProtoMessage()               {}
func (*ZcServicePoint) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ZcServicePoint) GetZsType() ZcServiceType {
	if m != nil {
		return m.ZsType
	}
	return ZcServiceType_zcloudInvalidSrv
}

func (m *ZcServicePoint) GetNameOrIp() string {
	if m != nil {
		return m.NameOrIp
	}
	return ""
}

func (m *ZcServicePoint) GetCredential() string {
	if m != nil {
		return m.Credential
	}
	return ""
}

func init() {
	proto.RegisterType((*UUIDandVersion)(nil), "UUIDandVersion")
	proto.RegisterType((*Adapter)(nil), "Adapter")
	proto.RegisterType((*ZcServicePoint)(nil), "ZcServicePoint")
	proto.RegisterEnum("ZCioType", ZCioType_name, ZCioType_value)
	proto.RegisterEnum("ZcServiceType", ZcServiceType_name, ZcServiceType_value)
}

func init() { proto.RegisterFile("devcommon.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x6b, 0xe2, 0x40,
	0x14, 0xc6, 0x8d, 0x8a, 0x9a, 0xb7, 0x6b, 0x76, 0x76, 0xd8, 0x43, 0x58, 0x70, 0x77, 0xf1, 0xb0,
	0x14, 0x0f, 0x11, 0xda, 0x6b, 0x29, 0xad, 0x5a, 0xda, 0x1c, 0xd4, 0xa2, 0xb5, 0x07, 0xe9, 0x65,
	0xcc, 0x4c, 0x75, 0xc0, 0xcc, 0x0c, 0x93, 0x49, 0xc0, 0xfc, 0xf3, 0x2d, 0x99, 0x24, 0xd2, 0xde,
	0xde, 0xf7, 0xbd, 0xef, 0xfd, 0xe6, 0x83, 0x81, 0x1f, 0x94, 0x65, 0x91, 0x8c, 0x63, 0x29, 0x02,
	0xa5, 0xa5, 0x91, 0xc3, 0x1b, 0xf0, 0x36, 0x9b, 0x70, 0x46, 0x04, 0x7d, 0x61, 0x3a, 0xe1, 0x52,
	0x60, 0x0c, 0xed, 0x34, 0xe5, 0xd4, 0x77, 0xfe, 0x39, 0x17, 0xee, 0xca, 0xce, 0xd8, 0x87, 0x6e,
	0x56, 0xae, 0xfd, 0xa6, 0xb5, 0x6b, 0x39, 0xbc, 0x86, 0xee, 0x1d, 0x25, 0xca, 0x30, 0x8d, 0x07,
	0xd0, 0x36, 0x27, 0xc5, 0xec, 0xa1, 0x77, 0xe9, 0x06, 0xdb, 0x29, 0x97, 0xcf, 0x27, 0xc5, 0x56,
	0xd6, 0x2e, 0xb8, 0x82, 0xc4, 0xac, 0x02, 0xd8, 0x79, 0x68, 0xc0, 0xdb, 0x46, 0x6b, 0xa6, 0x33,
	0x1e, 0xb1, 0x27, 0xc9, 0x85, 0xc1, 0xff, 0xa1, 0x93, 0x27, 0xc5, 0x95, 0xdf, 0xb2, 0x18, 0x2f,
	0x38, 0x07, 0x2c, 0xab, 0xda, 0xe2, 0xdf, 0xd0, 0x5b, 0x90, 0x98, 0x2d, 0x75, 0xa8, 0xaa, 0xa6,
	0x67, 0x8d, 0xff, 0x00, 0x4c, 0x35, 0xa3, 0x4c, 0x18, 0x4e, 0x8e, 0xd5, 0x7b, 0x9f, 0x9c, 0xd1,
	0x2b, 0xf4, 0xea, 0x6e, 0xf8, 0x1b, 0x74, 0x8b, 0x79, 0x21, 0x15, 0x6a, 0xd4, 0xe2, 0xde, 0x1c,
	0x90, 0x53, 0x8b, 0xcd, 0x7a, 0x82, 0x9a, 0xb5, 0x98, 0x2e, 0xe7, 0xa8, 0x85, 0xbf, 0x97, 0xf7,
	0x8f, 0xb3, 0x79, 0x88, 0xda, 0xd8, 0x03, 0xb7, 0x50, 0x4b, 0x73, 0x60, 0x1a, 0xbd, 0x3b, 0xa3,
	0x07, 0xe8, 0x7f, 0xa9, 0x8c, 0x7f, 0x01, 0xca, 0xa3, 0xa3, 0x4c, 0x69, 0x28, 0x32, 0x72, 0xe4,
	0x74, 0xad, 0x33, 0xd4, 0xc0, 0x7d, 0x70, 0x63, 0xa2, 0x8a, 0x1c, 0xd3, 0xc8, 0xc1, 0x3f, 0xa1,
	0x9f, 0xa4, 0x4a, 0x49, 0x6d, 0x2a, 0xab, 0x39, 0xb9, 0x85, 0xbf, 0x91, 0x8c, 0x83, 0x9c, 0x51,
	0x46, 0x49, 0x60, 0x09, 0x41, 0x9a, 0x94, 0xe0, 0xf2, 0xf7, 0xb6, 0x83, 0x3d, 0x37, 0x87, 0x74,
	0x17, 0x44, 0x32, 0x1e, 0x97, 0xb9, 0x31, 0x51, 0x7c, 0x9c, 0x47, 0x52, 0xbc, 0xf1, 0xfd, 0xae,
	0x63, 0x53, 0x57, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x70, 0x60, 0xab, 0xf6, 0x01, 0x00,
	0x00,
}
