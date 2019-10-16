// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v1alpha/deps/deps.proto

package deps

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Dependency struct {
	Organization         *string  `protobuf:"bytes,1,opt,name=organization" json:"organization,omitempty"`
	Module               *string  `protobuf:"bytes,2,opt,name=module" json:"module,omitempty"`
	VersionConstraint    *string  `protobuf:"bytes,3,opt,name=versionConstraint" json:"versionConstraint,omitempty"`
	Scopes               []string `protobuf:"bytes,4,rep,name=scopes" json:"scopes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dependency) Reset()         { *m = Dependency{} }
func (m *Dependency) String() string { return proto.CompactTextString(m) }
func (*Dependency) ProtoMessage()    {}
func (*Dependency) Descriptor() ([]byte, []int) {
	return fileDescriptor_edce5762dbbc1533, []int{0}
}
func (m *Dependency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dependency.Unmarshal(m, b)
}
func (m *Dependency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dependency.Marshal(b, m, deterministic)
}
func (m *Dependency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dependency.Merge(m, src)
}
func (m *Dependency) XXX_Size() int {
	return xxx_messageInfo_Dependency.Size(m)
}
func (m *Dependency) XXX_DiscardUnknown() {
	xxx_messageInfo_Dependency.DiscardUnknown(m)
}

var xxx_messageInfo_Dependency proto.InternalMessageInfo

func (m *Dependency) GetOrganization() string {
	if m != nil && m.Organization != nil {
		return *m.Organization
	}
	return ""
}

func (m *Dependency) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

func (m *Dependency) GetVersionConstraint() string {
	if m != nil && m.VersionConstraint != nil {
		return *m.VersionConstraint
	}
	return ""
}

func (m *Dependency) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

type DependencyManagementFile struct {
	Language             *string       `protobuf:"bytes,1,opt,name=language" json:"language,omitempty"`
	System               *string       `protobuf:"bytes,2,opt,name=system" json:"system,omitempty"`
	Organization         *string       `protobuf:"bytes,5,opt,name=organization" json:"organization,omitempty"`
	Module               *string       `protobuf:"bytes,6,opt,name=module" json:"module,omitempty"`
	Version              *string       `protobuf:"bytes,7,opt,name=version" json:"version,omitempty"`
	Dependencies         []*Dependency `protobuf:"bytes,8,rep,name=dependencies" json:"dependencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DependencyManagementFile) Reset()         { *m = DependencyManagementFile{} }
func (m *DependencyManagementFile) String() string { return proto.CompactTextString(m) }
func (*DependencyManagementFile) ProtoMessage()    {}
func (*DependencyManagementFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_edce5762dbbc1533, []int{1}
}
func (m *DependencyManagementFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DependencyManagementFile.Unmarshal(m, b)
}
func (m *DependencyManagementFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DependencyManagementFile.Marshal(b, m, deterministic)
}
func (m *DependencyManagementFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DependencyManagementFile.Merge(m, src)
}
func (m *DependencyManagementFile) XXX_Size() int {
	return xxx_messageInfo_DependencyManagementFile.Size(m)
}
func (m *DependencyManagementFile) XXX_DiscardUnknown() {
	xxx_messageInfo_DependencyManagementFile.DiscardUnknown(m)
}

var xxx_messageInfo_DependencyManagementFile proto.InternalMessageInfo

func (m *DependencyManagementFile) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *DependencyManagementFile) GetSystem() string {
	if m != nil && m.System != nil {
		return *m.System
	}
	return ""
}

func (m *DependencyManagementFile) GetOrganization() string {
	if m != nil && m.Organization != nil {
		return *m.Organization
	}
	return ""
}

func (m *DependencyManagementFile) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

func (m *DependencyManagementFile) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *DependencyManagementFile) GetDependencies() []*Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func init() {
	proto.RegisterType((*Dependency)(nil), "cloud.deps.api.v1alpha.deps.Dependency")
	proto.RegisterType((*DependencyManagementFile)(nil), "cloud.deps.api.v1alpha.deps.DependencyManagementFile")
}

