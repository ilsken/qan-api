// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vschema.proto

/*
Package vschema is a generated protocol buffer package.

It is generated from these files:
	vschema.proto

It has these top-level messages:
	Keyspace
	Vindex
	Table
	ColumnVindex
	AutoIncrement
	SrvVSchema
*/
package vschema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Keyspace is the vschema for a keyspace.
type Keyspace struct {
	// If sharded is false, vindexes and tables are ignored.
	Sharded  bool               `protobuf:"varint,1,opt,name=sharded" json:"sharded,omitempty"`
	Vindexes map[string]*Vindex `protobuf:"bytes,2,rep,name=vindexes" json:"vindexes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Tables   map[string]*Table  `protobuf:"bytes,3,rep,name=tables" json:"tables,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Keyspace) Reset()                    { *m = Keyspace{} }
func (m *Keyspace) String() string            { return proto.CompactTextString(m) }
func (*Keyspace) ProtoMessage()               {}
func (*Keyspace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Keyspace) GetSharded() bool {
	if m != nil {
		return m.Sharded
	}
	return false
}

func (m *Keyspace) GetVindexes() map[string]*Vindex {
	if m != nil {
		return m.Vindexes
	}
	return nil
}

func (m *Keyspace) GetTables() map[string]*Table {
	if m != nil {
		return m.Tables
	}
	return nil
}

// Vindex is the vindex info for a Keyspace.
type Vindex struct {
	// The type must match one of the predefined
	// (or plugged in) vindex names.
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// params is a map of attribute value pairs
	// that must be defined as required by the
	// vindex constructors. The values can only
	// be strings.
	Params map[string]string `protobuf:"bytes,2,rep,name=params" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// A lookup vindex can have an owner table defined.
	// If so, rows in the lookup table are created or
	// deleted in sync with corresponding rows in the
	// owner table.
	Owner string `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
}

func (m *Vindex) Reset()                    { *m = Vindex{} }
func (m *Vindex) String() string            { return proto.CompactTextString(m) }
func (*Vindex) ProtoMessage()               {}
func (*Vindex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Vindex) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Vindex) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Vindex) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

// Table is the table info for a Keyspace.
type Table struct {
	// If the table is a sequence, type must be
	// "sequence". Otherwise, it should be empty.
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// column_vindexes associates columns to vindexes.
	ColumnVindexes []*ColumnVindex `protobuf:"bytes,2,rep,name=column_vindexes,json=columnVindexes" json:"column_vindexes,omitempty"`
	// auto_increment is specified if a column needs
	// to be associated with a sequence.
	AutoIncrement *AutoIncrement `protobuf:"bytes,3,opt,name=auto_increment,json=autoIncrement" json:"auto_increment,omitempty"`
}

func (m *Table) Reset()                    { *m = Table{} }
func (m *Table) String() string            { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()               {}
func (*Table) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Table) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Table) GetColumnVindexes() []*ColumnVindex {
	if m != nil {
		return m.ColumnVindexes
	}
	return nil
}

func (m *Table) GetAutoIncrement() *AutoIncrement {
	if m != nil {
		return m.AutoIncrement
	}
	return nil
}

