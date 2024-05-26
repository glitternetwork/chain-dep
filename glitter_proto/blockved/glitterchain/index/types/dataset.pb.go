// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blockved/glitterchain/index/dataset.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/regen-network/cosmos-proto"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BondStatus is the status of a validator.
type BondStatus int32

const (
	// UNSPECIFIED defines an invalid validator status.
	Unspecified BondStatus = 0
	// UNBONDED defines a validator that is not bonded.
	Unbonded BondStatus = 1
	// UNBONDING defines a validator that is unbonding.
	Unbonding BondStatus = 2
	// BONDED defines a validator that is bonded.
	Bonded BondStatus = 3
)

var BondStatus_name = map[int32]string{
	0: "BOND_STATUS_UNSPECIFIED",
	1: "BOND_STATUS_UNBONDED",
	2: "BOND_STATUS_UNBONDING",
	3: "BOND_STATUS_BONDED",
}

var BondStatus_value = map[string]int32{
	"BOND_STATUS_UNSPECIFIED": 0,
	"BOND_STATUS_UNBONDED":    1,
	"BOND_STATUS_UNBONDING":   2,
	"BOND_STATUS_BONDED":      3,
}

func (x BondStatus) String() string {
	return proto.EnumName(BondStatus_name, int32(x))
}

func (BondStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_30db0b27d43e4978, []int{0}
}

type ServiceStatus int32

const (
	unset     ServiceStatus = 0
	preparing ServiceStatus = 1
	ready     ServiceStatus = 2
)

var ServiceStatus_name = map[int32]string{
	0: "Service_Status_Unset",
	1: "Service_Status_Preparing",
	2: "Service_Status_Ready",
}

var ServiceStatus_value = map[string]int32{
	"Service_Status_Unset":     0,
	"Service_Status_Preparing": 1,
	"Service_Status_Ready":     2,
}

func (x ServiceStatus) String() string {
	return proto.EnumName(ServiceStatus_name, int32(x))
}

func (ServiceStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_30db0b27d43e4978, []int{1}
}

type Table struct {
	DatasetName string    `protobuf:"bytes,1,opt,name=dataset_name,json=datasetName,proto3" json:"dataset_name,omitempty"`
	TableName   string    `protobuf:"bytes,2,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	Meta        string    `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
	UpdateTime  time.Time `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3,stdtime" json:"update_time" yaml:"update_time"`
}

func (m *Table) Reset()         { *m = Table{} }
func (m *Table) String() string { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()    {}
func (*Table) Descriptor() ([]byte, []int) {
	return fileDescriptor_30db0b27d43e4978, []int{0}
}
func (m *Table) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Table) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Table.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Table) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Table.Merge(m, src)
}
func (m *Table) XXX_Size() int {
	return m.Size()
}
func (m *Table) XXX_DiscardUnknown() {
	xxx_messageInfo_Table.DiscardUnknown(m)
}

var xxx_messageInfo_Table proto.InternalMessageInfo

func (m *Table) GetDatasetName() string {
	if m != nil {
		return m.DatasetName
	}
	return ""
}

func (m *Table) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *Table) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *Table) GetUpdateTime() time.Time {
	if m != nil {
		return m.UpdateTime
	}
	return time.Time{}
}

