// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: warden/async/v1beta1/future.proto

package v1beta1

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

type FutureVoteType int32

const (
	// Unspecified vote type.
	FutureVoteType_VOTE_TYPE_UNSPECIFIED FutureVoteType = 0
	// Vote to approve the result of the Future.
	FutureVoteType_VOTE_TYPE_VERIFIED FutureVoteType = 1
	// Vote to reject the result of the Future.
	FutureVoteType_VOTE_TYPE_REJECTED FutureVoteType = 2
)

var FutureVoteType_name = map[int32]string{
	0: "VOTE_TYPE_UNSPECIFIED",
	1: "VOTE_TYPE_VERIFIED",
	2: "VOTE_TYPE_REJECTED",
}

var FutureVoteType_value = map[string]int32{
	"VOTE_TYPE_UNSPECIFIED": 0,
	"VOTE_TYPE_VERIFIED":    1,
	"VOTE_TYPE_REJECTED":    2,
}

func (x FutureVoteType) String() string {
	return proto.EnumName(FutureVoteType_name, int32(x))
}

func (FutureVoteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9657bb2a8acf3037, []int{0}
}

// Future defines a future that will be executed in the future, asynchronously.
// One validator will add the result of the Future in the blockchain. Other
// validators will then be able to vote on the execution of the Future, whether
// they were able to verify the result or not.
type Future struct {
	// Unique ID of the Future.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Creator of the Future.
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// Unique name of the handler to be used to execute the Future.
	Handler string `protobuf:"bytes,3,opt,name=handler,proto3" json:"handler,omitempty"`
	// Input data to be used by the handler to execute the Future.
	// The actual format is determined by the handler being used.
	Input []byte `protobuf:"bytes,4,opt,name=input,proto3" json:"input,omitempty"`
}

func (m *Future) Reset()         { *m = Future{} }
func (m *Future) String() string { return proto.CompactTextString(m) }
func (*Future) ProtoMessage()    {}
func (*Future) Descriptor() ([]byte, []int) {
	return fileDescriptor_9657bb2a8acf3037, []int{0}
}
func (m *Future) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Future) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Future.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Future) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Future.Merge(m, src)
}
func (m *Future) XXX_Size() int {
	return m.Size()
}
func (m *Future) XXX_DiscardUnknown() {
	xxx_messageInfo_Future.DiscardUnknown(m)
}

var xxx_messageInfo_Future proto.InternalMessageInfo

func (m *Future) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Future) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Future) GetHandler() string {
	if m != nil {
		return m.Handler
	}
	return ""
}

func (m *Future) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