// ColumnVindex is used to associate a column to a vindex.
type ColumnVindex struct {
	Column string `protobuf:"bytes,1,opt,name=column" json:"column,omitempty"`
	// The name must match a vindex defined in Keyspace.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ColumnVindex) Reset()                    { *m = ColumnVindex{} }
func (m *ColumnVindex) String() string            { return proto.CompactTextString(m) }
func (*ColumnVindex) ProtoMessage()               {}
func (*ColumnVindex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ColumnVindex) GetColumn() string {
	if m != nil {
		return m.Column
	}
	return ""
}

func (m *ColumnVindex) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Autoincrement is used to designate a column as auto-inc.
type AutoIncrement struct {
	Column string `protobuf:"bytes,1,opt,name=column" json:"column,omitempty"`
	// The sequence must match a table of type SEQUENCE.
	Sequence string `protobuf:"bytes,2,opt,name=sequence" json:"sequence,omitempty"`
}

func (m *AutoIncrement) Reset()                    { *m = AutoIncrement{} }
func (m *AutoIncrement) String() string            { return proto.CompactTextString(m) }
func (*AutoIncrement) ProtoMessage()               {}
func (*AutoIncrement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AutoIncrement) GetColumn() string {
	if m != nil {
		return m.Column
	}
	return ""
}

func (m *AutoIncrement) GetSequence() string {
	if m != nil {
		return m.Sequence
	}
	return ""
}

// SrvVSchema is the roll-up of all the Keyspace schema for a cell.
type SrvVSchema struct {
	// keyspaces is a map of keyspace name -> Keyspace object.
	Keyspaces map[string]*Keyspace `protobuf:"bytes,1,rep,name=keyspaces" json:"keyspaces,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SrvVSchema) Reset()                    { *m = SrvVSchema{} }
func (m *SrvVSchema) String() string            { return proto.CompactTextString(m) }
func (*SrvVSchema) ProtoMessage()               {}
func (*SrvVSchema) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SrvVSchema) GetKeyspaces() map[string]*Keyspace {
	if m != nil {
		return m.Keyspaces
	}
	return nil
}

func init() {
	proto.RegisterType((*Keyspace)(nil), "vschema.Keyspace")
	proto.RegisterType((*Vindex)(nil), "vschema.Vindex")
	proto.RegisterType((*Table)(nil), "vschema.Table")
	proto.RegisterType((*ColumnVindex)(nil), "vschema.ColumnVindex")
	proto.RegisterType((*AutoIncrement)(nil), "vschema.AutoIncrement")
	proto.RegisterType((*SrvVSchema)(nil), "vschema.SrvVSchema")
}

func init() { proto.RegisterFile("vschema.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xd1, 0x6a, 0xd4, 0x40,
	0x14, 0x86, 0x99, 0xc4, 0x4d, 0xb3, 0x27, 0x26, 0xd5, 0xa1, 0x96, 0x10, 0x11, 0x97, 0xa0, 0xb8,
	0x57, 0xb9, 0xd8, 0x22, 0x68, 0x45, 0x51, 0x8a, 0x17, 0x45, 0x41, 0x49, 0xa5, 0xb7, 0x65, 0x9a,
	0x3d, 0xd0, 0xd2, 0xcd, 0x24, 0x66, 0x92, 0x68, 0x5e, 0xc5, 0x1b, 0xc1, 0x37, 0xf0, 0x0d, 0xa5,
	0x93, 0xc9, 0x74, 0xb2, 0x1b, 0xef, 0xe6, 0xe7, 0x9c, 0xff, 0x9f, 0x6f, 0xe6, 0xcc, 0x80, 0xdf,
	0x8a, 0xec, 0x0a, 0x73, 0x96, 0x94, 0x55, 0x51, 0x17, 0x74, 0x4f, 0xc9, 0xf8, 0xaf, 0x05, 0xee,
	0x27, 0xec, 0x44, 0xc9, 0x32, 0xa4, 0x21, 0xec, 0x89, 0x2b, 0x56, 0xad, 0x71, 0x1d, 0x92, 0x05,
	0x59, 0xba, 0xe9, 0x20, 0xe9, 0x1b, 0x70, 0xdb, 0x6b, 0xbe, 0xc6, 0x9f, 0x28, 0x42, 0x6b, 0x61,
	0x2f, 0xbd, 0xd5, 0xd3, 0x64, 0x48, 0x1c, 0xec, 0xc9, 0xb9, 0xea, 0xf8, 0xc8, 0xeb, 0xaa, 0x4b,
	0xb5, 0x81, 0xbe, 0x04, 0xa7, 0x66, 0x97, 0x1b, 0x14, 0xa1, 0x2d, 0xad, 0x4f, 0x76, 0xad, 0xdf,
	0x64, 0xbd, 0x37, 0xaa, 0xe6, 0xe8, 0x33, 0xf8, 0xa3, 0x44, 0xfa, 0x00, 0xec, 0x1b, 0xec, 0x24,
	0xda, 0x3c, 0xbd, 0x5d, 0xd2, 0xe7, 0x30, 0x6b, 0xd9, 0xa6, 0xc1, 0xd0, 0x5a, 0x90, 0xa5, 0xb7,
	0xda, 0xd7, 0xc1, 0xbd, 0x31, 0xed, 0xab, 0xc7, 0xd6, 0x2b, 0x12, 0x9d, 0x82, 0x67, 0x6c, 0x32,
	0x91, 0xf5, 0x6c, 0x9c, 0x15, 0xe8, 0x2c, 0x69, 0x33, 0xa2, 0xe2, 0x3f, 0x04, 0x9c, 0x7e, 0x03,
	0x4a, 0xe1, 0x5e, 0xdd, 0x95, 0xa8, 0x72, 0xe4, 0x9a, 0x1e, 0x81, 0x53, 0xb2, 0x8a, 0xe5, 0xc3,
	0x4d, 0x3d, 0xde, 0xa2, 0x4a, 0xbe, 0xca, 0xaa, 0x3a, 0x6c, 0xdf, 0x4a, 0x0f, 0x60, 0x56, 0xfc,
	0xe0, 0x58, 0x85, 0xb6, 0x4c, 0xea, 0x45, 0xf4, 0x1a, 0x3c, 0xa3, 0x79, 0x02, 0xfa, 0xc0, 0x84,
	0x9e, 0x9b, 0x90, 0xbf, 0x08, 0xcc, 0x24, 0xf9, 0x24, 0xe3, 0x3b, 0xd8, 0xcf, 0x8a, 0x4d, 0x93,
	0xf3, 0x8b, 0xad, 0xb1, 0x3e, 0xd2, 0xb0, 0x27, 0xb2, 0xae, 0x2e, 0x32, 0xc8, 0x0c, 0x85, 0x82,
	0xbe, 0x85, 0x80, 0x35, 0x75, 0x71, 0x71, 0xcd, 0xb3, 0x0a, 0x73, 0xe4, 0xb5, 0xe4, 0xf6, 0x56,
	0x87, 0xda, 0xfe, 0xa1, 0xa9, 0x8b, 0xd3, 0xa1, 0x9a, 0xfa, 0xcc, 0x94, 0xf1, 0x31, 0xdc, 0x37,
	0xe3, 0xe9, 0x21, 0x38, 0xfd, 0x06, 0x0a, 0x52, 0xa9, 0x5b, 0x74, 0xce, 0xf2, 0xe1, 0x74, 0x72,
	0x1d, 0x9f, 0x80, 0x3f, 0xca, 0xfe, 0xaf, 0x39, 0x02, 0x57, 0xe0, 0xf7, 0x06, 0x79, 0x36, 0x04,
	0x68, 0x1d, 0xff, 0x26, 0x00, 0x67, 0x55, 0x7b, 0x7e, 0x26, 0x61, 0xe9, 0x7b, 0x98, 0xdf, 0xa8,
	0xa7, 0x28, 0x42, 0x22, 0x2f, 0x22, 0xd6, 0x27, 0xb9, 0xeb, 0xd3, 0xef, 0x55, 0x0d, 0xef, 0xce,
	0x14, 0x7d, 0x81, 0x60, 0x5c, 0x9c, 0x18, 0xd6, 0x8b, 0xf1, 0x0b, 0x7b, 0xb8, 0xf3, 0x0d, 0x8c,
	0xf9, 0x5d, 0x3a, 0xf2, 0xa3, 0x1e, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x47, 0x7e, 0x27,
	0xb9, 0x03, 0x00, 0x00,
}