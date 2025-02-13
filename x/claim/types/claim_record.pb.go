// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/claim/v1beta1/claim_record.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type Action int32

const (
	ActionInitialClaim  Action = 0
	ActionBidNFT        Action = 1
	ActionMintNFT       Action = 2
	ActionVote          Action = 3
	ActionDelegateStake Action = 4
)

var Action_name = map[int32]string{
	0: "ActionInitialClaim",
	1: "ActionBidNFT",
	2: "ActionMintNFT",
	3: "ActionVote",
	4: "ActionDelegateStake",
}

var Action_value = map[string]int32{
	"ActionInitialClaim":  0,
	"ActionBidNFT":        1,
	"ActionMintNFT":       2,
	"ActionVote":          3,
	"ActionDelegateStake": 4,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8a75b9157744df4f, []int{0}
}

type ClaimRecord struct {
	// address of claim user
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// total initial claimable amount for the user
	InitialClaimableAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=initial_claimable_amount,json=initialClaimableAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_claimable_amount" yaml:"initial_claimable_amount"`
	// true if action is completed
	// index of bool in array refers to action enum #
	ActionCompleted []bool `protobuf:"varint,4,rep,packed,name=action_completed,json=actionCompleted,proto3" json:"action_completed,omitempty" yaml:"action_completed"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a75b9157744df4f, []int{0}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func (m *ClaimRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClaimRecord) GetInitialClaimableAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.InitialClaimableAmount
	}
	return nil
}

func (m *ClaimRecord) GetActionCompleted() []bool {
	if m != nil {
		return m.ActionCompleted
	}
	return nil
}

func init() {
	proto.RegisterEnum("publicawesome.stargaze.claim.v1beta1.Action", Action_name, Action_value)
	proto.RegisterType((*ClaimRecord)(nil), "publicawesome.stargaze.claim.v1beta1.ClaimRecord")
}

func init() {
	proto.RegisterFile("stargaze/claim/v1beta1/claim_record.proto", fileDescriptor_8a75b9157744df4f)
}

var fileDescriptor_8a75b9157744df4f = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x41, 0x6e, 0xd4, 0x30,
	0x14, 0x4d, 0x66, 0x46, 0x05, 0x5c, 0x28, 0xc1, 0xa0, 0x76, 0x98, 0x85, 0x33, 0x8a, 0x58, 0x0c,
	0x88, 0xda, 0x2a, 0x5d, 0x20, 0xb1, 0x6b, 0x06, 0x21, 0x81, 0x54, 0x16, 0x53, 0xc4, 0x82, 0xcd,
	0xc8, 0x49, 0xbe, 0x82, 0xd5, 0x24, 0x8e, 0x62, 0x4f, 0xa1, 0x9c, 0x80, 0x25, 0x77, 0x60, 0x83,
	0x58, 0x72, 0x8a, 0x2e, 0xbb, 0x64, 0x35, 0xa0, 0x99, 0x1b, 0xf4, 0x04, 0x28, 0xb6, 0x03, 0x15,
	0x52, 0x57, 0xb6, 0x9f, 0xff, 0x7b, 0xff, 0xfd, 0xaf, 0x87, 0x1e, 0x2a, 0xcd, 0x9b, 0x9c, 0x7f,
	0x02, 0x96, 0x16, 0x5c, 0x94, 0xec, 0x64, 0x2f, 0x01, 0xcd, 0xf7, 0xec, 0x6b, 0xde, 0x40, 0x2a,
	0x9b, 0x8c, 0xd6, 0x8d, 0xd4, 0x12, 0x3f, 0xa8, 0x17, 0x49, 0x21, 0x52, 0xfe, 0x01, 0x94, 0x2c,
	0x81, 0x76, 0x44, 0x6a, 0x4a, 0xa9, 0x23, 0x8e, 0xee, 0xe5, 0x32, 0x97, 0x86, 0xc0, 0xda, 0x9b,
	0xe5, 0x8e, 0x48, 0x2a, 0x55, 0x29, 0x15, 0x4b, 0xb8, 0x82, 0x7f, 0x3d, 0xa4, 0xa8, 0xec, 0x7f,
	0xf4, 0xa3, 0x87, 0x36, 0xa7, 0xad, 0xce, 0xcc, 0x74, 0xc4, 0x8f, 0xd1, 0x35, 0x9e, 0x65, 0x0d,
	0x28, 0x35, 0xf4, 0xc7, 0xfe, 0xe4, 0x46, 0x8c, 0x2f, 0x96, 0xe1, 0xd6, 0x29, 0x2f, 0x8b, 0x67,
	0x91, 0xfb, 0x88, 0x66, 0x5d, 0x09, 0xfe, 0xe6, 0xa3, 0xa1, 0xa8, 0x84, 0x16, 0xbc, 0x98, 0x1b,
	0x37, 0x3c, 0x29, 0x60, 0xce, 0x4b, 0xb9, 0xa8, 0xf4, 0xb0, 0x37, 0xee, 0x4f, 0x36, 0x9f, 0xdc,
	0xa7, 0xd6, 0x01, 0x6d, 0x1d, 0x74, 0x66, 0xe9, 0x54, 0x8a, 0x2a, 0x3e, 0x3a, 0x5b, 0x86, 0xde,
	0xc5, 0x32, 0x0c, 0xad, 0xfc, 0x55, 0x42, 0xd1, 0xf7, 0x5f, 0xe1, 0x24, 0x17, 0xfa, 0xfd, 0x22,
	0xa1, 0xa9, 0x2c, 0x99, 0x9b, 0xc8, 0x1e, 0xbb, 0x2a, 0x3b, 0x66, 0xfa, 0xb4, 0x06, 0x65, 0x34,
	0xd5, 0x6c, 0xdb, 0xc9, 0x4c, 0x3b, 0x95, 0x03, 0x23, 0x82, 0x5f, 0xa1, 0x80, 0xa7, 0x5a, 0xc8,
	0x6a, 0x9e, 0xca, 0xb2, 0x2e, 0x40, 0x43, 0x36, 0x1c, 0x8c, 0xfb, 0x93, 0xeb, 0x71, 0xe8, 0x6c,
	0xec, 0xb8, 0x29, 0xff, 0xab, 0x8a, 0x66, 0xb7, 0x2d, 0x34, 0xed, 0x90, 0x47, 0x1a, 0x6d, 0x1c,
	0x18, 0x08, 0x6f, 0x23, 0x6c, 0x6f, 0x2f, 0x2f, 0x75, 0x0d, 0x3c, 0x1c, 0xa0, 0x9b, 0x16, 0x8f,
	0x45, 0xf6, 0xfa, 0xc5, 0x9b, 0xc0, 0xc7, 0x77, 0xd0, 0x2d, 0x8b, 0x1c, 0x8a, 0x4a, 0xb7, 0x50,
	0x0f, 0x6f, 0x21, 0x64, 0xa1, 0xb7, 0x52, 0x43, 0xd0, 0xc7, 0x3b, 0xe8, 0xae, 0x7d, 0x3f, 0x87,
	0x02, 0x72, 0xae, 0xe1, 0x48, 0xf3, 0x63, 0x08, 0x06, 0xa3, 0xc1, 0xe7, 0xaf, 0xc4, 0x8b, 0x0f,
	0xcf, 0x56, 0xc4, 0x3f, 0x5f, 0x11, 0xff, 0xf7, 0x8a, 0xf8, 0x5f, 0xd6, 0xc4, 0x3b, 0x5f, 0x13,
	0xef, 0xe7, 0x9a, 0x78, 0xef, 0xf6, 0x2f, 0x6d, 0xc7, 0x66, 0x65, 0xd7, 0x85, 0x85, 0xfd, 0x4d,
	0xd9, 0xc9, 0x53, 0xf6, 0xd1, 0x45, 0xcd, 0xac, 0x2b, 0xd9, 0x30, 0x01, 0xd8, 0xff, 0x13, 0x00,
	0x00, 0xff, 0xff, 0xbf, 0xd4, 0xc8, 0x64, 0x89, 0x02, 0x00, 0x00,
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionCompleted) > 0 {
		for iNdEx := len(m.ActionCompleted) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionCompleted[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintClaimRecord(dAtA, i, uint64(len(m.ActionCompleted)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.InitialClaimableAmount) > 0 {
		for iNdEx := len(m.InitialClaimableAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialClaimableAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaimRecord(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaimRecord(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaimRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaimRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaimRecord(uint64(l))
	}
	if len(m.InitialClaimableAmount) > 0 {
		for _, e := range m.InitialClaimableAmount {
			l = e.Size()
			n += 1 + l + sovClaimRecord(uint64(l))
		}
	}
	if len(m.ActionCompleted) > 0 {
		n += 1 + sovClaimRecord(uint64(len(m.ActionCompleted))) + len(m.ActionCompleted)*1
	}
	return n
}

func sovClaimRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaimRecord(x uint64) (n int) {
	return sovClaimRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaimRecord
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
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
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
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaimableAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaimRecord
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
				return ErrInvalidLengthClaimRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaimRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialClaimableAmount = append(m.InitialClaimableAmount, types.Coin{})
			if err := m.InitialClaimableAmount[len(m.InitialClaimableAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
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
				m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaimRecord
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaimRecord
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaimRecord
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionCompleted) == 0 {
					m.ActionCompleted = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaimRecord
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
					m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionCompleted", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaimRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaimRecord
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
func skipClaimRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaimRecord
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
					return 0, ErrIntOverflowClaimRecord
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
					return 0, ErrIntOverflowClaimRecord
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
				return 0, ErrInvalidLengthClaimRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaimRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaimRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaimRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaimRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaimRecord = fmt.Errorf("proto: unexpected end of group")
)