// FutureResult defines the result of the execution of a Future.
type FutureResult struct {
	// ID of the Future this result is for.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Output of the Future.
	// The actual format is determined by the handler being used.
	Output []byte `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	// Address of the validator that submitted the result.
	Submitter []byte `protobuf:"bytes,3,opt,name=submitter,proto3" json:"submitter,omitempty"`
}

func (m *FutureResult) Reset()         { *m = FutureResult{} }
func (m *FutureResult) String() string { return proto.CompactTextString(m) }
func (*FutureResult) ProtoMessage()    {}
func (*FutureResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_9657bb2a8acf3037, []int{1}
}
func (m *FutureResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FutureResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FutureResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FutureResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FutureResult.Merge(m, src)
}
func (m *FutureResult) XXX_Size() int {
	return m.Size()
}
func (m *FutureResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FutureResult.DiscardUnknown(m)
}

var xxx_messageInfo_FutureResult proto.InternalMessageInfo

func (m *FutureResult) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FutureResult) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *FutureResult) GetSubmitter() []byte {
	if m != nil {
		return m.Submitter
	}
	return nil
}

// FutureVote defines a vote on a Future execution.
type FutureVote struct {
	// ID of the Future this vote is for.
	FutureId uint64 `protobuf:"varint,1,opt,name=future_id,json=futureId,proto3" json:"future_id,omitempty"`
	// Address of the voter.
	Voter []byte `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`
	// Vote type.
	Vote FutureVoteType `protobuf:"varint,3,opt,name=vote,proto3,enum=warden.async.v1beta1.FutureVoteType" json:"vote,omitempty"`
}

func (m *FutureVote) Reset()         { *m = FutureVote{} }
func (m *FutureVote) String() string { return proto.CompactTextString(m) }
func (*FutureVote) ProtoMessage()    {}
func (*FutureVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_9657bb2a8acf3037, []int{2}
}
func (m *FutureVote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FutureVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FutureVote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FutureVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FutureVote.Merge(m, src)
}
func (m *FutureVote) XXX_Size() int {
	return m.Size()
}
func (m *FutureVote) XXX_DiscardUnknown() {
	xxx_messageInfo_FutureVote.DiscardUnknown(m)
}

var xxx_messageInfo_FutureVote proto.InternalMessageInfo

func (m *FutureVote) GetFutureId() uint64 {
	if m != nil {
		return m.FutureId
	}
	return 0
}

func (m *FutureVote) GetVoter() []byte {
	if m != nil {
		return m.Voter
	}
	return nil
}

func (m *FutureVote) GetVote() FutureVoteType {
	if m != nil {
		return m.Vote
	}
	return FutureVoteType_VOTE_TYPE_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("warden.async.v1beta1.FutureVoteType", FutureVoteType_name, FutureVoteType_value)
	proto.RegisterType((*Future)(nil), "warden.async.v1beta1.Future")
	proto.RegisterType((*FutureResult)(nil), "warden.async.v1beta1.FutureResult")
	proto.RegisterType((*FutureVote)(nil), "warden.async.v1beta1.FutureVote")
}

func init() { proto.RegisterFile("warden/async/v1beta1/future.proto", fileDescriptor_9657bb2a8acf3037) }

var fileDescriptor_9657bb2a8acf3037 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x4f, 0xea, 0x40,
	0x14, 0xed, 0xf4, 0xf1, 0x78, 0x8f, 0x1b, 0x42, 0xc8, 0x84, 0x47, 0xfa, 0xf2, 0x5e, 0x1a, 0x24,
	0x2e, 0x88, 0x89, 0x6d, 0xd0, 0x8d, 0x6b, 0x61, 0x48, 0x70, 0xa1, 0x64, 0xac, 0x24, 0xea, 0x02,
	0xfb, 0x31, 0x4a, 0x13, 0xe8, 0x34, 0xed, 0x14, 0xed, 0xbf, 0xf0, 0x67, 0xb9, 0x64, 0xe9, 0xd2,
	0xc0, 0x1f, 0x31, 0xed, 0x14, 0x09, 0xea, 0x6e, 0xce, 0xb9, 0xe7, 0x9e, 0x73, 0x67, 0xe6, 0xc2,
	0xde, 0xa3, 0x1d, 0x79, 0x2c, 0x30, 0xed, 0x38, 0x0d, 0x5c, 0x73, 0xd1, 0x75, 0x98, 0xb0, 0xbb,
	0xe6, 0x7d, 0x22, 0x92, 0x88, 0x19, 0x61, 0xc4, 0x05, 0xc7, 0x0d, 0x29, 0x31, 0x72, 0x89, 0x51,
	0x48, 0xda, 0x0e, 0x94, 0x07, 0xb9, 0x0a, 0xd7, 0x40, 0xf5, 0x3d, 0x0d, 0xb5, 0x50, 0xa7, 0x44,
	0x55, 0xdf, 0xc3, 0x1a, 0xfc, 0x72, 0x23, 0x66, 0x0b, 0x1e, 0x69, 0x6a, 0x0b, 0x75, 0x2a, 0x74,
	0x03, 0xb3, 0xca, 0xd4, 0x0e, 0xbc, 0x19, 0x8b, 0xb4, 0x1f, 0xb2, 0x52, 0x40, 0xdc, 0x80, 0x9f,
	0x7e, 0x10, 0x26, 0x42, 0x2b, 0xb5, 0x50, 0xa7, 0x4a, 0x25, 0x68, 0x5b, 0x50, 0x95, 0x19, 0x94,
	0xc5, 0xc9, 0x4c, 0x7c, 0x49, 0x6a, 0x42, 0x99, 0x27, 0x22, 0x6b, 0x53, 0xf3, 0xb6, 0x02, 0xe1,
	0xff, 0x50, 0x89, 0x13, 0x67, 0xee, 0x0b, 0x51, 0x24, 0x55, 0xe9, 0x96, 0x68, 0xa7, 0x00, 0xd2,
	0x75, 0xcc, 0x05, 0xc3, 0xff, 0xa0, 0x22, 0x6f, 0x3b, 0xf9, 0xb0, 0xfe, 0x2d, 0x89, 0xa1, 0x97,
	0x8d, 0xb5, 0xe0, 0x99, 0x89, 0xf4, 0x97, 0x00, 0x9f, 0x40, 0x29, 0x3b, 0xe4, 0xce, 0xb5, 0xa3,
	0x7d, 0xe3, 0xbb, 0xf7, 0x31, 0xb6, 0x11, 0x56, 0x1a, 0x32, 0x9a, 0x77, 0x1c, 0xdc, 0x42, 0x6d,
	0x97, 0xc7, 0x7f, 0xe1, 0xcf, 0xf8, 0xc2, 0x22, 0x13, 0xeb, 0x7a, 0x44, 0x26, 0x57, 0xe7, 0x97,
	0x23, 0xd2, 0x1b, 0x0e, 0x86, 0xa4, 0x5f, 0x57, 0x70, 0x13, 0xf0, 0xb6, 0x34, 0x26, 0x54, 0xf2,
	0x68, 0x97, 0xa7, 0xe4, 0x8c, 0xf4, 0x2c, 0xd2, 0xaf, 0xab, 0xa7, 0x77, 0x2f, 0x2b, 0x1d, 0x2d,
	0x57, 0x3a, 0x7a, 0x5b, 0xe9, 0xe8, 0x79, 0xad, 0x2b, 0xcb, 0xb5, 0xae, 0xbc, 0xae, 0x75, 0xe5,
	0x66, 0xf0, 0xe0, 0x8b, 0x69, 0xe2, 0x18, 0x2e, 0x9f, 0x9b, 0x72, 0xd8, 0xc3, 0xfc, 0x6b, 0x5d,
	0x3e, 0x2b, 0xf0, 0x27, 0x68, 0x3e, 0x15, 0x0b, 0x21, 0xd2, 0x90, 0xc5, 0x9b, 0xb5, 0x70, 0xca,
	0xb9, 0xec, 0xf8, 0x3d, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x03, 0xa0, 0x36, 0x35, 0x02, 0x00, 0x00,
}

func (m *Future) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Future) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Future) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Input) > 0 {
		i -= len(m.Input)
		copy(dAtA[i:], m.Input)
		i = encodeVarintFuture(dAtA, i, uint64(len(m.Input)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Handler) > 0 {
		i -= len(m.Handler)
		copy(dAtA[i:], m.Handler)
		i = encodeVarintFuture(dAtA, i, uint64(len(m.Handler)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintFuture(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintFuture(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FutureResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FutureResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FutureResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Submitter) > 0 {
		i -= len(m.Submitter)
		copy(dAtA[i:], m.Submitter)
		i = encodeVarintFuture(dAtA, i, uint64(len(m.Submitter)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Output) > 0 {
		i -= len(m.Output)
		copy(dAtA[i:], m.Output)
		i = encodeVarintFuture(dAtA, i, uint64(len(m.Output)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintFuture(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FutureVote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FutureVote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FutureVote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Vote != 0 {
		i = encodeVarintFuture(dAtA, i, uint64(m.Vote))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintFuture(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x12
	}
	if m.FutureId != 0 {
		i = encodeVarintFuture(dAtA, i, uint64(m.FutureId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintFuture(dAtA []byte, offset int, v uint64) int {
	offset -= sovFuture(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Future) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovFuture(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovFuture(uint64(l))
	}
	l = len(m.Handler)
	if l > 0 {
		n += 1 + l + sovFuture(uint64(l))
	}
	l = len(m.Input)
	if l > 0 {
		n += 1 + l + sovFuture(uint64(l))
	}
	return n
}

func (m *FutureResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovFuture(uint64(m.Id))
	}
	l = len(m.Output)
	if l > 0 {
		n += 1 + l + sovFuture(uint64(l))
	}
	l = len(m.Submitter)
	if l > 0 {
		n += 1 + l + sovFuture(uint64(l))
	}
	return n
}

func (m *FutureVote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FutureId != 0 {
		n += 1 + sovFuture(uint64(m.FutureId))
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovFuture(uint64(l))
	}
	if m.Vote != 0 {
		n += 1 + sovFuture(uint64(m.Vote))
	}
	return n
}

func sovFuture(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFuture(x uint64) (n int) {
	return sovFuture(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Future) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFuture
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
			return fmt.Errorf("proto: Future: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Future: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFuture
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
					return ErrIntOverflowFuture
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
				return ErrInvalidLengthFuture
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFuture
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFuture
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
				return ErrInvalidLengthFuture
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFuture
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handler = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Input", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFuture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFuture
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFuture
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Input = append(m.Input[:0], dAtA[iNdEx:postIndex]...)
			if m.Input == nil {
				m.Input = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFuture(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFuture
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
func (m *FutureResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFuture
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
			return fmt.Errorf("proto: FutureResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FutureResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFuture
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
				return fmt.Errorf("proto: wrong wireType = %d for field Output", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFuture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFuture
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFuture
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Output = append(m.Output[:0], dAtA[iNdEx:postIndex]...)
			if m.Output == nil {
				m.Output = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Submitter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFuture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFuture
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFuture
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Submitter = append(m.Submitter[:0], dAtA[iNdEx:postIndex]...)
			if m.Submitter == nil {
				m.Submitter = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFuture(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFuture
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
func (m *FutureVote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFuture
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
			return fmt.Errorf("proto: FutureVote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FutureVote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FutureId", wireType)
			}
			m.FutureId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFuture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FutureId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFuture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFuture
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFuture
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voter = append(m.Voter[:0], dAtA[iNdEx:postIndex]...)
			if m.Voter == nil {
				m.Voter = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			m.Vote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFuture
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vote |= FutureVoteType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFuture(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFuture
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
func skipFuture(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFuture
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
					return 0, ErrIntOverflowFuture
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
					return 0, ErrIntOverflowFuture
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
				return 0, ErrInvalidLengthFuture
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFuture
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFuture
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFuture        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFuture          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFuture = fmt.Errorf("proto: unexpected end of group")
)