type Dataset struct {
	DatasetName     string                                 `protobuf:"bytes,1,opt,name=dataset_name,json=datasetName,proto3" json:"dataset_name,omitempty"`
	WorkStatus      ServiceStatus                          `protobuf:"varint,2,opt,name=work_status,json=workStatus,proto3,enum=blockved.glitterchain.index.ServiceStatus" json:"work_status,omitempty"`
	OwnerAddress    string                                 `protobuf:"bytes,3,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	OwnerStake      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=owner_stake,json=ownerStake,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"owner_stake"`
	Hosts           string                                 `protobuf:"bytes,5,opt,name=hosts,proto3" json:"hosts,omitempty"`
	ManageAddresses string                                 `protobuf:"bytes,6,opt,name=manage_addresses,json=manageAddresses,proto3" json:"manage_addresses,omitempty"`
	Description     string                                 `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	TableMap        map[string]*Table                      `protobuf:"bytes,8,rep,name=table_map,json=tableMap,proto3" json:"table_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UpdateTime      time.Time                              `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3,stdtime" json:"update_time" yaml:"update_time"`
}

func (m *Dataset) Reset()         { *m = Dataset{} }
func (m *Dataset) String() string { return proto.CompactTextString(m) }
func (*Dataset) ProtoMessage()    {}
func (*Dataset) Descriptor() ([]byte, []int) {
	return fileDescriptor_30db0b27d43e4978, []int{1}
}
func (m *Dataset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Dataset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Dataset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Dataset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dataset.Merge(m, src)
}
func (m *Dataset) XXX_Size() int {
	return m.Size()
}
func (m *Dataset) XXX_DiscardUnknown() {
	xxx_messageInfo_Dataset.DiscardUnknown(m)
}

var xxx_messageInfo_Dataset proto.InternalMessageInfo

func (m *Dataset) GetDatasetName() string {
	if m != nil {
		return m.DatasetName
	}
	return ""
}

func (m *Dataset) GetWorkStatus() ServiceStatus {
	if m != nil {
		return m.WorkStatus
	}
	return unset
}

func (m *Dataset) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *Dataset) GetHosts() string {
	if m != nil {
		return m.Hosts
	}
	return ""
}

func (m *Dataset) GetManageAddresses() string {
	if m != nil {
		return m.ManageAddresses
	}
	return ""
}

func (m *Dataset) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Dataset) GetTableMap() map[string]*Table {
	if m != nil {
		return m.TableMap
	}
	return nil
}

func (m *Dataset) GetUpdateTime() time.Time {
	if m != nil {
		return m.UpdateTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterEnum("blockved.glitterchain.index.BondStatus", BondStatus_name, BondStatus_value)
	proto.RegisterEnum("blockved.glitterchain.index.ServiceStatus", ServiceStatus_name, ServiceStatus_value)
	proto.RegisterType((*Table)(nil), "blockved.glitterchain.index.Table")
	proto.RegisterType((*Dataset)(nil), "blockved.glitterchain.index.Dataset")
	proto.RegisterMapType((map[string]*Table)(nil), "blockved.glitterchain.index.Dataset.TableMapEntry")
}

func init() {
	proto.RegisterFile("blockved/glitterchain/index/dataset.proto", fileDescriptor_30db0b27d43e4978)
}

var fileDescriptor_30db0b27d43e4978 = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xf3, 0x44,
	0x14, 0x8d, 0xf3, 0xd3, 0xaf, 0x19, 0x37, 0x7c, 0xd6, 0x28, 0x08, 0x7f, 0x46, 0x38, 0x26, 0x95,
	0xaa, 0xb4, 0x50, 0x5b, 0x0a, 0x9b, 0xaa, 0xbb, 0x86, 0x04, 0x54, 0x21, 0xd2, 0x2a, 0x3f, 0x1b,
	0x58, 0x58, 0x13, 0x7b, 0xea, 0x5a, 0x89, 0x67, 0x2c, 0xcf, 0xa4, 0x25, 0x6f, 0x80, 0x22, 0x16,
	0x7d, 0x81, 0xac, 0x78, 0x07, 0x16, 0x3c, 0x41, 0x97, 0x5d, 0x22, 0x16, 0x05, 0xda, 0x37, 0xe0,
	0x09, 0x90, 0x67, 0x26, 0x22, 0x2d, 0x28, 0x62, 0xc1, 0xca, 0x73, 0xcf, 0x3d, 0xf7, 0xce, 0x19,
	0x9f, 0x3b, 0x03, 0x0e, 0x27, 0x33, 0x1a, 0x4c, 0x6f, 0x70, 0xe8, 0x45, 0xb3, 0x98, 0x73, 0x9c,
	0x05, 0xd7, 0x28, 0x26, 0x5e, 0x4c, 0x42, 0xfc, 0x9d, 0x17, 0x22, 0x8e, 0x18, 0xe6, 0x6e, 0x9a,
	0x51, 0x4e, 0xe1, 0x87, 0x6b, 0xaa, 0xbb, 0x49, 0x75, 0x05, 0xd5, 0x7a, 0x17, 0x50, 0x96, 0x50,
	0xe6, 0x0b, 0xaa, 0x27, 0x03, 0x59, 0x67, 0xd5, 0x23, 0x1a, 0x51, 0x89, 0xe7, 0x2b, 0x85, 0xbe,
	0x8b, 0x28, 0x8d, 0x66, 0xd8, 0x13, 0xd1, 0x64, 0x7e, 0xe5, 0x21, 0xb2, 0x50, 0xa9, 0xc6, 0xeb,
	0x14, 0x8f, 0x13, 0xcc, 0x38, 0x4a, 0x52, 0x49, 0x68, 0xfe, 0xac, 0x81, 0xca, 0x08, 0x4d, 0x66,
	0x18, 0x7e, 0x0c, 0xf6, 0x94, 0x48, 0x9f, 0xa0, 0x04, 0x9b, 0x9a, 0xa3, 0xb5, 0xaa, 0x03, 0x5d,
	0x61, 0x7d, 0x94, 0x60, 0xf8, 0x11, 0x00, 0x3c, 0xe7, 0x4a, 0x42, 0x51, 0x10, 0xaa, 0x02, 0x11,
	0x69, 0x08, 0xca, 0x09, 0xe6, 0xc8, 0x2c, 0x89, 0x84, 0x58, 0xc3, 0x6f, 0x81, 0x3e, 0x4f, 0x43,
	0xc4, 0xb1, 0x9f, 0xef, 0x6c, 0x96, 0x1d, 0xad, 0xa5, 0xb7, 0x2d, 0x57, 0xca, 0x72, 0xd7, 0xb2,
	0xdc, 0xd1, 0x5a, 0x56, 0xc7, 0xbe, 0x7f, 0x6c, 0x14, 0xfe, 0x7c, 0x6c, 0xc0, 0x05, 0x4a, 0x66,
	0xa7, 0xcd, 0x8d, 0xe2, 0xe6, 0xdd, 0x6f, 0x0d, 0x6d, 0x00, 0x24, 0x92, 0x17, 0x34, 0xff, 0x28,
	0x83, 0x37, 0x5d, 0xa9, 0xef, 0xbf, 0xc8, 0xff, 0x0a, 0xe8, 0xb7, 0x34, 0x9b, 0xfa, 0x8c, 0x23,
	0x3e, 0x67, 0x42, 0xff, 0x7b, 0xed, 0x23, 0x77, 0x8b, 0x17, 0xee, 0x10, 0x67, 0x37, 0x71, 0x80,
	0x87, 0xa2, 0x62, 0x00, 0xf2, 0x72, 0xb9, 0x86, 0xfb, 0xa0, 0x46, 0x6f, 0x09, 0xce, 0x7c, 0x14,
	0x86, 0x19, 0x66, 0x4c, 0x9d, 0x7a, 0x4f, 0x80, 0x67, 0x12, 0x83, 0x17, 0x40, 0x97, 0x24, 0xc6,
	0xd1, 0x54, 0x9e, 0xbe, 0xda, 0x71, 0xf3, 0x13, 0xfe, 0xfa, 0xd8, 0x38, 0x88, 0x62, 0x7e, 0x3d,
	0x9f, 0xb8, 0x01, 0x4d, 0x94, 0xcb, 0xea, 0x73, 0xcc, 0xc2, 0xa9, 0xc7, 0x17, 0x29, 0x66, 0x6e,
	0x17, 0x07, 0x03, 0x20, 0x5a, 0x0c, 0xf3, 0x0e, 0xb0, 0x0e, 0x2a, 0xd7, 0x94, 0x71, 0x66, 0x56,
	0xc4, 0x6e, 0x32, 0x80, 0x87, 0xc0, 0x48, 0x10, 0x41, 0x11, 0x5e, 0x8b, 0xc1, 0xcc, 0xdc, 0x11,
	0x84, 0xb7, 0x12, 0x3f, 0x5b, 0xc3, 0xd0, 0x01, 0x7a, 0x88, 0x59, 0x90, 0xc5, 0x29, 0x8f, 0x29,
	0x31, 0xdf, 0xa8, 0xbf, 0xf4, 0x37, 0x04, 0x2f, 0x80, 0xb4, 0xd4, 0x4f, 0x50, 0x6a, 0xee, 0x3a,
	0xa5, 0x96, 0xde, 0x6e, 0x6f, 0xfd, 0x47, 0xca, 0x01, 0x57, 0x8c, 0xd1, 0xd7, 0x28, 0xed, 0x11,
	0x9e, 0x2d, 0x06, 0xbb, 0x5c, 0x85, 0xaf, 0x47, 0xa0, 0xfa, 0x7f, 0x8e, 0x80, 0xe5, 0x83, 0xda,
	0x8b, 0x7d, 0xa1, 0x01, 0x4a, 0x53, 0xbc, 0x50, 0xf6, 0xe7, 0x4b, 0x78, 0x02, 0x2a, 0x37, 0x68,
	0x36, 0x97, 0x03, 0xab, 0xb7, 0x9b, 0x5b, 0x0f, 0x23, 0x9a, 0x0d, 0x64, 0xc1, 0x69, 0xf1, 0x44,
	0x3b, 0xfa, 0x49, 0x03, 0xa0, 0x43, 0x49, 0xa8, 0x6c, 0xff, 0x14, 0x7c, 0xd0, 0xb9, 0xe8, 0x77,
	0xfd, 0xe1, 0xe8, 0x6c, 0x34, 0x1e, 0xfa, 0xe3, 0xfe, 0xf0, 0xb2, 0xf7, 0xf9, 0xf9, 0x17, 0xe7,
	0xbd, 0xae, 0x51, 0xb0, 0xde, 0x2e, 0x57, 0x8e, 0x3e, 0x26, 0x2c, 0xc5, 0x41, 0x7c, 0x15, 0xe3,
	0x10, 0x1e, 0x80, 0xfa, 0x4b, 0x76, 0x1e, 0xf5, 0xba, 0x86, 0x66, 0xed, 0x2d, 0x57, 0xce, 0xee,
	0x98, 0x4c, 0x28, 0x09, 0x71, 0x08, 0x5b, 0xe0, 0xfd, 0x7f, 0xf2, 0xce, 0xfb, 0x5f, 0x1a, 0x45,
	0xab, 0xb6, 0x5c, 0x39, 0x55, 0x49, 0x8c, 0x49, 0x04, 0x9b, 0x00, 0x6e, 0x32, 0x55, 0xbf, 0x92,
	0x05, 0x96, 0x2b, 0x67, 0xa7, 0x23, 0xba, 0x59, 0xe5, 0xef, 0x7f, 0xb4, 0x0b, 0x47, 0x3f, 0x68,
	0xa0, 0xf6, 0x62, 0x7c, 0xe1, 0x3e, 0xa8, 0x2b, 0xc0, 0x97, 0x88, 0x3f, 0x26, 0x0c, 0x73, 0xa3,
	0x60, 0x55, 0x97, 0x2b, 0xa7, 0x32, 0xcf, 0x03, 0xf8, 0x09, 0x30, 0x5f, 0x91, 0x2e, 0x33, 0x9c,
	0xa2, 0x2c, 0x26, 0x91, 0xa1, 0x49, 0x35, 0xe9, 0x1a, 0xf8, 0x97, 0x8e, 0x03, 0x8c, 0xc2, 0x85,
	0x51, 0x94, 0x1d, 0xb3, 0x3c, 0x90, 0x72, 0x3a, 0xc9, 0xfd, 0x93, 0xad, 0x3d, 0x3c, 0xd9, 0xda,
	0xef, 0x4f, 0xb6, 0x76, 0xf7, 0x6c, 0x17, 0x1e, 0x9e, 0xed, 0xc2, 0x2f, 0xcf, 0x76, 0xe1, 0x9b,
	0xe1, 0xc6, 0x3d, 0x50, 0x8e, 0x10, 0xcc, 0xf3, 0xab, 0xe6, 0x09, 0x67, 0x8e, 0x43, 0x9c, 0xae,
	0x13, 0xea, 0x51, 0xdc, 0xf6, 0xd2, 0x8a, 0x8b, 0x33, 0xd9, 0x11, 0xc4, 0xcf, 0xfe, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x4a, 0xba, 0x14, 0x24, 0x95, 0x05, 0x00, 0x00,
}

func (m *Table) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Table) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Table) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.UpdateTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdateTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintDataset(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if len(m.Meta) > 0 {
		i -= len(m.Meta)
		copy(dAtA[i:], m.Meta)
		i = encodeVarintDataset(dAtA, i, uint64(len(m.Meta)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TableName) > 0 {
		i -= len(m.TableName)
		copy(dAtA[i:], m.TableName)
		i = encodeVarintDataset(dAtA, i, uint64(len(m.TableName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DatasetName) > 0 {
		i -= len(m.DatasetName)
		copy(dAtA[i:], m.DatasetName)
		i = encodeVarintDataset(dAtA, i, uint64(len(m.DatasetName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Dataset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Dataset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Dataset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.UpdateTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdateTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintDataset(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x4a
	if len(m.TableMap) > 0 {
		for k := range m.TableMap {
			v := m.TableMap[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintDataset(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintDataset(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintDataset(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintDataset(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ManageAddresses) > 0 {
		i -= len(m.ManageAddresses)
		copy(dAtA[i:], m.ManageAddresses)
		i = encodeVarintDataset(dAtA, i, uint64(len(m.ManageAddresses)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Hosts) > 0 {
		i -= len(m.Hosts)
		copy(dAtA[i:], m.Hosts)
		i = encodeVarintDataset(dAtA, i, uint64(len(m.Hosts)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size := m.OwnerStake.Size()
		i -= size
		if _, err := m.OwnerStake.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDataset(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintDataset(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.WorkStatus != 0 {
		i = encodeVarintDataset(dAtA, i, uint64(m.WorkStatus))
		i--
		dAtA[i] = 0x10
	}
	if len(m.DatasetName) > 0 {
		i -= len(m.DatasetName)
		copy(dAtA[i:], m.DatasetName)
		i = encodeVarintDataset(dAtA, i, uint64(len(m.DatasetName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDataset(dAtA []byte, offset int, v uint64) int {
	offset -= sovDataset(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Table) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DatasetName)
	if l > 0 {
		n += 1 + l + sovDataset(uint64(l))
	}
	l = len(m.TableName)
	if l > 0 {
		n += 1 + l + sovDataset(uint64(l))
	}
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovDataset(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdateTime)
	n += 1 + l + sovDataset(uint64(l))
	return n
}

func (m *Dataset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DatasetName)
	if l > 0 {
		n += 1 + l + sovDataset(uint64(l))
	}
	if m.WorkStatus != 0 {
		n += 1 + sovDataset(uint64(m.WorkStatus))
	}
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovDataset(uint64(l))
	}
	l = m.OwnerStake.Size()
	n += 1 + l + sovDataset(uint64(l))
	l = len(m.Hosts)
	if l > 0 {
		n += 1 + l + sovDataset(uint64(l))
	}
	l = len(m.ManageAddresses)
	if l > 0 {
		n += 1 + l + sovDataset(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovDataset(uint64(l))
	}
	if len(m.TableMap) > 0 {
		for k, v := range m.TableMap {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovDataset(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovDataset(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovDataset(uint64(mapEntrySize))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdateTime)
	n += 1 + l + sovDataset(uint64(l))
	return n
}

func sovDataset(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDataset(x uint64) (n int) {
	return sovDataset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Table) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataset
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Table: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Table: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DatasetName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DatasetName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TableName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.UpdateTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDataset
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Dataset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDataset
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Dataset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Dataset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DatasetName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DatasetName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkStatus", wireType)
			}
			m.WorkStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WorkStatus |= ServiceStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerStake", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OwnerStake.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hosts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hosts = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ManageAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ManageAddresses = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDataset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TableMap == nil {
				m.TableMap = make(map[string]*Table)
			}
			var mapkey string
			var mapvalue *Table
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowDataset
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDataset
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthDataset
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthDataset
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowDataset
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthDataset
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthDataset
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Table{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipDataset(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthDataset
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.TableMap[mapkey] = mapvalue
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDataset
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDataset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.UpdateTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDataset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDataset
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDataset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDataset
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDataset
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDataset
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDataset
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDataset
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDataset        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDataset          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDataset = fmt.Errorf("proto: unexpected end of group")
)
