// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: warden/warden/v1beta3/space.proto

package v1beta3

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// Space is a collection of users (called owners) that manages a set of Keys.
type Space struct {
	// Unique ID of the space.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Address of the creator of the space.
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// List of owners of the space.
	Owners []string `protobuf:"bytes,3,rep,name=owners,proto3" json:"owners,omitempty"`
	// Optional ID of the Rule to be applied to every *admin* operation.
	// If not specified, the default Rule is used.
	//
	// Admin operations are:
	// - warden.warden.Msg.AddSpaceOwner
	// - warden.warden.Msg.RemoveSpaceOwner
	// - warden.warden.Msg.UpdateSpace
	//
	// The default Rule is to allow any operation when at least one of its
	// owner approves it.
	AdminRuleId uint64 `protobuf:"varint,5,opt,name=admin_rule_id,json=adminRuleId,proto3" json:"admin_rule_id,omitempty"`
	// Optional ID of the Rule to be applied to every *sign* operation.
	// If not specified, the default Rule is used.
	//
	// Sign operations are:
	// - warden.warden.Msg.NewKeyRequest
	// - warden.warden.Msg.NewSignRequest
	// - warden.warden.Msg.UpdateKey
	//
	// The default Rule is to allow any operation when at least one of its
	// owner approves it.
	SignRuleId uint64 `protobuf:"varint,6,opt,name=sign_rule_id,json=signRuleId,proto3" json:"sign_rule_id,omitempty"`
	// Version of the space. Every time the Space is updated, this number gets increasead by one.
	Nonce uint64 `protobuf:"varint,7,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Optional ID of the Rule to be applied to every approve vote on *admin* operation.
	// If not specified, the default Rule is used.
	//
	// Admin operations are:
	// - warden.warden.Msg.AddSpaceOwner
	// - warden.warden.Msg.RemoveSpaceOwner
	// - warden.warden.Msg.UpdateSpace
	//
	// The default Rule is to allow any operation when at least one of its
	// owner approves it.
	ApproveAdminRuleId uint64 `protobuf:"varint,8,opt,name=approve_admin_rule_id,json=approveAdminRuleId,proto3" json:"approve_admin_rule_id,omitempty"`
	// Optional ID of the Rule to be applied to every reject vote on *admin* operation.
	// If not specified, the default Rule is used.
	//
	// Admin operations are:
	// - warden.warden.Msg.AddSpaceOwner
	// - warden.warden.Msg.RemoveSpaceOwner
	// - warden.warden.Msg.UpdateSpace
	//
	// The default Rule is to allow any operation when at least one of its
	// owner approves it.
	RejectAdminRuleId uint64 `protobuf:"varint,9,opt,name=reject_admin_rule_id,json=rejectAdminRuleId,proto3" json:"reject_admin_rule_id,omitempty"`
	// Optional ID of the Rule to be applied to every approve vote on *sign* operation.
	// If not specified, the default Rule is used.
	//
	// Sign operations are:
	// - warden.warden.Msg.NewKeyRequest
	// - warden.warden.Msg.NewSignRequest
	// - warden.warden.Msg.UpdateKey
	//
	// The default Rule is to allow any operation when at least one of its
	// owner approves it.
	ApproveSignRuleId uint64 `protobuf:"varint,10,opt,name=approve_sign_rule_id,json=approveSignRuleId,proto3" json:"approve_sign_rule_id,omitempty"`
	// Optional ID of the Rule to be applied to every reject vote on *sign* operation.
	// If not specified, the default Rule is used.
	//
	// Sign operations are:
	// - warden.warden.Msg.NewKeyRequest
	// - warden.warden.Msg.NewSignRequest
	// - warden.warden.Msg.UpdateKey
	//
	// The default Rule is to allow any operation when at least one of its
	// owner approves it.
	RejectSignRuleId uint64 `protobuf:"varint,11,opt,name=reject_sign_rule_id,json=rejectSignRuleId,proto3" json:"reject_sign_rule_id,omitempty"`
}

func (m *Space) Reset()         { *m = Space{} }
func (m *Space) String() string { return proto.CompactTextString(m) }
func (*Space) ProtoMessage()    {}
func (*Space) Descriptor() ([]byte, []int) {
	return fileDescriptor_da76bdfe3c1772b2, []int{0}
}
func (m *Space) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Space) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Space.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Space) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Space.Merge(m, src)
}
func (m *Space) XXX_Size() int {
	return m.Size()
}
func (m *Space) XXX_DiscardUnknown() {
	xxx_messageInfo_Space.DiscardUnknown(m)
}

var xxx_messageInfo_Space proto.InternalMessageInfo

func (m *Space) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Space) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Space) GetOwners() []string {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *Space) GetAdminRuleId() uint64 {
	if m != nil {
		return m.AdminRuleId
	}
	return 0
}

func (m *Space) GetSignRuleId() uint64 {
	if m != nil {
		return m.SignRuleId
	}
	return 0
}

func (m *Space) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Space) GetApproveAdminRuleId() uint64 {
	if m != nil {
		return m.ApproveAdminRuleId
	}
	return 0
}

func (m *Space) GetRejectAdminRuleId() uint64 {
	if m != nil {
		return m.RejectAdminRuleId
	}
	return 0
}

func (m *Space) GetApproveSignRuleId() uint64 {
	if m != nil {
		return m.ApproveSignRuleId
	}
	return 0
}

func (m *Space) GetRejectSignRuleId() uint64 {
	if m != nil {
		return m.RejectSignRuleId
	}
	return 0
}

func init() {
	proto.RegisterType((*Space)(nil), "warden.warden.v1beta3.Space")
}

func init() { proto.RegisterFile("warden/warden/v1beta3/space.proto", fileDescriptor_da76bdfe3c1772b2) }

var fileDescriptor_da76bdfe3c1772b2 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x4e, 0x02, 0x31,
	0x10, 0x86, 0xd9, 0x45, 0x40, 0x06, 0x35, 0x5a, 0xc1, 0xf4, 0xd4, 0xac, 0x9c, 0xb8, 0xc0, 0x86,
	0xf0, 0x04, 0x7a, 0x31, 0x5e, 0x97, 0x9b, 0x17, 0x52, 0xda, 0x09, 0xd6, 0xc0, 0x76, 0xd3, 0x2d,
	0xa0, 0x6f, 0xe1, 0x63, 0x79, 0xe4, 0xe8, 0xd1, 0xc0, 0xc9, 0xb7, 0x30, 0x74, 0xbb, 0x41, 0x38,
	0x75, 0xff, 0xf9, 0xbf, 0x7f, 0xff, 0x49, 0x06, 0xee, 0xd7, 0xdc, 0x48, 0x4c, 0x63, 0xff, 0xac,
	0x86, 0x53, 0xb4, 0x7c, 0x14, 0xe7, 0x19, 0x17, 0x38, 0xc8, 0x8c, 0xb6, 0x9a, 0x74, 0x0a, 0x6f,
	0xe0, 0x1f, 0x8f, 0x74, 0x7f, 0x43, 0xa8, 0x8d, 0xf7, 0x18, 0xb9, 0x82, 0x50, 0x49, 0x1a, 0x44,
	0x41, 0xef, 0x2c, 0x09, 0x95, 0x24, 0x14, 0x1a, 0xc2, 0x20, 0xb7, 0xda, 0xd0, 0x30, 0x0a, 0x7a,
	0xcd, 0xa4, 0x94, 0xe4, 0x0e, 0xea, 0x7a, 0x9d, 0xa2, 0xc9, 0x69, 0x35, 0xaa, 0xf6, 0x9a, 0x89,
	0x57, 0xa4, 0x0b, 0x97, 0x5c, 0x2e, 0x54, 0x3a, 0x31, 0xcb, 0x39, 0x4e, 0x94, 0xa4, 0x35, 0xf7,
	0xb3, 0x96, 0x1b, 0x26, 0xcb, 0x39, 0x3e, 0x4b, 0x12, 0xc1, 0x45, 0xae, 0x66, 0x07, 0xa4, 0xee,
	0x10, 0xd8, 0xcf, 0x3c, 0xd1, 0x86, 0x5a, 0xaa, 0x53, 0x81, 0xb4, 0xe1, 0xac, 0x42, 0x90, 0x21,
	0x74, 0x78, 0x96, 0x19, 0xbd, 0xc2, 0xc9, 0x71, 0xc7, 0xb9, 0xa3, 0x88, 0x37, 0x1f, 0xfe, 0x55,
	0xc5, 0xd0, 0x36, 0xf8, 0x86, 0xc2, 0x9e, 0x24, 0x9a, 0x2e, 0x71, 0x53, 0x78, 0x27, 0x81, 0xb2,
	0xe3, 0x68, 0x47, 0x28, 0x02, 0xde, 0x1b, 0x1f, 0x56, 0xed, 0xc3, 0xad, 0x6f, 0x38, 0xe2, 0x5b,
	0x8e, 0xbf, 0x2e, 0xac, 0x03, 0xfe, 0xc8, 0xbf, 0xb6, 0x2c, 0xd8, 0x6c, 0x59, 0xf0, 0xb3, 0x65,
	0xc1, 0xe7, 0x8e, 0x55, 0x36, 0x3b, 0x56, 0xf9, 0xde, 0xb1, 0xca, 0xcb, 0xd3, 0x4c, 0xd9, 0xd7,
	0xe5, 0x74, 0x20, 0xf4, 0xc2, 0xdf, 0xb0, 0xef, 0xae, 0x26, 0xf4, 0xdc, 0xeb, 0x13, 0x19, 0xbf,
	0x97, 0x1f, 0xf6, 0x23, 0xc3, 0xbc, 0xbc, 0xf8, 0xb4, 0xee, 0xb8, 0xd1, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x34, 0xbf, 0xba, 0x4a, 0x11, 0x02, 0x00, 0x00,
}

func (m *Space) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Space) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Space) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RejectSignRuleId != 0 {
		i = encodeVarintSpace(dAtA, i, uint64(m.RejectSignRuleId))
		i--
		dAtA[i] = 0x58
	}
	if m.ApproveSignRuleId != 0 {
		i = encodeVarintSpace(dAtA, i, uint64(m.ApproveSignRuleId))
		i--
		dAtA[i] = 0x50
	}
	if m.RejectAdminRuleId != 0 {
		i = encodeVarintSpace(dAtA, i, uint64(m.RejectAdminRuleId))
		i--
		dAtA[i] = 0x48
	}
	if m.ApproveAdminRuleId != 0 {
		i = encodeVarintSpace(dAtA, i, uint64(m.ApproveAdminRuleId))
		i--
		dAtA[i] = 0x40
	}
	if m.Nonce != 0 {
		i = encodeVarintSpace(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x38
	}
	if m.SignRuleId != 0 {
		i = encodeVarintSpace(dAtA, i, uint64(m.SignRuleId))
		i--
		dAtA[i] = 0x30
	}
	if m.AdminRuleId != 0 {
		i = encodeVarintSpace(dAtA, i, uint64(m.AdminRuleId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Owners) > 0 {
		for iNdEx := len(m.Owners) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Owners[iNdEx])
			copy(dAtA[i:], m.Owners[iNdEx])
			i = encodeVarintSpace(dAtA, i, uint64(len(m.Owners[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSpace(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintSpace(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSpace(dAtA []byte, offset int, v uint64) int {
	offset -= sovSpace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Space) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSpace(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSpace(uint64(l))
	}
	if len(m.Owners) > 0 {
		for _, s := range m.Owners {
			l = len(s)
			n += 1 + l + sovSpace(uint64(l))
		}
	}
	if m.AdminRuleId != 0 {
		n += 1 + sovSpace(uint64(m.AdminRuleId))
	}
	if m.SignRuleId != 0 {
		n += 1 + sovSpace(uint64(m.SignRuleId))
	}
	if m.Nonce != 0 {
		n += 1 + sovSpace(uint64(m.Nonce))
	}
	if m.ApproveAdminRuleId != 0 {
		n += 1 + sovSpace(uint64(m.ApproveAdminRuleId))
	}
	if m.RejectAdminRuleId != 0 {
		n += 1 + sovSpace(uint64(m.RejectAdminRuleId))
	}
	if m.ApproveSignRuleId != 0 {
		n += 1 + sovSpace(uint64(m.ApproveSignRuleId))
	}
	if m.RejectSignRuleId != 0 {
		n += 1 + sovSpace(uint64(m.RejectSignRuleId))
	}
	return n
}

func sovSpace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSpace(x uint64) (n int) {
	return sovSpace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Space) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpace
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
			return fmt.Errorf("proto: Space: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Space: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
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
				return ErrInvalidLengthSpace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owners", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
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
				return ErrInvalidLengthSpace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owners = append(m.Owners, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminRuleId", wireType)
			}
			m.AdminRuleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdminRuleId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignRuleId", wireType)
			}
			m.SignRuleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignRuleId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApproveAdminRuleId", wireType)
			}
			m.ApproveAdminRuleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApproveAdminRuleId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RejectAdminRuleId", wireType)
			}
			m.RejectAdminRuleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RejectAdminRuleId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApproveSignRuleId", wireType)
			}
			m.ApproveSignRuleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApproveSignRuleId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RejectSignRuleId", wireType)
			}
			m.RejectSignRuleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RejectSignRuleId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSpace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSpace
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
func skipSpace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSpace
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
					return 0, ErrIntOverflowSpace
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
					return 0, ErrIntOverflowSpace
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
				return 0, ErrInvalidLengthSpace
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSpace
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSpace
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSpace        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSpace          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSpace = fmt.Errorf("proto: unexpected end of group")
)