func init() { proto.RegisterFile("v1alpha/deps/deps.proto", fileDescriptor_edce5762dbbc1533) }

var fileDescriptor_edce5762dbbc1533 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x55, 0x08, 0xb4, 0xc5, 0x74, 0xc1, 0x03, 0x58, 0xb0, 0x44, 0x19, 0x20, 0x03, 0x38, 0x82,
	0x4f, 0x00, 0xc4, 0x82, 0x58, 0x3a, 0xb2, 0x1d, 0xc9, 0x29, 0xb5, 0x94, 0xdc, 0x59, 0xb1, 0x53,
	0xa9, 0x7c, 0x04, 0xff, 0xca, 0x1f, 0xa0, 0xba, 0x09, 0x25, 0xaa, 0xd4, 0xc5, 0xd2, 0x7b, 0xef,
	0xfc, 0xf4, 0xde, 0x9d, 0xb8, 0x5c, 0x3d, 0x40, 0x6d, 0x97, 0x90, 0x97, 0x68, 0x5d, 0x78, 0xb4,
	0x6d, 0xd9, 0xb3, 0xbc, 0x2e, 0x6a, 0xee, 0x4a, 0x1d, 0x18, 0xb0, 0x46, 0xf7, 0x73, 0x81, 0x48,
	0xbf, 0x23, 0x21, 0x5e, 0xd0, 0x22, 0x95, 0x48, 0xc5, 0x5a, 0xa6, 0x62, 0xce, 0x6d, 0x05, 0x64,
	0xbe, 0xc0, 0x1b, 0x26, 0x15, 0x25, 0x51, 0x76, 0xba, 0x18, 0x71, 0xf2, 0x42, 0x4c, 0x1a, 0x2e,
	0xbb, 0x1a, 0xd5, 0x51, 0x50, 0x7b, 0x24, 0xef, 0xc4, 0xf9, 0x0a, 0x5b, 0x67, 0x98, 0x9e, 0x99,
	0x9c, 0x6f, 0xc1, 0x90, 0x57, 0x71, 0x18, 0xd9, 0x17, 0x36, 0x2e, 0xae, 0x60, 0x8b, 0x4e, 0x1d,
	0x27, 0xf1, 0xc6, 0x65, 0x8b, 0xd2, 0x9f, 0x48, 0xa8, 0x5d, 0xa0, 0x77, 0x20, 0xa8, 0xb0, 0x41,
	0xf2, 0xaf, 0xa6, 0x46, 0x79, 0x25, 0x66, 0x35, 0x50, 0xd5, 0x41, 0x85, 0x7d, 0xb4, 0x3f, 0x1c,
	0x0c, 0xd7, 0xce, 0x63, 0x33, 0xc4, 0xda, 0xa2, 0xbd, 0x4a, 0x27, 0x07, 0x2b, 0x4d, 0x46, 0x95,
	0x94, 0x98, 0xf6, 0xc9, 0xd5, 0x34, 0x08, 0x03, 0x94, 0x6f, 0x62, 0x5e, 0x0e, 0x29, 0x0d, 0x3a,
	0x35, 0x4b, 0xe2, 0xec, 0xec, 0xf1, 0x56, 0x1f, 0xd8, 0xb5, 0xde, 0xd5, 0x5a, 0x8c, 0x3e, 0x3f,
	0x65, 0x1f, 0x37, 0x95, 0xf1, 0xcb, 0xee, 0x53, 0x17, 0xdc, 0x84, 0xd3, 0xdd, 0x07, 0x9f, 0x1c,
	0xac, 0xc9, 0xff, 0x9f, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0x20, 0x8b, 0x30, 0xd9, 0xe5, 0x01,
	0x00, 0x00,
}
