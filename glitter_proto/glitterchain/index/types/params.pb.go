// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: glitterchain/index/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Params defines the parameters for the module.
type Params struct {
	FeePerDatasetPerSecond github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=fee_per_dataset_per_second,json=feePerDatasetPerSecond,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"fee_per_dataset_per_second"`
	GracePeriod            int64                                  `protobuf:"varint,2,opt,name=grace_period,json=gracePeriod,proto3" json:"grace_period,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b84119a8e057aa5, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetGracePeriod() int64 {
	if m != nil {
		return m.GracePeriod
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "glitterchain.index.Params")
}

func init() { proto.RegisterFile("glitterchain/index/params.proto", fileDescriptor_2b84119a8e057aa5) }

var fileDescriptor_2b84119a8e057aa5 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xcf, 0xc9, 0x2c,
	0x29, 0x49, 0x2d, 0x4a, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xcf, 0xcc, 0x4b, 0x49, 0xad, 0xd0, 0x2f,
	0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42, 0x56, 0xa0,
	0x07, 0x56, 0x20, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x96, 0xd6, 0x07, 0xb1, 0x20, 0x2a, 0x95,
	0xe6, 0x32, 0x72, 0xb1, 0x05, 0x80, 0xb5, 0x0a, 0x65, 0x71, 0x49, 0xa5, 0xa5, 0xa6, 0xc6, 0x17,
	0xa4, 0x16, 0xc5, 0xa7, 0x24, 0x96, 0x24, 0x16, 0xa7, 0x96, 0x80, 0xd9, 0xc5, 0xa9, 0xc9, 0xf9,
	0x79, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e, 0x7a, 0x27, 0xee, 0xc9, 0x33, 0xdc, 0xba,
	0x27, 0xaf, 0x96, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9c, 0x5f,
	0x9c, 0x9b, 0x5f, 0x0c, 0xa5, 0x74, 0x8b, 0x53, 0xb2, 0xf5, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0xf5,
	0x3c, 0xf3, 0x4a, 0x82, 0xc4, 0xd2, 0x52, 0x53, 0x03, 0x52, 0x8b, 0x5c, 0x20, 0xe6, 0x05, 0xa4,
	0x16, 0x05, 0x83, 0x4d, 0x13, 0x52, 0xe4, 0xe2, 0x49, 0x2f, 0x4a, 0x4c, 0x06, 0xdb, 0x96, 0x99,
	0x9f, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1c, 0xc4, 0x0d, 0x16, 0x0b, 0x00, 0x0b, 0x59, 0xb1,
	0xcc, 0x58, 0x20, 0xcf, 0xe0, 0x94, 0x72, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c,
	0x51, 0x5e, 0x48, 0x4e, 0x80, 0x7a, 0x37, 0x2f, 0xb5, 0xa4, 0x3c, 0xbf, 0x28, 0x5b, 0x1f, 0xec,
	0x6d, 0xdd, 0x94, 0xd4, 0x02, 0x98, 0x44, 0x3c, 0xd4, 0xdf, 0x98, 0xc1, 0x06, 0x76, 0x6a, 0x12,
	0x1b, 0x58, 0xde, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x17, 0x39, 0x2b, 0x42, 0x59, 0x01, 0x00,
	0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GracePeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.GracePeriod))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.FeePerDatasetPerSecond.Size()
		i -= size
		if _, err := m.FeePerDatasetPerSecond.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.FeePerDatasetPerSecond.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.GracePeriod != 0 {
		n += 1 + sovParams(uint64(m.GracePeriod))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeePerDatasetPerSecond", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeePerDatasetPerSecond.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GracePeriod", wireType)
			}
			m.GracePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GracePeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
