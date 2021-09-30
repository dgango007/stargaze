// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/claim/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Params defines the claim module's parameters.
type Params struct {
	AirdropEnabled     bool          `protobuf:"varint,1,opt,name=airdrop_enabled,json=airdropEnabled,proto3" json:"airdrop_enabled,omitempty"`
	AirdropStartTime   time.Time     `protobuf:"bytes,2,opt,name=airdrop_start_time,json=airdropStartTime,proto3,stdtime" json:"airdrop_start_time" yaml:"airdrop_start_time"`
	DurationUntilDecay time.Duration `protobuf:"bytes,3,opt,name=duration_until_decay,json=durationUntilDecay,proto3,stdduration" json:"duration_until_decay,omitempty" yaml:"duration_until_decay"`
	DurationOfDecay    time.Duration `protobuf:"bytes,4,opt,name=duration_of_decay,json=durationOfDecay,proto3,stdduration" json:"duration_of_decay,omitempty" yaml:"duration_of_decay"`
	// denom of claimable asset
	ClaimDenom string `protobuf:"bytes,5,opt,name=claim_denom,json=claimDenom,proto3" json:"claim_denom,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_c219c2c72539a013, []int{0}
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

func (m *Params) GetAirdropEnabled() bool {
	if m != nil {
		return m.AirdropEnabled
	}
	return false
}

func (m *Params) GetAirdropStartTime() time.Time {
	if m != nil {
		return m.AirdropStartTime
	}
	return time.Time{}
}

func (m *Params) GetDurationUntilDecay() time.Duration {
	if m != nil {
		return m.DurationUntilDecay
	}
	return 0
}

func (m *Params) GetDurationOfDecay() time.Duration {
	if m != nil {
		return m.DurationOfDecay
	}
	return 0
}

func (m *Params) GetClaimDenom() string {
	if m != nil {
		return m.ClaimDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "publicawesome.stargaze.claim.v1beta1.Params")
}

func init() {
	proto.RegisterFile("stargaze/claim/v1beta1/params.proto", fileDescriptor_c219c2c72539a013)
}

var fileDescriptor_c219c2c72539a013 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x3f, 0x6f, 0xd4, 0x30,
	0x18, 0xc6, 0xcf, 0x14, 0x2a, 0x70, 0x25, 0x0a, 0x51, 0x87, 0xf4, 0x2a, 0x39, 0xa7, 0x00, 0xe2,
	0x06, 0xb0, 0x29, 0x6c, 0x8c, 0xc7, 0x31, 0xc0, 0x02, 0x3a, 0x60, 0x61, 0x89, 0x9c, 0xc4, 0x17,
	0x2c, 0xc5, 0x71, 0x94, 0x38, 0x40, 0xf8, 0x08, 0x4c, 0x9d, 0x10, 0x1f, 0xa9, 0x63, 0x47, 0xa6,
	0x80, 0xee, 0x26, 0x18, 0xfb, 0x09, 0x90, 0xff, 0x55, 0xa2, 0x77, 0x12, 0x5b, 0xf2, 0x3c, 0xef,
	0xe3, 0xe7, 0xf7, 0x4a, 0x2f, 0xbc, 0xd3, 0x2a, 0xda, 0x14, 0xf4, 0x0b, 0x23, 0x59, 0x49, 0xb9,
	0x20, 0x1f, 0x8f, 0x53, 0xa6, 0xe8, 0x31, 0xa9, 0x69, 0x43, 0x45, 0x8b, 0xeb, 0x46, 0x2a, 0x19,
	0xdc, 0xad, 0xbb, 0xb4, 0xe4, 0x19, 0xfd, 0xc4, 0x5a, 0x29, 0x18, 0xf6, 0x11, 0x6c, 0x22, 0xd8,
	0x45, 0xc6, 0x07, 0x85, 0x2c, 0xa4, 0x09, 0x10, 0xfd, 0x65, 0xb3, 0x63, 0x54, 0x48, 0x59, 0x94,
	0x8c, 0x98, 0xbf, 0xb4, 0x5b, 0x92, 0xbc, 0x6b, 0xa8, 0xe2, 0xb2, 0x72, 0x7e, 0x74, 0xd9, 0x57,
	0x5c, 0xb0, 0x56, 0x51, 0x51, 0xdb, 0x81, 0xf8, 0xf7, 0x0e, 0xdc, 0x7d, 0x6d, 0x68, 0x82, 0xfb,
	0x70, 0x9f, 0xf2, 0x26, 0x6f, 0x64, 0x9d, 0xb0, 0x8a, 0xa6, 0x25, 0xcb, 0x43, 0x30, 0x01, 0xd3,
	0xeb, 0x8b, 0x9b, 0x4e, 0x7e, 0x6e, 0xd5, 0x40, 0xc2, 0xc0, 0x0f, 0x6a, 0x58, 0x95, 0xe8, 0x47,
	0xc3, 0x2b, 0x13, 0x30, 0xdd, 0x7b, 0x3c, 0xc6, 0xb6, 0x11, 0xfb, 0x46, 0xfc, 0xd6, 0x37, 0xce,
	0xee, 0x9d, 0x0e, 0xd1, 0xe8, 0x7c, 0x88, 0x0e, 0x7b, 0x2a, 0xca, 0xa7, 0xf1, 0xe6, 0x1b, 0xf1,
	0xc9, 0xcf, 0x08, 0x2c, 0x6e, 0x39, 0xe3, 0x8d, 0xd6, 0x75, 0x3a, 0xf8, 0x06, 0xe0, 0x81, 0x5f,
	0x2c, 0xe9, 0x2a, 0xc5, 0xcb, 0x24, 0x67, 0x19, 0xed, 0xc3, 0x1d, 0xd3, 0x79, 0xb8, 0xd1, 0x39,
	0x77, 0xc3, 0xb3, 0x17, 0xba, 0xf2, 0xcf, 0x10, 0xa1, 0x6d, 0xf1, 0x07, 0x52, 0x70, 0xc5, 0x44,
	0xad, 0xfa, 0xf3, 0x21, 0x3a, 0xb2, 0x50, 0xdb, 0xe6, 0xe2, 0xef, 0x1a, 0x2b, 0xf0, 0xd6, 0x3b,
	0xed, 0xcc, 0xb5, 0x11, 0x7c, 0x05, 0xf0, 0xf6, 0x45, 0x42, 0x2e, 0x1d, 0xd5, 0xd5, 0xff, 0x51,
	0x3d, 0x73, 0x54, 0x47, 0x1b, 0xd9, 0x7f, 0x90, 0xc2, 0x4b, 0x48, 0x7e, 0xc8, 0xf2, 0xec, 0x7b,
	0xfd, 0xd5, 0xd2, 0xc2, 0x44, 0x70, 0xcf, 0x9c, 0x4c, 0x92, 0xb3, 0x4a, 0x8a, 0xf0, 0xda, 0x04,
	0x4c, 0x6f, 0x2c, 0xa0, 0x91, 0xe6, 0x5a, 0x99, 0xbd, 0x3c, 0x5d, 0x21, 0x70, 0xb6, 0x42, 0xe0,
	0xd7, 0x0a, 0x81, 0x93, 0x35, 0x1a, 0x9d, 0xad, 0xd1, 0xe8, 0xc7, 0x1a, 0x8d, 0xde, 0x3f, 0x2a,
	0xb8, 0xfa, 0xd0, 0xa5, 0x38, 0x93, 0x82, 0xd8, 0x6b, 0x7c, 0xe8, 0xce, 0x91, 0x5c, 0x5c, 0xf0,
	0x67, 0x77, 0xc3, 0xaa, 0xaf, 0x59, 0x9b, 0xee, 0x9a, 0xad, 0x9e, 0xfc, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xf9, 0x0a, 0x06, 0xad, 0xe2, 0x02, 0x00, 0x00,
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
	if len(m.ClaimDenom) > 0 {
		i -= len(m.ClaimDenom)
		copy(dAtA[i:], m.ClaimDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ClaimDenom)))
		i--
		dAtA[i] = 0x2a
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationOfDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationUntilDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.AirdropStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintParams(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	if m.AirdropEnabled {
		i--
		if m.AirdropEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
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
	if m.AirdropEnabled {
		n += 2
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay)
	n += 1 + l + sovParams(uint64(l))
	l = len(m.ClaimDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AirdropEnabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.AirdropStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationUntilDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationUntilDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationOfDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationOfDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimDenom", wireType)
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
			m.ClaimDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
